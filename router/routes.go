package routes

import (
	"sethjoe/pkg/handlers"
)

func UserRoutesfunc(w http.ResponseWriter, r *http.Request){
	router.HandleFunc("/register", handlers.UserSignUP).Methods(http.MethodGet)
    router.HandleFunc("/login", handlers.Login).Methods(http.MethodPut)
    router.HandleFunc("/admin/addproduct", handlers.ProductViewerAdmin).Methods(http.MethodPost)
	router.HandleFunc("/users/productview", handlers.SearchProduct).Methods(http.MethodGet
	router.HandleFunc("/users/search", handlers.SearchProductByQuery).Methods(http.MethodGet)
}


