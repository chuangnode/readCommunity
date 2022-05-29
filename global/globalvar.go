package global

import (
	"crypto/rsa"
	ut "github.com/go-playground/universal-translator"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	DB     *gorm.DB
	Trans  ut.Translator
	Pubkey *rsa.PublicKey
	Lg     *zap.Logger
)
