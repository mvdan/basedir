/* Copyright (c) 2015, Daniel Martí <mvdan@mvdan.cc> */
/* See LICENSE for licensing information */

package basedir

import (
	"errors"
	"os/user"
)

func userHomeDir() (string, error) {
	curUser, err := user.Current()
	if err != nil {
		return "", errors.New("unable to determine your home dir")
	}
	return curUser.HomeDir, nil
}

// Cache returns the base cache directory and an error, if any.
func Cache() (string, error) {
	return cache()
}

// CacheList returns the base cache directory list and an error, if any. The
// list is ordered by priority and the first element, if present, is the same
// directory returned by Cache.
func CacheList() ([]string, error) {
	return cacheList()
}

// Config returns the base config directory and an error, if any.
func Config() (string, error) {
	return config()
}

// ConfigList returns the base config directory list and an error, if any. The
// list is ordered by priority and the first element, if present, is the same
// directory returned by Config.
func ConfigList() ([]string, error) {
	return configList()
}

// Data returns the base data directory and an error, if any.
func Data() (string, error) {
	return data()
}

// DataList returns the base data directory list and an error, if any. The
// list is ordered by priority and the first element, if present, is the same
// directory returned by Data.
func DataList() ([]string, error) {
	return dataList()
}
