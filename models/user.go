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
    Company     Company `json:"company"`
}

var (
    ipDB *dbs.IpDB
    firstnameDBs map[string]dbs.FirstnameDB
)

func (user *User) LoadInfos() error {

    err := user.Email.LoadInfos()

    if (err != nil) {
        log.Println(err)
    }

    if (user.Email.Domain.Company.Country != "") {
        user.Company = user.Email.Domain.Company
        if (user.Country == "") {
            user.Country = user.Company.Country
        }
    }

    if user.Country == "" && user.Ip != "" {
        var err error
        user.Country, err = ipDB.GetCountryIso(user.Ip)
        if err != nil {
            return err
        }
    }

    if user.Country != "" {
        if firstnameDBs[user.Country] != nil {
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
    }

    return nil
}

func (user *User) GetFirstnameFromEmail() string {
    var retainedName string

    if user.Country != "" {
        var probName []rune
        probName = make([]rune, len(user.Firstname))

        var retainedSize int = 0

        for _, char := range user.Email.GetUserPart() {
            if (char == '@') {
                break
            }
            probName = append(probName, char)

            yearAndSex := firstnameDBs[user.Country].GetName(string(probName))
            if yearAndSex.Count > retainedSize {
                retainedSize = yearAndSex.Count
                retainedName = string(probName)
            }
        }
    }

    return retainedName
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

func CloseUser() {
    ipDB.Close()
}
