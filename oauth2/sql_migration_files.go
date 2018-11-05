// Code generated by go-bindata. (@generated) DO NOT EDIT.
// sources:
// migrations/sql/shared/1.sql
// migrations/sql/shared/2.sql
// migrations/sql/shared/3.sql
// migrations/sql/shared/4.sql
// migrations/sql/mysql/.gitkeep
// migrations/sql/mysql/5.sql
// migrations/sql/mysql/6.sql
// migrations/sql/mysql/7.sql
// migrations/sql/postgres/.gitkeep
// migrations/sql/postgres/5.sql
// migrations/sql/postgres/6.sql
// migrations/sql/postgres/7.sql
// migrations/sql/tests/.gitkeep
// migrations/sql/tests/1_test.sql
// migrations/sql/tests/2_test.sql
// migrations/sql/tests/3_test.sql
// migrations/sql/tests/4_test.sql
// migrations/sql/tests/5_test.sql
// migrations/sql/tests/6_test.sql
// migrations/sql/tests/7_test.sql
package oauth2

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

var _migrationsSqlShared1Sql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x94\x31\x6f\xf2\x30\x10\x86\x67\xfb\x57\xdc\x08\xfa\x60\x41\x62\x62\xca\x57\x8c\x84\x9a\x02\x0a\x41\x2a\x53\x74\xb2\x8f\xc4\x52\x63\xa7\xb6\x53\xda\x7f\x5f\x85\xa6\x82\x22\x22\x31\x30\x26\xeb\xf3\xde\x1b\xeb\xb1\x75\xe3\x31\xfc\x2b\x75\xee\x30\x10\xec\x2a\xfe\x94\x88\x28\x15\x90\x46\xff\x63\x01\xcb\x05\xac\xd6\x29\x88\xd7\xe5\x36\xdd\x42\xf1\xa5\x1c\x66\x16\xeb\x50\x4c\x32\x94\x92\xbc\x87\x01\x67\x5e\xe7\x06\x43\xed\x08\x4e\x1f\xfb\x40\x27\x0b\x74\x83\xc9\x74\x3a\x3c\x8d\xaf\x76\x71\x0c\x9b\x64\xf9\x12\x25\x7b\x78\x16\xfb\x11\x67\x8e\xde\x6b\xf2\x21\xd3\x0a\x80\x01\xdc\x1c\x39\xc7\x48\x65\x18\x00\x58\xd0\x25\xf9\x80\x65\x75\xae\x9d\x8b\x45\xb4\x8b\x53\x30\xf6\x38\x18\x8e\x38\x93\x6f\x9a\x4c\xdb\xcb\x00\x02\x7d\x86\xcb\x42\x2f\x6d\x45\x0d\x62\xcd\x51\xaf\x69\xee\xd0\x34\x3f\xfb\x49\xb1\x6b\x7c\xb0\xae\xcc\x14\x06\xec\xe8\x26\xef\xb5\x35\xbf\x89\x3f\x98\x0f\x67\x77\x9b\x75\x74\x70\xe4\x8b\x5e\xed\xe3\xd5\x4a\xab\xa8\xf7\xfa\x78\xaf\x56\x2b\xd9\x7b\xbd\xdf\x2b\xbf\x5c\xba\x73\x7b\x34\x7c\x9e\xac\x37\xad\xe7\x1b\x6b\x76\xd6\xc9\xdb\x65\xd1\x1d\x68\x9e\x7c\x37\x6d\x2e\x6e\xc6\xbf\x03\x00\x00\xff\xff\x30\xe3\x7b\xfb\x03\x06\x00\x00")

func migrationsSqlShared1SqlBytes() ([]byte, error) {
	return bindataRead(
		_migrationsSqlShared1Sql,
		"migrations/sql/shared/1.sql",
	)
}

