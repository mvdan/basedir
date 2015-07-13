/* Copyright (c) 2015, Daniel Martí <mvdan@mvdan.cc> */
/* See LICENSE for licensing information */

package appdir

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

func Cache() (string, error) {
	return cache()
}

func CacheList() ([]string, error) {
	return cacheList()
}

func Config() (string, error) {
	return config()
}

func ConfigList() ([]string, error) {
	return configList()
}

func Data() (string, error) {
	return data()
}

func DataList() ([]string, error) {
	return dataList()
}
