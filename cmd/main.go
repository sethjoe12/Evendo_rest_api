package main

import (
    "log"
    "net/http"

    "github.com/gorilla/mux"
    "sethjoe/pkg/handlers"
)

func main() {
    router := mux.NewRouter()
    //for user 
    router.HandleFunc("/register", handlers.UserSignUP).Methods(http.MethodGet)
    router.HandleFunc("/login", handlers.Login).Methods(http.MethodPut)


    ///for admin
    




    
    
    //for user/seller
    router.HandleFunc("/seller", handlers.ViewAllProduct).Methods(http.MethodGet)//wiew all products
    router.HandleFunc("/seller", handlers.AddProduct).Methods(http.MethodPost)//add products
    router.HandleFunc("/seller/id", handlers.DeleteProduct).Methods(http.MethodDelete)//delete products
    //for user/buyer
    router.HandleFunc("/Order", handlers.ViewProducts).Methods(http.MethodGet)//view prducts
    router.HandleFunc("/Order/{id}", handlers.OrderProduct).Methods(http.MethodGet)//order product//use methid[add/get]/use product id 
    router.HandleFunc("/Order/{id}", handlers.ViewOrder).Methods(http.MethodPut)//update order
    
    

        
    
    log.Println("API is running!")
    http.ListenAndServe(":8080", router)
}
