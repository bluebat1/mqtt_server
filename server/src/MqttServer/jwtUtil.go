package MqttServer

import (
	"errors"
	"math/rand"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWT struct {
	User  string
	Token string
	jwt.RegisteredClaims
}

// 随机字符串
var letters = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randStr(str_len int) string {
	rand_bytes := make([]rune, str_len)
	for i := range rand_bytes {
		rand_bytes[i] = letters[rand.Intn(len(letters))]
	}
	return string(rand_bytes)
}

// 签名密钥
var jwtSignKey string

func JwtInit() error {
	jwtSignKey = randStr(64)
	return nil
}

// new a jwt token
func JwtNew(user string, token string) (string, error) {
	claim := JWT{
		User:  user,
		Token: token,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "Auth_Server",                                   // 签发者
			Subject:   token,                                           // 签发对象
			Audience:  jwt.ClaimStrings{"web"},                         //签发受众
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),   //过期时间
			NotBefore: jwt.NewNumericDate(time.Now().Add(time.Second)), //最早使用时间
			IssuedAt:  jwt.NewNumericDate(time.Now()),                  //签发时间
			ID:        randStr(10),                                     // wt ID, 类似于盐值
		},
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claim).SignedString([]byte(jwtSignKey))
	return token, err
}

// 验签 jwt
func JwtParse(token string) (*JWT, error) {
	jwtToken, err := jwt.ParseWithClaims(token, &JWT{}, func(token *jwt.Token) (interface{}, error) {
		//返回签名密钥
		return []byte(jwtSignKey), nil
	})
	if err != nil {
		return nil, err
	}

	if !jwtToken.Valid {
		return nil, errors.New("claim invalid")
	}

	claims, ok := jwtToken.Claims.(*JWT)
	if !ok {
		return nil, errors.New("invalid claim type")
	}

	return claims, nil
}
