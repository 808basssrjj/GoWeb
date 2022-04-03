package jwt

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

// MyClaims 自定义声明结构体并内嵌jwt.StandardClaims
// jwt包自带的jwt.StandardClaims只包含了官方字段
// 如果想要保存更多信息，都可以添加到这个结构体中
type MyClaims struct {
	UserID int64 `json:"user_id"`
	jwt.StandardClaims
}

// TokenExpireDuration 过期时间
const TokenExpireDuration = time.Hour * 2

// mySecret Signature密钥 是对前两部分的签名，防止数据篡改。
var mySecret = []byte("gophergogogo")

//GenToken 生成JWT
//func GenToken(userID int64) (string, error) {
//	// 创建一个我们自己的声明
//	c := MyClaims{
//		userID, // 自定义字段
//		jwt.StandardClaims{
//			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(), // 过期时间
//			Issuer:    "znnn",                                     // 签发人
//		},
//	}
//	// 使用指定的签名方法创建签名对象
//	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
//	// 使用指定的secret签名并获得完整的编码后的字符串token
//	return token.SignedString(MySecret)
//}
//
//// ParseToken 解析JWT
//func ParseToken(tokenString string) (*MyClaims, error) {
//	mc := new(MyClaims)
//	// 解析token
//	token, err := jwt.ParseWithClaims(tokenString, mc, func(token *jwt.Token) (i interface{}, err error) {
//		return MySecret, nil
//	})
//	if err != nil {
//		return nil, err
//	}
//	if token.Valid { // 校验token
//		return mc, nil
//	}
//	return nil, errors.New("invalid token")
//}

func keyFunc(_ *jwt.Token) (i interface{}, err error) {
	return mySecret, nil
}

// GenToken 生成access token 和 refresh token
func GenToken(userID int64) (aToken, rToken string, err error) {
	// 创建一个我们自己的声明
	c := MyClaims{
		userID, // 自定义字段
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(), // 过期时间
			Issuer:    "bluebell",                                 // 签发人
		},
	}
	// 加密并获得完整的编码后的字符串token
	aToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, c).SignedString(mySecret)

	// refresh token 不需要存任何自定义数据
	rToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Hour * 120).Unix(), // 过期时间
		Issuer:    "bluebell",                             // 签发人
	}).SignedString(mySecret)
	// 使用指定的secret签名并获得完整的编码后的字符串token
	return
}

// ParseToken 解析JWT
func ParseToken(tokenString string) (claims *MyClaims, err error) {
	mc := new(MyClaims)
	token, err := jwt.ParseWithClaims(tokenString, mc, keyFunc)
	if err != nil {
		return nil, err
	}
	if token.Valid { // 校验token
		return mc, nil
	}
	return nil, errors.New("invalid token")
}

// RefreshToken 刷新AccessToken
func RefreshToken(aToken, rToken string) (newAToken, newRToken string, err error) {
	// refresh token无效直接返回
	if _, err = jwt.Parse(rToken, keyFunc); err != nil {
		return
	}

	// 从旧access token中解析出claims数据
	var claims MyClaims
	_, err = jwt.ParseWithClaims(aToken, &claims, keyFunc)
	v, _ := err.(*jwt.ValidationError)

	// 当access token是过期错误 并且 refresh token没有过期时就创建一个新的access token
	if v.Errors == jwt.ValidationErrorExpired {
		return GenToken(claims.UserID)
	}
	return
}
