package jwt

import (
	"bookstore-go/global"
	"context"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte("bookstore-go-secret")

const (
	AccessTokenExpireDuration  = 2 * time.Hour
	RefreshTokenExpireDuration = 7 * 24 * time.Hour
)

type Claims struct {
	UserID    int    `json:"user_id"`
	Username  string `json:"username"`
	TokenType string `json:"token_type"`
	jwt.RegisteredClaims
}

type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int64  `json:"expires_in"`
}

func GenerateToken(userID int, username string) (*TokenResponse, error) {
	accessClaims := Claims{
		UserID:    userID,
		Username:  username,
		TokenType: "access",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(AccessTokenExpireDuration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}
	jwtAccessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	jwtAccessTokenString, err := jwtAccessToken.SignedString(jwtSecret)
	if err != nil {
		return nil, err
	}

	refreshClaims := Claims{
		UserID:    userID,
		Username:  username,
		TokenType: "refresh",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(RefreshTokenExpireDuration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}
	jwtRefreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	jwtRefreshTokenString, err := jwtRefreshToken.SignedString(jwtSecret)
	if err != nil {
		return nil, err
	}
	SetTokenToRedis(userID, jwtAccessTokenString, jwtRefreshTokenString)
	return &TokenResponse{
		AccessToken:  jwtAccessTokenString,
		RefreshToken: jwtRefreshTokenString,
		ExpiresIn:    int64(AccessTokenExpireDuration.Seconds()),
	}, nil
}

func SetTokenToRedis(userID int, accessToken, refreshToken string) error {
	ctx := context.Background()
	userKey := fmt.Sprintf("user:%d", userID)
	err := global.RedisClient.HMSet(ctx, userKey, "access_token", accessToken, "access_token_expire", AccessTokenExpireDuration, "refresh_token", refreshToken, "refresh_token_expire", AccessTokenExpireDuration).Err()
	if err != nil {
		return err
	}
	return global.RedisClient.Expire(ctx, userKey, RefreshTokenExpireDuration).Err()
}

func IsTokenValid(userID uint, token string, tokenType string) bool {
	ctx := context.Background()
	userKey := fmt.Sprintf("user:%d", userID)
	var redisToken string
	var err error
	if tokenType == "access" {
		redisToken, err = global.RedisClient.HGet(ctx, userKey, "access_token").Result()
	} else if tokenType == "refresh" {
		redisToken, err = global.RedisClient.HGet(ctx, userKey, "refresh_token").Result()
	}
	if err != nil {
		return false
	}
	return token == redisToken
}
