package authentication

import (
	"bufio"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"github.com/alamin-mahamud/golang-jwt-authentication-api-sample/settings"
	"github.com/alamin-mahamud/golang-jwt-authentication-api-sample/pkg/models"
	"github.com/dgrijalva/jwt-go"

	"golang.org/x/crypto/bcrypt"
	"os"
	"time"
)

type JWTAuthenticationBackend struct {
	privateKey *rsa.PrivateKey
	publicKey *rsa.PublicKey
}

const (
	tokenDuration = 72
	expireOffset = 3600
)

var authenticationBackendInstance *JWTAuthenticationBackend = nil

func InitJWTAuthenticationBackend() *JWTAuthenticationBackend{
	if authenticationBackendInstance == nil {
		authenticationBackendInstance = &JWTAuthenticationBackend{
			privateKey: getPrivateKey(),
			publicKey: getPublicKey(),
		}
	}
	return authenticationBackendInstance
}
func (backend *JWTAuthenticationBackend) GenerateToken(userUUID string) (string, error){
	token := jwt.New(jwt.SigningMethodRS512)
	token.Claims = jwt.MapClaims{
		"exp": time.Now().Add(time.Hour * time.Duration(settings.Get().JWTExpirationDelta)),
		"iat": time.Now().Unix(),
		"sub": userUUID,
	}
	tokenString, err := token.SignedString(backend.privateKey)
	if err != nil {
		panic(err)
		return "", err
	}

	return tokenString, nil
}
func (backend *JWTAuthenticationBackend) Authenticate(user *models.User) bool {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("testing"), 10)
	testUser := models.User{
		UUID: uuid.New(),
		Username: "alamin",
		Password: string(hashedPassword),
	}
	return user.Username == testUser.Username &&
	bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(testUser.Password)) == nil
}
func (backend *JWTAuthenticationBackend) getTokenRemainingValidity(timestamp interface{}) int {
	if validity, ok := timestamp.(float64); ok {
		tm := time.Unix(int64(validity), 0)
		remainer := tm.Sub(time.Now())
		if remainer > 0 {
			return int(remainer.Seconds() + expireOffset)
		}
	}
	return expireOffset
}
func (backend *JWTAuthenticationBackend) LogOut(tokenString string, token *jwt.Token) error {
	// log out
	return nil
}
func IsInBlackList(token string) bool {
	return true
}
func getPrivateKey() *rsa.PrivateKey{
	privateKeyFile, err := os.Open(settings.Get().PrivateKeyPath)
	defer privateKeyFile.Close()
	checkErrPanic(err)

	data := commonFileOperation(privateKeyFile)

	privateKeyImported, err := x509.ParsePKCS1PrivateKey(data.Bytes)
	checkErrPanic(err)

	return privateKeyImported
}
func getPublicKey() *rsa.PublicKey{
	publicKeyFile, err := os.Open(settings.Get().PublicKeyPath)
	defer publicKeyFile.Close()
	checkErrPanic(err)

	data := commonFileOperation(publicKeyFile)
	publicKeyImported, err := x509.ParsePKIXPublicKey(data.Bytes)
	checkErrPanic(err)

	rsaPub, ok := publicKeyImported.(*rsa.PublicKey)
	if !ok {
		panic("Could not convert to appropriate type")
	}
	return rsaPub
}
func commonFileOperation(f *os.File) *pem.Block{
	pemFileInfo, _ := f.Stat()
	var size int64 = pemFileInfo.Size()
	pemBytes := make([]byte, size)

	buffer := bufio.NewReader(f)
	_, err = buffer.Read(pemBytes)
	data, _ := pem.Decode([]byte(pemBytes))
	return data
}
func checkErrPanic(err error) {
	if err != nil {
		panic(err)
	}
}