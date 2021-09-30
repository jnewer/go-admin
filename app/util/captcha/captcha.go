package captcha

import (
	"fmt"
	"github.com/mojocn/base64Captcha"
	e2 "pear-admin-go/app/global/e"
)

// 设置自带的store
var store = base64Captcha.DefaultMemStore

//生成验证码
func CaptMake() (id, b64s string, err error) {
	var driver base64Captcha.Driver
	var driverString base64Captcha.DriverString

	// 配置验证码信息
	captchaConfig := base64Captcha.DriverString{
		Height:          e2.ImgHeight,
		Width:           e2.ImgWidth,
		NoiseCount:      0,     // 干扰字母
		ShowLineOptions: 1 | 3, // 干扰线
		Length:          e2.ImgKeyLength,
		Source:          "qwertyuioplkjhgfdsazxcvbnm",
	}

	driverString = captchaConfig
	driver = driverString.ConvertFonts()
	captcha := base64Captcha.NewCaptcha(driver, store)
	lid, lb64s, lerr := captcha.Generate()
	return lid, lb64s, lerr
}

//验证captcha是否正确
func CaptVerify(id string, capt string) bool {
	fmt.Println("id:" + id)
	fmt.Println("capt:" + capt)
	if store.Verify(id, capt, false) {
		return true
	} else {
		return false
	}
}
