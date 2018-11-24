package models

import (
    "log"

    "github.com/PuKoren/convenient-api/dbs"
)

var (
    ipDB *dbs.IpDB
    dbPublicDomains *dbs.PublicMailDB
    firstnameDBs map[string]dbs.FirstnameDB
    lastnameDBs map[string]dbs.LastnameDB
)

func Init() error {
    firstnameDBs = make(map[string]dbs.FirstnameDB)
    lastnameDBs = make(map[string]dbs.LastnameDB)

    firstnameDBs["FR"] = &dbs.FirstnameDB_FR{}
    lastnameDBs["FR"] = &dbs.LastnameDB_FR{}

    ipDB = &dbs.IpDB{}
    ipDB.Init()

    for _, v := range firstnameDBs {
        err := v.Init()
        if err != nil {
            log.Fatal(err)
        }
    }

    for _, v := range lastnameDBs {
        err := v.Init()
        if err != nil {
            log.Fatal(err)
        }
    }

    dbPublicDomains = &dbs.PublicMailDB{}
    err := dbPublicDomains.Init()

    if err != nil {
        return err
    }

    return nil
}

func Close() {
    ipDB.Close()
}