package dbs

import (
    "runtime"
    "path/filepath"
)

type DB interface {
    Init() error
    Close() error
}


var (
    _, b, _, _  = runtime.Caller(0)
    basepath    = filepath.Dir(b)
)