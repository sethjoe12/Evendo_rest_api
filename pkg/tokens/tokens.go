package token

import (
	"fmt"
	"net/http"
	"github.com/dgrijalva/jwt-go"
	"encoding/json"
	"sethjoe/pkg/models"
	"github.com/mitchellh/mapstructure"
)

func UserEntry(w http.ResponseWriter, req *http.Request) {
	var username models.User
	_ = json.NewDecoder(req.Body).Decode(&username)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username.Username,
	})
	tokenString, error := token.SignedString([]byte("secret"))
	if error != nil {
		fmt.Println(error)
	}
	json.NewEncoder(w).Encode(models.JwtToken{Token: tokenString})
  }
  
  func ProtectedEndpoint(w http.ResponseWriter, req *http.Request) {
	params := req.URL.Query()
	token, _ := jwt.Parse(params["token"][0], func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("There was an error")
		}
		return []byte("secret"), nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		var username models.User
		mapstructure.Decode(claims, &username)
		json.NewEncoder(w).Encode(username)
	} else {
		json.NewEncoder(w).Encode(models.Exception{Message: "Invalid authorization token"})
	}
  }
  
  func UserTokenLogin(w http.ResponseWriter, req *http.Request){
	fmt.Fprint(w, "Hi! Good Day!")
	
  }

  ///this will display all data inputed by the users.
 // func UserTokenLogin(w http.ResponseWriter, req *http.Request) {
	//decoded := context.Get(req, "decoded")
	//var username models.User
	//mapstructure.Decode(decoded.(jwt.MapClaims), &username)
	//json.NewEncoder(w).Encode(username)
  //}
  