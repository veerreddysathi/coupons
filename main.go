package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/login", LoginHandler)
	http.HandleFunc("/user/getCoupon", GetCouponForUser)
	http.HandleFunc("/brand/generateCoupons", GenerateCouponsHandler)

	log.Fatal(http.ListenAndServe(":3000", nil))
}
