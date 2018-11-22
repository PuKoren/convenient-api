package models

import (
    "testing"

    "github.com/stretchr/testify/assert"
)

func Test_Main(t *testing.T) {
    InitUser()
}

func TestUserModel_ExtractFirstNameFromEmail(t *testing.T) {
    var user = User{
        Country: "FR",
        Email: Email{ String: "jonathan@toto.com" } }

    result := user.GetFirstnameFromEmail().Firstname

    assert.Equal(t, "jonathan", result, "Extracted firstname should be jonathan")
}


func TestUserModel_ExtractFirstNameFromEmailWithName(t *testing.T) {
    var user = User{
        Country: "FR",
        Email: Email{ String: "jonathanmuller@toto.com" } }

    result := user.GetFirstnameFromEmail().Firstname

    assert.Equal(t, "jonathan", result, "Extracted firstname should be jonathan")
}

func TestUserModel_ExtractFirstNameFromEmailWithNameReversed(t *testing.T) {
    var user = User{
        Country: "FR",
        Email: Email{ String: "mullerjonathan@toto.com" } }

    result := user.GetFirstnameFromEmail().Firstname

    assert.Equal(t, "jonathan", result, "Extracted firstname should be jonathan")
}

func TestUserModel_CannotExtractFirstNameFromShortEmailWithName(t *testing.T) {
    var user = User{
        Country: "FR",
        Email: Email{ String: "jbalanciaga@toto.com" } }

    result := user.GetFirstnameFromEmail().Firstname

    assert.Equal(t, "", result, "Should fail to extract firstname")
}
