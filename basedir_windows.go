// Copyright (c) 2015, Daniel Martí <mvdan@mvdan.cc>
// See LICENSE for licensing information

package basedir

import (
	"os"
)

var (
	cacheDir = firstGetenv("TEMP", "TMP")
	dataDir  = os.Getenv("APPDATA")
)

func cache() (string, error) {
	return cacheDir, nil
}

func data() (string, error) {
	return dataDir, nil
}
