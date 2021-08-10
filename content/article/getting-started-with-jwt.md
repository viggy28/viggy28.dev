---
date: 2021-08-09T00:20:08-04:00
description: "Things that I learnt about JWT"
featured_image: "/images/jwt_logo.svg"
tags: ["JWT", "Go"]
title: "Getting Started with JWT"
---

Disclaimer: Whatever I write here is my learnings. So, don't take these words as granted or reference. Also, please lmk if there are any errors.

I have often come across this term JWT. I knew the definition - Json Web Token, but honestly that's all I knew for a veryyyy longgg time. Recently, I came across a problem where I have been told that JWT might be a good solution for it. That piqued my interest.

From what I understood, a JWT implementation simply encrypts some plain text and returns the encrypted token of the plain text.

The problem I had was how to verify that the corresponding user is the one who is requesting the resource. Authentication solves only whether the user is legit or not but it doesn't solve whether the user is the right owner to access a resource. I solved that using JWT.

JWT uses private and public key cryptography.

```go
func main() {
    signBytes, _ := ioutil.ReadFile("app.rsa")
	signKey, _ := jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	verifyBytes, _ := ioutil.ReadFile("app.rsa.pub")
	verifyKey, _ := jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
	jwt, _ := stringToJWT("hello world")
}
```

```go
func stringToJWT(text string) (string, error) {
	// Declare the expiration time of the token
	// here, we have kept it as 2 days
	log.Println("string to JWTify:", text)
	expirationTime := time.Now().Add(48 * time.Hour)
	// Create the JWT claims, which includes the text and expiry time
	claims := &claims{
		Text: text,
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: expirationTime.Unix(),
		},
	}
	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodPS512, claims)
	// Create the JWT string
	jwt, err := token.SignedString(signKey)
	if err != nil {
		return "", err
	}
	return jwt, nil
}
```

`app.rsa` - private key file
`app.rsa.pub` - public key file

You can generate these files using OpenSSL.

**To create a private key**
```
viggy28@viggy28s-MacBook-Pro spinup % openssl genrsa -out /tmp/spinuplocal/app.rsa 4096 
Generating RSA private key, 4096 bit long modulus
...++
...................++
e is 65537 (0x10001)
```

**To create a public key**
```
viggy28@viggy28s-MacBook-Pro spinup % openssl rsa -in /tmp/spinuplocal/app.rsa -pubout > /tmp/spinuplocal/app.rsa.pub
writing RSA key
```
