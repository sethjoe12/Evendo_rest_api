package handlers

import (
	"net/http"
	"io/ioutil"
	"log"
	"sethjoe/pkg/models"
	"sethjoe/pkg/mocks"
	"encoding/json"
	"math/rand"
	"strconv"
	"github.com/gorilla/mux"
)

func AdminAddProduct(w http.ResponseWriter, r *http.Request) {
    
    defer r.Body.Close()
    body, err := ioutil.ReadAll(r.Body)

    if err != nil {
        log.Fatalln(err)
    }

    var Product models.Product
    json.Unmarshal(body, &Product)
    Product.Id = rand.Intn(100)
    mocks.Product = append(mocks.Product, Product)
    w.Header().Add("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode("Created")
}

func AdminDeleteProduct(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id, _ := strconv.Atoi(vars["id"])

    for index, Product := range mocks.Product {
        if Product.Id == id {
            
            mocks.Product = append(mocks.Product[:index], mocks.Product[index+1:]...)

            w.Header().Add("Content-Type", "application/json")
            w.WriteHeader(http.StatusOK)
            json.NewEncoder(w).Encode("Deleted")
            break
        }
    }
}
