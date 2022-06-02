// Code generated by go-bindata. DO NOT EDIT.
// sources:
// default.json (17.888kB)
// fuse-container.json (18.247kB)
// allow-perf-syscalls.json (18.022kB)

package seccomp

import (
	"bytes"
	"compress/gzip"
	"crypto/sha256"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("read %q: %w", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %w", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes  []byte
	info   os.FileInfo
	digest [sha256.Size]byte
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _defaultJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x5b\x5b\x6f\x5b\xb9\x11\x7e\xdf\x5f\x21\xe8\x39\x28\x7c\xab\xd7\xce\x9b\xaa\xb8\x8d\xb1\x76\x9c\xda\x29\x76\x17\x45\x40\xd0\x3c\x73\x8e\x58\xf1\x66\x5e\x64\x0b\x41\xfe\x7b\xc1\x73\x24\x59\xf2\x0c\x4f\xbc\x2b\xa7\x51\x01\x3f\x24\xb6\xbf\x8f\x97\xe1\x90\x9c\x19\x0e\x79\xbe\xfc\x34\x18\x0c\x86\xdc\x8b\xc9\x25\x77\xc3\xb7\x83\x7f\xe7\xbf\x07\x83\xc1\x97\xc5\xcf\x25\x2b\x23\x88\x98\x3c\x0c\xdf\x0e\x86\x37\xe3\xcb\x8f\x6c\x74\x3d\x7e\xcf\x7e\x3b\x39\x66\xc7\x47\xc3\x37\x6b\x85\x43\xba\x1d\xad\x95\x0f\x6b\x6d\x2e\x4a\x6c\x54\x5f\xaf\xfb\x84\x3c\x3c\x18\x3e\x72\x9f\x17\xbf\x7e\x7d\xf3\x87\x04\x1c\xe5\xff\xb7\x90\x70\x74\x7d\xb9\xbd\x10\x97\xe7\x1f\x6f\xb6\x90\x21\x57\x2f\xab\xa9\x6b\xfc\xc3\x4b\x28\xeb\xb1\xa9\xef\x29\xea\xcb\xc8\x79\x76\xb1\xa5\x46\xcf\x2e\xfa\x05\xcd\x1d\xbc\x98\x56\x97\x8d\x7d\x6f\x81\xb7\x97\xf6\xe6\xf0\x74\xef\xb7\x3f\x2d\x67\xae\x4d\xc9\x90\x7f\x7e\x6e\x1b\x1d\x56\x50\xf3\xa4\xe2\x48\x44\x69\xcd\x63\xdf\xe3\x4f\xec\xec\xfa\xfa\xc3\xd5\x70\xa3\xd4\x99\xf7\xc6\x5e\x43\x1c\xbe\x1d\xec\x77\x44\x98\x07\xc1\x95\x0a\x25\x53\x85\x9b\x1d\x5d\x5c\x5c\xfd\xba\x31\x22\x69\x84\x4a\x55\x3b\x92\x2f\x9b\x23\x11\xdc\xe1\xf1\x65\x62\x3c\xfa\xc8\x6e\x7e\xbf\x61\xa3\x77\x97\xe7\x1f\x86\x1b\xf4\xe7\xc7\xbf\xbe\xae\xf7\x62\xb8\xa6\x94\x75\xeb\xea\xa7\xf3\x28\x94\x35\xf0\x14\xac\xb9\xb1\x51\xd6\x73\x26\x8d\x8c\x88\x0c\xc2\x9a\x5a\x36\x18\xd7\x36\x19\xa2\xb8\x75\x60\x30\xea\xa4\x98\x3e\x45\x95\xb5\xd3\xe4\x58\x25\xac\x9d\x4a\x24\x15\xd9\xbc\xb6\x33\x60\x24\x93\x75\xc0\xa2\x65\x13\x6e\x2a\x05\x8c\xa3\x02\x59\x2e\x16\x3d\xa0\x8e\x1c\xf8\x9a\xc1\x0c\x4c\x64\x94\xec\x77\xc9\x46\x2e\xa2\x7a\x8a\x07\x88\x95\xd5\x5c\x9a\xdc\x33\x41\x4e\x6c\x88\x05\xca\x04\x04\xce\x83\xb2\x48\xc7\x89\x1c\x69\x87\x1e\x20\xd8\x84\x09\xf7\xf0\xec\x7d\xf9\x9c\xe5\x5b\x58\x58\x5c\x08\x70\x48\xae\x0e\x3d\xa2\xe0\x80\x06\xcc\xab\xff\x44\xa9\xe1\x01\xe1\x8a\x7b\xfd\x14\xbc\x95\xa6\x42\x98\x47\xeb\x49\x70\xd7\x00\x12\x2b\x6f\x34\x02\x9d\x54\xd2\x63\x50\x5b\xd4\x91\x98\xd8\x7b\xb4\x28\x5a\xf0\x10\xcd\x81\x50\x56\x4c\xd9\x62\x6c\xbd\xe4\x31\xd2\x53\x47\x37\x10\xb3\xe1\xeb\xe1\xd8\x37\xaa\x97\xbb\x5e\x90\xa5\xba\x86\x1b\x1b\x14\x80\xfb\x06\x5d\x16\x20\x50\x1d\x07\x60\x9e\x9b\x06\x53\xd6\x18\x10\x78\x62\xac\x9b\xb3\x5a\xaa\x52\x2d\x0f\x78\x6b\x57\x09\xc9\x5c\x25\x87\x66\xa7\x4a\xee\xf0\x29\x06\xce\x2a\xc5\xda\x56\x51\x5f\xeb\xdc\x7e\x81\xc4\x76\x61\x45\x30\xab\xd0\x62\xea\x48\x77\xcf\xb1\xa1\x5d\xa3\x90\xe4\x1d\x57\xae\x95\x19\xb2\xbb\x6c\xd5\xea\x12\x8c\xbb\x79\x00\x31\xc3\x6a\x68\x51\xac\x75\x78\x20\xc4\x79\x90\x91\x35\xde\xe2\x09\xa9\x3b\x3b\x80\x9b\x59\x11\x48\x9c\x9a\x57\x33\x19\x88\xa5\xb6\x22\x18\xc5\x29\x65\x05\x31\x9b\x2b\x2f\xa7\x39\x36\x1d\x35\x69\x0f\x6a\xd2\x20\x74\x28\x31\x10\xd2\x52\xd4\x05\x53\xd1\xe1\x54\x2b\x06\xaf\xa9\x16\x24\xc6\x5a\xf1\xc8\xc3\xdc\x08\x44\x34\x10\x1f\x78\x8c\x78\x44\x4a\x86\x12\x63\xb1\x87\xae\x2d\xa1\x2a\x0f\xd9\x0b\xd3\x8d\x84\x52\xbf\x21\x12\x43\xcd\x20\x31\xaa\x0c\x17\x89\x1a\x19\xc7\x05\x4c\x95\xa7\x54\x13\x7d\x32\xe4\xfa\x58\x12\x44\x43\x29\x62\x4f\xd5\x82\x05\x5b\x58\xa7\x8c\x13\x8b\xbd\x81\x28\x5c\xa2\xd0\x7b\xb4\xce\x1a\x88\x15\x98\x88\xc6\xbb\xc4\x71\xb7\x0d\x44\x68\x24\xd5\x50\x86\xf1\x1a\xcc\x44\xa2\xcb\xa7\x42\x79\xba\xf9\x52\xeb\xad\x25\xa0\x06\xd0\x11\x64\x1d\x99\x55\x87\x16\x50\x03\xd1\x01\x78\x2a\xa0\xca\x14\x2d\x96\x6b\x3c\xb2\x43\x19\xa6\x0b\x17\x60\x2f\xad\x97\x71\x4e\x50\x9e\x9b\xca\xa2\x60\xa5\xf3\xd2\xb4\x40\x1d\x41\x0e\xdb\x43\xa0\xa7\xa2\x23\xe8\x3a\x4a\x6a\x6c\x83\x1b\x88\xcc\xdb\xdb\x14\x22\xcb\xbb\x9d\xaa\x97\x02\xc7\xae\xb5\x81\x18\x48\x09\x82\x15\xd3\x82\xe2\x33\x65\x71\x28\x98\x45\x88\x13\x0f\xbc\x62\xdc\x03\x27\xe8\x48\xf6\x94\xa7\xde\xd6\x15\xa7\xb4\x4d\x6b\xa7\xa4\x1a\xd2\x0a\xc9\x85\x07\xe0\x55\xc5\xee\x79\x14\x93\x52\x01\xea\x20\xb4\xce\xa1\x68\x60\x49\x7a\x5d\x68\xd8\x32\xc1\x8d\x00\x64\xd9\xa5\x25\x42\x08\x69\x59\x05\x21\x7a\x8b\xf4\x20\x6d\x8e\xe2\x5a\xff\x8d\xb6\x95\xb4\xcc\x3d\x8f\x2c\x58\x2d\x69\xf3\x62\x67\x44\x0c\xbd\x60\x88\x38\xba\x43\xb1\xbb\xcf\x78\xba\x25\x56\xa7\xb4\x2c\x79\x69\x1a\x06\x26\xe2\x6d\xbe\x62\x3d\x34\x32\xf4\x15\xa0\x7b\x75\xc8\xde\x4f\xa5\x42\xfa\x55\xa4\xa7\x56\x05\x4f\xad\x4a\xcb\x49\x49\x83\xcf\xb4\xd2\x4c\xb1\xd9\xcf\xfb\x10\x9f\x2c\x8b\xbe\x58\x15\x19\xa6\x54\x00\xc0\xbd\xf6\xf8\x64\xba\x42\xc9\x51\x2b\xca\x51\x2b\xda\x51\xeb\x2e\x0c\x43\x30\xe8\x5b\xee\xbd\xc4\x93\xa7\x41\xd7\x55\x21\xd8\xd6\xd2\x08\xeb\x31\x3c\x25\xe2\xb2\x16\xc4\x62\xea\xa9\xc1\xd1\x5a\x0b\x12\x65\xa9\x70\xa7\x05\xd1\xfc\xb7\x28\xc7\xab\x48\x6b\x8e\x16\x60\xc6\x70\x03\xce\xdb\x48\x1c\x75\xf4\x5d\xde\x6b\x01\x22\x35\x11\xfa\x8e\x75\x36\x85\x20\xa8\x34\x85\xbe\x6b\x77\x75\xe5\x41\x80\xc4\x41\xfc\x13\xba\x60\x01\x96\xa5\x02\xe0\x03\xf7\x3a\x57\xae\x9e\x0c\xb5\x2b\xb4\x07\x4a\x59\xa1\x21\x8c\x9f\x0e\x0d\x61\x81\x74\x68\xbc\x98\x11\x68\x20\x24\xa5\xe2\x3e\x9d\x0c\x39\xe5\x1d\x4c\x4d\x6f\x32\x84\xcc\xc5\x43\xb2\x81\xfb\x45\xe0\x8a\xf6\xac\x81\xfb\x00\x8a\x58\x00\xd4\x3c\x66\x8c\x4e\x5e\x11\x27\x24\xc7\x13\xde\x7f\x4e\x56\x75\x45\xae\x91\x8e\x69\x27\x30\xc8\xc6\x70\x34\x66\x27\x1d\xd1\x9c\x03\xdc\xb1\xc5\x0a\x73\x45\xb0\xb0\x5c\x9c\x27\xa6\xdf\xe5\xa8\x81\x2a\x0b\xbc\x42\xf3\xdf\xa1\x58\xba\x2e\x30\x22\x5a\xe9\xe6\x01\x5d\x05\x2d\xf1\x92\xa0\xf7\x5e\x52\xe7\x82\x0e\xc7\x52\x75\x30\x12\x2b\x0b\x4b\x61\x7c\x52\x20\xa8\xad\xb4\xc4\xf1\x1a\x21\x35\xe4\x01\xef\x9a\x8c\xd5\x1e\x47\xae\x19\xd7\x3a\xa0\x34\xe4\x12\x2f\x68\xa7\xa5\xa9\x5a\x9a\xbb\x2e\x93\xe3\x78\x83\x73\x5a\x3d\x2e\xcb\x03\x15\x6e\x76\x28\x35\xee\x0e\x27\xf4\x1d\x22\xf7\x91\x2d\x12\xf9\x88\xd6\x84\x6b\xf1\x01\xee\x10\x16\xf3\x76\x59\x24\x4c\x49\xce\x81\xa9\xa4\xc1\x2a\xe8\x48\x6f\x85\xe6\x01\xcf\x64\xcb\xde\x25\x48\x20\x4d\x6d\x69\xda\x43\x4c\xbe\xd0\x6b\x48\xc1\x11\x86\xba\x23\x5b\x5b\x4d\x65\x8c\x9e\xd0\xa5\x49\x8d\x2c\x36\x7d\xe2\x05\x31\x81\x2a\xfb\x2f\x5e\xd7\x39\x24\x46\x7e\xea\xb1\x00\x31\xc3\x2b\xd2\x71\xcf\xd1\x42\x5c\xb1\x6c\x79\xfa\x62\x9a\xa3\x03\x38\x55\x4a\x22\x55\xad\x4a\xb5\xbf\x24\x85\xa3\x92\xae\x84\xf7\x6d\x53\x32\x47\xa5\x33\x6c\x19\xc9\x42\x05\xdd\x75\x65\xc3\xb7\x54\x53\x70\xfc\x2b\xb2\x47\x35\xe1\x9b\xc3\x99\x4b\xc0\xe9\xc0\x00\x42\x58\x8d\x5c\x17\xed\x9c\x02\x68\xf2\xd6\x43\x13\xde\x39\x80\xb6\x44\xb3\xba\x5d\x65\x7d\x4c\x49\x83\xc4\xb2\xce\x58\xb6\x26\x25\x9c\x6e\x85\xb2\x67\x2d\x4e\xc3\x11\xaf\x73\x88\x35\x75\x9a\x5f\xe2\xf8\xbc\xd0\x32\xc4\x69\x75\x89\x93\x35\xe8\x1e\x4a\xed\xd3\x79\x95\x50\xce\xab\x84\x52\x5e\x25\xd0\xc9\x93\x50\x4e\x7c\x04\x88\x9e\xca\x32\x2d\x71\xb2\x73\x3a\x23\x12\xca\x19\x91\x50\xca\x88\x84\x72\x46\xa4\xa5\x4a\x55\x8a\x35\xc8\x1c\x4a\xe8\xcf\xa1\x04\x32\x55\x12\x8a\xf9\x90\xd0\x9f\x0f\x69\x69\x59\x31\x5e\x55\x9e\xb8\x2b\x0b\x64\xf2\x23\x14\x92\x1f\xa5\x93\x5d\x98\x68\xec\x3b\xc3\x84\xdc\xe4\x13\x5d\x51\x45\xa9\x9d\x3f\x49\xb1\x22\x0e\xd3\xd9\x61\xaa\x18\x22\xc7\x21\x77\x17\x7a\xe2\x5b\x89\x25\x8e\xf7\x71\xd9\x87\x16\x3d\x64\x9e\x08\x42\xda\x16\xa5\x82\x81\x8e\x71\x1c\x47\x04\xc1\x29\x29\xb0\xd5\x21\x02\x7d\xfa\x94\x4c\x27\xad\x4b\x39\xeb\x8c\x63\x4f\x37\xd7\x54\x28\xb8\x80\x09\x41\x88\xf3\x4f\xc6\x7a\xee\xd5\x32\x4d\x48\x39\x0f\x94\xef\x8f\xf8\x12\x3d\x36\x54\xa2\x85\xba\x8b\x6c\xcd\x50\x21\x09\xd0\x71\x15\x28\x28\x71\x0d\x44\x3b\x03\xef\x13\x9a\xf1\x15\x5f\xee\xb4\x78\x01\xda\xd1\xa1\xaf\x6e\xe8\xad\x5b\x4c\x6b\x2c\xd9\x3e\xb9\x1e\xe9\x72\xeb\x7d\xb2\x3d\xd2\x74\x7d\x34\xad\x91\x9c\xab\xc2\xbd\x48\xf9\x5a\x24\x15\x53\xd0\x89\xda\xab\x89\x0a\xea\xe9\x7c\x41\x87\xe2\x85\x9d\x28\x1d\xb4\xa0\x21\x2e\x5b\x56\x44\x21\xca\x48\xa4\x72\x66\xd4\xa5\xd7\x4c\xd3\x66\x20\x47\xd0\xa8\xdd\x0c\x62\x7b\x9d\x51\xe2\x8a\xa1\x3d\x2b\x92\xe0\xec\x45\x9f\x72\x94\x5f\x22\x69\x69\x7e\x01\x6f\x40\xe5\x16\x8e\xfe\x72\xb2\xd6\xed\x73\x5e\x19\x65\xd3\x0c\x21\xb0\x99\x66\x85\x33\xfa\x8a\x2f\x1c\x97\xa3\xe7\xe2\x65\xdf\xad\x70\xdf\x60\x49\x37\x87\xdd\xe9\xa4\x82\x87\xe1\xdb\xc1\xde\x1b\x44\x59\xb7\xea\x22\xff\x3b\xfb\xe7\x10\x97\x99\x71\x95\x20\x57\xdf\x60\xbe\xae\x0d\xe4\x39\xea\x03\x1f\xac\xe1\x2a\x47\x5a\xff\xbf\x2a\x38\x79\x55\xc1\xfe\xe1\xfe\xde\xcf\x07\xaf\x7a\xc8\x7a\x38\x79\xdd\x12\x83\xa3\x83\xd3\xa3\xd3\xe3\x9f\x0f\x4e\xff\xba\x8b\xba\x28\xbb\x03\xee\xc5\x84\x90\xa4\x15\xc6\x89\xe3\x23\x05\x5b\x3d\x4a\x7d\x12\x87\xe2\xe3\xcb\x3d\x77\xc2\x9a\x08\x0f\xf1\xc7\x0f\x18\xbf\x45\x5c\xa0\x1b\xef\x9f\xff\xb8\x12\xb8\xd7\xac\xef\x11\x53\xe6\x9f\x11\xb0\xf7\xe9\xf1\xd6\x03\x9f\x3a\x2b\xf1\xfb\x51\xc1\xc5\x04\x6a\x95\x02\xba\xa5\x6e\xcf\xa1\x2a\xec\x80\xde\x35\xbe\x02\xc8\xf8\xc3\xc6\x2b\xf9\x3f\xa3\x77\x31\x61\xdd\xb5\xc3\x4e\x8f\x91\x40\x4f\x8e\xb7\x1a\xb9\xb6\x95\xac\xe7\x4c\x55\x3b\xb0\xab\xc2\xe1\xe9\x1e\x31\xc4\x0c\x3f\x6c\x67\x5b\x0e\x4f\xf7\x98\x13\x92\x69\x2d\x2d\xa3\xae\x5b\x36\x4b\x90\xb1\x77\x5b\xc4\x27\x93\x0f\x07\x4c\x9a\x10\xfd\xff\x48\x63\xbd\x5f\x04\xbc\x1b\x8d\xd9\xf5\xd9\xe8\x1d\xbb\x39\x1b\x5d\x8f\xdf\x6f\xa5\xa6\xf6\x25\xfc\xed\x7c\xed\xa9\xfc\x0e\x0c\xf0\xf5\x93\x87\xd7\x4f\x1e\x76\xe1\x93\x87\x17\x0f\x12\x2f\x47\x37\xbf\x9c\xbd\xeb\x8f\x15\x0f\xf6\xf7\x8f\xf6\x8e\xf7\x0e\x4e\x9e\x77\x8e\x82\x87\xef\x6b\x75\x51\xc0\xf0\xdd\xf7\x6d\xb7\x49\x7f\xe0\x3c\xee\xff\x88\x79\x14\x56\x6b\x30\x31\x77\x93\xa7\x61\xd0\xde\xf8\x41\x04\x3f\xb0\xbe\x02\x2f\x4d\x33\xa8\xad\x1f\xb4\xca\x19\xc8\x30\xa8\x64\x5d\x83\x87\xcd\xfd\xd1\xb3\x1a\x5e\x72\xe2\x76\xcb\xd3\xbf\xfc\x7a\xd9\xc2\x6d\xfd\xed\xea\xea\xd3\x56\xa3\xf1\x70\x6b\xed\xce\x78\xe1\xf1\xfb\xeb\x6d\x07\x24\x26\x7e\x87\x06\x74\x79\xf5\xee\x5f\x17\x67\x5b\x0d\xa8\xbb\x92\x60\xda\x56\x09\xdf\x40\xe7\x18\xa2\x40\xd5\xeb\xdc\x8e\xa8\xe3\xe3\x68\x3c\xde\x6e\x7a\xb9\x10\x3b\x33\xb9\x1f\x3f\x5d\x8f\xc6\xdb\x4d\xee\x54\xe0\x37\x11\xdd\x53\xb9\x06\x88\xaf\xa7\x96\x09\xed\xc2\xf3\xd7\x5d\xcc\x87\x6f\xa1\xdf\xeb\xd1\xaf\xe7\x57\x5b\xa9\x57\x5a\x47\xbc\x70\x77\xe0\xf5\xae\x8c\xf1\xd3\xf9\xe5\x76\x2b\x28\xf4\x7c\xba\x10\xca\x1f\x64\x2e\x2f\xf5\x76\x45\x0d\x9f\x7e\x67\xe3\xab\x0f\x7f\x3f\xff\xc7\x56\xca\x98\x4d\xb8\x69\x92\xdb\x95\x51\x7d\x38\xdf\xd2\x3c\x34\x10\x99\x06\xed\xac\x92\x02\xbf\xc6\xa6\xbe\x4d\x0e\x1b\x35\x76\x43\x0f\x17\x57\xdb\xcd\xea\xe2\xb8\x46\x0c\xe6\xa7\xf6\x8f\xaf\x3f\xfd\x37\x00\x00\xff\xff\xb5\x98\x78\xcb\xe0\x45\x00\x00")

func defaultJsonBytes() ([]byte, error) {
	return bindataRead(
		_defaultJson,
		"default.json",
	)
}

func defaultJson() (*asset, error) {
	bytes, err := defaultJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "default.json", size: 17888, mode: os.FileMode(0664), modTime: time.Unix(1654125544, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xcf, 0x25, 0xb7, 0xde, 0x80, 0x4b, 0xd1, 0x3a, 0x9e, 0xd, 0x4a, 0x13, 0x75, 0xae, 0xef, 0xda, 0x23, 0xc1, 0x5e, 0x3f, 0xab, 0x31, 0x6f, 0x61, 0xbd, 0xf, 0x80, 0x19, 0xde, 0x38, 0xf7, 0xea}}
	return a, nil
}

var _fuseContainerJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x5b\x5b\x6f\x5b\xb9\x11\x7e\xcf\xaf\x10\xf4\x1c\x14\xbe\xd5\x6b\xe7\x4d\x75\xdc\xc6\x58\x3b\x4e\xed\x14\xbb\x8b\x22\x20\x68\x9e\x39\x47\xac\x78\x33\x2f\xb2\x85\x20\xff\xbd\xe0\x39\x92\x2c\x69\x86\x8a\xb3\xc7\x69\x54\x20\x0f\xbb\x49\xbe\x8f\x97\xe1\x0c\x39\x33\x1c\x1e\x7d\x7e\x35\x18\x0c\x86\xdc\x8b\xf1\x15\x77\xc3\x37\x83\x7f\xe7\x7f\x0f\x06\x83\xcf\xf3\x3f\x17\xac\x8c\x20\x62\xf2\x30\x7c\x33\x18\xde\x9e\x5d\x7d\x60\xa3\x9b\xb3\x77\xec\xf7\x93\x63\x76\x7c\x34\x7c\xbd\xd2\x38\xa4\xbb\xd1\x4a\xfb\xb0\x32\xe6\xbc\xc5\x5a\xf7\xd5\xbe\x1b\xe4\xe1\xc1\xf0\x89\xfb\x34\xff\xeb\x97\xd7\xdf\x24\xe0\x28\xff\xbf\x87\x84\xa3\x9b\xab\xfe\x42\x5c\x5d\x7c\xb8\xed\x21\x43\xee\x5e\x56\x53\x37\xf8\xfb\x97\x50\xd6\xd3\x50\xdf\x53\xd4\x97\x91\xf3\xfc\xb2\xa7\x46\xcf\x2f\xb7\x0b\x9a\x27\x78\x31\xad\x2e\x06\xfb\xde\x02\xf7\x97\xf6\xf6\xf0\x74\xef\xf7\x3f\x2d\x67\xee\x4d\xc9\x90\xff\xfc\xd4\x0e\x3a\xac\xa0\xe6\x49\xc5\x91\x88\xd2\x9a\xa7\xb9\xcf\x3e\xb2\xf3\x9b\x9b\xf7\xd7\xc3\xb5\x56\xe7\xde\x1b\x7b\x03\x71\xf8\x66\xb0\xdf\x11\x61\x16\x04\x57\x2a\x94\x5c\x15\x1e\x76\x74\x79\x79\xfd\xdb\xda\x8a\x0c\xd7\xd4\x32\x84\xb2\x06\x36\x75\x5c\x07\x61\x4d\x2d\x1b\x8c\x6b\x9b\x4c\xc4\xb0\x75\x60\x30\xea\xa4\x98\x6c\xa2\x13\x98\x89\xa8\x36\x51\x72\x58\x6d\xa7\xc0\x48\xc6\xc9\xa9\x8d\xcc\x5b\x8b\x98\x44\xb6\xef\xd0\x03\x04\x9b\x30\xe6\x1e\x9e\xbd\x7f\x9e\xa3\x66\x69\x84\x4a\x55\xab\xe9\xcf\x1b\x9a\xe6\x0e\xeb\x3f\x13\x67\xa3\x0f\xec\xf6\x8f\x5b\x36\x7a\x7b\x75\xf1\x7e\xb8\x46\x7f\x7a\xfa\xd7\x97\x67\x18\xf3\xce\xd5\x9b\x6b\xa4\xed\xcb\x8d\x8d\xb2\x9e\x31\x69\x24\x61\xcd\xef\x66\x7c\x65\xed\x24\x39\x56\x09\x6b\x27\x12\x49\xf5\x8d\x9b\x20\xeb\x80\x45\xcb\xc6\xdc\x54\x0a\x18\x47\x0d\xb2\x5c\x2c\x7a\x40\x13\x39\xf0\x35\x83\x29\x98\xc8\x28\xd9\xef\x93\x8d\x9c\xd8\xa4\x01\x62\x65\x35\x97\x26\xcf\x4c\x90\x63\x1b\x62\x81\x32\x01\x81\xb3\xa0\x2c\xd2\xf1\x0f\xdf\xbe\x85\x8d\xc5\x85\x00\x87\xe4\xea\xd0\x23\x0a\x0e\x68\xc1\xbc\xfa\x4f\x94\x1a\x1e\x11\xae\xb8\xd7\x9b\xe0\x9d\x34\x15\xc2\x3c\xda\x4f\x82\xbb\x06\x90\x58\xf9\xa0\x11\xe8\xb8\x92\x1e\x83\xda\xa2\x89\xc4\xd8\x3e\xa0\x4d\xd1\x82\x87\xc8\x06\x42\x59\x31\x61\xf3\xb5\x6d\x25\x8f\x91\x9e\x3a\xba\x81\x98\xe3\xcb\x16\x8e\x7d\xa5\x7b\x79\xea\x39\x59\xea\x6b\xb8\xb1\x41\x01\xb8\xaf\xd0\x65\x01\x02\x35\x71\x00\xe6\xb9\x69\x30\x65\x8d\x01\x81\x0d\x63\xdd\x8c\xd5\x52\x95\x7a\x79\xc0\x47\xbb\x4a\x48\xe6\x2a\x39\x64\x9d\x2a\xb9\xc3\x4d\x0c\x9c\x55\x8a\xb5\xa3\xa2\xb9\x56\xb9\xfd\x02\x89\xfd\xc2\x92\x60\x56\xa1\xcd\xd4\x91\xee\x81\x63\x47\xbb\x42\x21\xc9\x3b\xae\xdc\x2b\x33\xe4\x74\xd9\xab\xd5\x25\x18\x4f\xf3\x08\x62\x8a\xd5\xd0\xa2\x58\xeb\xf0\x48\x88\xf3\x28\x23\x6b\xbc\xc5\x06\xa9\x3b\x3f\x80\x87\x59\x12\x48\x9c\x9a\x57\x53\x19\x88\xad\xb6\x24\x18\xc5\x29\x65\x05\x61\xcd\x65\x94\xd3\x1c\xbb\x8e\x9a\xf4\x07\x35\xe9\x10\x3a\x94\x58\x08\xe9\x29\xea\x82\xab\xe8\x70\x6a\x14\x83\xf7\x54\x0b\x12\x6b\xad\x78\xe4\x61\x66\x04\x22\x1a\x88\x8f\x3c\x46\xbc\x22\x25\x43\x89\xb1\x38\x42\xd7\x96\x50\x95\x87\x1c\x85\xe9\x41\x42\x69\xde\x10\x89\xa5\x66\x90\x58\x55\x86\x8b\x44\x8d\x9c\xe3\x1c\xa6\xda\x53\xaa\x89\x3e\x19\x72\x7f\x2c\x08\x62\xa0\x14\x71\xa4\x6a\xc1\x82\x2f\xac\x53\xc6\x89\xcd\xde\x40\x14\x2e\x51\xe8\x03\xda\x67\x0d\xc4\x0a\x4c\x44\xeb\x5d\xe0\x78\xda\x06\x22\x34\x92\x1a\x28\xc3\x78\x0f\x66\x22\xd1\xed\x53\xa1\x3d\x3d\x7c\x69\xf4\xd6\x13\x50\x0b\xe8\x08\xb2\x8f\xcc\xaa\x43\x1b\xa8\x81\xe8\x00\x3c\x95\x50\x65\x8a\x16\xcb\x35\x1e\xf9\xa1\x0c\xd3\x8d\x0b\xb0\x97\xd6\xcb\x38\x23\x28\xcf\x4d\x65\x51\xb2\xd2\x45\x69\x5a\xa0\x8e\x20\x97\xed\x21\xd0\xa6\xe8\x08\xba\x8f\x92\x1a\xfb\xe0\x06\xf2\x65\xe8\x2e\x85\xc8\xf2\x69\xa7\xfa\xa5\xc0\x71\x68\x6d\x20\x06\x52\x82\x60\xc5\xa4\xa0\xf8\x4c\x59\x9c\x0a\x66\x11\xe2\xd8\x03\xaf\x18\xf7\xc0\x09\x3a\x92\x33\x65\xd3\xdb\xba\xe2\x94\xb6\x69\xed\x94\x54\x43\x7a\x21\x39\x8f\x00\xbc\xaa\xd8\x03\x8f\x62\x5c\x6a\x40\x5d\x84\x56\x39\x94\x0d\x2c\x48\xaf\x0b\x03\x5b\x26\xb8\x11\x80\x3c\xbb\xb4\x44\x0a\x21\x2d\xab\x20\x44\x6f\x91\x1e\xa4\xcd\x59\x5c\x1b\xbf\xd1\xb1\x92\x96\xb9\xe7\x91\x05\xaf\x25\x6d\xde\xec\x8c\xc8\xa1\xe7\x0c\x91\x47\x77\x28\x0e\xf7\x19\x4f\x77\xc4\xee\x94\x96\x25\x2f\x4d\xc3\xc0\x44\x7c\xcc\x97\xac\x87\x46\x86\x6d\x0d\xe8\x59\x1d\xf2\xf7\x13\xa9\x90\x7e\x15\x19\xa9\x55\x21\x52\xab\xd2\x76\x52\xd2\xe0\x3b\xad\x34\x13\xec\xf6\xf3\x39\xc4\x37\xcb\x62\x2c\x56\x45\x86\x29\x15\x00\xf0\xac\x5b\x62\x32\xdd\xa1\x14\xa8\x15\x15\xa8\x15\x1d\xa8\x75\x97\x86\x21\x18\xf4\x1d\xf7\x5e\x62\xe3\x69\xd0\x75\x55\x48\xb6\xb5\x34\xc2\x7a\x0c\x4f\x88\xbc\xac\x05\xb1\x98\x7a\x62\x70\xb6\xd6\x82\x44\x5b\x2a\xdd\x69\x41\x64\xff\x16\xe5\x78\x17\x69\xcd\xd1\x06\xcc\x18\x1e\xc0\x79\x1b\x89\xab\x8e\xbe\xcf\x67\x2d\x40\xa4\x0c\xa1\xef\x59\xe7\x53\x08\x82\x2a\x53\xe8\xfb\xf6\x54\x57\x1e\x04\x48\x9c\xc4\x6f\xd0\x05\x0f\xb0\x68\x15\x00\x5f\xb8\x57\xb9\x72\xf7\x64\xa8\x53\xa1\x3d\x50\xca\x0a\x0d\x55\xfc\x0b\x0d\xe1\x81\x74\x68\xbc\x98\x12\x68\x20\x24\xa5\xf2\x3e\x9d\x0c\x69\xf2\x0e\xa6\xcc\x9b\x0c\x21\x73\xf1\x92\x6c\xe0\x61\x9e\xb8\xa2\x33\x6b\xe0\x21\x80\x22\x36\x00\x65\xc7\x8c\xd1\xc5\x2b\xe2\x86\xe4\x78\xc2\xe7\xcf\xc9\xaa\xae\xc8\x3d\xd2\x31\xad\x01\x83\x6c\x0c\x47\x6b\x76\xd2\x11\xc3\x39\xc0\x13\x5b\xac\x30\x57\x04\x0b\xdb\xc5\x79\xc2\xfc\x2e\x67\x0d\x54\x5b\xe0\x15\xb2\x7f\x87\x62\xe9\xba\xc4\x88\x18\xa5\xb3\x03\x7a\x71\x5b\xe0\x25\x41\x1f\xbc\xa4\xee\x05\x1d\x8e\xa5\xea\x60\x24\x56\x16\x96\xc2\xf8\xb8\x40\x50\x47\x69\x81\xe3\x3d\x42\x6a\xc8\x03\x3e\x35\x19\xab\x3d\xce\x5c\x33\xae\x75\x40\x65\xc8\x05\x5e\xd0\x4e\x4b\x53\xbd\x34\x77\x5d\x25\xc7\xf1\x06\xd7\xb4\xb6\x84\x2c\x0f\x54\xba\xd9\xa1\xd4\xba\x3b\x9c\xd0\x77\x88\xdc\x47\x36\x7f\x2f\x41\xb4\x26\x42\x8b\x0f\x70\x8f\xb0\x98\x8f\xcb\xbc\x60\x4a\x72\x0e\x4c\x25\x0d\x56\x41\x47\x7a\x2b\x34\x0f\xd8\x92\x2d\x7b\x9f\x20\x81\x34\xb5\xa5\x69\x0f\x31\xf9\xc2\xac\x21\x05\x47\x38\xea\x8e\x6c\x7d\x35\x55\x31\xda\xa0\x4b\x46\x8d\x2c\x36\xdb\xc4\x0b\x62\x0c\x55\x8e\x5f\xbc\xae\x73\x4a\x8c\xe2\xd4\x53\x03\xc2\xc2\x4b\xd2\x71\xcf\xd1\x46\x5c\xb2\x6c\x71\xfb\x62\x9a\xa3\x0b\x38\xd5\x4a\x22\x55\x2d\x5b\xb5\x7f\x49\x0a\x67\x25\x5d\x0b\xef\xdb\xa1\x64\xce\x4a\xa7\xd8\x33\x92\x8d\x0a\xba\xeb\xda\x86\xaf\xa9\xa6\x10\xf8\x97\xe4\x16\xd5\x84\xaf\x2e\x67\x26\x01\x97\x03\x03\x08\x61\x35\x0a\x5d\x74\x70\x0a\xa0\xc9\x57\x0f\x4d\x44\xe7\x00\xda\x12\xc3\xea\x76\x97\x6d\x63\x4a\x1a\x24\xb6\x75\xc6\xb2\x37\x29\xe1\xf4\x28\x94\x3f\x6b\x71\x1a\x8e\x78\x9f\x43\xac\xa9\xdb\xfc\x02\xc7\xf7\x85\x96\x21\x6e\xab\x0b\x9c\xec\x41\xcf\x50\x1a\x9f\xae\xab\x84\x72\x5d\x25\x94\xea\x2a\x81\x2e\x9e\x84\x72\xe1\x23\x40\xf4\x54\x95\x69\x81\x93\x93\xd3\x15\x91\x50\xae\x88\x84\x52\x45\x24\x94\x2b\x22\x2d\x55\xea\x52\xec\x41\xd6\x50\xc2\xf6\x1a\x4a\x20\x4b\x25\xa1\x58\x0f\x09\xdb\xeb\x21\x2d\x2d\x2b\xc6\xab\xca\x13\x6f\x65\x81\x2c\x7e\x84\x42\xf1\xa3\x74\xb3\x0b\x63\x8d\x63\x67\x18\x93\x87\x7c\xac\x2b\xaa\x29\x75\xf2\xc7\x29\x56\xc4\x65\x3a\x07\x4c\x15\x43\xe4\x38\xe5\xee\x52\x4f\xfc\x2a\xb1\xc0\xf1\x39\x2e\xc7\xd0\x62\x84\xcc\x86\x20\xa4\x6d\x51\x2a\x19\xe8\x18\xc7\x71\x46\x10\x9c\x92\x02\x7b\x1d\x22\xd1\xa7\x6f\xc9\x74\xd1\xba\x54\xb3\xce\x38\x8e\x74\x33\x4d\xa5\x82\x73\x98\x10\x84\xb8\xff\x64\x6c\xcb\xbb\x5a\xa6\x09\x29\x67\x81\x8a\xfd\x11\x3f\xa2\xc7\x86\x2a\xb4\x50\x6f\x91\xad\x1b\x2a\x14\x01\x3a\xae\x02\x05\x25\xae\x81\x68\xa7\xe0\x7d\x42\x16\x5f\xf2\xe5\x49\x8b\x0f\xa0\x1d\x1d\xb6\xf5\x0d\x5b\xfb\x16\xcb\x1a\x0b\x76\x9b\x5c\x4f\x74\x79\xf4\x6d\xb2\x3d\xd1\x74\x7f\x64\xd6\x48\xda\xaa\xf0\x2e\x52\x7e\x16\x49\xc5\x12\x74\xa2\xce\x6a\xa2\x92\x7a\xba\x5e\xd0\xa1\x78\x63\x27\x4a\x07\x2d\x68\x88\xc7\x96\x25\x51\xc8\x32\x12\xa9\x9c\x29\xf5\xe8\x35\xd5\xb4\x1b\xc8\x19\x34\x1a\x37\x83\xd8\x5f\x67\x94\x78\x62\x68\xef\x8a\x24\x38\xfd\x1f\x7d\x89\xa4\xa5\xf9\x15\xbc\x01\x95\x47\x38\xfa\xcb\xc9\xca\xb4\xcf\xf9\xca\x28\xbb\x66\x08\x81\x4d\x35\x2b\xdc\xd1\x97\x7c\xe1\xba\x1c\x3d\x17\x2f\xfb\xdd\x0a\xf7\x0d\x96\x74\x7d\xd9\x9d\x4e\x2a\x78\x1c\xbe\x19\xec\xbd\x46\x94\x75\xcb\x29\xf2\x7f\xe7\xff\x1c\xe2\x36\x53\xae\x12\xe4\xee\x6b\xcc\x97\x95\x85\x3c\x47\x7d\xe0\x83\x35\x5c\xe5\x4c\xeb\xff\x57\x05\x27\x3f\x55\xb0\x7f\xb8\xbf\xf7\xcb\xc1\x4f\x3d\x64\x3d\x9c\xfc\x3c\x12\x83\xa3\x83\xd3\xa3\xd3\xe3\x5f\x0e\x4e\xff\xba\x8b\xba\x28\x87\x03\xee\xc5\x98\x90\xa4\x15\xc6\x89\xe3\x23\x05\xbd\x3e\x4a\xdd\xc8\x43\xf1\xf5\xe5\x81\x3b\x61\x4d\x84\xc7\xf8\xe3\x17\x8c\xbf\x45\x9c\xa3\x6b\x9f\x99\x7f\xbb\x12\xb8\xd7\x6c\xdb\x47\x4c\x99\x7f\x46\xc2\xbe\x4d\x8f\x77\x1e\xf8\xc4\x59\x89\xbf\x1f\x15\x5c\x8c\xa1\x56\x29\xa0\x57\xea\xf6\x1e\xaa\xc2\x0e\xe8\x5d\xe3\x27\x80\x8c\x3f\xae\xfd\x18\xe1\xcf\xe8\x5d\x8c\x59\xf7\xec\xb0\xd3\x6b\x24\xd0\x93\xe3\x5e\x2b\xd7\xb6\x92\xf5\x8c\xa9\x6a\x07\x4e\x55\x38\x3c\xdd\x23\x96\x98\xe1\xc7\x7e\xbe\xe5\xf0\x74\x8f\x39\x21\x99\xd6\xd2\x32\xea\xb9\x65\xbd\x05\x99\x7b\xb7\x4d\x7c\x32\xf9\x72\xc0\xa4\x09\xd1\xef\xc2\x2f\x02\xde\x8e\xce\xd8\xcd\xf9\xe8\x2d\xbb\x3d\x1f\xdd\x9c\xbd\xeb\xa5\xa6\xf6\x4b\xf8\xbb\xd9\xca\xa7\xf2\x3b\xb0\xc0\x9f\x3f\x79\xf8\xf9\x93\x87\x5d\xf8\xc9\xc3\x8b\x27\x89\x57\xa3\xdb\x5f\xcf\xdf\x6e\xcf\x15\x0f\xf6\xf7\x8f\xf6\x8e\xf7\x0e\x4e\x9e\x77\x8f\x82\xc7\xef\xeb\x75\x51\xc2\xf0\xdd\xcf\x6d\x77\x48\x7f\xa0\x1d\xf7\x7f\x84\x1d\x85\xd5\x1a\x4c\xcc\xd3\x64\x33\x0c\xda\x17\x3f\x88\xe0\x07\xd6\x57\xe0\xa5\x69\x06\xb5\xf5\x83\x56\x39\x03\x19\x06\x95\xac\x6b\xf0\xb0\x7e\x3e\xb6\xec\x86\x97\x34\xdc\x6e\x45\xfa\x97\xdf\x2f\x3d\xc2\xd6\xdf\xae\xaf\x3f\xf6\x5a\x8d\x87\x3b\x6b\x77\x26\x0a\x9f\xbd\xbb\xe9\xbb\x20\x31\xf6\x3b\xb4\xa0\xab\xeb\xb7\xff\xba\x3c\xef\xb5\xa0\xee\x49\x82\x69\x5b\x25\xfc\x02\x9d\x73\x88\x02\x55\xaf\x72\x3b\xa2\x8e\x0f\xa3\xb3\xb3\x7e\xe6\xe5\x42\xec\x8c\x71\x3f\x7c\xbc\x19\x9d\xf5\x33\xee\x44\xe0\x6f\x22\xba\x4f\xe5\x1a\x20\x7e\x3d\xb5\x28\x68\x17\x3e\x7f\xdd\xc5\x7a\x78\x0f\xfd\xde\x8c\x7e\xbb\xb8\xee\xa5\x5e\x69\x1d\xf1\x85\xbb\x03\xaf\x77\x65\x8d\x1f\x2f\xae\xfa\xed\xa0\xb0\xe5\xa7\x0b\xa1\xfc\x83\xcc\xc5\xa3\xde\xae\xa8\xe1\xe3\x1f\xec\xec\xfa\xfd\xdf\x2f\xfe\xd1\x4b\x19\xd3\x31\x37\x4d\x72\xbb\xb2\xaa\xf7\x17\x3d\xdd\x43\x03\x91\x69\xd0\xce\x2a\x29\xf0\xd7\xd8\xd4\x6f\x93\xc3\x5a\x8f\xdd\xd0\xc3\xe5\x75\x3f\xab\xce\xaf\x6b\xc4\x62\x5e\xb5\xff\xf8\xf2\xea\xbf\x01\x00\x00\xff\xff\xc1\x02\xbc\xe5\x47\x47\x00\x00")

func fuseContainerJsonBytes() ([]byte, error) {
	return bindataRead(
		_fuseContainerJson,
		"fuse-container.json",
	)
}

func fuseContainerJson() (*asset, error) {
	bytes, err := fuseContainerJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "fuse-container.json", size: 18247, mode: os.FileMode(0664), modTime: time.Unix(1654125544, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x52, 0x92, 0x1e, 0xfa, 0xd, 0x1f, 0xd1, 0x83, 0x91, 0xc7, 0x68, 0x53, 0xd0, 0xe0, 0xa5, 0x96, 0x39, 0x56, 0x2d, 0x9, 0x24, 0x4d, 0xd8, 0x6e, 0xed, 0xa6, 0x2b, 0xcb, 0xbd, 0x13, 0x8d, 0x1e}}
	return a, nil
}

