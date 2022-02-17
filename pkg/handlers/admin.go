package handlers

import (
	 "net/http"
     "github.com/dgrijalva/jwt-go"
	 "encoding/json"
	 "time"
	 "fmt"

)


var jwtkey = []byte("secretkey")
var Users = map[string]string{
	"users1":"password1",
	"users2":"password2",
}
type Credential struct{
	Username string `json:"username"`
	Password string `json:"password"`
}
type Claims struct{
	Username string `json:"username"`
	jwt.StandardClaims 

}

func AdminLogin(w http.ResponseWriter, r *http.Request){
         var credentials Credential
		 err := json.NewDecoder(r.Body).Decode(&credentials)
		 if err != nil{
			 w.WriteHeader(http.StatusBadRequest)
			 return
        }

		expectedPassword, ok := Users[credentials.Username]
		if !ok || expectedPassword != credentials.Password {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		expirationTime := time.Now().Add(time.Minute *5)

		claims := &Claims{
			Username:credentials.Username,
			StandardClaims: jwt.StandardClaims{
				ExpiresAt:expirationTime.Unix(),				
			},
		}

		token :=jwt.NewWithClaims(jwt.SigningMethodHS256,claims)
		tokenString,err:= token.SignedString(jwtkey)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		http.SetCookie(w,
		&http.Cookie{
			Name:  "token",
			Value:  tokenString,
			Expires: expirationTime,
		})

}
func Home(w http.ResponseWriter, r *http.Request){
	cookie, err := r.Cookie("token")
	if err != nil{
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return

	}
	
	tokenStr := cookie.Value

	claims := &Claims{}

	tkn,err := jwt.ParseWithClaims(tokenStr, claims,
	func(t *jwt.Token)(interface{}, error){
		return jwtkey ,nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
		  return
		}

		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if !tkn.Valid{
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	w.Write([]byte(fmt.Sprintf("Hello, %s",claims.Username)))

}
func Refresh(w http.ResponseWriter, r *http.Request){

}