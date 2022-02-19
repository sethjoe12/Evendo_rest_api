package models

import (

	"time"
	"github.com/dgrijalva/jwt-go"
  )


type Order struct{
    Id          int    `json:"id"`
	ProductName  string `json:"productName"`
	ProductPrice int    `json:"productPrice"`
	Quantity int    `json:"quantity"`
	Created_At      time.Time          `json:"created_at"`
	Updated_At      time.Time          `json:"updtaed_at"`
	User_ID         string             `json:"user_id"`

}
type Product struct {
	Id  int `bson:"_id"`
	ProductPrice int           `json:"productprice"`
	ProductName string          `json:"productname"`
	Quantity int                `json:"quantity"`
      
}
type User struct {
	Id int `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
    Role string `json:"role"`  
}

type Claims struct {
	UserID   string `json:"userId"`
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
	jwt.StandardClaims
}


