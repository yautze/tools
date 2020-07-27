package token

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var (
	key    []byte                   // hash key (salt)
	method = jwt.SigningMethodHS512 // hash method
	issuer string                   // jwt的發行人
)

// JWT -
type JWT interface {
	Create(body interface{}) JWT
	Parse(token string, body interface{}) JWT
	Error() error
	Valid() error
	Get() string
}

type jt struct {
	t   *jwt.Token
	jwt string
	err error
}

// SetUp -
func SetUp(skey, iss string) {
	key = []byte(skey) // set key
	issuer = iss       // set issuer
}

// NewJwtMap - type MapClaims map[string]interface{}
func NewJwtMap(sub string, exp int64, nbf int64) jwt.MapClaims {
	j := jwt.MapClaims{
		"iss": issuer,
		"sub": sub, // sub - 標題
		"iat": time.Now().Unix(),
	}
	// exp - 過期時間timestamp, 不設置 : 0
	if exp != 0 {
		j["exp"] = exp
	}
	// nbf - 生效時間timestamp, 不設置 : 0
	if nbf != 0 {
		j["nbf"] = nbf
	}

	return j
}

// NewJWT -
func NewJWT() JWT {
	return new(jt)
}

// Create - create
func (j *jt) Create(body interface{}) JWT {
	switch body.(type) {
	case jwt.MapClaims:
		j.t = jwt.NewWithClaims(method, body.(jwt.MapClaims))
	case jwt.Claims:
		j.t = jwt.NewWithClaims(method, body.(jwt.Claims))
	default:
		j.err = fmt.Errorf("wrong data type , need (JwtMap / JwtClaims)")
	}
	j.jwt, j.err = j.t.SignedString(key)
	return j
}

// 解密 token 內容
func (j *jt) Parse(token string, body interface{}) JWT {
	j.t, j.err = jwt.ParseWithClaims(token, body.(jwt.Claims), keyLookup)

	if j.err == nil {
		j.err = j.Valid()
	}

	return j
}

// check which type error
func (j *jt) Valid() error {
	switch j.t.Claims.(type) {
	case jwt.MapClaims:
		t := j.t.Claims.(jwt.MapClaims)
		if !t.VerifyIssuer(issuer, true) {
			return fmt.Errorf("the issuer is not  %v", issuer)
		}
		if t.Valid() != nil {
			return fmt.Errorf("valid token failed")
		}
	case jwt.Claims:
		if j.t.Claims.Valid() != nil {
			return fmt.Errorf("valid token failed")
		}
	default:
		return fmt.Errorf("valid token failed")
	}

	return nil
}

// Error - return err
func (j *jt) Error() error {
	return j.err
}

// Get - get jwt (access token)
func (j *jt) Get() string {
	return j.jwt
}

// 解密 token 內的 info (map[string]string)
func keyLookup(t *jwt.Token) (interface{}, error) {
	if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("jwt signature method faile")
	}
	return key, nil
}