var _allowPerfSyscallsJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x5b\x5b\x6f\x5b\xb9\x11\x7e\xcf\xaf\x10\xf4\x1c\x14\xbe\xd5\x6b\xe7\x4d\x75\xdc\xc6\x58\x3b\x4e\xed\x14\xbb\x8b\x22\x20\x68\x9e\x39\x47\xac\x78\x33\x2f\xb2\x85\x20\xff\xbd\xe0\x39\x92\x2c\x69\x86\x8a\x37\xc7\x69\x54\x20\x0f\xbb\x49\xbe\x8f\x97\xe1\x0c\x39\x33\x1c\x1e\x7d\x7e\x35\x18\x0c\x86\xdc\x8b\xf1\x15\x77\xc3\x37\x83\x7f\xe7\x7f\x0f\x06\x83\xcf\xf3\x3f\x17\xac\x8c\x20\x62\xf2\x30\x7c\x33\x18\xde\x9e\x5d\x7d\x60\xa3\x9b\xb3\x77\xec\xf7\x93\x63\x76\x7c\x34\x7c\xbd\xd2\x38\xa4\xbb\xd1\x4a\xfb\xb0\x32\xe6\xbc\xc5\x5a\xf7\xd5\xbe\x1b\xe4\xe1\xc1\xf0\x89\xfb\x34\xff\xeb\x97\xd7\x7f\x4a\xc0\x51\xfe\x7f\x0f\x09\x47\x37\x57\xfd\x85\xb8\xba\xf8\x70\xdb\x43\x86\xdc\xbd\xac\xa6\x6e\xf0\xf7\x2f\xa1\xac\xa7\xa1\xbe\xa7\xa8\x2f\x23\xe7\xf9\x65\x4f\x8d\x9e\x5f\x6e\x17\x34\x4f\xf0\x62\x5a\x5d\x0c\xf6\xbd\x05\xee\x2f\xed\xed\xe1\xe9\xde\xef\xdf\x2c\x67\xee\x4d\xc9\x90\xff\xfc\xd4\x0e\x3a\xac\xa0\xe6\x49\xc5\x91\x88\xd2\x9a\xa7\xb9\xcf\x3e\xb2\xf3\x9b\x9b\xf7\xd7\xc3\xb5\x56\xe7\xde\x1b\x7b\x03\x71\xf8\x66\xb0\xdf\x11\x61\x16\x04\x57\x2a\x94\x5c\x15\x1e\x76\x74\x79\x79\xfd\xdb\xda\x8a\x0c\xd7\xd4\x32\xee\x5c\xbd\xa9\x61\x07\xbe\x66\x30\x05\x13\x99\x75\x60\x9e\xad\xde\xe7\x48\x21\x8d\x50\xa9\x6a\x05\xf9\xbc\x3e\xab\xe0\x0e\x8b\x97\x89\xb3\xd1\x07\x76\xfb\xc7\x2d\x1b\xbd\xbd\xba\x78\x3f\x5c\xa3\x3f\x3d\xfd\xeb\xcb\xb7\xad\x55\x28\x6b\x60\x13\xac\xb9\xb1\x51\xd6\x33\x26\x8d\x8c\x88\x0c\xc2\x9a\x5a\x36\x18\xd7\x36\x19\xa2\x79\xab\x43\x84\x3a\x29\x26\x9b\xa8\xb2\x76\x92\x1c\xab\x84\xb5\x13\x89\xa4\x22\x87\xd7\x76\x0a\x8c\x64\xb2\x0e\x58\xb4\x6c\xcc\x4d\xa5\x80\x71\xd4\x20\xcb\xc5\xa2\x07\x34\xd1\xa6\xfd\x37\xe8\xfb\x64\x23\x17\x51\x6d\xe2\x01\x62\x65\x35\x97\x26\xcf\x4c\x90\x63\x1b\x62\x81\x32\x01\x81\xb3\xa0\x2c\xd2\x71\x22\x57\xda\xa1\x07\x08\x36\x61\xcc\x3d\xbc\xe8\xf6\x2d\x6c\x2c\x2e\x04\x38\x24\x57\x87\x1e\x51\x70\x40\x0b\xe6\xd5\x7f\xa2\xd4\xf0\x88\x70\xc5\xbd\xde\x04\xef\xa4\xa9\x10\xe6\xd1\x7e\x12\xdc\x35\x80\xc4\xca\x07\x8d\x40\xc7\x95\xf4\x18\xd4\x16\x4d\x24\xc6\xf6\x01\x6d\x8a\x16\x3c\x44\x36\x10\xca\x8a\x09\x9b\xaf\x6d\x2b\x79\x8c\xf4\xd4\xd1\x0d\xc4\xec\x7e\xb7\x70\xec\x2b\xdd\xcb\x53\xcf\xc9\x52\x5f\xc3\x8d\x0d\x0a\xc0\x7d\x85\x2e\x0b\x10\xa8\x89\x03\x30\xcf\x4d\x83\x29\x6b\x0c\x08\x6c\x18\xeb\x66\xac\x96\xaa\xd4\xcb\x03\x3e\xda\x55\x42\x32\x57\xc9\x21\xeb\x54\xc9\x1d\x6e\x62\xe0\xac\x52\xac\x1d\x15\xcd\xb5\xca\xed\x17\x48\xec\x17\x96\x04\xb3\x0a\x6d\xa6\x8e\x74\x0f\x1c\x3b\xda\x15\x0a\x49\xde\x71\xe5\x5e\x99\x21\xa7\xcb\x5e\xad\x2e\xc1\x78\x9a\x47\x10\x53\xac\x86\x16\xc5\x5a\x87\x47\x42\x9c\x47\x19\x59\xe3\x2d\x36\x48\xdd\xf9\x01\x3c\xcc\x92\x40\xe2\xd4\xbc\x9a\xca\x40\x6c\xb5\x25\xc1\x28\x4e\x29\x2b\x08\x6b\x2e\xa3\x9c\xe6\xd8\x75\xd4\xa4\x3f\xa8\x49\x87\xd0\xa1\xc4\x42\x48\x4f\x51\x17\x5c\x45\x87\x53\xa3\x18\xbc\xa7\x5a\x90\x58\x6b\xc5\x23\x0f\x33\x23\x10\xd1\x40\x7c\xe4\x31\xe2\x15\x29\x19\x4a\x8c\xc5\x11\xba\xb6\x84\xaa\x3c\xe4\x28\x4c\x0f\x12\x4a\xf3\x86\x48\x2c\x35\x83\xc4\xaa\x32\x5c\x24\x6a\xe4\x1c\xe7\x30\xd5\x9e\x52\x4d\xf4\xc9\x90\xfb\x63\x41\x10\x03\xa5\x88\x23\x55\x0b\x16\x7c\x61\x9d\x32\x4e\x6c\xf6\x06\xa2\x70\x89\x42\x1f\xd0\x3e\x6b\x20\x56\x60\x22\x5a\xef\x02\xc7\xd3\x36\x10\xa1\x91\xd4\x40\x19\xc6\x7b\x30\x13\x89\x6e\x9f\x0a\xed\xe9\xe1\x4b\xa3\xb7\x9e\x80\x5a\x40\x47\x90\x7d\x64\x56\x1d\xda\x40\x0d\x44\x07\xe0\xa9\x84\x2a\x53\xb4\x58\xae\xf1\xc8\x0f\x65\x98\x6e\x5c\x80\xbd\xb4\x5e\xc6\x19\x41\x79\x6e\x2a\x8b\x92\x95\x2e\x4a\xd3\x02\x75\x04\xb9\x6c\x0f\x81\x36\x45\x47\xd0\x7d\x94\xd4\xd8\x07\x37\x10\x99\xb7\x77\x29\x44\x96\x4f\x3b\xd5\x2f\x05\x8e\x43\x6b\x03\x31\x90\x12\x04\x2b\x26\x05\xc5\x67\xca\xe2\x54\x30\x8b\x10\xc7\x1e\x78\xc5\xb8\x07\x4e\xd0\x91\x9c\x29\x9b\xde\xd6\x15\xa7\xb4\x4d\x6b\xa7\xa4\x1a\xd2\x0b\xc9\x79\x04\xe0\x55\xc5\x1e\x78\x14\xe3\x52\x03\xea\x22\xb4\xca\xa1\x6c\x60\x41\x7a\x5d\x18\xd8\x32\xc1\x8d\x00\xe4\xd9\xa5\x25\x52\x08\x69\x59\x05\x21\x7a\x8b\xf4\x20\x6d\xce\xe2\xda\xf8\x8d\x8e\x95\xb4\xcc\x3d\x8f\x2c\x78\x2d\x69\xf3\x66\x67\x44\x0e\x3d\x67\x88\x3c\xba\x43\x71\xb8\xcf\x78\xba\x23\x76\xa7\xb4\x2c\x79\x69\x1a\x06\x26\xe2\x63\xbe\x64\x3d\x34\x32\x6c\x6b\x40\xcf\xea\x90\xbf\x9f\x48\x85\xf4\xab\xc8\x48\xad\x0a\x91\x5a\x95\xb6\x93\x92\x06\xdf\x69\xa5\x99\x60\xb7\x9f\xcf\x21\xbe\x59\x16\x63\xb1\x2a\x32\x4c\xa9\x00\x80\x67\xdd\x12\x93\xe9\x0e\xa5\x40\xad\xa8\x40\xad\xe8\x40\xad\xbb\x34\x0c\xc1\xa0\xef\xb8\xf7\x12\x1b\x4f\x83\xae\xab\x42\xb2\xad\xa5\x11\xd6\x63\x78\x42\xe4\x65\x2d\x88\xc5\xd4\x13\x83\xb3\xb5\x16\x24\xda\x52\xe9\x4e\x0b\x22\xfb\xb7\x28\xc7\xbb\x48\x6b\x8e\x36\x60\xc6\xf0\x00\xce\xdb\x48\x5c\x75\xf4\x7d\x3e\x6b\x01\x22\x65\x08\x7d\xcf\x3a\x9f\x42\x10\x54\x99\x42\xdf\xb7\xa7\xba\xf2\x20\x40\xe2\x24\x7e\x83\x2e\x78\x80\x45\xab\x00\xf8\xc2\xbd\xca\x95\xbb\x27\x43\x9d\x0a\xed\x81\x52\x56\x68\x08\xe7\xa7\x43\x43\x78\x20\x1d\x1a\x2f\xa6\x04\x1a\x08\x49\xa9\xbc\x4f\x27\x43\x9a\xbc\x83\x29\xf3\x26\x43\xc8\x5c\xbc\x24\x1b\x78\x98\x27\xae\xe8\xcc\x1a\x78\x08\xa0\x88\x0d\x40\xd9\x31\x63\x74\xf1\x8a\xb8\x21\x39\x9e\xf0\xf9\x73\xb2\xaa\x2b\x72\x8f\x74\x4c\x6b\xc0\x20\x1b\xc3\xd1\x9a\x9d\x74\xc4\x70\x0e\xf0\xc4\x16\x2b\xcc\x15\xc1\xc2\x76\x71\x9e\x30\xbf\xcb\x59\x03\xd5\x16\x78\x85\xec\xdf\xa1\x58\xba\x2e\x31\x22\x46\xe9\xec\x80\x1e\xa4\x16\x78\x49\xd0\x07\x2f\xa9\x7b\x41\x87\x63\xa9\x3a\x18\x89\x95\x85\xa5\x30\x3e\x2e\x10\xd4\x51\x5a\xe0\x78\x8f\x90\x1a\xf2\x80\x4f\x4d\xc6\x6a\x8f\x33\xd7\x8c\x6b\x1d\x50\x19\x72\x81\x17\xb4\xd3\xd2\x54\x2f\xcd\x5d\x57\xc9\x71\xbc\xc1\x35\xad\x2d\x21\xcb\x03\x95\x6e\x76\x28\xb5\xee\x0e\x27\xf4\x1d\x22\xf7\x91\xcd\x9f\x13\x10\xad\x89\xd0\xe2\x03\xdc\x23\x2c\xe6\xe3\x32\x2f\x98\x92\x9c\x03\x53\x49\x83\x55\xd0\x91\xde\x0a\xcd\x03\xb6\x64\xcb\xde\x27\x48\x20\x4d\x6d\x69\xda\x43\x4c\xbe\x30\x6b\x48\xc1\x11\x8e\xba\x23\x5b\x5f\x4d\x55\x8c\x36\xe8\x92\x51\x23\x8b\xcd\x36\xf1\x82\x18\x43\x95\xe3\x17\xaf\xeb\x9c\x12\xa3\x38\xf5\xd4\x80\xb0\xf0\x92\x74\xdc\x73\xb4\x11\x97\x2c\x5b\xdc\xbe\x98\xe6\xe8\x02\x4e\xb5\x92\x48\x55\xcb\x56\xed\x5f\x92\xc2\x59\x49\xd7\xc2\xfb\x76\x28\x99\xb3\xd2\x29\xf6\x8c\x64\xa3\x82\xee\xba\xb6\xe1\x6b\xaa\x29\x04\xfe\x25\xb9\x45\x35\xe1\xab\xcb\x99\x49\xc0\xe5\xc0\x00\x42\x58\x8d\x42\x17\x1d\x9c\x02\x68\xf2\xd5\x43\x13\xd1\x39\x80\xb6\xc4\xb0\xba\xdd\x65\xdb\x98\x92\x06\x89\x6d\x9d\xb1\xec\x4d\x4a\x38\x3d\x0a\xe5\xcf\x5a\x9c\x86\x23\xde\xe7\x10\x6b\xea\x36\xbf\xc0\xf1\x7d\xa1\x65\x88\xdb\xea\x02\x27\x7b\xd0\x33\x94\xc6\xa7\xeb\x2a\xa1\x5c\x57\x09\xa5\xba\x4a\xa0\x8b\x27\xa1\x5c\xf8\x08\x10\x3d\x55\x65\x5a\xe0\xe4\xe4\x74\x45\x24\x94\x2b\x22\xa1\x54\x11\x09\xe5\x8a\x48\x4b\x95\xba\x14\x7b\x90\x35\x94\xb0\xbd\x86\x12\xc8\x52\x49\x28\xd6\x43\xc2\xf6\x7a\x48\x4b\xcb\x8a\xf1\xaa\xf2\xc4\x5b\x59\x20\x8b\x1f\xa1\x50\xfc\x28\xdd\xec\xc2\x58\xe3\xd8\x19\xc6\xe4\x21\x1f\xeb\x8a\x6a\x4a\x9d\xfc\x71\x8a\x15\x71\x99\xce\x01\x53\xc5\x10\x39\x4e\xb9\xbb\xd4\x13\xbf\x4a\x2c\x70\x7c\x8e\xcb\x31\xb4\x18\x21\xb3\x21\x08\x69\x5b\x94\x4a\x06\x3a\xc6\x71\x9c\x11\x04\xa7\xa4\xc0\x5e\x87\x48\xf4\xe9\x5b\x32\x5d\xb4\x2e\xd5\xac\x33\x8e\x23\xdd\x4c\x53\xa9\xe0\x1c\x26\x04\x21\xee\x3f\x19\xdb\xf2\xae\x96\x69\x42\xca\x59\xa0\x62\x7f\xc4\x8f\xe8\xb1\xa1\x0a\x2d\xd4\x5b\x64\xeb\x86\x0a\x45\x80\x8e\xab\x40\x41\x89\x6b\x20\xda\x29\x78\x9f\x90\xc5\x97\x7c\x79\xd2\xe2\x03\x68\x47\x87\x6d\x7d\xc3\xd6\xbe\xc5\xb2\xc6\x82\xdd\x26\xd7\x13\x5d\x1e\x7d\x9b\x6c\x4f\x34\xdd\x1f\x99\x35\x92\xb6\x2a\xbc\x8b\x94\x9f\x45\x52\xb1\x04\x9d\xa8\xb3\x9a\xa8\xa4\x9e\xae\x17\x74\x28\xde\xd8\x89\xd2\x41\x0b\x1a\xe2\xb1\x65\x49\x14\xb2\x8c\x44\x2a\x67\x4a\x3d\x7a\x4d\x35\xed\x06\x72\x06\x8d\xc6\xcd\x20\xf6\xd7\x19\x25\x9e\x18\xda\xbb\x22\x09\x4e\xff\x47\x5f\x22\x69\x69\x7e\x05\x6f\x40\xe5\x11\x8e\xfe\x72\xb2\x32\xed\x73\xbe\x32\xca\xae\x19\x42\x60\x53\xcd\x0a\x77\xf4\x25\x5f\xb8\x2e\x47\xcf\xc5\xcb\x7e\xb7\xc2\x7d\x83\x25\x5d\x5f\x76\xa7\x93\x0a\x1e\x87\x6f\x06\x7b\xaf\x11\x65\xdd\x72\x8a\xfc\xdf\xf9\x3f\x87\xb8\xcd\x94\xab\x04\xb9\xfb\x1a\xf3\x65\x65\x21\xcf\x51\x1f\xf8\x60\x0d\x57\x39\xd3\xfa\xff\x55\xc1\xc9\x4f\x15\xec\x1f\xee\xef\xfd\x72\xf0\x53\x0f\x59\x0f\x27\x3f\x8f\xc4\xe0\xe8\xe0\xf4\xe8\xf4\xf8\x97\x83\xd3\xbf\xee\xa2\x2e\xca\xe1\x80\x7b\x31\x26\x24\x69\x85\x71\xe2\xf8\x48\x41\xaf\x8f\x52\x37\xf2\x50\x7c\x7d\x79\xe0\x4e\x58\x13\xe1\x31\xfe\xf8\x05\xe3\x6f\x11\xe7\xe8\xda\x57\xd8\x7f\x5e\x09\xdc\x6b\xb6\xed\x23\xa6\xcc\x3f\x23\x61\xdf\xa6\xc7\x3b\x0f\x7c\xe2\xac\xc4\xdf\x8f\x0a\x2e\xc6\x50\xab\x14\xd0\x2b\x75\x7b\x0f\x55\x61\x07\xf4\xae\xf1\x13\x40\xc6\x1f\xd7\xbe\xd5\xff\x16\xbd\x8b\x31\xeb\x9e\x1d\x76\x7a\x8d\x04\x7a\x72\xdc\x6b\xe5\xda\x56\xb2\x9e\x31\x55\xed\xc0\xa9\x0a\x87\xa7\x7b\xc4\x12\x33\xfc\xd8\xcf\xb7\x1c\x9e\xee\x31\x27\x24\xd3\x5a\x5a\x46\x3d\xb7\xac\xb7\x20\x73\xef\xb6\x89\x4f\x26\x5f\x0e\x98\x34\x21\xfa\x5d\xf8\x45\xc0\xdb\xd1\x19\xbb\x39\x1f\xbd\x65\xb7\xe7\xa3\x9b\xb3\x77\xbd\xd4\xd4\x7e\x09\x7f\x37\x5b\xf9\x54\x7e\x07\x16\xf8\xf3\x27\x0f\x3f\x7f\xf2\xb0\x0b\x3f\x79\x78\xf1\x24\xf1\x6a\x74\xfb\xeb\xf9\xdb\xed\xb9\xe2\xc1\xfe\xfe\xd1\xde\xf1\xde\xc1\xc9\xf3\xee\x51\xf0\xf8\x7d\xbd\x2e\x4a\x18\xbe\xfb\xb9\xed\x0e\xe9\x0f\xb4\xe3\xfe\x8f\xb0\xa3\xb0\x5a\x83\x89\x79\x9a\x6c\x86\x41\xfb\xe2\x07\x11\xfc\xc0\xfa\x0a\xbc\x34\xcd\xa0\xb6\x7e\xd0\x2a\x67\x20\xc3\xa0\x92\x75\x0d\x1e\xd6\xcf\xc7\x96\xdd\xf0\x92\x86\xdb\xad\x48\xff\xf2\xfb\xa5\x47\xd8\xfa\xdb\xf5\xf5\xc7\x5e\xab\xf1\x70\x67\xed\xce\x44\xe1\xb3\x77\x37\x7d\x17\x24\xc6\x7e\x87\x16\x74\x75\xfd\xf6\x5f\x97\xe7\xbd\x16\xd4\x3d\x49\x30\x6d\xab\x84\x5f\xa0\x73\x0e\x51\xa0\xea\x55\x6e\x47\xd4\xf1\x61\x74\x76\xd6\xcf\xbc\x5c\x88\x9d\x31\xee\x87\x8f\x37\xa3\xb3\x7e\xc6\x9d\x08\xfc\x4d\x44\xf7\xa9\x5c\x03\xc4\xaf\xa7\x16\x05\xed\xc2\xe7\xaf\xbb\x58\x0f\xef\xa1\xdf\x9b\xd1\x6f\x17\xd7\xbd\xd4\x2b\xad\x23\xbe\x70\x77\xe0\xf5\xae\xac\xf1\xe3\xc5\x55\xbf\x1d\x14\xb6\xfc\x74\x21\x94\x7f\x90\xb9\x78\xd4\xdb\x15\x35\x7c\xfc\x83\x9d\x5d\xbf\xff\xfb\xc5\x3f\x7a\x29\x63\x3a\xe6\xa6\x49\x6e\x57\x56\xf5\xfe\xa2\xa7\x7b\x68\x20\x32\x0d\xda\x59\x25\x05\xfe\x1a\x9b\xfa\x6d\x72\x58\xeb\xb1\x1b\x7a\xb8\xbc\xee\x67\xd5\xf9\x75\x8d\x58\xcc\xab\xf6\x1f\x5f\x5e\xfd\x37\x00\x00\xff\xff\xce\x8c\x12\x26\x66\x46\x00\x00")

