package models

import (
    "strings"

    "github.com/PuKoren/convenient-api/dbs"
)

var (
    dbPublicDomains *dbs.PublicMailDB
)

type Email struct {
    String  string  `json:"string"`

    Domain  Domain  `json:"domain"`
}

func (email *Email) LoadInfos () error {

    if email.String == "" {
        return nil
    }

    domain := strings.Split(email.String, "@")[1]

    if dbPublicDomains.IsPublicProvider(domain) {
        return nil
    }

    email.Domain = Domain{ Name: domain }
    err := email.Domain.LoadInfos()

    if err != nil {
        return err
    }

    return nil
}

func InitEmail() error {
    dbPublicDomains = &dbs.PublicMailDB{}
    err := dbPublicDomains.Init()

    if err != nil {
        return err
    }

    return nil
}
