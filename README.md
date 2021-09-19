
# Coupons Service
A simple coupon generator service where brands can generate coupons for their promotions, and users can claim a coupon code.


## Setup in Local

To deploy this project make sure you have golang set up in your environment. If not you can follow the instructions in the below links

https://golang.org/doc/install

https://golang.org/doc/tutorial/getting-started

Once you have Golang set up, go to the root fo the project folder and run the below commands

```bash
  go build
```

This will generate an executable file that you can run later to have the project up and running. Depending on which OS you are running on, use either of the below commands

```bash
  ./coupons (For MAC)
  ./coupons.exe (For Windows)
```
 
### Base URL

http://localhost:3000/

## Requests/End points

### Login
**Description**: Autheticates a user using JWT

**Path**: /login

**Request Type**: POST

**Body Parameters**:
```json
  {
    "username": "user1",
    "password": "password1"
  }
```

**Response**: 200 Upon success

### Get Coupons
**Description**: Return a coupon code of the desired brand category, if the user is logged in or else return an error

**Path**: /user/getCoupon

**Request Type**: POST

**Body Parameters**:
```json
  {
	  "brand": "zara",
	  "category": "autumn sale"
  }
```
#### Success Response
**Status Code**: 200
**Response**:
```json
  {
    "username": "user1",
    "brand": "zara",
    "category": "autumn sale",
    "code": "F4APQLBT"
  }
```
#### Failure Response
**Status Code**: 401
**Response**: Unauthorized user

### Generate Brand Coupons
**Description**: Generate a given number of coupons for a brand in a desired category/store

**Path**: /brand/generateCoupons

**Request Type**: POST

**Body Parameters**:
```json
  {
	  "brand": "wilson",
	  "category": "federer campaign",
	  "count": 20
  }
```
#### Success Response
**Status Code**: 200
**Response**:
```json
  {
    "brand": "wilson",
    "category": "pro federer",
    "count": 20,
    "coupons": [
        "FPLLNGZI",
        "EYOH43E0",
        "133OLS6K",
        "1HH2GDNY",
        "XXVI7HVS",
        "ZWK1B182",
        "TVJZJPEZ",
        "I4HX9GVM",
        "KIR0XCTA",
        "0OPSB5QI",
        "PJZB3H3X",
        "9KCEGTA5",
        "M1ZCV5DR",
        "XCKN42GB",
        "50ANXNDS",
        "CKJDWGFW",
        "5JAPZ01Z",
        "ICAPY9EQ",
        "IXUC9UEH",
        "Q235V48C"
    ],
    "redeemed_coupons": null,
    "remaining_coupons": 20
  }
```
#### Failure Response
**Status Code**: 400
**Response**: Bad Request