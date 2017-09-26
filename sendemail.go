package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"gopkg.in/gomail.v2"
	"log"
	"os"
	"strconv"
	"time"
)

var (
	HOST     string
	PORT     int
	USER     string
	PASSWORD string
	TO       string
	ATTACH   string
	FILENAME string
	NAME     string
)

func init() {
	date := time.Now().Format("2006-01-02")
	err := godotenv.Load("email.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	HOST = os.Getenv("HOST")
	PORT, _ = strconv.Atoi(os.Getenv("PORT"))
	USER = os.Getenv("USER")         //发送邮件的邮箱
	PASSWORD = os.Getenv("PASSWORD") //发送邮件邮箱的密码
	TO = os.Getenv("TO")
	NAME = os.Getenv("NAME")
	FILENAME = "周报-" + NAME + "-" + date + ".doc"
	ATTACH = os.Getenv("ATTACH") + "\\" + FILENAME
}

func main() {
	s := `
//
//                       _oo0oo_
//                      o8888888o
//                      88" . "88
//                      (| -_- |)
//                      0\  =  /0
//                    ___/'---'\___
//                  .' \\|     |// '.
//                 / \\|||  :  |||// \
//                / _||||| -:- |||||- \
//               |   | \\\  -  /// |   |
//               | \_|  ''\---/''  |_/ |
//               \  .-\__  '-'  ___/-. /
//             ___'. .'  /--.--\  '. .'___
//          ."" '<  '.___\_<|>_/___.' >' "".
//         | | :  '- \'.;'\ _ /';.'/ - '' : | |
//         \  \ '_.   \_ __\ /__ _/   .-'' /  /
//     ====='-.____'.___ \_____/___.-'___.-'=====
//                       '=---='
//
//     
//     ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
//
//               佛祖保佑         永无BUG
//
//     Copyright © 2017 - 2017  YaoBin All Rights Reserved
//`
	fmt.Println(s)

	m := gomail.NewMessage()
	m.SetHeader("From", USER)
	_, err := os.Stat(ATTACH)
	if err != nil {
		// no such file or dir
		m.SetHeader("To", USER)
		m.SetHeader("Subject", "主人，主人！！！")
		m.SetBody("text/html", "您忘了写周报了/(ㄒoㄒ)/~~")

		d := gomail.NewDialer(HOST, PORT, USER, PASSWORD)

		// Send the email to Bob, Cora and Dan.
		if err := d.DialAndSend(m); err != nil {
			panic(err)
		}
	} else {
		m.SetHeader("To", TO)
		// m.SetAddressHeader("Cc", "dan@example.com", "Dan")
		// m.SetHeader("Subject", "Hello!")
		// m.SetBody("text/html", "Hello <b>Bob</b> and <i>Cora</i>!")
		m.SetHeader("Subject", "周报")
		h := make(map[string][]string, 0)
		h["Content-Type"] = []string{`application/octet-stream; charset=utf-8; name="` + FILENAME + `"`} //要设置这个，否则中文会乱码
		fileSetting := gomail.SetHeader(h)
		m.Attach(ATTACH, fileSetting)

		d := gomail.NewDialer(HOST, PORT, USER, PASSWORD)

		if err := d.DialAndSend(m); err != nil {
			panic(err)
		}

		n := gomail.NewMessage()
		n.SetHeader("From", USER)
		n.SetHeader("To", USER)
		n.SetHeader("Subject", "主人，主人！！！")
		n.SetBody("text/html", "您的周报发送成功了O(∩_∩)O~")

		// Send the email to Bob, Cora and Dan.
		if err := d.DialAndSend(n); err != nil {
			panic(err)
		}
		fmt.Println("发送成功")
	}
}
