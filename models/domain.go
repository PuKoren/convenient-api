package models

import (
    "github.com/domainr/whois"
)

type Domain struct {
    Country string  `json:"country"`
    Name    string  `json:"name"`
    Whois   string  `json:"whois"`
}

func (domain *Domain) LoadInfos() error {

    if domain.Name == "" {
        return nil
    }

    request, err := whois.NewRequest(domain.Name)

    if err != nil {
        return err
    }

    response, err := whois.DefaultClient.Fetch(request)

    if err != nil {
        return err
    }

    domain.Whois = string(response.Body)

    return nil
}
