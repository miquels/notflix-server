//go:build (!freebsd && !linux)

package main

import (
        "syscall"
        "time"
)

func (fi *FileInfo) setCreatetime() () {
	fi.createtime = fi.modtime
}
