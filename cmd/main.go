package main

import (
    "log"
    "net/http"
    "sethjoe/middleware"
    "github.com/gorilla/mux"
    "sethjoe/pkg/handlers"
    //"sethjoe/router"
   // "fmt"
)

func main() {
    router := mux.NewRouter()
   
    //for admin
    
    router.Handle("/admin/addproducts", basicauthmiddleware.BasicAuthMiddleware(http.HandlerFunc(handlers.AdminAddProduct))).Methods("GET")
	router.Handle("/admin/deleteproducts/{id}", basicauthmiddleware.BasicAuthMiddleware(http.HandlerFunc(handlers.AdminDeleteProduct))).Methods("DELETE")
    router.Handle("/admin/viewAllProducts", basicauthmiddleware.BasicAuthMiddleware(http.HandlerFunc(handlers.AdminViewProducts))).Methods("GET")
    router.Handle("/admin/viewOrderbyID/{id}", basicauthmiddleware.BasicAuthMiddleware(http.HandlerFunc(handlers.AdminViewOrder))).Methods("GET")

    
    router.HandleFunc("/register", handlers.UserSignUP).Methods(http.MethodGet)
    router.HandleFunc("/login", handlers.Login).Methods(http.MethodPut) 
    
    router.HandleFunc("/seller", handlers.ViewAllProduct).Methods(http.MethodGet)//wiew all products
    router.HandleFunc("/seller", handlers.AddProduct).Methods(http.MethodPost)//add products
    router.HandleFunc("/seller/id", handlers.DeleteProduct).Methods(http.MethodDelete)//delete products
    
    router.HandleFunc("/ViewALllProducts", handlers.ViewALllProducts).Methods(http.MethodGet)//view prducts
    router.HandleFunc("/Order", handlers.Order).Methods(http.MethodGet)//order product//use methid[add/get]/use product id 

    
    

        
    
    log.Println("API is running!")
    http.ListenAndServe(":8080", router)
}
