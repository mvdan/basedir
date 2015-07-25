// Copyright (c) 2015, Daniel Martí <mvdan@mvdan.cc>
// See LICENSE for licensing information

// +build darwin dragonfly freebsd linux netbsd openbsd

package basedir

import (
	"os"
	"path/filepath"
)

type envDir struct {
	val, def string
}

var (
	cacheSet = envDir{
		val: os.Getenv("XDG_CACHE_HOME"),
		def: ".cache",
	}
	configSet = envDir{
		val: os.Getenv("XDG_CONFIG_HOME"),
		def: ".config",
	}
	dataSet = envDir{
		val: os.Getenv("XDG_DATA_HOME"),
		def: ".local/share",
	}
)

func (ds envDir) dir() (string, error) {
	if ds.val != "" {
		return ds.val, nil
	}
	home, err := homeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, ds.def), nil
}

func cache() (string, error) {
	return cacheSet.dir()
}

func config() (string, error) {
	return configSet.dir()
}

func data() (string, error) {
	return dataSet.dir()
}
