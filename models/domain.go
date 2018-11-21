package models

import (
    "github.com/domainr/whois"
    "github.com/likexian/whois-parser-go"
)

type Domain struct {
    Country string  `json:"country"`
    Name    string  `json:"name"`
    Whois   string  `json:"-"`
    Company Company `json:"-"`
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

    parsedResponse, err := whois_parser.Parse(domain.Whois)

    if (err != nil) {
        return err
    }

    r := parsedResponse.Registrant
    domain.Company = Company { Name: r.Name, Organization: r.Organization, Phone: r.Phone, Country: r.Country  }
    domain.Country = r.Country

    return nil
}