func migrationsSqlShared1Sql() (*asset, error) {
	bytes, err := migrationsSqlShared1SqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "migrations/sql/shared/1.sql", size: 1539, mode: os.FileMode(438), modTime: time.Unix(1541453045, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _migrationsSqlShared2Sql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x90\xb1\x0a\xc2\x30\x14\x45\xf7\x7e\xc5\xdd\xaa\x48\x97\x42\x27\xa7\x68\xea\x14\x5b\x29\xc9\x5c\x62\x1a\x4d\x05\x8d\xbc\xb4\x8a\x7f\x2f\x14\x04\x07\x45\x0b\xfd\x80\x73\x2e\xf7\x24\x09\x16\xe7\xf6\x48\xba\xb3\x50\xd7\x88\x09\x99\x57\x90\x6c\x25\x72\xb8\x47\x43\xba\xf6\xba\xef\x5c\x5a\x6b\x63\x6c\x08\x60\x9c\x23\xf4\xfb\x93\x35\x1d\x6e\x9a\x8c\xd3\x34\x4b\xb3\x6c\x8e\xa2\x94\x28\x94\x10\xe0\xf9\x86\x29\x21\x11\xc7\xcb\xef\x36\xb2\x07\xb2\xc1\x4d\xa5\x33\xbe\xb1\x53\xb9\x7c\xdb\x98\x91\xae\xe8\x3d\x22\xf7\xf7\xcb\xcf\x8c\xe0\x55\xb9\xc3\xba\x14\x6a\x5b\xbc\x86\xfe\xc8\x35\x8e\x1a\xaa\x8c\x43\x86\xf3\x1f\x91\x67\x00\x00\x00\xff\xff\xa3\x05\x9b\x27\x28\x02\x00\x00")

func migrationsSqlShared2SqlBytes() ([]byte, error) {
	return bindataRead(
		_migrationsSqlShared2Sql,
		"migrations/sql/shared/2.sql",
	)
}

func migrationsSqlShared2Sql() (*asset, error) {
	bytes, err := migrationsSqlShared2SqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "migrations/sql/shared/2.sql", size: 552, mode: os.FileMode(438), modTime: time.Unix(1541453045, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _migrationsSqlShared3Sql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x91\xb1\x6e\xc2\x30\x10\x86\x67\xfb\x29\x6e\x24\x2a\x2c\x48\x4c\x4c\x69\x63\xa4\xa8\x29\xa0\x90\x48\x65\x8a\xae\xc9\x35\x71\xdb\xd8\xa9\x7d\x29\xed\xdb\x57\x01\x2a\x10\x50\xaf\xdf\x7f\xdf\xd9\xfe\x27\x13\xb8\x6b\x75\xed\x90\x09\xf2\x4e\x3e\xa4\x2a\xcc\x14\x64\xe1\x7d\xa2\x20\x5e\xc0\x72\x95\x81\x7a\x8e\x37\xd9\x06\x9a\x9f\xca\x61\x61\xb1\xe7\x66\x5a\x74\xef\x25\xc1\x48\x0a\xaf\x6b\x83\xdc\x3b\x82\xfd\x11\x5f\xe8\xca\x06\xdd\x68\x3a\x9b\x05\xfb\xe1\x65\x9e\x24\xb0\x4e\xe3\xa7\x30\xdd\xc2\xa3\xda\x8e\xa5\x70\xf4\xd9\x93\xe7\x42\x57\x00\x02\xe0\xe6\xc8\x29\x46\x55\x81\x0c\x20\x58\xb7\xe4\x19\xdb\xee\xa4\x8d\xd4\x22\xcc\x93\x0c\x8c\xdd\x8d\x82\xb1\x14\xe5\x87\x26\x73\xf4\x0a\x00\xa6\x6f\x3e\x17\xfa\xd2\x76\x34\x20\x31\x5c\xf5\x92\xd6\x0e\xcd\xb0\xec\x90\x12\x97\xf8\xd5\xba\xb6\xa8\x90\xf1\x1f\x37\x79\xaf\xad\xf9\x4b\x5c\xe1\xfe\xe5\x8d\x4a\x86\xc3\xea\x9b\x2f\x96\xc1\x5c\xca\xf3\x36\x22\xbb\x33\x32\x4a\x57\xeb\x63\x1b\x57\xff\x3f\x97\xbf\x01\x00\x00\xff\xff\x58\x5f\xd7\xfa\xbd\x01\x00\x00")

func migrationsSqlShared3SqlBytes() ([]byte, error) {
	return bindataRead(
		_migrationsSqlShared3Sql,
		"migrations/sql/shared/3.sql",
	)
}

func migrationsSqlShared3Sql() (*asset, error) {
	bytes, err := migrationsSqlShared3SqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "migrations/sql/shared/3.sql", size: 445, mode: os.FileMode(438), modTime: time.Unix(1541453045, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _migrationsSqlShared4Sql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\xd0\xb1\x0e\x82\x30\x10\xc6\xf1\x9d\xa7\xf8\x76\xc3\xe2\xea\x54\x6c\x9d\x4e\x6a\x48\x3b\x93\xa6\x9c\x42\x8c\x96\x14\xd4\xf8\xf6\x26\xc4\xc1\x41\x23\xa4\x0f\xf0\xfd\xef\xf2\xcb\x73\xac\x2e\xdd\x29\xba\x91\x61\xfb\x4c\x90\x51\x15\x8c\x28\x48\xa1\x7d\x36\xd1\xd5\xc1\xdd\xc6\x76\x5d\x3b\xef\x79\x18\x20\xa4\x84\xf3\x63\x77\x67\x14\x5a\x13\x4a\x6d\x50\x5a\x22\x48\xb5\x13\x96\x0c\x4c\x65\xd5\xe6\x77\x26\xf2\x31\xf2\xd0\x26\x77\x7c\x68\x38\x39\x12\xba\xc6\x27\x47\xfa\xb3\x9f\xfd\x49\xf6\xa9\x2d\xc3\xe3\xfa\xd7\x5b\x56\xfa\x80\xad\x26\xbb\x2f\xdf\x07\x66\xe0\x2e\x1a\x4d\x92\x8b\x16\x13\xdb\xa2\xc5\x64\xf4\x6d\xf1\x0a\x00\x00\xff\xff\xcb\x7a\x0d\xdc\x7e\x02\x00\x00")

func migrationsSqlShared4SqlBytes() ([]byte, error) {
	return bindataRead(
		_migrationsSqlShared4Sql,
		"migrations/sql/shared/4.sql",
	)
}

func migrationsSqlShared4Sql() (*asset, error) {
	bytes, err := migrationsSqlShared4SqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "migrations/sql/shared/4.sql", size: 638, mode: os.FileMode(438), modTime: time.Unix(1541453045, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _migrationsSqlMysqlGitkeep = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00")

func migrationsSqlMysqlGitkeepBytes() ([]byte, error) {
	return bindataRead(
		_migrationsSqlMysqlGitkeep,
		"migrations/sql/mysql/.gitkeep",
	)
}

func migrationsSqlMysqlGitkeep() (*asset, error) {
	bytes, err := migrationsSqlMysqlGitkeepBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "migrations/sql/mysql/.gitkeep", size: 0, mode: os.FileMode(438), modTime: time.Unix(1541453045, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _migrationsSqlMysql5Sql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd2\xd5\x55\xd0\xce\xcd\x4c\x2f\x4a\x2c\x49\x55\x08\x2d\xe0\x72\x0e\x72\x75\x0c\x71\x55\x08\xf5\xf3\x0c\x0c\x75\x55\xf0\xf4\x73\x71\x8d\x50\xc8\xa8\x4c\x29\x4a\x8c\xcf\x4f\x2c\x2d\xc9\x30\x8a\x4f\x4c\x4e\x4e\x2d\x2e\x8e\x2f\x4a\x2d\x2c\x4d\x2d\x2e\x89\xcf\x4c\x89\xcf\x4c\xa9\x50\xf0\xf7\xc3\xa6\x4a\x41\x03\xa1\x4c\xd3\x9a\xb0\xd9\x45\xa9\x69\x45\xa9\xc5\x19\x84\x0c\x87\x2a\x43\x33\x9d\x0b\xd9\x27\x2e\xf9\xe5\x79\x5c\x2e\x41\xfe\x01\x94\x7a\xc1\x1a\xa7\x29\xa4\x39\xd6\x9a\x0b\x10\x00\x00\xff\xff\x5b\x41\x11\xcf\x69\x01\x00\x00")

func migrationsSqlMysql5SqlBytes() ([]byte, error) {
	return bindataRead(
		_migrationsSqlMysql5Sql,
		"migrations/sql/mysql/5.sql",
	)
}

func migrationsSqlMysql5Sql() (*asset, error) {
	bytes, err := migrationsSqlMysql5SqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "migrations/sql/mysql/5.sql", size: 361, mode: os.FileMode(438), modTime: time.Unix(1541453045, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _migrationsSqlMysql6Sql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd2\xd5\x55\xd0\xce\xcd\x4c\x2f\x4a\x2c\x49\x55\x08\x2d\xe0\x72\x0e\x72\x75\x0c\x71\x55\xf0\xf4\x73\x71\x8d\x50\xc8\xa8\x4c\x29\x4a\x8c\xcf\x4f\x2c\x2d\xc9\x30\x8a\x4f\x4c\x4e\x4e\x2d\x2e\x8e\x2f\x4a\x2d\x2c\x4d\x2d\x2e\x49\x4d\x89\x4f\x2c\x89\xcf\x4c\xa9\x50\xf0\xf7\xc3\xa6\x4e\x41\x03\x59\xa1\xa6\x35\x17\x17\xb2\x45\x2e\xf9\xe5\x79\x5c\x2e\x41\xfe\x01\x94\x5b\x64\xcd\x05\x08\x00\x00\xff\xff\xf1\x5f\x27\x42\xc2\x00\x00\x00")

func migrationsSqlMysql6SqlBytes() ([]byte, error) {
	return bindataRead(
		_migrationsSqlMysql6Sql,
		"migrations/sql/mysql/6.sql",
	)
}

func migrationsSqlMysql6Sql() (*asset, error) {
	bytes, err := migrationsSqlMysql6SqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "migrations/sql/mysql/6.sql", size: 194, mode: os.FileMode(438), modTime: time.Unix(1541453045, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _migrationsSqlMysql7Sql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x95\x4d\x6a\xc3\x30\x10\x85\xf7\x39\x85\x76\x59\x94\x6c\xba\x0d\x5d\xb8\x95\x0b\x05\x27\x0e\xa9\x0c\xed\xca\x08\x69\x1a\x9b\x52\x3b\x95\x6d\x4a\x6f\x1f\x18\xb0\x10\xfe\xc1\x99\xf1\x01\x9e\xde\xa7\x6f\x64\xcf\x6e\x27\x1e\x7e\xca\x8b\xd3\x2d\x88\xec\xba\x89\x12\x15\x9f\x85\x8a\x9e\x93\x58\x14\xff\xd6\xe9\xbc\xd6\x5d\x5b\x3c\xe6\xda\x18\x68\x1a\x11\x49\x29\x1c\xfc\x76\xd0\xb4\x60\x73\xdd\xd9\x12\x2a\x03\x42\xc5\x1f\x4a\x1c\xb3\x24\xd9\x6f\xb2\x93\x8c\xd4\x74\xf8\x3d\x56\x13\xe1\xa7\xed\x76\xbf\xd8\x7b\x48\xe5\xdb\xeb\xe7\x7c\x75\xda\xd7\xdf\x73\x81\x8b\xd3\x15\x13\x7f\x18\xa5\xc0\xcf\xd4\x7a\xf4\xf9\x73\x1c\x7c\x39\x68\x0a\xa6\xfd\x3e\xcd\xd1\xdf\x67\xd7\xfb\x0f\xef\x40\x1d\x40\x78\x03\xda\x04\x06\xfc\xfc\x11\x98\xda\x02\xd3\x3f\x46\x39\xf2\x31\xb8\xde\xbc\x47\xa7\x6a\xf7\xe0\x34\xe7\x21\x36\x5f\x78\x5d\x5a\xc3\x14\x8e\x51\x8e\x70\x0c\xae\x17\xee\xd1\xa9\xc2\x3d\x38\x4d\x78\x88\xcd\x17\x7e\xfd\x36\xdc\x17\x8e\x51\x8e\x70\x0c\xae\x17\xee\xd1\xa9\xc2\x3d\x38\x4d\x78\x88\xbd\x28\x3c\xdc\xb0\xb2\xfe\xab\x16\xd7\x85\x3c\xa7\x27\xf1\x92\x26\xd9\xe1\x38\x21\x65\x79\xdd\x84\xf9\x21\xdd\x3d\x5b\x86\xdd\x3f\x75\x00\x05\x00\x7f\x1c\xec\xf6\x51\x9a\x52\x8d\x9f\x10\xbb\x7a\x94\xa6\x54\xe3\x63\x62\x57\x8f\xd2\xe3\xea\x5b\x00\x00\x00\xff\xff\xd3\x68\x22\x03\xe3\x09\x00\x00")

func migrationsSqlMysql7SqlBytes() ([]byte, error) {
	return bindataRead(
		_migrationsSqlMysql7Sql,
		"migrations/sql/mysql/7.sql",
	)
}

func migrationsSqlMysql7Sql() (*asset, error) {
	bytes, err := migrationsSqlMysql7SqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "migrations/sql/mysql/7.sql", size: 2531, mode: os.FileMode(438), modTime: time.Unix(1541453045, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _migrationsSqlPostgresGitkeep = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00")

func migrationsSqlPostgresGitkeepBytes() ([]byte, error) {
	return bindataRead(
		_migrationsSqlPostgresGitkeep,
		"migrations/sql/postgres/.gitkeep",
	)
}

func migrationsSqlPostgresGitkeep() (*asset, error) {
	bytes, err := migrationsSqlPostgresGitkeepBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "migrations/sql/postgres/.gitkeep", size: 0, mode: os.FileMode(438), modTime: time.Unix(1541453045, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _migrationsSqlPostgres5Sql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd2\xd5\x55\xd0\xce\xcd\x4c\x2f\x4a\x2c\x49\x55\x08\x2d\xe0\x72\x0e\x72\x75\x0c\x71\x55\x08\xf5\xf3\x0c\x0c\x75\x55\xf0\xf4\x73\x71\x8d\x50\xc8\xa8\x4c\x29\x4a\x8c\xcf\x4f\x2c\x2d\xc9\x30\x8a\x4f\x4c\x4e\x4e\x2d\x2e\x8e\x2f\x4a\x2d\x2c\x4d\x2d\x2e\x89\xcf\x4c\x89\xcf\x4c\xa9\x50\xf0\xf7\xc3\xa6\x4a\x41\x03\xa1\x4c\xd3\x9a\xb0\xd9\x45\xa9\x69\x45\xa9\xc5\x19\x84\x0c\x87\x2a\x43\x33\x9d\x0b\xd9\x27\x2e\xf9\xe5\x79\x5c\x2e\x41\xfe\x01\x44\x7b\xc1\x1a\xa7\x72\xec\xae\xb2\xe6\x02\x04\x00\x00\xff\xff\x35\xd0\xf0\x59\x3a\x01\x00\x00")

func migrationsSqlPostgres5SqlBytes() ([]byte, error) {
	return bindataRead(
		_migrationsSqlPostgres5Sql,
		"migrations/sql/postgres/5.sql",
	)
}

func migrationsSqlPostgres5Sql() (*asset, error) {
	bytes, err := migrationsSqlPostgres5SqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "migrations/sql/postgres/5.sql", size: 314, mode: os.FileMode(438), modTime: time.Unix(1541453045, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _migrationsSqlPostgres6Sql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd2\xd5\x55\xd0\xce\xcd\x4c\x2f\x4a\x2c\x49\x55\x08\x2d\xe0\x72\x0e\x72\x75\x0c\x71\x55\xf0\xf4\x73\x71\x8d\x50\xc8\xa8\x4c\x29\x4a\x8c\xcf\x4f\x2c\x2d\xc9\x30\x8a\x4f\x4c\x4e\x4e\x2d\x2e\x8e\x2f\x4a\x2d\x2c\x4d\x2d\x2e\x49\x4d\x89\x4f\x2c\x89\xcf\x4c\xa9\x50\xf0\xf7\xc3\xa6\x4e\x41\x03\x59\xa1\xa6\x35\x17\x17\xb2\x45\x2e\xf9\xe5\x79\x5c\x2e\x41\xfe\x01\x24\x58\x64\xcd\x05\x08\x00\x00\xff\xff\xd2\x18\x3e\xa9\xab\x00\x00\x00")

func migrationsSqlPostgres6SqlBytes() ([]byte, error) {
	return bindataRead(
		_migrationsSqlPostgres6Sql,
		"migrations/sql/postgres/6.sql",
	)
}

func migrationsSqlPostgres6Sql() (*asset, error) {
	bytes, err := migrationsSqlPostgres6SqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "migrations/sql/postgres/6.sql", size: 171, mode: os.FileMode(438), modTime: time.Unix(1541453045, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _migrationsSqlPostgres7Sql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x93\xb1\x4e\x86\x30\x14\x46\x77\x9e\xe2\x6e\xff\x60\x58\x5c\x9d\xaa\xad\x53\x05\x43\xda\xc4\x8d\x34\xed\x15\x88\x91\x62\x0b\x31\xbe\xbd\x09\x03\x21\x82\xa1\x94\x3e\xc0\x3d\xe7\xa4\xf9\x9a\xe7\x70\xf7\xd9\x35\x4e\x8d\x08\x72\xc8\x08\x17\xac\x02\x41\x1e\x39\x83\xf6\xc7\x38\x55\x5b\x35\x8d\xed\x7d\xad\xb4\x46\xef\x81\x50\x0a\x0e\xbf\x26\xf4\x23\x9a\x5a\x4d\xa6\xc3\x5e\x23\x08\xf6\x26\xa0\x90\x9c\x03\x65\xcf\x44\x72\x01\xb7\xdb\x43\x10\xac\x71\xaa\x0f\x40\xfd\xcf\x72\xf8\xee\xd0\xb7\x89\xca\xd6\xb4\xcb\x69\xda\x1a\x4c\xd4\xb5\xa0\x2e\x47\xd9\xce\xe8\x44\x51\x0b\xea\x72\xd4\xf0\xa1\x53\xbd\xd4\x82\x0a\x8d\xca\xb2\xf5\x1f\xa0\xf6\xbb\x3f\x1c\x2e\xad\xca\x57\x78\x2a\xb9\x7c\x29\x76\x8a\x8f\x87\xbf\xbe\xff\x9b\x19\xb2\xf5\x68\xff\x1e\xe0\x4c\xc0\x3c\xc3\x68\xfb\xe6\xfa\x8c\x7a\x1e\x5b\xb4\x7a\x73\x7d\x46\x3d\x4f\x2a\x5a\xbd\xb9\xde\x51\xff\x06\x00\x00\xff\xff\xf4\xd7\x34\xf2\x86\x05\x00\x00")

func migrationsSqlPostgres7SqlBytes() ([]byte, error) {
	return bindataRead(
		_migrationsSqlPostgres7Sql,
		"migrations/sql/postgres/7.sql",
	)
}

func migrationsSqlPostgres7Sql() (*asset, error) {
	bytes, err := migrationsSqlPostgres7SqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "migrations/sql/postgres/7.sql", size: 1414, mode: os.FileMode(438), modTime: time.Unix(1541453045, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _migrationsSqlTestsGitkeep = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00")

func migrationsSqlTestsGitkeepBytes() ([]byte, error) {
	return bindataRead(
		_migrationsSqlTestsGitkeep,
		"migrations/sql/tests/.gitkeep",
	)
}

func migrationsSqlTestsGitkeep() (*asset, error) {
	bytes, err := migrationsSqlTestsGitkeepBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "migrations/sql/tests/.gitkeep", size: 0, mode: os.FileMode(438), modTime: time.Unix(1541453045, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _migrationsSqlTests1_testSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xdc\x93\x41\x4b\x03\x31\x10\x85\xcf\xcd\xaf\x98\x5b\x76\x31\x7b\xa8\x57\x4f\x82\x3d\x14\x64\x0b\xb6\xd5\x63\x18\x92\x69\x36\x60\x93\x3a\x93\x45\x44\xfc\xef\xd2\xcd\xa2\x9e\xbc\xb7\x87\xc0\xbc\xf7\xe0\x3d\xbe\x43\xba\x0e\x6e\x8e\x31\x30\x16\x82\xfd\x49\xad\xfb\xed\xea\x69\x07\xeb\x7e\xb7\x51\x8b\xe1\xc3\x33\xda\x8c\x63\x19\x6e\x2d\x3a\x47\x22\xd0\x48\x0c\x09\xcb\xc8\x64\x80\xe9\x6d\x24\x29\x36\xfa\x9f\x9b\xbc\xc5\x62\xc0\xbd\x46\x4a\x35\x10\x97\x4f\x64\x20\x30\xa6\x73\x3a\xcb\x43\xe6\xa3\xf5\x58\xd0\x80\x90\x48\xcc\x69\x52\xad\x7a\xbe\x7f\xdc\xaf\xb6\x6a\xd1\xe8\x65\x27\x31\x68\x03\x7a\xd9\xcd\xe5\xda\x40\xbf\x79\x69\xda\xc9\xab\x13\x35\x9f\x4a\xeb\x39\xef\xfc\x5a\xe7\xf7\xf9\xa5\xdb\x3b\xf5\x0f\x1c\xd3\x81\x49\x86\x2b\xa5\x73\xd9\xd3\x95\xa2\xe5\xe8\xdd\x45\xa3\xfd\xfd\x7f\x0f\xf9\x3d\xa9\xef\x00\x00\x00\xff\xff\x72\x0a\x2b\x58\x91\x03\x00\x00")

func migrationsSqlTests1_testSqlBytes() ([]byte, error) {
	return bindataRead(
		_migrationsSqlTests1_testSql,
		"migrations/sql/tests/1_test.sql",
	)
}

func migrationsSqlTests1_testSql() (*asset, error) {
	bytes, err := migrationsSqlTests1_testSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "migrations/sql/tests/1_test.sql", size: 913, mode: os.FileMode(438), modTime: time.Unix(1541453045, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _migrationsSqlTests2_testSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xe4\x93\xbf\x4e\xc3\x30\x10\xc6\xe7\xfa\x29\x6e\x73\x22\x9c\x25\x2b\x13\x12\x1d\x2a\xa1\x54\xa2\x2d\x8c\xd6\x61\x5f\x13\x23\x6a\x97\x3b\x47\x08\x21\xde\x1d\xb5\x0e\x7f\x26\x1e\xa0\x19\x2c\xdd\x7d\x9f\xe5\x9f\x7e\x83\x9b\x06\xae\x0e\xa1\x67\xcc\x04\xbb\xa3\x5a\x75\x9b\xe5\xfd\x16\x56\xdd\x76\xad\x16\xc3\xbb\x67\xb4\x09\xc7\x3c\xb4\x16\x9d\x23\x11\xa8\x24\xf4\x11\xf3\xc8\x64\x80\xe9\x75\x24\xc9\x36\xf8\x9f\x99\xbc\xc5\x6c\xc0\xbd\x04\x8a\xa5\x10\x97\x8e\x64\xa0\x67\x8c\xa7\x76\x5a\xf7\x89\x0f\xd6\x63\x46\x03\x42\x22\x21\xc5\xef\x6d\x7c\x7a\x26\x97\x6b\xf5\x70\x73\xb7\x5b\x6e\xd4\xa2\xd2\x6d\x23\xa1\xd7\x06\x74\xdb\x4c\x14\x6d\xa0\x5b\x3f\x56\xf5\x39\x2b\xac\xd2\x9f\x5f\x2f\xe3\x04\xfc\x8d\x4e\xe7\xe3\x73\xba\x57\x28\xba\xbe\x56\xff\x28\x33\xed\x99\x64\x98\x95\xb3\x4b\x9e\x66\x25\x9c\x82\x77\x17\x28\xfc\xf7\x5f\xdf\xa6\xb7\xa8\xbe\x02\x00\x00\xff\xff\x68\x3d\x71\xc0\xe9\x03\x00\x00")

func migrationsSqlTests2_testSqlBytes() ([]byte, error) {
	return bindataRead(
		_migrationsSqlTests2_testSql,
		"migrations/sql/tests/2_test.sql",
	)
}

func migrationsSqlTests2_testSql() (*asset, error) {
	bytes, err := migrationsSqlTests2_testSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "migrations/sql/tests/2_test.sql", size: 1001, mode: os.FileMode(438), modTime: time.Unix(1541453045, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _migrationsSqlTests3_testSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xe4\x94\xcf\x4a\x03\x31\x10\x87\xcf\xcd\x53\xcc\x2d\xbb\x98\xbd\xd8\xa3\x27\xc1\x1e\x0a\xb2\x05\xdb\xea\x31\x8c\xc9\x74\x37\x6a\x93\x35\x93\x45\x44\x7c\x77\xe9\x26\xfe\x39\xf9\x00\xdd\x43\x60\x66\x7e\x21\x1f\x1f\x81\x69\x1a\xb8\x38\xba\x2e\x62\x22\xd8\x0f\x62\xdd\x6e\x57\x77\x3b\x58\xb7\xbb\x8d\x58\xf4\xef\x36\xa2\x0e\x38\xa6\xfe\x52\xa3\x31\xc4\x0c\x15\xbb\xce\x63\x1a\x23\x29\x88\xf4\x3a\x12\x27\xed\xec\x4f\x4d\x56\x63\x52\x60\x5e\x1c\xf9\x1c\xb0\x09\x03\x29\xe8\x22\xfa\x53\x5a\xda\x43\x88\x47\x6d\x31\xa1\x02\x26\x66\x17\xfc\x77\x37\x3e\x3e\x91\x49\xb5\xb8\xbf\xbe\xdd\xaf\xb6\x62\x51\xc9\x65\xc3\xae\x93\x0a\xe4\xb2\x29\x14\xa9\xa0\xdd\x3c\x54\xf5\x34\xcb\xac\x9c\x4f\xaf\xe7\xb2\x00\x7f\x47\xa7\xf3\xf1\x59\xee\x65\x8a\xac\xaf\xc4\x3f\xca\x91\x0e\x91\xb8\x9f\x95\xb3\x09\x96\x66\x25\x1c\x9c\x35\xb3\x12\x1e\x9e\xcd\x39\xfe\xf0\xdf\x45\x76\x13\xde\xbc\xf8\x0a\x00\x00\xff\xff\x83\x53\x8b\x8b\xda\x04\x00\x00")

func migrationsSqlTests3_testSqlBytes() ([]byte, error) {
	return bindataRead(
		_migrationsSqlTests3_testSql,
		"migrations/sql/tests/3_test.sql",
	)
}

func migrationsSqlTests3_testSql() (*asset, error) {
	bytes, err := migrationsSqlTests3_testSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "migrations/sql/tests/3_test.sql", size: 1242, mode: os.FileMode(438), modTime: time.Unix(1541453045, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _migrationsSqlTests4_testSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xe4\x94\x4f\x4b\x03\x31\x10\xc5\xcf\xcd\xa7\x98\x5b\x76\x31\x7b\x91\xde\x3c\x09\xf6\x50\x90\x2d\xd8\x56\x8f\x61\x4c\xa6\xbb\x51\x9b\xac\x99\x44\x11\xf1\xbb\x4b\x37\xeb\x9f\x93\x1f\xc0\x3d\x04\x66\xde\x0b\xfc\x78\x2f\x90\xa6\x81\xb3\xa3\xeb\x22\x26\x82\xfd\x20\xd6\xed\x76\x75\xb3\x83\x75\xbb\xdb\x88\x45\xff\x66\x23\xea\x80\x39\xf5\xe7\x1a\x8d\x21\x66\xa8\xd8\x75\x1e\x53\x8e\xa4\x20\xd2\x73\x26\x4e\xda\xd9\xef\x99\xac\xc6\xa4\xc0\x3c\x39\xf2\xc5\x60\x13\x06\x52\xd0\x45\xf4\x27\x77\x5a\x0f\x21\x1e\xb5\xc5\x84\x0a\x98\x98\x5d\xf0\x5f\x5b\xbe\x7f\x20\x93\x14\xa0\x49\xee\x85\x6a\x71\x7b\x79\xbd\x5f\x6d\xc5\xa2\x92\xcb\x86\x5d\x27\x15\xc8\x65\x33\xd1\xa4\x82\x76\x73\x57\xd5\xa3\x56\x98\xc5\x1f\x29\x65\x9c\xc0\x3f\xd2\xe9\xbc\x7f\x4c\xf7\x0a\x4d\x2a\x48\x31\x53\x7d\x21\xfe\x28\x20\xd2\x21\x12\xf7\x33\x6e\xc0\x04\x4b\x33\x8e\x1f\x9c\x35\x33\x8e\x3f\x3c\x9a\xff\xff\xfa\xbf\xbf\xc3\xab\xf0\xea\xc5\x67\x00\x00\x00\xff\xff\xb0\xf6\x29\xba\x20\x05\x00\x00")

func migrationsSqlTests4_testSqlBytes() ([]byte, error) {
	return bindataRead(
		_migrationsSqlTests4_testSql,
		"migrations/sql/tests/4_test.sql",
	)
}

func migrationsSqlTests4_testSql() (*asset, error) {
	bytes, err := migrationsSqlTests4_testSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "migrations/sql/tests/4_test.sql", size: 1312, mode: os.FileMode(438), modTime: time.Unix(1541453045, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _migrationsSqlTests5_testSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xe4\x94\x4f\x4b\x03\x31\x10\xc5\xcf\xcd\xa7\x98\x5b\x76\x31\x7b\x11\x7a\xf2\x24\xd8\x43\x41\xb6\x60\x5b\x3d\x86\x31\x99\xee\x46\x6d\xb2\x66\x12\x45\xc4\xef\x2e\xdd\xac\x7f\x4e\x7e\x00\xf7\x10\x98\x79\x2f\xf0\xe3\xbd\x40\x9a\x06\xce\x8e\xae\x8b\x98\x08\xf6\x83\x58\xb7\xdb\xd5\xcd\x0e\xd6\xed\x6e\x23\x16\xfd\x9b\x8d\xa8\x03\xe6\xd4\x9f\x6b\x34\x86\x98\xa1\x62\xd7\x79\x4c\x39\x92\x82\x48\xcf\x99\x38\x69\x67\xbf\x67\xb2\x1a\x93\x02\xf3\xe4\xc8\x17\x83\x4d\x18\x48\x41\x17\xd1\x9f\xdc\x69\x3d\x84\x78\xd4\x16\x13\x2a\x60\x62\x76\xc1\x7f\x6d\xf9\xfe\x81\x4c\x52\x80\x26\xb9\x17\xaa\xc5\xed\xe5\xf5\x7e\xb5\x15\x8b\x4a\x2e\x1b\x76\x9d\x54\x20\x97\xcd\x44\x93\x0a\xda\xcd\x5d\x55\x8f\x5a\x61\x16\x7f\xa4\x94\x71\x02\xff\x48\xa7\xf3\xfe\x31\xdd\x2b\x34\xa9\x20\xc5\x4c\xf5\x85\xf8\xa3\x80\x48\x87\x48\xdc\xcf\xb8\x01\x13\x2c\xcd\x38\x7e\x70\xd6\xcc\x38\xfe\xf0\x68\xfe\xff\xeb\xff\xfe\x0e\xaf\xc2\xab\x17\x9f\x01\x00\x00\xff\xff\xbb\x53\x16\x3f\x20\x05\x00\x00")

func migrationsSqlTests5_testSqlBytes() ([]byte, error) {
	return bindataRead(
		_migrationsSqlTests5_testSql,
		"migrations/sql/tests/5_test.sql",
	)
}

func migrationsSqlTests5_testSql() (*asset, error) {
	bytes, err := migrationsSqlTests5_testSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "migrations/sql/tests/5_test.sql", size: 1312, mode: os.FileMode(438), modTime: time.Unix(1541453045, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _migrationsSqlTests6_testSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xe4\x94\x4f\x4b\x03\x31\x10\xc5\xcf\xcd\xa7\x98\x5b\x76\x31\x7b\xf1\xd0\x8b\x27\xc1\x1e\x0a\xb2\x05\xdb\xea\x31\x8c\xc9\x74\x37\x6a\x93\x35\x93\x28\x22\x7e\x77\xe9\x66\xfd\x73\xf2\x03\xb8\x87\xc0\xcc\x7b\x81\x1f\xef\x05\xd2\x34\x70\x76\x74\x5d\xc4\x44\xb0\x1f\xc4\xba\xdd\xae\x6e\x76\xb0\x6e\x77\x1b\xb1\xe8\xdf\x6c\x44\x1d\x30\xa7\xfe\x5c\xa3\x31\xc4\x0c\x15\xbb\xce\x63\xca\x91\x14\x44\x7a\xce\xc4\x49\x3b\xfb\x3d\x93\xd5\x98\x14\x98\x27\x47\xbe\x18\x6c\xc2\x40\x0a\xba\x88\xfe\xe4\x4e\xeb\x21\xc4\xa3\xb6\x98\x50\x01\x13\xb3\x0b\xfe\x6b\xcb\xf7\x0f\x64\x92\x02\x34\xc9\xbd\x50\x2d\x6e\x2f\xaf\xf7\xab\xad\x58\x54\x72\xd9\xb0\xeb\xa4\x02\xb9\x6c\x26\x9a\x54\xd0\x6e\xee\xaa\x7a\xd4\x0a\xb3\xf8\x23\xa5\x8c\x13\xf8\x47\x3a\x9d\xf7\x8f\xe9\x5e\xa1\x49\x05\x29\x66\xaa\x2f\xc4\x1f\x05\x44\x3a\x44\xe2\x7e\xc6\x0d\x98\x60\x69\xc6\xf1\x83\xb3\x66\xc6\xf1\x87\x47\xf3\xff\x5f\xff\xf7\x77\x78\x15\x5e\xbd\xf8\x0c\x00\x00\xff\xff\xe7\xba\x27\x6b\x20\x05\x00\x00")

func migrationsSqlTests6_testSqlBytes() ([]byte, error) {
	return bindataRead(
		_migrationsSqlTests6_testSql,
		"migrations/sql/tests/6_test.sql",
	)
}

func migrationsSqlTests6_testSql() (*asset, error) {
	bytes, err := migrationsSqlTests6_testSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "migrations/sql/tests/6_test.sql", size: 1312, mode: os.FileMode(438), modTime: time.Unix(1541453045, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _migrationsSqlTests7_testSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x95\x31\x4f\xc3\x30\x10\x85\xe7\xfa\x57\xdc\x96\x46\x38\x0b\x4b\x07\x26\x24\x3a\x54\x42\xa9\x44\x5b\x18\xa3\xc3\xbe\x26\x06\x6a\x07\x9f\x0d\x42\x88\xff\x8e\x52\xa7\x90\x2e\x88\x21\x63\x86\x48\x77\xef\xc5\x4f\x4f\x9f\x2c\xb9\x28\xe0\xe2\x60\x6a\x8f\x81\x60\xd7\x8a\x55\xb9\x59\xde\x6d\x61\x55\x6e\xd7\x62\xd6\x7c\x68\x8f\x95\xc3\x18\x9a\xcb\x0a\x95\x22\x66\x98\xb3\xa9\x2d\x86\xe8\x49\x82\xa7\xd7\x48\x1c\x2a\xa3\x7f\x66\xd2\x15\x06\x09\xea\xc5\x90\x4d\x06\x2b\xd7\x92\x84\xda\xa3\xed\xdc\x7e\xdd\x3b\x7f\xa8\x34\x06\x94\xc0\xc4\x6c\x9c\x3d\x6d\xf1\xf1\x89\x54\x90\x80\x2a\x98\x37\x3a\x0b\x8e\xda\x90\x55\x83\xb0\x93\x92\x8b\xfb\xeb\xdb\xdd\x72\x23\x66\xf3\x6c\x51\xb0\xa9\x33\x09\xd9\xa2\xe8\x8f\x66\x12\xca\xf5\xc3\x3c\x3f\x6a\xa9\x59\xf2\x8f\x5d\xd2\xd8\x27\xfe\x4a\xdd\xf7\xf9\xd5\xff\x97\x3a\x65\x12\x82\x8f\x34\x8c\x26\x5d\x60\xd4\xe7\x11\x9d\x90\x5f\x89\x3f\x50\x7a\xda\x7b\xe2\x66\x62\x39\x02\x4b\xe5\x34\x4d\x20\x47\x00\xe9\x8c\x56\x13\xc8\x11\x40\xb6\xcf\x6a\xba\x91\xff\x07\x39\x7c\x80\x6e\xdc\xbb\x15\xdf\x01\x00\x00\xff\xff\xfa\x55\x0f\x31\x92\x06\x00\x00")

func migrationsSqlTests7_testSqlBytes() ([]byte, error) {
	return bindataRead(
		_migrationsSqlTests7_testSql,
		"migrations/sql/tests/7_test.sql",
	)
}

func migrationsSqlTests7_testSql() (*asset, error) {
	bytes, err := migrationsSqlTests7_testSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "migrations/sql/tests/7_test.sql", size: 1682, mode: os.FileMode(438), modTime: time.Unix(1541453045, 0)}
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
	"migrations/sql/shared/1.sql":      migrationsSqlShared1Sql,
	"migrations/sql/shared/2.sql":      migrationsSqlShared2Sql,
	"migrations/sql/shared/3.sql":      migrationsSqlShared3Sql,
	"migrations/sql/shared/4.sql":      migrationsSqlShared4Sql,
	"migrations/sql/mysql/.gitkeep":    migrationsSqlMysqlGitkeep,
	"migrations/sql/mysql/5.sql":       migrationsSqlMysql5Sql,
	"migrations/sql/mysql/6.sql":       migrationsSqlMysql6Sql,
	"migrations/sql/mysql/7.sql":       migrationsSqlMysql7Sql,
	"migrations/sql/postgres/.gitkeep": migrationsSqlPostgresGitkeep,
	"migrations/sql/postgres/5.sql":    migrationsSqlPostgres5Sql,
	"migrations/sql/postgres/6.sql":    migrationsSqlPostgres6Sql,
	"migrations/sql/postgres/7.sql":    migrationsSqlPostgres7Sql,
	"migrations/sql/tests/.gitkeep":    migrationsSqlTestsGitkeep,
	"migrations/sql/tests/1_test.sql":  migrationsSqlTests1_testSql,
	"migrations/sql/tests/2_test.sql":  migrationsSqlTests2_testSql,
	"migrations/sql/tests/3_test.sql":  migrationsSqlTests3_testSql,
	"migrations/sql/tests/4_test.sql":  migrationsSqlTests4_testSql,
	"migrations/sql/tests/5_test.sql":  migrationsSqlTests5_testSql,
	"migrations/sql/tests/6_test.sql":  migrationsSqlTests6_testSql,
	"migrations/sql/tests/7_test.sql":  migrationsSqlTests7_testSql,
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
	"migrations": &bintree{nil, map[string]*bintree{
		"sql": &bintree{nil, map[string]*bintree{
			"mysql": &bintree{nil, map[string]*bintree{
				".gitkeep": &bintree{migrationsSqlMysqlGitkeep, map[string]*bintree{}},
				"5.sql":    &bintree{migrationsSqlMysql5Sql, map[string]*bintree{}},
				"6.sql":    &bintree{migrationsSqlMysql6Sql, map[string]*bintree{}},
				"7.sql":    &bintree{migrationsSqlMysql7Sql, map[string]*bintree{}},
			}},
			"postgres": &bintree{nil, map[string]*bintree{
				".gitkeep": &bintree{migrationsSqlPostgresGitkeep, map[string]*bintree{}},
				"5.sql":    &bintree{migrationsSqlPostgres5Sql, map[string]*bintree{}},
				"6.sql":    &bintree{migrationsSqlPostgres6Sql, map[string]*bintree{}},
				"7.sql":    &bintree{migrationsSqlPostgres7Sql, map[string]*bintree{}},
			}},
			"shared": &bintree{nil, map[string]*bintree{
				"1.sql": &bintree{migrationsSqlShared1Sql, map[string]*bintree{}},
				"2.sql": &bintree{migrationsSqlShared2Sql, map[string]*bintree{}},
				"3.sql": &bintree{migrationsSqlShared3Sql, map[string]*bintree{}},
				"4.sql": &bintree{migrationsSqlShared4Sql, map[string]*bintree{}},
			}},
			"tests": &bintree{nil, map[string]*bintree{
				".gitkeep":   &bintree{migrationsSqlTestsGitkeep, map[string]*bintree{}},
				"1_test.sql": &bintree{migrationsSqlTests1_testSql, map[string]*bintree{}},
				"2_test.sql": &bintree{migrationsSqlTests2_testSql, map[string]*bintree{}},
				"3_test.sql": &bintree{migrationsSqlTests3_testSql, map[string]*bintree{}},
				"4_test.sql": &bintree{migrationsSqlTests4_testSql, map[string]*bintree{}},
				"5_test.sql": &bintree{migrationsSqlTests5_testSql, map[string]*bintree{}},
				"6_test.sql": &bintree{migrationsSqlTests6_testSql, map[string]*bintree{}},
				"7_test.sql": &bintree{migrationsSqlTests7_testSql, map[string]*bintree{}},
			}},
		}},
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
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}