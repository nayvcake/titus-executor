#define _GNU_SOURCE
#include <errno.h>
#include <stdio.h>
#include <stdlib.h>
/* getaddrinfo */
#include <arpa/inet.h>
#include <netdb.h>
#include <string.h>
/* setns */
#include <sched.h>
/* for mount syscalls */
#include <asm-generic/unistd.h>
#include <linux/mount.h>
/* fcntl */
#include "scm_rights.h"
#include <assert.h>
#include <fcntl.h>
#include <sys/wait.h>
#include <unistd.h>
/* mkdir */
#include <linux/limits.h>
#include <signal.h>
#include <sys/stat.h>

#include "common.h"

#define E(x)                                                                   \
	do {                                                                   \
		if ((x) == -1) {                                               \
			perror(#x);                                            \
			exit(1);                                               \
		}                                                              \
	} while (0)

static void check_messages(int fd)
{
	char buf[4096];
	int err, n;
	err = errno;
	for (;;) {
		n = read(fd, buf, sizeof(buf));
		if (n < 0)
			break;
		n -= 2;
		switch (buf[0]) {
		case 'e':
			fprintf(stderr, "Error: %*.*s\n", n, n, buf + 2);
			break;
		case 'w':
			fprintf(stderr, "Warning: %*.*s\n", n, n, buf + 2);
			break;
		case 'i':
			fprintf(stderr, "Info: %*.*s\n", n, n, buf + 2);
			break;
		}
	}
	errno = err;
}

static __attribute__((noreturn)) void mount_error(int fd, const char *s)
{
	check_messages(fd);
	fprintf(stderr, "titus-mount-block-device mount error on '%s': %m\n", s);
	exit(1);
}

static inline int pidfd_open(pid_t pid, unsigned int flags)
{
	return syscall(__NR_pidfd_open, pid, flags);
}

static inline int fsopen(const char *fs_name, unsigned int flags)
{
	return syscall(__NR_fsopen, fs_name, flags);
}

static inline int fsmount(int fsfd, unsigned int flags, unsigned int ms_flags)
{
	return syscall(__NR_fsmount, fsfd, flags, ms_flags);
}

static inline int fsconfig(int fsfd, unsigned int cmd, const char *key,
			   const void *val, int aux)
{
	return syscall(__NR_fsconfig, fsfd, cmd, key, val, aux);
}

static inline int move_mount(int from_dfd, const char *from_pathname,
			     int to_dfd, const char *to_pathname,
			     unsigned int flags)
{
	return syscall(__NR_move_mount, from_dfd, from_pathname, to_dfd,
		       to_pathname, flags);
}

#define E_fsconfig(fd, cmd, key, val, aux)                                     \
	do {                                                                   \
		if (fsconfig(fd, cmd, key, val, aux) == -1)                    \
			mount_error(fd, key ?: "create");                      \
	} while (0)

static void process_option(char *option, int fsfd)
{
	/* Splits up a k=v string and runs fsconfig on it.
           We supply all the inputs here, so it is safe, but parsing things
           like this is a little dangerous */
	char *key, *value, *saveptr;
	key = strtok_r(option, "=", &saveptr);
	option = NULL;
	value = strtok_r(option, "=", &saveptr);
	fprintf(stderr, "titus-mount-block-device: Setting filesystem mount option %s=%s\n",
		key, value);
	E_fsconfig(fsfd, FSCONFIG_SET_STRING, key, value, 0);
}

static void do_fsconfigs(int fsfd, char *options)
{
	char *str1, *token, *saveptr;
	/* Mount options come in in the classic comma-separated key=value pairs
	   we need to split them up and pass them in for fsconfig to handle one at a time */
	for (str1 = options;; str1 = NULL) {
		token = strtok_r(str1, ",", &saveptr);
		if (token == NULL)
			break;
		process_option(token, fsfd);
	}
	/* This last FSCONFIG_CMD_CREATE fsconfig call actually creates the superblock */
	E_fsconfig(fsfd, FSCONFIG_CMD_CREATE, NULL, NULL, 0);
}

static int setup_fsfd_in_namespaces(int sk, int nsfd, const char *fstype)
{
	int fsfd, ret;
	int usernsfd = openat(nsfd, "user", O_RDONLY);
	int mnt_fd = openat(nsfd, "mnt", O_RDONLY);
	assert(usernsfd != -1);
	assert(mnt_fd != -1);
	ret = setns(usernsfd, CLONE_NEWUSER);
	if (ret == -1) {
		perror("setns user");
		return 1;
	}
	ret = setns(mnt_fd, CLONE_NEWNS);
	if (ret == -1) {
		perror("setns mnt");
		return 1;
	}
	fsfd = fsopen(fstype, 0);
	if (fsfd == -1) {
		perror("fsopen");
		return 1;
	}
	assert(send_fd(sk, fsfd) == 0);
	return 0;
}

static int fork_and_get_fsfd(long int nsfd, const char *fstype)
{
	int sk_pair[2], ret, fsfd, status;
	pid_t worker;

	if (socketpair(PF_LOCAL, SOCK_SEQPACKET, 0, sk_pair) < 0) {
		perror("socketpair");
		exit(1);
	}
	worker = fork();
	if (worker < 0) {
		perror("fork");
		exit(1);
	}
	if (worker == 0) {
		close(sk_pair[0]);
		ret = setup_fsfd_in_namespaces(sk_pair[1], nsfd, fstype);
		close(sk_pair[1]);
		exit(ret);
	}
	close(sk_pair[1]);
	fsfd = recv_fd(sk_pair[0]);
	assert(fsfd >= 0);
	if (waitpid(worker, &status, 0) != worker) {
		perror("waitpid");
		exit(1);
	}
	if (!WIFEXITED(status) || WEXITSTATUS(status)) {
		fprintf(stderr, "worker exited nonzero\n");
		exit(1);
	}
	close(sk_pair[0]);
	return fsfd;
}

static void switch_namespaces(int nsfd)
{
	int ret;
	int mnt_fd = openat(nsfd, "mnt", O_RDONLY);
	assert(mnt_fd != -1);
	ret = setns(mnt_fd, CLONE_NEWNS);
	if (ret == -1) {
		perror("setns mnt");
		exit(1);
	}
}

static void mount_and_move(int fsfd, const char *target, unsigned long flags)
{
	int mfd = fsmount(fsfd, 0, flags);
	if (mfd < 0) {
		mount_error(fsfd, "fsmount");
		E(close(fsfd));
	}
	E(move_mount(mfd, "", AT_FDCWD, target, MOVE_MOUNT_F_EMPTY_PATH));
	E(close(mfd));
}

int main(int argc, char *argv[])
{
	int nsfd, fsfd;
	long int container_pid;
	unsigned long flags_ul;
	/*
	 * We do this because parsing args is a bigger pain than passing
	 * via environment variable, although passing via environment
	 * variable has a "cost" in that they are limited in size
	 */
	const char *target = getenv("MOUNT_TARGET");
	const char *flags = getenv("MOUNT_FLAGS");
	char *options = getenv("MOUNT_OPTIONS");
	const char *fstype = getenv("MOUNT_FSTYPE");
	const char *pid1dir = getenv("TITUS_PID_1_DIR");

	errno = 0;
	assert(errno == 0);

	if (!(target && flags && options && fstype && pid1dir)) {
		fprintf(stderr,
			"Usage: must provide MOUNT_TARGET, MOUNT_FLAGS, and MOUNT_OPTIONS, MOUNT_FSTYPE, TITUS_PID_1_DIR env vars");
		return 1;
	}

	errno = 0;
	flags_ul = strtoul(flags, NULL, 10);
	if (errno) {
		perror("flags");
		return 1;
	}

	/* This nsfd is used to extract other namespaces to switch to, similar to a pidfd */
	char pid1_ns_path[PATH_MAX];
	snprintf(pid1_ns_path, PATH_MAX - 1, "%s/ns", pid1dir);
	nsfd = open(pid1_ns_path, O_PATH);
	if (nsfd == -1) {
		perror("pidfd_open");
		return 1;
	}

	/* First we need to get a fsfd, but it must be created inside the user namespace */
	fsfd = fork_and_get_fsfd(nsfd, fstype);

	fprintf(stderr, "titus-mount-block-device: user-inputed options: %s\n", options);

	/* Now we can do the fs_config calls and actual mount
	This isn't possible *inside* the conatiner, because it can't see the device */
	do_fsconfigs(fsfd, options);

	/* And the last step we can finally switch namespaces so the mount
	looks right from the container's perspective */
	switch_namespaces(nsfd);
	close(nsfd);
	mkdir_p(target);
	mount_and_move(fsfd, target, flags_ul);

	fprintf(stderr, "titus-mount-block-device: All done, mounted on %s\n", target);
	return 0;
}
