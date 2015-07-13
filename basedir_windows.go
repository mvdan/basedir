// Copyright (c) 2015, Daniel Martí <mvdan@mvdan.cc>
// See LICENSE for licensing information

package basedir

import (
	"os"
)

var (
	cacheDir  = firstGetenv("TEMP", "TMP")
	configDir = os.Getenv("APPDATA")
	dataDir   = os.Getenv("APPDATA")
)

func cache() (string, error) {
	return cacheDir, nil
}

func cacheList() ([]string, error) {
	return []string{cacheDir}, nil
}

func config() (string, error) {
	return configDir, nil
}

func configList() ([]string, error) {
	return []string{configDir}, nil
}

func data() (string, error) {
	return configDir, nil
}

func dataList() ([]string, error) {
	return []string{configDir}, nil
}
