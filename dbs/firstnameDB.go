package dbs

import (
    "os"
    "log"
    "bufio"
    "strings"
    "strconv"
)

type FirstnameDB interface {
    Init() error
    Close() error

    GetNameBirthyear(name string) int
    GetNameSex(name string) string
    GetName(name string) YearAndSex
}

type YearAndSex struct {
    Year int
    Sex string
    Count int
}

type FirstnameDB_FR struct {
    names map[string]*YearAndSex
}

func (db *FirstnameDB_FR) Init() error {
    db.names = make(map[string]*YearAndSex)

    file, err := os.Open("dbs/data/nat2017.txt")
    if err != nil {
        return err
    }

    defer file.Close()

    scanner := bufio.NewScanner(file)
    // skip first line
    scanner.Scan()

    for scanner.Scan() {
        exp := strings.Split(scanner.Text(), "\t")

        if len(exp) < 3 {
            log.Println(exp)
            continue
        }

        year, _ := strconv.Atoi(exp[2])
        count, _ := strconv.Atoi(exp[3])
        sex := exp[0]
        name := strings.ToLower(exp[1])

        if db.names[name] == nil {
            db.names[name] = &YearAndSex { year, sex, count }
        }

        if db.names[name].Count < count {
            db.names[name].Year = year
            db.names[name].Count = count
        }
    }

    log.Println("Firstname DB Loaded [FR].")

    return scanner.Err()
}

func (db *FirstnameDB_FR) GetNameBirthyear(name string) int {
    yands := db.names[strings.ToLower(name)]

    if yands == nil {
        return 0
    }

    return yands.Year
}

func (db *FirstnameDB_FR) GetNameSex(name string) string {
    yands := db.names[strings.ToLower(name)]

    if yands == nil {
        return ""
    }

    if yands.Sex == "1" {
        return "M"
    }

    if yands.Sex == "2" {
        return "F"
    }

    return ""
}

func (db *FirstnameDB_FR) GetName(name string) YearAndSex {
    yands := db.names[strings.ToLower(name)]
    if yands != nil {
        return *yands
    }

    return YearAndSex {}
}

func (db *FirstnameDB_FR) Close() error {
    return nil
}
