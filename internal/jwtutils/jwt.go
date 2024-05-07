package jwtutils

import (
	"errors"
	"time"

	"galen-gvm/global"
	"galen-gvm/model/system"
	"galen-gvm/utils"

	"github.com/gofrs/uuid/v5"
	"github.com/golang-jwt/jwt/v5"
	"go.uber.org/zap"
)

// Custom claims structure
type CustomClaims struct {
	BaseClaims
	BufferTime int64
	jwt.RegisteredClaims
}

type BaseClaims struct {
	UUID        uuid.UUID
	ID          uint
	Username    string
	NickName    string
	AuthorityId uint
}

type JWT struct {
	SigningKey []byte
}

var (
	TokenInvalid = errors.New("Couldn't handle this token:")
)

func NewJwt() *JWT {
	key := global.GVA_CONFIG.Jwt.SigningKey
	if key == "" {
		key = global.DefaultJwtSigningKey
	}
	return &JWT{
		SigningKey: []byte(key),
	}
}

// 创建声明信息
func (j *JWT) CreateClaims(user *system.SysUser) CustomClaims {
	bt, _ := utils.ParseDuration(global.GVA_CONFIG.Jwt.BufferTime)
	et, _ := utils.ParseDuration(global.GVA_CONFIG.Jwt.ExpiresTime)
	issuer := global.GVA_CONFIG.Jwt.Issuer
	if issuer == "" {
		issuer = global.DefaultJwtIssuer
	}
	return CustomClaims{
		BaseClaims: BaseClaims{
			UUID:        user.UUID,
			ID:          user.ID,
			Username:    user.Username,
			NickName:    user.NickName,
			AuthorityId: user.AuthorityId,
		},
		BufferTime: int64(bt / time.Second),
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    issuer,                                    // 签发人
			Audience:  jwt.ClaimStrings{"GVM"},                   // 受众
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(et)),    // 过期时间
			NotBefore: jwt.NewNumericDate(time.Now().Add(-1000)), //生效时间
		},
	}

}

// 解析
func (j *JWT) ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return j.SigningKey, nil
	})
	if err != nil {
		global.GVA_LOG.Error("ParseWithClaims err:", zap.Error(err))
		return nil, err
	}
	if token != nil {
		if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
			return claims, nil
		}
		return nil, TokenInvalid

	}
	return nil, TokenInvalid
}

// 创建
func (j *JWT) CreateToken(claims CustomClaims) (string, error) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return t.SignedString(j.SigningKey)
}
