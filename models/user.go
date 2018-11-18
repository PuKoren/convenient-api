package models

import (
    "log"

    "github.com/PuKoren/convenient-api/dbs"
)

type User struct {
    Firstname   string  `json:"firstname"`
    Lastname    string  `json:"lastname"`
    Country     string  `json:"country"`
    Birthyear   int     `json:"birthyear"`
    Sex         string  `json:"sex"`
    Ip          string  `json:"ip"`

    Email       Email   `json:"email"`
}

var (
    ipDB *dbs.IpDB
    firstnameDBs map[string]dbs.FirstnameDB
)

func (user *User) LoadInfos() error {

    user.Email.LoadInfos()

    if user.Country == "" && user.Ip != "" {
        var err error
        user.Country, err = ipDB.GetCountryIso(user.Ip)
        if err != nil {
            return err
        }
    }

    if user.Country != "" {
        if user.Firstname == "" && user.Email.String != "" {
            user.Firstname = user.GetFirstnameFromEmail()
        }

        if user.Firstname != "" {
            if  user.Birthyear == 0 {
                user.Birthyear = firstnameDBs[user.Country].GetNameBirthyear(user.Firstname)
            }

            if user.Sex == "" {
                user.Sex = firstnameDBs[user.Country].GetNameSex(user.Firstname)
            }
        }
    }

    return nil
}

func (user *User) GetFirstnameFromEmail() string {
    if user.Country != "" {

    }

    return ""
}

func InitUser() error {
    firstnameDBs = make(map[string]dbs.FirstnameDB)
    firstnameDBs["FR"] = &dbs.FirstnameDB_FR{}

    ipDB = &dbs.IpDB{}
    ipDB.Init()

    for _, v := range firstnameDBs {
        err := v.Init()
        if err != nil {
            log.Fatal(err)
        }
    }

    log.Println("User model loaded.")

    return nil
}