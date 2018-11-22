package dbs

import (
    "log"
    "net"
    "fmt"

    "github.com/oschwald/geoip2-golang"
)

type IpDB struct {
    geoDb *geoip2.Reader
}

func (db *IpDB) Init() error {
    if db.geoDb == nil {
        var err error

        db.geoDb, err = geoip2.Open(fmt.Sprintf("%s/data/GeoLite2-Country.mmdb", basepath))

        if err != nil {
            log.Println(err)
            return err
        }

        log.Println("Geo DB Ready.")
    }

    return nil
}

func (db *IpDB) GetCountryIso(ip string) (string, error) {
    parsedIp := net.ParseIP(ip)
    record, err := db.geoDb.Country(parsedIp)

    if err != nil {
        log.Println(err)
        return "", err
    }

    return record.Country.IsoCode, nil
}

func (db *IpDB) Close() error {
    defer db.geoDb.Close()

    return nil
}
