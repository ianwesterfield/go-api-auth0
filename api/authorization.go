package main

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"net/http"

	goAuth0 "github.com/auth0-community/go-auth0"
	jose "gopkg.in/square/go-jose.v2"
)

// LoadPublicKey - gets the pem file contents
func LoadPublicKey(data []byte) (interface{}, error) {
	input := data

	block, _ := pem.Decode(data)
	if block != nil {
		input = block.Bytes
	}

	// Try to load SubjectPublicKeyInfo
	pub, err0 := x509.ParsePKIXPublicKey(input)
	if err0 == nil {
		return pub, nil
	}

	cert, err1 := x509.ParseCertificate(input)
	if err1 == nil {
		return cert.PublicKey, nil
	}

	return nil, fmt.Errorf("square/go-jose: parse error, got '%s' and '%s'", err0, err1)
}

// auth0 - middleware to handle authorization with Auth0
func auth0(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		pem, err := ioutil.ReadFile(configuration.Auth0.Certificate)
		checkErr(err)

		secret, err := LoadPublicKey(pem)
		checkErr(err)

		secretProvider := goAuth0.NewKeyProvider(secret)
		audience := configuration.Auth0.Audience
		configuration := goAuth0.NewConfiguration(secretProvider, []string{audience}, configuration.Auth0.Issuer, jose.RS256)
		validator := goAuth0.NewValidator(configuration, nil)
		token, err := validator.ValidateRequest(r)

		if err != nil {
			fmt.Println("Token is not valid:", token)
			w.WriteHeader(http.StatusUnauthorized)
			w.Header().Set("Content-Type", "text/plain")
			w.Write([]byte("Unauthorized"))
		} else {
			next.ServeHTTP(w, r)
		}
	})
}
