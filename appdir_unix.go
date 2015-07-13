/* Copyright (c) 2015, Daniel Martí <mvdan@mvdan.cc> */
/* See LICENSE for licensing information */

// +build darwin dragonfly freebsd linux nacl netbsd openbsd solaris

package appdir

import (
	"os"
	"path/filepath"
)

type dirSet struct {
	dirVar, dirDef   string
	dirsVar, dirsDef string
}

var (
	cacheSet = dirSet{
		dirVar:  "XDG_CACHE_HOME",
		dirDef:  ".cache",
		dirsVar: "",
		dirsDef: "",
	}
	configSet = dirSet{
		dirVar:  "XDG_CONFIG_HOME",
		dirDef:  ".config",
		dirsVar: "XDG_CONFIG_DIRS",
		dirsDef: "/etc/xdg",
	}
	dataSet = dirSet{
		dirVar:  "XDG_DATA_HOME",
		dirDef:  ".local/share",
		dirsVar: "XDG_DATA_DIRS",
		dirsDef: "/usr/local/share:/usr/share",
	}
)

func (ds dirSet) dir() (string, error) {
	if dir := os.Getenv(ds.dirVar); dir != "" {
		return dir, nil
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
	if ds.dirsVar == "" {
		return dirs, nil
	}
	extra := ds.dirsDef
	if v := os.Getenv(ds.dirsVar); v != "" {
		extra = v
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
