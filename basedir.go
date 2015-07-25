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

// Cache returns the base cache directory and an error, if any.
func Cache() (string, error) {
	return clean(cache())
}

// Config returns the base config directory and an error, if any.
func Config() (string, error) {
	return clean(config())
}

// Data returns the base data directory and an error, if any.
func Data() (string, error) {
	return clean(data())
}
