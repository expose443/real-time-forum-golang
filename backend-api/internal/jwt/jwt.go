package jwt

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

func CreateJWT(secretKey string, claims map[string]interface{}) (string, error) {
	header := map[string]string{
		"alg": "HS256",
		"typ": "JWT",
	}

	claims["iat"] = time.Now()
	headerBytes, err := json.Marshal(header)
	if err != nil {
		return "", err
	}
	fmt.Println(headerBytes)
	encodedHeader := base64.RawURLEncoding.EncodeToString(headerBytes)
	fmt.Println(encodedHeader)
	claimsBytes, err := json.Marshal(claims)
	if err != nil {
		return "", err
	}
	encodedClaims := base64.RawURLEncoding.EncodeToString(claimsBytes)
	fmt.Println(encodedClaims)

	data := []byte(encodedHeader + "." + encodedClaims)
	h := hmac.New(sha256.New, []byte(secretKey))
	_, err = h.Write(data)
	if err != nil {
		return "", err
	}
	signature := base64.RawURLEncoding.EncodeToString(h.Sum(nil))
	token := strings.Join([]string{encodedHeader, encodedClaims, signature}, ".")

	return token, nil
}

func VerifyJWT(secretKey string, token string) (bool, map[string]interface{}, error) {
	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		return false, nil, nil
	}

	headerBytes, err := base64.RawURLEncoding.DecodeString(parts[0])
	if err != nil {
		return false, nil, err
	}
	var header map[string]string
	err = json.NewEncoder(bytes.NewBuffer(headerBytes)).Encode(&header)
	if err != nil {
		return false, nil, err
	}

	var claims map[string]interface{}
	claimsBytes, err := base64.RawURLEncoding.DecodeString(parts[1])
	if err != nil {
		return false, nil, err
	}
	err = json.NewEncoder(bytes.NewBuffer(claimsBytes)).Encode(&claims)
	if err != nil {
		return false, nil, err
	}

	if header["alg"] != "HS256" {
		return false, nil, nil
	}

	data := []byte(parts[0] + "." + parts[1])
	signature, err := base64.RawURLEncoding.DecodeString(parts[2])
	if err != nil {
		return false, nil, err
	}
	h := hmac.New(sha256.New, []byte(secretKey))
	_, err = h.Write(data)
	if err != nil {
		return false, nil, err
	}
	if !hmac.Equal(signature, h.Sum(nil)) {
		return false, nil, nil
	}

	exp, ok := claims["exp"].(float64)
	if !ok {
		return false, nil, nil
	}
	if int64(exp) < time.Now().Unix() {
		return false, nil, nil
	}
	return true, claims, nil
}
