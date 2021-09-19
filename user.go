package main

import (
	// "crypto/rand"
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("secret_key")

var users = map[string]string{
	"user1": "password1",
	"user2": "password2",
}

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type GetCouponRequestParams struct {
	Brand    string `json:"brand"`
	Category string `json:"category"`
}

type GetCouponResponse struct {
	Username string `json:"username"`
	Brand    string `json:"brand"`
	Category string `json:"category"`
	Code     string `json:"code"`
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// Login functionality for the user using JWT Token
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var credentials Credentials
	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	expectedPassword, ok := users[credentials.Username]

	if !ok || expectedPassword != credentials.Password {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	expirationTime := time.Now().Add(time.Minute * 5)

	claims := &Claims{
		Username: credentials.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.SetCookie(w,
		&http.Cookie{
			Name:    "token",
			Value:   tokenString,
			Expires: expirationTime,
		})

}

// Fetch a discount code for the brand and category requested by the user
func GetCouponForUser(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	tokenStr := cookie.Value
	claims := &Claims{}

	tkn, err := jwt.ParseWithClaims(tokenStr, claims,
		func(t *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if !tkn.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	var params GetCouponRequestParams
	err = json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	rObj := &GetCouponResponse{Brand: params.Brand, Category: params.Category, Username: claims.Username, Code: getCoupon(params, claims.Username)}

	response, err := json.Marshal(rObj)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)

}

func getCoupon(r GetCouponRequestParams, u string) string {
	var coupon string
	var position int
	file, _ := ioutil.ReadFile("coupons.json")

	data := CouponJSON{}
	_ = json.Unmarshal([]byte(file), &data)

	for _, c := range data.Coupons {
		if c.Brand == r.Brand && c.Category == r.Category {
			position = rand.Intn(len(c.Coupons))
			coupon = c.Coupons[position]
			//c.Coupons = append(c.Coupons[:position], c.Coupons[position+1:]...)
			//c.Coupons = c.Coupons[:len(c.Coupons)-1]
			//c.RedeemedCoupons = append(c.RedeemedCoupons, RedeemedItem{Username: u, Code: coupon})
		}
	}

	//NotifyCouponStore(r.Brand, r.Category, u, data, coupon)

	return coupon
}
