package app

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"strings"

	api2captcha "github.com/2captcha/2captcha-go"
	"github.com/cavaliergopher/grab/v3"
)

type captchaurl struct {
	Error struct {
		Url string `json:"captcha_img"`
		Sid string `json:"captcha_sid"`
	}
}

func random() string {

	var numbers = []rune("0123456789")

	numb := make([]rune, 10)
	for i := range numb {
		numb[i] = numbers[rand.Intn(len(numbers))]
	}
	return string(numb)
}
func DownloadCaptcha(url string) (string, error) {

	name := random()
	resp, err := grab.Get("./temp", url)
	if err != nil {
		fmt.Println("Captcha Error")
		os.Exit(0)
		return "", nil
	}

	os.Rename(resp.Filename, "./temp/captcha"+name+".jpg")

	return "./temp/captcha" + name + ".jpg", nil

}

func SolveCaptcha(captchaurl string) (string, error) {

	path, err := DownloadCaptcha(captchaurl)
	if err != nil {
		fmt.Println("Captcha Error")
		os.Exit(0)
	}

	client := api2captcha.NewClient(Anticaptcha)

	cap := api2captcha.Normal{
		File: path,
	}

	code, err := client.Solve(cap.ToRequest())
	if err != nil {
		return "123", nil
	}

	return code, nil

}

var FindAntiCaptcha captchaurl

func Findcaptcha(liness []byte, url string) {

	lines := strings.Split(string(liness), "\n\r")
	for _, line := range lines {
		if strings.Contains(line, "Captcha needed") {
			json.Unmarshal(liness, &FindAntiCaptcha)

			code, err := SolveCaptcha(FindAntiCaptcha.Error.Url)
			if err != nil {
				return
			}

			newurl := url + "&captcha_sid=" + FindAntiCaptcha.Error.Sid + "&captcha_key=" + code

			solved++
			Request(newurl)
		}
	}
}
