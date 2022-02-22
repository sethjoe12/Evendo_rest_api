package models

import (

	"time"
	"github.com/dgrijalva/jwt-go"
	
  )


type Order struct{
    OrderId          int    `json:"id"`
	ProductId  int `json:"productid"`
	Quantity int    `json:"quantity"`
	Created_At      time.Time          `json:"created_at"`
	Updated_At      time.Time          `json:"updtaed_at"`

}
type Product struct {
	ProductId  int `bson:"_productid"`
	ProductName string          `json:"productname"`
	ProductPrice int           `json:"productprice"`
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


