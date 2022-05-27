package global

import (
	"crypto/rsa"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
	Pubkey *rsa.PublicKey
)
