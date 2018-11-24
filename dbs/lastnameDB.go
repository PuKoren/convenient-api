package dbs

import (
    "os"
    "log"
    "fmt"
    "bufio"
    "strings"
    "strconv"
)

type LastnameDB interface {
    Init() error
    Close() error

    Exists(name string) bool
    GetCount(name string) int
}

type LastnameDB_FR struct {
    names map[string]int
}

func (db *LastnameDB_FR) Init() error {
    db.names = make(map[string]int)

    file, err := os.Open(fmt.Sprintf("%s/data/noms2008nat_txt.txt", basepath))
    if err != nil {
        return err
    }

    defer file.Close()

    scanner := bufio.NewScanner(file)
    // skip first line
    scanner.Scan()

    for scanner.Scan() {
        exp := strings.Split(scanner.Text(), "\t")

        if len(exp) < 11 {
            log.Println(exp)
            continue
        }

        name := exp[0]

        count := 0
        for i := 1; i < 11; i++ {
            if v, err := strconv.Atoi(exp[i]); err != nil {
                count += v
            }
        }

        db.names[strings.ToLower(name)] = count
    }

    log.Println("Lastname DB Loaded [FR].")

    return scanner.Err()
}

func (db *LastnameDB_FR) Exists(name string) bool {
    if _, ok := db.names[strings.ToLower(name)]; ok {
        return true
    } else {
        return false
    }
}

func (db *LastnameDB_FR) GetCount(name string) int {
    if val, ok := db.names[strings.ToLower(name)]; ok {
        return val
    } else {
        return 0
    }
}

func (db *LastnameDB_FR) Close() error {
    return nil
}
