package main

import (
	// "crypto/rand"
	"encoding/json"
	"math/rand"
	"net/http"
)

type GenerateCouponsRequest struct {
	Brand    string `json:"brand"`
	Category string `json:"category"`
	Count    int    `json:"count"`
}

type Coupon struct {
	Brand            string         `json:"brand"`
	Category         string         `json:"category"`
	Count            int            `json:"count"`
	Coupons          []string       `json:"coupons"`
	RedeemedCoupons  []RedeemedItem `json:"redeemed_coupons"`
	RemainingCoupons int            `json:"remaining_coupons"`
}

type RedeemedItem struct {
	Username string `json:"username"`
	Code     string `json:"code"`
}

type CouponJSON struct {
	Coupons []Coupon `json:"coupons"`
}

// Main Handler function that accepts the request to generate coupon codes
func GenerateCouponsHandler(w http.ResponseWriter, r *http.Request) {
	var params GenerateCouponsRequest
	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	response, err := json.Marshal(generateCoupons(params))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

// This function generates the coupon struct for a requested category of a brand with the coupon codes
func generateCoupons(r GenerateCouponsRequest) Coupon {
	coupons := &Coupon{Brand: r.Brand, Category: r.Category, Count: r.Count, Coupons: []string{}}
	for i := 0; i < r.Count; i++ {
		c := RandomString(8)
		coupons.Coupons = append(coupons.Coupons, c)
	}

	return *coupons
}

// This function generates the discount codes
func RandomString(n int) string {
	var letter = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	b := make([]rune, n)
	for i := range b {
		b[i] = letter[rand.Intn(len(letter))]
	}
	return string(b)
}

// Function to update the coupons and notify the brand
//func NotifyCouponStore(brand string, category string, user string, data CouponJSON, code string) {
//	item := RedeemedItem{Username: user, Code: code}
//	for _, c := range data.Coupons {
//		if c.Brand == brand && c.Category == category {
//			c.RedeemedCoupons = append(c.RedeemedCoupons, item)
//		}
//	}
//	e := os.Remove("coupons.json")
//	if e != nil {
//		log.Fatal(e)
//	}
//	file, _ := json.MarshalIndent(data, "", " ")
//
//	_ = ioutil.WriteFile("coupons.json", file, 0644)
//
//	return
//}