func allowPerfSyscallsJsonBytes() ([]byte, error) {
	return bindataRead(
		_allowPerfSyscallsJson,
		"allow-perf-syscalls.json",
	)
}

func allowPerfSyscallsJson() (*asset, error) {
	bytes, err := allowPerfSyscallsJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "allow-perf-syscalls.json", size: 18022, mode: os.FileMode(0664), modTime: time.Unix(1654125544, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x78, 0x83, 0x8d, 0xf5, 0xe1, 0x37, 0x7e, 0x4f, 0x53, 0x1a, 0xaf, 0xc0, 0x6b, 0xc, 0x97, 0x82, 0x3f, 0x1b, 0x94, 0x10, 0x31, 0xa3, 0x41, 0x68, 0xc4, 0xbc, 0xb, 0x70, 0x2e, 0xac, 0x82, 0x32}}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// AssetString returns the asset contents as a string (instead of a []byte).
func AssetString(name string) (string, error) {
	data, err := Asset(name)
	return string(data), err
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// MustAssetString is like AssetString but panics when Asset would return an
// error. It simplifies safe initialization of global variables.
func MustAssetString(name string) string {
	return string(MustAsset(name))
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetDigest returns the digest of the file with the given name. It returns an
// error if the asset could not be found or the digest could not be loaded.
func AssetDigest(name string) ([sha256.Size]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s can't read by error: %v", name, err)
		}
		return a.digest, nil
	}
	return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s not found", name)
}

// Digests returns a map of all known files and their checksums.
func Digests() (map[string][sha256.Size]byte, error) {
	mp := make(map[string][sha256.Size]byte, len(_bindata))
	for name := range _bindata {
		a, err := _bindata[name]()
		if err != nil {
			return nil, err
		}
		mp[name] = a.digest
	}
	return mp, nil
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"default.json":             defaultJson,
	"fuse-container.json":      fuseContainerJson,
	"allow-perf-syscalls.json": allowPerfSyscallsJson,
}

// AssetDebug is true if the assets were built with the debug flag enabled.
const AssetDebug = false

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"},
// AssetDir("data/img") would return []string{"a.png", "b.png"},
// AssetDir("foo.txt") and AssetDir("notexist") would return an error, and
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		canonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(canonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{
	"allow-perf-syscalls.json": &bintree{allowPerfSyscallsJson, map[string]*bintree{}},
	"default.json":             &bintree{defaultJson, map[string]*bintree{}},
	"fuse-container.json":      &bintree{fuseContainerJson, map[string]*bintree{}},
}}

// RestoreAsset restores an asset under the given directory.
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	return os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
}

// RestoreAssets restores an asset under the given directory recursively.
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(canonicalName, "/")...)...)
}
