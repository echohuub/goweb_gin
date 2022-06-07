package tool

import (
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"image/color"
)

type CaptchaResult struct {
	Id           string `json:"id"`
	Base64Blob   string `json:"base_64_blob"`
	VertifyValue string `json:"code"`
}

func GenerateCaptcha(ctx *gin.Context) CaptchaResult {
	parameters := base64Captcha.ConfigCharacter{
		Height:                 30,
		Width:                  60,
		Mode:                   3,
		ComplexOfNoiseText:     0,
		ComplexOfNoiseDot:      0,
		IsShowHollowLine:       false,
		IsShowNoiseDot:         false,
		IsShowNoiseText:        false,
		IsShowSlimeLine:        false,
		IsShowSineLine:         false,
		ChineseCharacterSource: "",
		SequencedCharacters:    nil,
		UseCJKFonts:            false,
		CaptchaLen:             4,
		BgHashColor:            "",
		BgColor: &color.RGBA{
			R: 3,
			G: 102,
			B: 214,
			A: 254,
		},
	}
	captchaId, instance := base64Captcha.GenerateCaptcha("", parameters)
	base64Blob := base64Captcha.CaptchaWriteToBase64Encoding(instance)
	return CaptchaResult{Id: captchaId, Base64Blob: base64Blob}
}

func VerifyCaptcha(id string, value string) bool {
	return base64Captcha.VerifyCaptcha(id, value)
}
