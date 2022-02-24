package routes

import (
    "log"
    "net/http"
    "sethjoe/middleware"
	"sethjoe/pkg/tokens"
    "github.com/gorilla/mux"
    "sethjoe/pkg/handlers"
  
)

func Routes() {
    router := mux.NewRouter()
   
    //for admin purposes
    router.Handle("/admin", middleware.BasicAuthMiddleware(http.HandlerFunc(handlers.AdminEntry))).Methods("GET")
    router.Handle("/admin/addproducts", middleware.BasicAuthMiddleware(http.HandlerFunc(handlers.AdminAddProduct))).Methods("POST")
	router.Handle("/admin/deleteproducts/{id}", middleware.BasicAuthMiddleware(http.HandlerFunc(handlers.AdminDeleteProduct))).Methods("DELETE")
    router.Handle("/admin/viewAllProducts", middleware.BasicAuthMiddleware(http.HandlerFunc(handlers.AdminViewProducts))).Methods("GET")
	router.Handle("/admin/viewAlluser", middleware.BasicAuthMiddleware(http.HandlerFunc(handlers.ViewAlluser))).Methods("GET")
    router.Handle("/admin/viewOrderbyID/{id}", middleware.BasicAuthMiddleware(http.HandlerFunc(handlers.AdminViewOrder))).Methods("GET")

    
    router.HandleFunc("/UserSignup", token.UserEntry).Methods(http.MethodPost)
    router.HandleFunc("/login", middleware.UserMiddleware(token.UserTokenLogin)).Methods(http.MethodGet) 
    
    router.HandleFunc("/seller", handlers.ViewAllProduct).Methods(http.MethodGet)//wiew all products
    router.HandleFunc("/seller/addproduct", middleware.UserMiddleware(handlers.AddProduct)).Methods(http.MethodGet)//add products
    router.HandleFunc("/seller/deleteproduct/{id}",middleware.UserMiddleware(handlers.DeleteProduct)).Methods(http.MethodDelete)//delete products
    
    router.HandleFunc("/Order/ViewALllProducts", handlers.OrderViewALllProducts).Methods(http.MethodGet)//view prducts
    router.HandleFunc("/Order", middleware.BuyerMiddleware(handlers.Order)).Methods(http.MethodPost)//order product//use methid[add/get]/use product id 
    router.HandleFunc("/Order/ViewOrders", handlers.ViewOrders).Methods(http.MethodGet)
    
    

        
    
    log.Println("API is running!")
    http.ListenAndServe(":8080", router)
}