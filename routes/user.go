package routes

import (
    "net/http"
    "encoding/json"

    "github.com/gorilla/mux"

    "github.com/PuKoren/convenient-api/models"
)

func GetUser(w http.ResponseWriter, r *http.Request) {

    w.Header().Set("Content-Type", "application/json")

    user := models.User{}
    json.NewDecoder(r.Body).Decode(&user)

    if user.Ip == "" {
        user.Ip = r.Header.Get("X-Forwarded-For")
    }

    err := user.LoadInfos()

    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        json.NewEncoder(w).Encode(err)
        return
    }

    json.NewEncoder(w).Encode(user)
}

func RegisterHandlersUser(router *mux.Router) {

    router.HandleFunc("/v1/user/", GetUser).Methods("POST", "OPTIONS")
}
