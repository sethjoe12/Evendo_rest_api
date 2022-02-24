package mocks

import "sethjoe/pkg/models"
var Username = []models.User{
    {
        Username:"jojo",
        Password:"gwapo",
        Role:"buyer",
        Gender:"male",
        Email:"sethjoe@yahoo.com",
        Status: "single",
        ContactNumber:12223343,
        
    },

}

var Product = []models.Product{
    {
        ProductId:     1,
        ProductName:  "shabu",
        ProductPrice: 1233334,
        Quantity: 12,
        
    },
	{
		ProductId:     3,
        ProductName:  "loso ne mako",
        ProductPrice: 1233334,
        Quantity: 12,
	},

}

var Order = []models.Order{
    {
        OrderId:   1,
        ProductId:     1,
        Quantity:  12,
        
    },

}