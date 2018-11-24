package models

import (
    "log"
    "strings"
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

    if user.Email.String != "" {
        userNames := user.GetFirstnameFromEmail()
        if user.Firstname == "" {
            user.Firstname = userNames.Firstname
        }
        if user.Lastname == "" {
            user.Lastname = userNames.Lastname
        }
    }

    if user.Country != "" && firstnameDBs[user.Country] != nil {
        if user.Firstname != "" {
            if  user.Birthyear == 0 {
                user.Birthyear = firstnameDBs[user.Country].GetNameBirthyear(user.Firstname)
            }

            if user.Sex == "" {
                user.Sex = firstnameDBs[user.Country].GetNameSex(user.Firstname)
            }
        }
    }

    user.Firstname = strings.Title(user.Firstname)
    user.Lastname = strings.Title(user.Lastname)

    return nil
}

type UserNames struct {
    Firstname   string
    Lastname    string
}

func (user *User) GetFirstnameFromEmail() UserNames {
    if firstnameDBs[user.Country] == nil {
        return UserNames{}
    }

    var retainedName string
    var retainedLastName string

    if user.Country != "" {
        var retainedSize int = 0

        userPart := user.Email.GetUserPart()

        probName := make([]rune, len(user.Firstname))
        probNameR := make([]rune, len(user.Firstname))
        for i, char := range userPart {
            probName    = append(probName, char)
            probNameR   = append([]rune{rune(userPart[len(userPart) -1 -i])}, probNameR...)

            yearAndSex := firstnameDBs[user.Country].GetName(string(probName))
            yearAndSexR := firstnameDBs[user.Country].GetName(string(probNameR))

            if yearAndSex.Count > retainedSize {
                retainedSize = yearAndSex.Count
                retainedName = string(probName)
            }

            if yearAndSexR.Count > retainedSize {
                retainedSize = yearAndSex.Count
                retainedName = string(probNameR)
            }
        }

        if lastnameDBs[user.Country] != nil {
            splitedPart := strings.Split(userPart, ".")
            if len(splitedPart) > 1 {
                if splitedPart[0] == retainedName {
                    retainedLastName = splitedPart[1]
                }
                if splitedPart[1] == retainedName {
                    retainedLastName = splitedPart[0]
                }

                if !lastnameDBs[user.Country].Exists(retainedLastName) {
                    retainedLastName = ""
                }
            }

            if retainedLastName == "" {
                retainedLastName = strings.Replace(userPart, retainedName, "", -1)
                if !lastnameDBs[user.Country].Exists(retainedLastName) {
                    retainedLastName = ""
                }
            }

            if lastnameDBs[user.Country].GetCount(retainedLastName) < lastnameDBs[user.Country].GetCount(retainedName) {
                log.Println("#### Swapping")
                swap := retainedName
                retainedName = retainedLastName
                retainedLastName = swap
            }
        }
    }

    if len(retainedName) < 2 {
        retainedName = ""
    }

    return UserNames{ Firstname: retainedName, Lastname: retainedLastName }
}
