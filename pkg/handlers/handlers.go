package handlers

import (
    "encoding/json"
    //"io/ioutil"
    //"log"
    //"math/rand"
    "net/http"
    "errors"
    "github.com/gorilla/mux"
   // "sethjoe/pkg/mocks"
    "sethjoe/pkg/models"
	//"sethjoe/pkg/response"
	//"sethjoe/pkg/handlers/token"
   // "strconv"
)

var (
    users []models.User
)


func Login(w http.ResponseWriter, r *http.Request){
    vars := mux.Vars(r)
    username := vars["username"]
    password := vars["password"]

    user,err := GetUser(username)
    if err != nil {
        w.Header().Add("Content-Type", "application/json")
        w.WriteHeader(400)

        json.NewEncoder(w).Encode("Invalid Credentials")
        return
    }
    if user.Password == password {
        w.Header().Add("Content-Type", "application/json")
        w.WriteHeader(200)

        json.NewEncoder(w).Encode("Success Login")
        return
    }
}

func UserSignUP(w http.ResponseWriter, r *http.Request){
    vars := mux.Vars(r)
    username := vars["username"]
    password := vars["password"]
    role := vars["role"]
    _,err := GetUser(username)
    if err != nil {
        if role != "buyer"{
            role = "seller"
        }
        id := len(users) + 1
        newUser := models.User{
            Id:id,
            Username:username,
            Password:password,
            Role:role,
        }
        users := append(users,newUser)
        w.Header().Add("Content-Type", "application/json")
        w.WriteHeader(200)

        json.NewEncoder(w).Encode(users)
        return
    }
    w.Header().Add("Content-Type", "application/json")
    w.WriteHeader(400)

    json.NewEncoder(w).Encode("Error")
    return
    
}

func GetUser(username string)(models.User,error){
    for _,u := range users {
        if u.Username == username {
            return u,nil
        }
    }
    return models.User{}, errors.New("No user found")
}