package dbs

import (
    "log"
    "net"

    "github.com/oschwald/geoip2-golang"
)

type IpDB struct {
    geoDb *geoip2.Reader
}

func (db *IpDB) Init() error {
    if db.geoDb == nil {
        var err error
        db.geoDb, err = geoip2.Open("dbs/data/GeoLite2-Country.mmdb")

        if err != nil {
            log.Println(err)
            return err
        }
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
