// Copyright (c) 2015, Daniel Martí <mvdan@mvdan.cc>
// See LICENSE for licensing information

package basedir

import (
	"errors"
	"os"
	"os/user"
	"path/filepath"
)

// Cache returns the base cache directory and an error, if any.
func Cache() (string, error) {
	return clean(cache())
}

// Data returns the base data directory and an error, if any.
func Data() (string, error) {
	return clean(data())
}

func firstGetenv(def string, evs ...string) string {
	for _, ev := range evs {
		if v := os.Getenv(ev); v != "" {
			return v
		}
	}
	home, err := homeDir()
	if err != nil {
		return ""
	}
	return filepath.Join(home, def)
}

func homeDir() (string, error) {
	curUser, err := user.Current()
	if err != nil {
		return "", errors.New("unable to determine your home dir")
	}
	return curUser.HomeDir, nil
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
