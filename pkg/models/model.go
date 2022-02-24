package models

import (

	"time"
	"github.com/dgrijalva/jwt-go"
	
  )


type Order struct{
    OrderId          int    `json:"id"`
	ProductId  int `json:"productid"`
	ProductName string          `json:"productname"`
	Quantity int    `json:"quantity"`
	Created_At      time.Time          `json:"ordered_at"`
	

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
    Gender string `json:"gender"`
    Email string `json:"email"`
    Status string `json:"Status"`
    Age int `json:"age"`
    ContactNumber int `json:"contactnumber"`
}

type Claims struct {
	UserID   string `json:"userId"`
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
	jwt.StandardClaims
}
type JwtToken struct {
    Token string `json:"token"`
}

type Exception struct {
    Message string `json:"message"`
}

