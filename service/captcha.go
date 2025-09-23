package service

import (
	"bookstore-go/global"
	"context"
	"fmt"
	"log"
	"time"

	"github.com/mojocn/base64Captcha"
)

type CaptchaService struct {
	store base64Captcha.Store
}

type CaptchaResponse struct {
	CaptchaID string `json:"captcha_id"`
	Image     string `json:"image"`
}

func NewCaptchaService() *CaptchaService {
	return &CaptchaService{
		store: base64Captcha.DefaultMemStore,
	}
}

func (s *CaptchaService) GenerateCaptcha() (*CaptchaResponse, error) {
	driver := base64Captcha.NewDriverDigit(80, 240, 5, 0.7, 80)
	base64Captcha := base64Captcha.NewCaptcha(driver, s.store)
	id, b64s, answer, err := base64Captcha.Generate()
	if err != nil {
		return nil, err
	}
	log.Println("captcha answer:", answer)
	redis_key := fmt.Sprintf("captcha_%s", id)
	err = global.RedisClient.Set(context.TODO(), redis_key, answer, 1*time.Minute).Err()
	if err != nil {
		return nil, err
	}
	return &CaptchaResponse{
		CaptchaID: id,
		Image:     b64s,
	}, nil
}

func (s *CaptchaService) VerifyCaptcha(captchaID, captcha string) bool {
	redis_key := fmt.Sprintf("captcha_%s", captchaID)
	answer, err := global.RedisClient.Get(context.TODO(), redis_key).Result()
	if err != nil {
		return false
	}
	return captcha == answer
}
