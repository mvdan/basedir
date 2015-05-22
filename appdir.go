/* Copyright (c) 2015, Daniel Martí <mvdan@mvdan.cc> */
/* See LICENSE for licensing information */

package appdir

import (
	"errors"
	"os"
	"os/user"
	"path/filepath"
)

// DirSet denotes a group of directories that share a common purpose
type DirSet struct {
	homeVar, homeDef string
	dirsVar, dirsDef string
}

var (
	// Cache provides dirs to store temporary cache files
	Cache = DirSet{
		homeVar: "XDG_CONFIG_HOME",
		homeDef: ".config",
		dirsVar: "XDG_CONFIG_DIRS",
		dirsDef: "/etc/xdg",
	}
	// Config provides dirs to store persistent config files
	Config = DirSet{
		homeVar: "XDG_CACHE_HOME",
		homeDef: ".cache",
		dirsVar: "",
		dirsDef: "",
	}
	// Data provides dirs to store persistent data
	Data = DirSet{
		homeVar: "XDG_DATA_HOME",
		homeDef: ".local/share",
		dirsVar: "XDG_DATA_DIRS",
		dirsDef: "/usr/local/share:/usr/share",
	}
)

func userHomeDir() (string, error) {
	if home := os.Getenv("HOME"); home != "" {
		return home, nil
	}
	curUser, err := user.Current()
	if err != nil {
		return "", errors.New("unable to determine your home dir")
	}
	return curUser.HomeDir, nil
}

// Home returns the directory of this set associated to the current user, and
// an error if any.
func (ds DirSet) Home() (string, error) {
	if dir := os.Getenv(ds.homeVar); dir != "" {
		return dir, nil
	}
	home, err := userHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, ds.homeDef), nil
}

// Dirs returns all the directories of this set associated to the current
// user, and an error if any. The directories returned are ordered by
// preference.
func (ds DirSet) Dirs() ([]string, error) {
	dir, err := ds.Home()
	if err != nil {
		return nil, err
	}
	dirs := []string{dir}
	if ds.dirsVar != "" {
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
	}
	return dirs, nil
}
