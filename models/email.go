package models

import (
    "strings"
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

    email.Domain = Domain{ Name: domain }
    err := email.Domain.LoadInfos()

    if err != nil {
        return err
    }

    return nil
}
