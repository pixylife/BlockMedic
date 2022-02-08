package env

import (
	"github.com/dgrijalva/jwt-go"
	gorm "github.com/jinzhu/gorm"
)

var RDB *gorm.DB

var RestPort = "8080"

var DBDIALET = "mysql"

var DBdb = "blockmedic"

var DBhost = "127.0.0.1"

var DBport = "3306"

var DBuser = "root"

var DBpwd = "sahan"

var StripeAPIKey = "pk_test_51IdOGzIc2dfDJNEgxjsDGTWzPiq2SjJyRkKeqRXiYQoytDe1Ye8QrTJSyoHtLMyaFKKbxeAIqxjxxR3n8alBjfGM00xiDA7cDq"

var StripeSecretKey = "sk_test_51IdOGzIc2dfDJNEgcotpLXZQcC9WtMhubRf3KO4q6oQ8gzY3hrDwJizcL8tWsTRjuBhmpbiLIHNYCiqU9wLLGBtH008jb3bIbc"

const (
	DOCTOR          = "ROLE_DOCTOR"
	PATIENT         = "ROLE_PATIENT"
	ADMIN           = "ROLE_ADMIN"
	STATUS_NEW      = "NEW"
	STATUS_REJECTED = "REJECTED"
	STATUS_APPROVED = "APPROVED"
)

type JwtCustomClaims struct {
	Sub  string
	Auth string
	Role string
	jwt.StandardClaims
}

type Token struct {
	Role string
	jwt.StandardClaims
}
