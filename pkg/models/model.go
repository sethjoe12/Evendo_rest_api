package models


type Product struct{
    Id          int    `json:"id"`
	ProductName  string `json:"productName"`
	ProductPrice int    `json:"productPrice"`
	Quantity int    `json:"quantity"`
}

type User struct {
	Id int `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Role string `json:"role"`
}