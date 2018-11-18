package models

type Domain struct {
    Country string  `json:"country"`
    Name    string  `json:"name"`
}

func (domain *Domain) LoadInfos() error {

    if domain.Name == "" {
        return nil
    }

    return nil
}
