package version

import "time"

// Version represents the version of pouchd.
const Version = "0.1.0"

// BuildTime is the time when this binary of daemon is built
var BuildTime = time.Date(2018, 1, 19, 8, 0, 0, 0, time.Local).Format(time.RFC3339Nano)

// APIVersion means the api version daemon serves
var APIVersion = "1.24"

// GOVersion is the go version to build Pouch
var GOVersion = "go1.9.1"

// GitCommit is the commit id to build Pouch
var GitCommit = "6be2080cd9837e9b8a0039c2d21521bb00a30c84"
