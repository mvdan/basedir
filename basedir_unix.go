// Copyright (c) 2015, Daniel Martí <mvdan@mvdan.cc>
// See LICENSE for licensing information

// +build darwin dragonfly freebsd linux netbsd openbsd

package basedir

type envDir struct {
	val, def string
}

var (
	cacheDir = firstGetenv(".cache", "XDG_CACHE_HOME")
	dataDir  = firstGetenv(".config", "XDG_CONFIG_HOME")
)

func cache() (string, error) {
	return cacheDir, nil
}

func data() (string, error) {
	return dataDir, nil
}
