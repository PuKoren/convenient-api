package models

type Email struct {
    String  string  `json:"string"`

    Domain  Domain  `json:"domain"`
}

func (email *Email) LoadInfos () error {

    if email.String == "" {
        return nil
    }

    email.Domain = Domain{}
    email.Domain.LoadInfos()

    return nil
}
