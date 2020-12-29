// Code generated by go-bindata. (@generated) DO NOT EDIT.

// Package config generated by go-bindata.// sources:
// configs/config.ini
// configs/migrations/01_initial.up.sql
package config

import (
	"bytes"
	"compress/gzip"
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
		return nil, fmt.Errorf("read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// ModTime return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _configIni = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x54\xdd\x6e\xea\x38\x10\xbe\xf7\x53\xcc\x65\xbb\x52\x03\xdd\x1f\x09\x15\x71\x91\x16\x4a\xd1\x06\x82\x48\xe9\xae\x16\xad\x90\x49\x86\xc4\xaa\x63\xa7\x9e\x09\x1c\xce\xd3\x1f\x39\x81\x8a\x56\xa7\x3d\x85\x1b\x34\xfe\x7e\x86\xf9\xc6\x5e\x11\xba\x1d\xba\xff\x45\x1f\xc2\x2c\x73\x48\x04\xd6\xc0\xbe\x50\x69\x01\x5c\x20\xb4\xc7\xa0\x15\x31\x1a\x0a\x44\x34\x49\x1e\x47\xb3\x75\x38\x1c\x2e\x46\x49\x02\x03\xe8\x06\xcd\x57\x88\x3e\xcc\xad\xe3\xcf\xd9\xf3\x78\xf1\x08\x03\xe8\x75\x7b\xd7\x42\xac\x32\xc9\x72\x23\x09\xbd\xf9\xf0\xf8\x1b\x08\x99\x95\xc9\x29\x80\x7b\xeb\xc0\xd8\x3d\x58\xa3\x0f\x50\x59\xe2\xdc\x21\xbd\x68\x50\x04\x54\x57\x95\x75\x8c\x59\x20\x1e\xe2\xc4\x4b\x6a\x9b\x4a\x5d\x58\xe2\x93\xc7\x5f\x7f\xfe\xf1\xbb\x58\x26\xa3\x05\x0c\x5e\xc9\x62\x1e\x26\xc9\x3f\xf1\x62\x78\x5e\x1b\xde\xae\x67\xe1\x74\x74\x5e\x12\x7d\x48\x92\x08\x4a\x9b\x21\xb0\x85\x0d\x42\x4d\x98\xc1\x5e\x71\x71\xd6\x48\x00\x4f\x52\xab\xac\x81\x11\x48\x87\x37\xa2\x0f\xbf\x41\xa6\x48\x6e\x34\xc2\x15\xcc\xac\x97\x69\x8a\x0e\x5f\x6a\xe5\x7c\x31\xd4\x7b\x79\xa0\x46\xff\x82\x9e\x55\x05\x3b\x74\x6a\xab\x52\xc9\xca\x9a\xcb\x06\xdc\x54\x0e\x57\xa9\x7c\x07\x6f\xeb\xc0\x85\xe4\x66\xba\x29\x3a\x6e\xa9\x08\x95\x43\x42\xc3\x98\xc1\xc6\x23\x50\xf4\x01\x4e\xe3\xdf\x4b\x02\x52\xb9\x69\x0f\x25\xb0\xab\xc9\x23\xef\xc2\x37\x7e\xdb\x5a\xeb\x2f\x3b\x2a\x6b\xde\x78\x36\x7e\x67\x91\x7f\xec\x09\xd2\x64\xe7\x48\x1f\x1a\x18\x59\xb6\x2d\x97\x92\xd3\x02\xa9\x01\x58\x83\xa0\xcc\xfb\xbf\x7a\x29\x92\x24\x5a\x4f\xe3\xa1\x8f\xec\x38\x6c\x21\x56\xda\xe6\xe4\x17\x69\xaa\x8c\x2a\xeb\x12\xb4\xcd\x41\xe3\x0e\x75\x20\xa2\x78\xbc\x8e\x46\x4f\xa3\xc8\xaf\x6b\x13\x6e\x61\x6b\x9d\x79\x08\xf9\x70\xf7\x4e\x31\xa3\xf1\x51\x6f\x95\xc6\x40\xdc\x4f\xa2\xd1\x3a\x8a\xc7\xe3\xc9\x6c\x0c\x03\xdf\x3b\x7e\x4e\x4b\xad\x21\xab\x11\x2e\x88\x33\x5b\xf3\x65\x20\xee\xe2\x59\x12\xff\x54\x65\x2e\xb9\x38\x59\x81\x32\x6c\xcf\xee\x4b\x23\x2d\xdd\xab\x76\xdb\xfb\x3c\x7c\x7c\x68\x56\x3c\xa7\x4e\x25\x0f\x25\x1a\xa6\x40\xdb\xdc\xab\x0d\x3e\xfa\xf8\x7e\x8f\x57\x09\xb6\xd6\xb5\x76\xda\xe6\xb9\x32\xf9\x2f\x78\x53\xf9\x0d\x48\x7d\x47\xb0\xdb\x53\x57\x2d\x7d\x83\x5b\xeb\x10\x14\x13\x38\xcb\xd2\xe7\xa9\x0c\x4c\x6f\x03\x31\x0d\xff\x5d\x27\x93\xff\x7c\x24\xd7\xdd\xa3\x84\xcc\xbf\xaa\x90\xc9\x03\xb5\x1a\xe1\xf8\x5c\xa2\x09\xd2\xd4\xe5\x06\x9d\x57\x3a\x31\x4e\x6a\xe4\xc7\xf8\x8c\x58\xb5\xdc\xdb\xf0\xee\xef\xe5\x3c\x69\xf9\x62\x55\xaa\xd4\x59\xbf\x63\x2a\xc5\x66\x33\x96\x4e\xb7\xa3\xb0\x5c\xa0\x6b\xaf\xf4\x1b\x50\xf3\x5c\x24\xeb\xe5\xc2\x2f\x4a\xc1\x5c\xdd\x74\x3a\xaf\xef\xca\x4d\xaf\xdb\xeb\x76\x64\xa5\x3a\xbb\x6b\xf1\x23\x00\x00\xff\xff\x0a\x97\x2e\xf5\x36\x05\x00\x00")

func configIniBytes() ([]byte, error) {
	return bindataRead(
		_configIni,
		"config.ini",
	)
}

func configIni() (*asset, error) {
	bytes, err := configIniBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "config.ini", size: 1334, mode: os.FileMode(420), modTime: time.Unix(1609248236, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _migrations01_initialUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x90\x4d\x4e\xc3\x30\x10\x85\xf7\x3e\xc5\x5b\x82\xd4\x9e\x80\x95\x0b\x53\x88\x70\x92\x2a\x99\x0a\xca\x26\x32\xcd\x20\x8a\x92\xda\xca\x0f\x52\x6e\x8f\x70\x42\x17\x2d\x48\x59\x30\x3b\xcb\xef\x7b\xf6\x7c\xcb\x25\xd8\xbe\x56\x02\xf7\x06\x6f\x87\x5a\x8e\x5d\x8b\x46\xf6\x72\xf8\x94\x12\x5e\x1a\xf8\xc6\x7d\xc8\xbe\x83\x7f\xb7\xad\xa8\xdb\x8c\x34\x13\x58\xaf\x0c\x9d\x00\x75\xa5\x00\xe0\x50\xe2\x67\x72\xca\x22\x6d\xb0\xc9\xa2\x58\x67\x3b\x3c\xd2\x6e\x11\x22\xa1\xa4\x08\xc1\x28\x61\xba\xa7\x0c\x17\x93\xa4\x8c\x64\x6b\xcc\x42\x05\xc4\xd6\xae\x3f\x76\xe1\x66\x06\x12\x1e\xb1\x83\x4c\x5f\x59\xa5\xa9\x21\x9d\xe0\x8e\xd6\x7a\x6b\x18\x6b\x6d\x72\x9a\x8a\x4b\xdb\x49\x51\xf6\x02\x80\xa3\x98\x72\xd6\xf1\x06\x4f\x11\x3f\x84\x23\x5e\xd2\x84\xce\x8a\x03\x31\xb6\xff\x45\xa8\xeb\x1b\xa5\x7e\x95\x3a\x72\x9d\x83\xd4\xbe\x72\x43\xb0\xeb\x7c\x75\xa6\xb4\x6f\xa5\x29\x26\x66\xbe\xd6\x40\x8d\x39\xa6\x67\xbe\x14\xf4\x1f\x5a\x67\x6c\x7f\x22\xbe\x35\x7c\x05\x00\x00\xff\xff\x8c\xc4\x81\x84\x5b\x02\x00\x00")

func migrations01_initialUpSqlBytes() ([]byte, error) {
	return bindataRead(
		_migrations01_initialUpSql,
		"migrations/01_initial.up.sql",
	)
}

func migrations01_initialUpSql() (*asset, error) {
	bytes, err := migrations01_initialUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "migrations/01_initial.up.sql", size: 603, mode: os.FileMode(420), modTime: time.Unix(1609248204, 0)}
	a := &asset{bytes: bytes, info: info}
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

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
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
	"config.ini":                   configIni,
	"migrations/01_initial.up.sql": migrations01_initialUpSql,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("nonexistent") would return an error
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
	"config.ini": &bintree{configIni, map[string]*bintree{}},
	"migrations": &bintree{nil, map[string]*bintree{
		"01_initial.up.sql": &bintree{migrations01_initialUpSql, map[string]*bintree{}},
	}},
}}

// RestoreAsset restores an asset under the given directory
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
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
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