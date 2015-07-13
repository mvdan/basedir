/* Copyright (c) 2015, Daniel Martí <mvdan@mvdan.cc> */
/* See LICENSE for licensing information */

// +build darwin dragonfly freebsd linux nacl netbsd openbsd solaris

package basedir

import (
	"os"
	"path/filepath"
)

type dirSet struct {
	dirVal, dirDef   string
	dirsVal, dirsDef string
}

var (
	cacheSet = dirSet{
		dirVal:  os.Getenv("XDG_CACHE_HOME"),
		dirDef:  ".cache",
		dirsVal: "",
		dirsDef: "",
	}
	configSet = dirSet{
		dirVal:  os.Getenv("XDG_CONFIG_HOME"),
		dirDef:  ".config",
		dirsVal: os.Getenv("XDG_CONFIG_DIRS"),
		dirsDef: "/etc/xdg",
	}
	dataSet = dirSet{
		dirVal:  os.Getenv("XDG_DATA_HOME"),
		dirDef:  ".local/share",
		dirsVal: os.Getenv("XDG_DATA_DIRS"),
		dirsDef: "/usr/local/share:/usr/share",
	}
)

func (ds dirSet) dir() (string, error) {
	if ds.dirVal != "" {
		return ds.dirVal, nil
	}
	home, err := userHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, ds.dirDef), nil
}

func (ds dirSet) dirs() ([]string, error) {
	dir, err := ds.dir()
	if err != nil {
		return nil, err
	}
	dirs := []string{dir}
	if ds.dirsVal == "" {
		return dirs, nil
	}
	extra := ds.dirsDef
	if ds.dirsVal != "" {
		extra = ds.dirsVal
	}
	dirs = append(dirs, filepath.SplitList(extra)...)
	return dirs, nil
}

func cache() (string, error) {
	return cacheSet.dir()
}

func cacheList() ([]string, error) {
	return cacheSet.dirs()
}

func config() (string, error) {
	return configSet.dir()
}

func configList() ([]string, error) {
	return configSet.dirs()
}

func data() (string, error) {
	return dataSet.dir()
}

func dataList() ([]string, error) {
	return dataSet.dirs()
}
