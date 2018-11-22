package dbs

import (
    "os"
    "log"
    "fmt"
    "bufio"
    "strings"
)

type PublicMailDB struct {
    domains map[string]*struct{}
}

func (db *PublicMailDB) Init() error {
    db.domains = make(map[string]*struct{})

    file, err := os.Open(fmt.Sprintf("%s/data/publicmails.txt", basepath))
    if err != nil {
        return err
    }

    defer file.Close()

    scanner := bufio.NewScanner(file)
    
    for scanner.Scan() {
        db.domains[scanner.Text()] = &struct{}{}
    }

    log.Println("Public mails DB Loaded.")

    return scanner.Err()
}

func (db *PublicMailDB) IsPublicProvider(domain string) bool {
    if db.domains[strings.ToLower(domain)] != nil {
        return true
    }

    return false
}

func (db *PublicMailDB) Close() error {
    return nil
}
