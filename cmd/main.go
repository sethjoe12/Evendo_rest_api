package main

import (
    "log"
    "net/http"
    "sethjoe/middleware"
    "github.com/gorilla/mux"
    "sethjoe/pkg/handlers"
    //"sethjoe/router"
)

func main() {
    router := mux.NewRouter()
    router.HandleFunc("/register", handlers.UserSignUP).Methods(http.MethodGet)
    router.HandleFunc("/login", handlers.Login).Methods(http.MethodPut)
    //for admin
    router.Handle("/admin/addproducts", basicauthmiddleware.BasicAuthMiddleware(http.HandlerFunc(handlers.AdminAddProduct))).Methods("GET")
	router.Handle("/admin/deleteproducts/{id}", basicauthmiddleware.BasicAuthMiddleware(http.HandlerFunc(handlers.AdminDeleteProduct))).Methods("DELETE")

    
    
    
    router.HandleFunc("/seller", handlers.ViewAllProduct).Methods(http.MethodGet)//wiew all products
    router.HandleFunc("/seller", handlers.AddProduct).Methods(http.MethodPost)//add products
    router.HandleFunc("/seller/id", handlers.DeleteProduct).Methods(http.MethodDelete)//delete products
    
    router.HandleFunc("/SeeAll", handlers.ViewProducts).Methods(http.MethodGet)//view prducts
    router.HandleFunc("/Order/", handlers.OrderProduct).Methods(http.MethodGet)//order product//use methid[add/get]/use product id 
    router.HandleFunc("/Order/{id}", handlers.ViewOrder).Methods(http.MethodPut)//update order
    
    

        
    
    log.Println("API is running!")
    http.ListenAndServe(":8080", router)
}
