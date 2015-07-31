// Code generated by go-bindata.
// sources:
// files/app.css
// files/app.js
// files/index.html
// DO NOT EDIT!

package static

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
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

var _filesAppCss = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x6c\x91\xdf\xaa\xa3\x30\x10\xc6\xaf\xeb\x53\x0c\x2c\x0b\x2d\x54\xb1\xdb\x85\x05\x65\xf7\xa2\xaf\xb1\xf4\x22\x9a\x34\x86\x8d\x19\x89\xe3\xb6\x9e\xd2\x77\x3f\x93\x28\x3d\xf6\xcf\x8d\x98\xe4\xfb\x7d\x33\xf3\x4d\x43\xad\xdd\x42\x85\x72\x84\x6b\xb2\x3a\x1b\x49\x4d\x01\xbb\x3c\xff\x5e\x26\xab\x4e\x48\x69\x9c\x2e\x20\xe7\x43\x2b\xbc\x36\x2e\xfe\xdf\x92\x24\x00\x8c\x0d\x44\xe8\xb6\x40\xea\x42\xc2\x2b\x11\x2c\x4e\xe8\x28\x3d\x89\xd6\xd8\xb1\x80\x16\x1d\xf6\x9d\xa8\x55\x84\x92\xac\x16\x5a\x05\xd1\xdd\xf9\x47\xde\x5d\xe2\xdb\xd2\xe2\xb1\x8b\x46\x19\xdd\x10\x1f\x27\xed\x54\xa0\x37\x1f\x8a\xaf\x7e\xce\x34\xd7\x70\x59\x4f\x82\x86\x9e\x0d\x00\x2a\x51\xff\xd3\x1e\x07\x27\x0b\xf8\x76\x38\x1c\x4a\xbe\xbb\xd7\xdc\x07\x88\x35\xe8\xa5\xf2\xa9\x17\xd2\x0c\x3d\x77\x32\x59\x2d\x9c\x0a\x87\xb4\xfe\x5b\xa3\x45\x7f\xdc\xc0\x15\xa4\xe9\x3b\x2b\x78\x2a\x87\x4e\x95\xf0\xa0\x9d\x64\xbf\xb5\x57\xca\x1d\x59\xbb\x6c\xc0\xeb\x6a\xbd\xdb\xff\xda\xf2\xb0\xfc\xe1\xbf\xcd\x7b\xd8\x2b\xf9\x82\x2a\x59\x42\x7c\x2d\xe0\xdc\x18\x52\xe5\x3b\xb0\xb2\x83\x7a\x21\x71\x14\x36\x3c\x3c\xf1\x5c\x79\x5e\xc3\x1f\x1e\xe8\x7f\x88\x7b\xda\x6c\x4a\xd8\x85\xcc\x63\xc6\xf7\xac\xe6\xf3\x94\x55\x01\xd4\x18\x07\x3d\x5a\x23\x81\x47\x1d\x63\xf6\x99\x45\x0d\x9d\x8f\x7b\x8d\x35\xd2\xb8\xf1\x10\x93\x6f\x85\x2d\x9f\xf6\xb9\x24\x32\xe3\x4e\x18\xb8\xb9\xc5\x57\xcf\xac\xf3\xc6\xd1\x42\xf2\x35\x57\x72\xfb\x0c\x00\x00\xff\xff\x46\x7f\x87\xbf\xbc\x02\x00\x00")

func filesAppCssBytes() ([]byte, error) {
	return bindataRead(
		_filesAppCss,
		"files/app.css",
	)
}

