package main

import (
	"fmt"
	"github.com/labstack/echo"
	"gopkg.in/gomail.v2"
	"math/rand"
	"time"
)

type DataCodeEmail struct {
	Greet string
}

type MsgSuccessCodeEmail struct {
	Success bool
	Data DataCodeEmail
}



func codeEmail(text echo.Context) error {

	getEmail := text.FormValue("email")

	sendEmail(getEmail)

	successJson := createSuccessJsonCodeEmail()

	return text.JSON(200,successJson)

}


func sendEmail(email string)  {

	mail := gomail.NewMessage()

	mail.SetHeader("From","CA15546466860@126.com")

	mail.SetHeader("To",email)

	mail.SetHeader("Subject","code")

	code := random()

	mail.SetBody("text/html",code)

	agent := gomail.NewDialer("smtp.126.com", 25, "CA15546466860@126.com", "TTUHLTNQWFXVZNMI")

	err := agent.DialAndSend(mail)
	if err != nil{
		fmt.Println(err.Error())
	}


	redisSave(code,email)

	}

func random() string {
	rand.Seed(time.Now().UnixNano())
	code := rand.Intn(900000)+100000
	neocode := fmt.Sprintf("%d", code)
	fmt.Println(neocode)
	return neocode
}

func createSuccessJsonCodeEmail() MsgSuccessCodeEmail {
	successData := new(DataCodeEmail)
	successData.Greet = "send already"
	successJson := new(MsgSuccessCodeEmail)
	successJson.Success = true
	successJson.Data = *successData
	return *successJson
}

func judgeCode(getCode string,code string) bool {
	judge := getCode==code
	return judge
}