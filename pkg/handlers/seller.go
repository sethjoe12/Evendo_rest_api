package handlers

import (
    "encoding/json"
    "io/ioutil"
    "log"
    "math/rand"
    "net/http"
    "github.com/gorilla/mux"
    "sethjoe/pkg/mocks"
    "sethjoe/pkg/models"
    "strconv"
)



func ViewAllProduct(w http.ResponseWriter, r *http.Request) {
    w.Header().Add("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(mocks.Product)
}
func AddProduct(w http.ResponseWriter, r *http.Request) {
    
    defer r.Body.Close()
    body, err := ioutil.ReadAll(r.Body)

    if err != nil {
        log.Fatalln(err)
    }

    var Product models.Product
    json.Unmarshal(body, &Product)
    Product.ProductId = rand.Intn(100)
    mocks.Product = append(mocks.Product, Product)
    w.Header().Add("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode("Created")
}
func DeleteProduct(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id, _ := strconv.Atoi(vars["id"])

    for index, Product := range mocks.Product {
        if Product.ProductId == id {
            
            mocks.Product = append(mocks.Product[:index], mocks.Product[index+1:]...)

            w.Header().Add("Content-Type", "application/json")
            w.WriteHeader(http.StatusOK)
            json.NewEncoder(w).Encode("Deleted")
            break
        }
    }
}

