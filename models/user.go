package models

type User struct {
    Firstname   string  `json:"firstname"`
    Lastname    string  `json:"lastname"`
    Country     string  `json:"country"`
    Birthyear   int     `json:"birthyear"`
    Sex         string  `json:"sex"`

    Email       Email   `json:"email"`
}

func (user *User) LoadInfos() error {

    user.Email = Email{}
    user.Email.LoadInfos()

    return nil
}
