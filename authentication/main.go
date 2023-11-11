package authentication

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	mathRand "math/rand"
	"net/smtp"
	"parties-app/backend/config"
	"time"
)

var (
	OtpClass Otp
)

func init() {
	OtpClass = NewOtp()
}

func LogInRequest(email string) error {
	code := generateEmailAuthenticationCode()

	OtpClass.Register(email, code)

	err := sendVerficationEmail(code, email)
	if err != nil {
		return err
	}

	return nil
}

func VeritifyEmail(code string, email string) error {
	otp, found := OtpClass.Get(email)
	if !found {
		return errors.New("0")
	}

	if otp.code == code {
		return nil
	} else {
		return errors.New("1")
	}
}

func GenerateRandomSalt(saltSize int) string {
	var salt = make([]byte, saltSize)

	_, err := rand.Read(salt[:])

	if err != nil {
		panic(err)
	}

	return base64.StdEncoding.EncodeToString(salt)
}

func sendVerficationEmail(code string, email string) error {
	from := config.SMTP_USERNAME
	password := config.SMTP_PASSWORD
	toList := []string{email}
	host := "sandbox.smtp.mailtrap.io"
	port := "2525"
	msg := fmt.Sprintf("Date: %v\nFrom: Parties Noreply <noreply@parties-test.thinh.tech>\nTo: User <%v>\nSubject: Your verification code!\n\nThis is your verification code %v", getCurrentDate(), email, code)
	body := []byte(msg)
	auth := smtp.PlainAuth("", from, password, host)
	err := smtp.SendMail(host+":"+port, auth, from, toList, body)
	if err != nil {
		return err
	}
	return nil
}

func generateEmailAuthenticationCode() string {
	low := 1000
	hi := 9999
	generatedCode := fmt.Sprintf("%v", low+mathRand.Intn(hi-low))
	return generatedCode
}

func getCurrentDate() string {
	// Get the current date and time
	currentTime := time.Now()

	// Format the current date and time using RFC2822 format
	rfc2822Format := currentTime.Format(time.RFC1123Z)

	return rfc2822Format
}
