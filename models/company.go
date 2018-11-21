package models

type Company struct {
    Name            string  `json:"name"`
    Organization    string  `json:"organization"`
    Phone           string  `json:"phone"`
    Country         string  `json:"country"`
}
