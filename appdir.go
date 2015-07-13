/* Copyright (c) 2015, Daniel Martí <mvdan@mvdan.cc> */
/* See LICENSE for licensing information */

package appdir

import (
	"errors"
	"os"
	"os/user"
	"path/filepath"
)

type dirSet struct {
	dirVar, dirDef   string
	dirsVar, dirsDef string
}

var (
	cache = dirSet{
		dirVar:  "XDG_CACHE_HOME",
		dirDef:  ".cache",
		dirsVar: "",
		dirsDef: "",
	}
	config = dirSet{
		dirVar:  "XDG_CONFIG_HOME",
		dirDef:  ".config",
		dirsVar: "XDG_CONFIG_DIRS",
		dirsDef: "/etc/xdg",
	}
	data = dirSet{
		dirVar:  "XDG_DATA_HOME",
		dirDef:  ".local/share",
		dirsVar: "XDG_DATA_DIRS",
		dirsDef: "/usr/local/share:/usr/share",
	}
)

func userHomeDir() (string, error) {
	curUser, err := user.Current()
	if err != nil {
		return "", errors.New("unable to determine your home dir")
	}
	return curUser.HomeDir, nil
}

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
	extra := os.Getenv(ds.dirsVar)
	if extra == "" {
		extra = ds.dirsDef
	}
	for _, path := range filepath.SplitList(extra) {
		if path == "" {
			continue
		}
		dirs = append(dirs, path)
	}
	return dirs, nil
}

func Cache() (string, error) {
	return cache.dir()
}

func CacheList() ([]string, error) {
	return cache.dirs()
}

func Config() (string, error) {
	return config.dir()
}

func ConfigList() ([]string, error) {
	return config.dirs()
}

func Data() (string, error) {
	return data.dir()
}

func DataList() ([]string, error) {
	return data.dirs()
}
