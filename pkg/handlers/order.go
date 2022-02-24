package handlers

import (
    "encoding/json"
    "io/ioutil"
    "log"
    "math/rand"
    "net/http"
  //  "github.com/gorilla/mux"
    "sethjoe/pkg/mocks"
    "sethjoe/pkg/models"
   // "strconv"
)



func OrderViewALllProducts(w http.ResponseWriter, r *http.Request) {
    w.Header().Add("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(mocks.Product)
}
func Order(w http.ResponseWriter, r *http.Request) {
    
    defer r.Body.Close()
    body, err := ioutil.ReadAll(r.Body)

    if err != nil {
        log.Fatalln(err)
    }

    var Order models.Order
    json.Unmarshal(body, &Order)
    Order.OrderId = rand.Intn(100)
    mocks.Order = append(mocks.Order, Order)
    w.Header().Add("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode("Created")
}
func ViewOrders(w http.ResponseWriter, r *http.Request) {
    w.Header().Add("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(mocks.Order)
}
