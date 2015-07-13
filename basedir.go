// Copyright (c) 2015, Daniel Martí <mvdan@mvdan.cc>
// See LICENSE for licensing information

package basedir

import (
	"errors"
	"os"
	"os/user"
)

func homeDir() (string, error) {
	curUser, err := user.Current()
	if err != nil {
		return "", errors.New("unable to determine your home dir")
	}
	return curUser.HomeDir, nil
}

func firstGetenv(evs ...string) string {
	for _, ev := range evs {
		if v := os.Getenv(ev); v != "" {
			return v
		}
	}
	return ""
}

func clean(dir string, err error) (string, error) {
	if err != nil {
		return "", err
	}
	if dir == "" {
		return "", errors.New("dir not found")
	}
	return dir, nil
}

func cleanList(dirs []string, err error) ([]string, error) {
	if err != nil {
		return nil, err
	}
	var clean []string
	for _, d := range dirs {
		if d == "" {
			continue
		}
		clean = append(clean, d)
	}
	if clean == nil {
		return nil, errors.New("dirs not found")
	}
	return clean, nil
}

// Cache returns the base cache directory and an error, if any.
func Cache() (string, error) {
	return clean(cache())
}

// CacheList returns the base cache directory list and an error, if any. The
// list is ordered by priority and the first element, if present, is the same
// directory returned by Cache.
func CacheList() ([]string, error) {
	return cleanList(cacheList())
}

// Config returns the base config directory and an error, if any.
func Config() (string, error) {
	return clean(config())
}

// ConfigList returns the base config directory list and an error, if any. The
// list is ordered by priority and the first element, if present, is the same
// directory returned by Config.
func ConfigList() ([]string, error) {
	return cleanList(configList())
}

// Data returns the base data directory and an error, if any.
func Data() (string, error) {
	return clean(data())
}

// DataList returns the base data directory list and an error, if any. The
// list is ordered by priority and the first element, if present, is the same
// directory returned by Data.
func DataList() ([]string, error) {
	return cleanList(dataList())
}
