package models

func Init() {
    InitUser()
    InitEmail()
}

func Close() {
    CloseUser()
}