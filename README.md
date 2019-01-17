# go-jwt
JWT Authentication Implementation in golang

## Model
```go
package models

type User struct {
	UUID     string `json:"uuid" form:"-"`
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}
```

## authentication
* Backend
```go
type JWTAuthenticationBackend struct {
	privateKey *rsa.PrivateKey
	publicKey *rsa.PublicKey
}

const (
	tokenDuration = 72
	expireOffset = 3600
)
```

using a package variable for the backendInstance
```go
var authenticationBackendInstance *JWTAuthenticationBackend = nil
```

initializing authentication Backend
```go
func InitJWTAuthenticationBackend() *JWTAuthenticationBackend{
	if authenticationBackendInstance == nil {
		authenticationBackendInstance = &JWTAuthenticationBackend{
			privateKey: getPrivateKey(),
			publicKey: getPublicKey(),
		}
	}
	
	return authenticationBackendInstance
}
```

get private key
```go
func getPrivateKey() *rsa.PrivateKey{
	privateKeyFile, err := os.Open(settings.Get().PrivateKeyPath)
	defer privateKeyFile.Close()
	checkErrPanic(err)
	
	data := commonFileOperation(privateKeyFile)
	privateKeyImported, err := x509.ParsePKCS1PrivateKey(data.Bytes)
	checkErrPanic(err)
	
	return privateKeyImported
}
```

get public key
```go
func getPublicKey *rsa.PublicKey{
	publicKeyFile, err := os.Open(settings.Get().PublicKeyPath)
    defer publicKeyFile.Close()
    checkErrPanic(err)

    data := commonFileOperation(publicKeyFile)
    publicKeyImported, err := x509.ParsePKIXPublicKey(data.Bytes)
    checkErrPanic(err)
    
    rsaPub, ok := publicKeyImported.(*rsa.PublicKey)
    if !ok {
    	panic("Could not convert to appropriate type.")
    }
    return rsaPub
}
```

common file operation
```go
func commonFileOperation(f *os.File) *pem.Block {
	pemFileInfo, _ := f.Stat()
	var size int64 = pemFileInfo.Size()
    pemBytes = make([]byte, size)
    
    buffer := bufio.NewReader(f)
    _, err = buffer.Read(pemBytes)
    return data
}
```

```go
func (backend *JWTAuthenticationBackend) GenerateToken(userUUID string)(string, error) {
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
}
```

```go
func (backend *JWTAuthenticationBackend) Authenticate(user *model.User) bool {
	return user.Username == dbUser.Username &&
		bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(dbUser.Password)) == nil
}
```

```go
func (backend *JWTAuthenticationBackend) Logout(tokenString string, token *jwt.Token) error {
	// Log out by doing necessary operations
	return nil
}
```

```go
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
```

```go
func IsInBlackList(token string) bool {
	// 1. get user name or uuid
	// 2. find username or uuid in blackList Array
	// ok := findInBlackList(username)
	var ok bool = false
	
	if ok {
		return true
	}
	
	return false
}
```

* Middleware
```go
func RequireTokenAuthentication(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
    authBackend := InitJWTAuthenticationBackend()
    token, err := request.ParseFromRequest(req, request.OAuth2Extractor, func(token *jwt.Token)(interface{}, error){
    	if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
    		return nil, fmt.Errorf("Unexpected signing method: %v", token)
    	} else {
    	    return authBackend.PublicKey, nil	
    	}
    })
    
    if err == nil && token.Valid && !authBackend.IsInBlacklist(req.Header.Get("Authorization")) {
    		next(rw, req)
    	} else {
    		rw.WriteHeader(http.StatusUnauthorized)
    }
}
```


POST http://localhost:12345/authenticate

``` json
{
	"username": "alamin-mahamud",
  	"password": "simple-password"
}
```

``` json
// Response
{
	"token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFsYW1pbi1tYWhhbXVkIn0.3pDO8lruVO2GAnABjknpMZK03XzsVktVBkNBmAbz-8I"
}

```
