// mail.go
package gostuff

import (
	"bufio"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"gopkg.in/gomail.v2"
	"log"
	"os"
)

//returns true if sucessful sent email
func Sendmail(target string, token string, name string) {

	m := gomail.NewMessage()
	m.SetHeader("From", "goplaychess@gmail.com", "Go Play Chess")
	m.SetHeader("To", target)
	m.SetHeader("Subject", "Welcome to Go Play Chess!")
	message := "Hello " + name + ",<br><br>Welcome to <b>Go Play Chess</b>! " +
		"<a href='https://localhost/activate?user=" + name + "&token=" + token + "'>Please click here to activate your account.</a>" +
		" Your token is " + token +
		"<br><br>Have fun!<br><br>GoPlayChess"
	m.SetBody("text/html", message)

	answer := mailConfig()

	d := gomail.NewPlainDialer("smtp.gmail.com", 587, "goplaychess@gmail.com", answer)

	if err := d.DialAndSend(m); err != nil {
		fmt.Println("error in Sendmail mail.go ", err)
	}
}

//this function is used to send player the activation token when five incorrect login attempts are made
func SendAttempt(target string, token string, name string, ip string) {
	m := gomail.NewMessage()
	m.SetHeader("From", "goplaychess@gmail.com", "Go Play Chess")
	m.SetHeader("To", target)
	m.SetHeader("Subject", "Go Play Chess Account Locked")
	message := "Hello " + name + ",<br><br>Your account on <b>Go Play Chess</b> has been locked because " +
		"there was at least five incorrect login attempts. The IP that tried to login your account was " + ip + "<br><br>Your reactivation token is " + token +
		"<br><br><a href='https://localhost/activate?user=" + name + "&token=" + token + "'>Please click here to activate your account.</a>" +
		"<br><br>Please reactivate your account.<br><br>GoPlayChess"
	m.SetBody("text/html", message)

	answer := mailConfig()

	d := gomail.NewPlainDialer("smtp.gmail.com", 587, "goplaychess@gmail.com", answer)

	if err := d.DialAndSend(m); err != nil {
		fmt.Println("error in SendAttempt mail.go ", err)
	}
}

//sends email to user of a token to reset his password
func SendForgot(target string, token string) {
	m := gomail.NewMessage()
	m.SetHeader("From", "goplaychess@gmail.com", "Go Play Chess")
	m.SetHeader("To", target)
	m.SetHeader("Subject", "Reset Password for Go Play Chess")
	message := "<a href='https://localhost/resetpass?token=" + token + "'>Please click here to type your token code and reset your password.</a>" +
		"Your token to reset your pass is : " + token +
		"<br><br>GoPlayChess"
	m.SetBody("text/html", message)

	answer := mailConfig()

	d := gomail.NewPlainDialer("smtp.gmail.com", 587, "goplaychess@gmail.com", answer)

	if err := d.DialAndSend(m); err != nil {
		fmt.Println("error in SendForgot mail.go ", err)
	}
}

//fetches pass for email account
func mailConfig() string {

	problem, err := os.OpenFile("logs/errors.txt", os.O_APPEND|os.O_WRONLY, 0666)
	defer problem.Close()
	log.SetOutput(problem)
	readFile, err := os.Open("secret/mailpass.txt")
	defer readFile.Close()
	if err != nil {
		fmt.Println("mailconfig mail.go ", err)
	}

	scanner := bufio.NewScanner(readFile)

	scanner.Scan()
	pass := scanner.Text()
	//decode
	ans, _ := hex.DecodeString(pass)
	result, _ := base64.StdEncoding.DecodeString(string(ans))
	answer := string(result)

	return answer
}