func filesAppCss() (*asset, error) {
	bytes, err := filesAppCssBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "files/app.css", size: 700, mode: os.FileMode(420), modTime: time.Unix(1426359669, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _filesAppJs = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x7c\x56\xdd\x6e\xe3\x44\x14\xbe\xb6\x9f\x62\x76\x90\x56\xce\x26\x72\x5a\x10\x37\x84\x82\x16\x58\x09\x50\x77\x17\x48\x25\x90\xca\x5e\x4c\xed\x93\x64\xb4\xf6\x8c\x77\x66\x9c\x34\xa0\xbe\x3b\xdf\x99\xb1\x53\x97\x76\xdb\x8b\xb6\x3e\xe7\x3b\xff\x7f\x93\xef\x95\x13\x95\x35\x1b\xbd\x15\x17\xa2\xb6\x55\xdf\x92\x09\xe5\xa7\x9e\xdc\x71\x4d\x0d\x55\xc1\xba\x42\x06\xba\x0d\xca\x91\x92\xb3\x55\x14\x68\xec\x73\xe8\x2f\xc0\x66\x64\x84\xb6\xea\xf6\xd2\x6e\xd7\xfa\x1f\x82\xc4\x97\xf4\x55\x52\xe0\xe9\x13\x3e\xcf\xd2\x47\xad\xf7\x1e\x5f\xd7\x1f\x06\x91\xca\xb6\x9d\x6e\x18\xbf\xe9\x4d\x15\xb4\x35\xc5\x4c\xfc\x9b\x67\xcc\xbb\xb1\xf5\x11\x8c\xe4\x71\xb9\x57\x4d\x4f\xab\x3c\x0b\xee\xc8\x80\xec\xd7\xf5\xfb\x77\x65\xa7\x9c\xa7\x82\x81\xf0\x21\xbb\x13\x95\x0a\xd5\xae\x20\xe7\xa2\x92\x4c\x35\xe4\x42\x21\xe5\x9c\x29\x00\x64\x8e\x42\xef\x0c\x43\xf3\x64\xe3\x76\xe7\x60\xc2\xd0\x41\xfc\xf5\xf6\xf2\xe7\x10\xba\x3f\x08\x11\xfa\x50\x30\x1c\xcc\xd2\x76\x64\x0a\xf9\xdb\xfb\xf5\x95\x5c\x08\xb9\x1c\xfc\x95\x27\xb6\x69\xac\xaa\x1f\xb9\x9f\xe9\x4d\xc1\x6c\x1f\x54\xe8\xbd\x78\x71\x81\x7c\x9c\x9d\x25\xd6\xe0\x15\xb3\x1d\xf9\xce\x1a\x4f\x57\xc8\x79\xf4\xef\x0e\x9e\x0d\x9a\x3d\x99\x7a\x8c\x0c\xb4\x3c\x5f\x2e\x2f\x9e\xf8\xc9\xf3\xd1\xb2\x48\xc6\x8a\x70\xec\x68\x21\x5a\xbf\x5d\x20\x75\x8d\x75\xa7\x7c\xa2\x68\xed\x33\xa5\x2c\xe5\x9c\x45\xe7\x72\xf0\x3a\xc6\x88\x38\x5e\xb0\xdc\x4c\xa4\xd4\x71\x35\xbc\x6d\xa8\x3c\x28\x87\xbc\x18\x2b\x90\x16\x16\x63\x30\x03\x4b\x6d\x0c\x39\x0e\x08\xa6\xe0\x44\xd2\x91\x1c\x41\x80\x11\xe2\x29\xbc\x0e\xc1\xe9\x9b\x3e\x50\x21\x23\x4f\x8e\xce\xa6\x60\x9f\x8e\x75\x08\x98\x63\xb1\x66\x6b\x6f\x69\x8f\x30\xa6\xb9\xa7\x18\x2b\xec\x51\x09\xd3\x29\xdd\xd6\xb4\xe4\xbd\xda\xd2\x40\x8c\x7d\x42\x8d\x27\x11\x71\x3e\xf8\x11\x37\xe4\x2f\xd1\x62\x8f\xdc\x8d\xc6\x06\x15\x53\x5b\xa3\x81\xd4\xde\x6f\x55\xd8\x95\xe8\xff\x02\x9f\x31\xf9\xa5\xae\x59\x47\xe7\xb4\x09\x8c\x2d\x43\x2c\xf1\xdd\x6a\x54\x39\xb4\xc6\x44\xa3\x1f\xbd\xf7\x65\xd5\x3b\x87\xd8\x38\x63\x83\x57\x72\x20\x21\x51\x27\x36\xda\xbf\xfa\x08\xb7\xd0\x98\x5b\x47\x64\x64\x2a\x82\xa7\x27\xc5\x50\xac\x30\xcc\x9b\x36\x5b\x99\x84\x8e\x71\x78\x4f\x68\xb4\x44\x4f\xd1\x84\xe9\xdb\xdf\xf9\xa3\x16\x73\x21\x85\x0e\xd4\x7a\x11\xb9\xb5\x5c\xb0\xfa\x09\xe0\x7b\x21\x6f\x30\x9b\x52\x7c\xf3\x94\x4a\xb6\xd8\x50\x60\xb9\xa4\xf6\x27\x6b\x68\xa2\x74\xc2\x1f\xf5\x46\xc4\x53\x5a\xc7\xec\x69\xb3\xb1\x0f\x32\x17\x52\x97\x6b\xcc\x12\x0f\x3c\xf3\xd9\x5c\x70\x13\xa1\x58\x8a\xe7\xa4\x22\xe0\x91\x58\xe2\x4e\xe5\xaa\xc6\x27\xd0\x38\x58\xc8\x27\x71\x25\x41\x2b\x7d\xd7\x68\xe8\xfa\x1b\xc5\xc0\x78\xef\x89\xf7\x53\x1c\xa4\xa8\xe6\x92\x91\x49\x41\x14\x7a\x64\xe7\x72\x50\xf5\xd0\x58\xc2\xb2\xb9\x0d\x26\x35\xa2\xe3\x4a\xc5\x9f\x6f\x13\xb3\x6c\xc8\x6c\xc3\x0e\x94\xf9\x3c\xf5\x73\x74\x0c\xa8\xc8\xbe\xd6\xbc\x70\x13\x31\xd8\x8e\xb7\x00\x16\xf1\xf5\xd9\x87\x55\xda\x55\x4c\x7b\xf9\x92\xb5\x62\x53\x8d\x7b\x0a\xc4\xc9\x34\xcf\xa1\x8a\xd1\x19\x36\x40\xd0\xa6\xa7\xb8\xae\x06\x9d\xd0\x36\xdd\x2c\x15\xee\x47\xa0\x37\x98\x75\x7c\x15\x12\xdc\xd8\x9a\x19\xfe\x29\xab\x46\x79\xff\x4e\xb5\x3c\x4c\x08\x6e\x24\x4f\xd7\x46\x13\x7d\xc5\x65\x29\x53\x52\x7e\x20\x84\x4d\x05\x60\x0b\xf6\x7e\x54\xe5\xcb\xde\xf8\x9d\xde\x04\xe6\xa4\x81\xcd\x0e\x3b\x6c\xe7\x22\x32\x53\x46\xc4\x77\x93\xb3\x94\x02\x63\xc5\x8e\x5a\xbb\xa7\x1f\x81\xae\x13\xba\xb3\x5d\x31\x1b\xa6\xfe\xb3\x0b\x28\x55\x2a\xa4\x53\x76\xf0\xff\xbb\x6f\xcb\xa5\xea\x83\x15\x1d\xa6\x2b\xc7\x1e\x08\xbf\x98\x40\x0e\x67\xab\x38\x6d\x68\x66\x15\xe3\x94\x1f\x3c\xa7\xfc\xe0\xe1\x8b\xaa\x8f\x6b\x8c\x0b\xc5\xf4\x9f\xf3\xd4\x83\x1c\x0f\x80\x64\x91\x17\xb1\xf7\x17\xe2\xfc\xeb\x57\xe7\x67\x38\x23\xb3\x93\x2d\x47\xa8\x86\xc1\x02\xcf\xef\x8d\x9c\x68\xc9\xd2\xfd\x1c\x46\x22\x10\x3c\xfa\xe3\x57\x5a\x04\x71\xd0\x38\xf8\x83\x1f\xae\xe1\x9f\x74\xb3\xb6\xd5\x47\x0a\x45\x63\x71\x54\x21\x55\x5a\xa7\xb7\xda\xc0\xdb\xae\x51\x15\xf6\xf6\x0e\xd7\x52\x2e\xe4\x01\x87\x82\xc7\x79\xc9\x2f\x01\xe8\x92\x73\x24\x24\x6e\x00\xc4\x60\x0d\xdf\xcf\xc7\xe7\xf1\x59\xa7\xe2\xaa\x98\x6c\xb4\x8c\x7b\x02\x81\xaf\xe2\x71\x1c\xf4\x56\x8d\xf5\x8f\x9f\x0d\xb1\x19\xfd\xb8\x8f\x9d\xed\x91\xc2\xb0\x84\xec\x0c\xbf\x56\x9f\x35\x5c\x6b\x7f\xb2\x2d\x0a\x5c\x3a\xbc\x31\xb4\x11\x08\x65\x2e\xfd\x8c\x11\x0e\x4e\x45\x5f\x50\xd7\x2b\xdd\x92\xed\x43\x71\x4a\x34\xba\x72\xf0\xf3\x15\xee\xfc\x03\x37\x9f\x39\x1c\xe9\x26\x47\xb9\xe1\x49\x93\x31\x6e\xf2\xae\xe1\xb3\x51\xab\xa0\xd2\xd3\x60\x7c\xdc\x0c\xd3\x79\x7a\xca\x70\xdb\xa7\x3b\x3d\x7b\x48\xe7\xfb\xc6\x6b\x25\x28\x53\x91\xdd\x88\xd7\xce\xa9\x63\xc4\x50\x89\x81\x7a\xa3\xa0\xed\xfe\x8e\x46\x23\xc3\xf1\xc8\xee\xc9\x45\xbc\xec\x08\xe8\x2e\x8e\x47\x9e\xff\x17\x00\x00\xff\xff\x8c\xd5\x29\xd6\x46\x0a\x00\x00")

func filesAppJsBytes() ([]byte, error) {
	return bindataRead(
		_filesAppJs,
		"files/app.js",
	)
}

func filesAppJs() (*asset, error) {
	bytes, err := filesAppJsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "files/app.js", size: 2630, mode: os.FileMode(420), modTime: time.Unix(1426442313, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _filesIndexHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x84\x92\xcf\x6e\x23\x21\x0c\xc6\xcf\x99\xa7\x40\x9c\x76\xa5\x68\xc8\x5e\x57\x64\x2e\x39\xec\x2b\xec\x95\x80\x9b\xa1\x21\x40\xb1\xc9\x1f\x55\x7d\xf7\x9a\x90\xaa\x91\x1a\x29\x37\xe6\xf3\xe7\xcf\x3f\x6b\xac\x67\x3a\x84\x69\xd0\x33\x18\x37\x0d\x0b\x4d\x9e\x02\x4c\x9b\x90\xaa\x13\xff\xd2\x7f\xad\xba\xc0\x95\xe0\xe3\x5e\x14\x08\x6b\x89\x74\x09\x80\x33\x00\x49\x31\x17\x78\x59\x4b\x93\xf3\x68\x11\x25\x07\xa9\x9e\xa4\xb7\xc9\x5d\xa6\x81\xfb\x9c\x3f\x0a\x1b\x0c\xe2\x5a\x5a\xb3\x03\xf6\x2c\x78\x0c\x9c\xc9\x14\x30\xd3\xfb\x20\x84\xcc\xc6\xee\x5b\xe9\xaf\x90\x3b\x4f\x73\xdd\x8e\x36\x1d\x54\x45\x28\xaa\x40\x4e\x72\xd9\x4c\x47\x28\xe8\x53\x6c\xa6\x3f\xe3\x6a\x5c\x75\x95\x81\xc0\xe0\xb5\xb5\x0b\x36\x45\xa4\x62\x7c\x24\x6c\xa2\x33\xe5\xe4\xe3\x92\xe1\xeb\x79\xc9\x2f\x97\x4e\x28\x87\x0f\xde\xeb\x8b\x60\x68\x3c\xf7\x90\x29\x52\x49\x01\xaf\xa0\x0b\xbd\xad\x44\x29\x8a\x14\x6d\xf0\x76\xdf\xca\x87\xec\x03\xfc\xfa\x2d\xa7\x4d\x7f\x6a\xd5\x3d\xdd\x8f\xd9\xc4\xbb\xa8\x08\x96\x98\x5a\x20\x19\xaa\x9c\xa9\x55\x33\x3c\xb0\xd6\x52\x20\xd2\x53\xdf\x5b\x85\x0a\xcf\xd3\x98\x2c\x00\x81\x7b\xe0\xd4\x8a\x97\xfd\xb1\x75\x48\xbb\xdb\xc2\xb9\x80\xf0\xee\xa6\x68\xc5\x9f\x77\x5d\xdf\xdd\x1a\x6d\xf1\x99\x81\x8b\xed\xff\xff\xb5\x4f\xb9\xaa\xed\x0e\xfa\x01\xf0\x3d\xb4\x03\xfb\x0c\x00\x00\xff\xff\x42\xc0\x40\x2c\x67\x02\x00\x00")

func filesIndexHtmlBytes() ([]byte, error) {
	return bindataRead(
		_filesIndexHtml,
		"files/index.html",
	)
}

func filesIndexHtml() (*asset, error) {
	bytes, err := filesIndexHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "files/index.html", size: 615, mode: os.FileMode(420), modTime: time.Unix(1426439668, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
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
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
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
	"files/app.css":    filesAppCss,
	"files/app.js":     filesAppJs,
	"files/index.html": filesIndexHtml,
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
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
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
	"files": &bintree{nil, map[string]*bintree{
		"app.css":    &bintree{filesAppCss, map[string]*bintree{}},
		"app.js":     &bintree{filesAppJs, map[string]*bintree{}},
		"index.html": &bintree{filesIndexHtml, map[string]*bintree{}},
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
	err = os.MkdirAll(_filePath(dir, path.Dir(name)), os.FileMode(0755))
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
		err = RestoreAssets(dir, path.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}