package dbs

type DB interface {
    Init() error
    Close() error
}
