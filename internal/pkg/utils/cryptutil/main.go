// 1.先生成私钥、公钥
// 2.使用公钥加密
// 3.使用私钥解密
package cryptutil

import (
	"crypto"
	"crypto/hmac"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/spf13/viper"
	"io"
	"readCommunity/global"
)

/*var (
	prvkey *rsa.PrivateKey
	pubkey []byte
)

func genrsakey() (privateKey *rsa.PrivateKey, publickey []byte) {
	var err error
	privateKey, err = rsa.GenerateKey(rand.Reader, 256)
	if err != nil {
		fmt.Printf("rsa.GenerateKey failed,err: %s\n", err)
	}
	// create privatekey
	pkcs1PrivateKey := x509.MarshalPKCS1PrivateKey(privateKey)
	privateBlock := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: pkcs1PrivateKey,
	}
	privatekey := pem.EncodeToMemory(privateBlock)

	// create publickey
	publicKey := &privateKey.PublicKey
	pkcs1PublicKey, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		fmt.Printf("x509.MarshalPKixPublicKey failed,err: %s\n", err)
	}
	publicBlock := &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: pkcs1PublicKey,
	}
	publickey = pem.EncodeToMemory(publicBlock)
	return
}

func init()  {
	prvkey, pubkey = genrsakey()
}

func GetKeyBytes() []byte {
	_, pubkey = genrsakey()
	return pubkey
}

func EncryptUtil(data []byte) string {
	publicBlock, _ := pem.Decode(pubkey)
	if publicBlock == nil {
		fmt.Println("pem.Decode failed")
	}
	parsePKIXPublicKey, err := x509.ParsePKIXPublicKey(publicBlock.Bytes)
	if err != nil {
		fmt.Printf("x509.parse failed,err: %s\n", err)
	}
	publicKey := parsePKIXPublicKey.(*rsa.PublicKey)
	//cryptedData, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, data)
	cryptedData, err := rsa.EncryptOAEP(sha256.New(),rand.Reader, publicKey, data, nil)
	if err != nil {
		fmt.Printf("rsa.encrypt failed, err: %s\n", err)
	}
	return hex.EncodeToString(cryptedData)
}

func DecryptUtil(data string) string {
	encrypt, _ := hex.DecodeString(data)
	decryptedData, err := rsa.DecryptOAEP(sha256.New(), rand.Reader, prvkey, encrypt, nil)
	if err != nil {
		fmt.Printf("rsa.decrypt failed, err: %v", decryptedData)
	}
	return hex.EncodeToString(decryptedData)
}*/
var (
	prvkey *rsa.PrivateKey
	ioReader io.Reader
)


func init() {
	ioReader = rand.Reader
	prvkey, global.Pubkey = generatekey()

}

func EncryptUtil(data, username string) string {
	key := viper.GetString("CryptKey")
	hash := hmac.New(sha256.New, []byte(key))
	hash.Write([]byte(data))
	hash.Write([]byte(username))
	return hex.EncodeToString(hash.Sum([]byte("lc")))
}

func Encrypt(data []byte) []byte {
	//fmt.Printf("pubkey encrypt, pubkey:%v\n", pubkey)
	encryptOAEP, err := rsa.EncryptOAEP(sha256.New(), ioReader, global.Pubkey, data, nil)
	if err != nil {
		fmt.Printf("encrypt failed,err: %v\n", err)
	}
	return encryptOAEP
}

func generatekey() (prvkey *rsa.PrivateKey, pubkey *rsa.PublicKey) {
	var err error
	prvkey, err = rsa.GenerateKey(ioReader, 1024)
	if err != nil {
		fmt.Printf("generatekey failed,err: %v\n", err)
	}
	pubkey = &prvkey.PublicKey
	return
}

func Decrypt(data string) string {
	/*decodeString, err:= hex.DecodeString(data)
	if err != nil {
		fmt.Printf("decode failed,err: %v\n", err)
	}
	decryptOAEP, err := rsa.DecryptOAEP(sha256.New(), rand.Reader, prvkey, decodeString, []byte(label))
	if err != nil {
		fmt.Printf("decrypt failed,err: %v\n", err)
	}
	return hex.EncodeToString(decryptOAEP)*/
	plaintext, err := prvkey.Decrypt(nil, []byte(data), &rsa.OAEPOptions{Hash: crypto.SHA256})
	if err != nil {
		fmt.Printf("decrypt failed,err:%v\n", err)
	}
	return string(plaintext)
}
