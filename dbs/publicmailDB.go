package dbs

import (
    "os"
    "bufio"
    "strings"
)

type PublicMailDB struct {
    domains map[string]*struct{}
}

func (db *PublicMailDB) Init() error {
    file, err := os.Open("dbs/publicmails.txt")
    if err != nil {
        return err
    }

    defer file.Close()

    scanner := bufio.NewScanner(file)
    
    for scanner.Scan() {
        db.domains[scanner.Text()] = &struct{}{}
    }

    return nil
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