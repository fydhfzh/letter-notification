package mailer

import (
	"crypto/tls"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/fydhfzh/letter-notification/dto"
	"github.com/fydhfzh/letter-notification/repository/letter_repository"
	gomail "gopkg.in/mail.v2"
)

var configSmtpHost = os.Getenv("CONFIG_SMTP_HOST")
var configSmtpPort = os.Getenv("CONFIG_SMTP_PORT")
var configAuthEmail = os.Getenv("CONFIG_AUTH_EMAIL")
var configAuthPassword = os.Getenv("CONFIG_AUTH_PASSWORD")

func schedule(letterID int, recipient []string, subject string, message string, mailTime time.Time, letterRepo letter_repository.LetterRepository) {
	time.Sleep(time.Until(mailTime))
	sendMail(recipient, subject, message)

	err := letterRepo.SetIsNotifyTrue(letterID)

	if err != nil {
		fmt.Printf("Mail is not sent")
	}

	time.Sleep(10 * time.Second)
}

func sendMail(recipient []string, subject string, message string) {
	mail := gomail.NewMessage()

	mail.SetHeader("From", configAuthEmail)
	mail.SetHeader("To", recipient...)
	mail.SetHeader("Subject", subject)
	mail.SetBody("text/plain", message)

	smtpPort, err := strconv.Atoi(configSmtpPort)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	d := gomail.NewDialer(configSmtpHost, smtpPort, configAuthEmail, configAuthPassword)

	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	if err := d.DialAndSend(mail); err != nil {
		fmt.Println(err.Error())
		return
	}

}

func SetSchedule(scheduler dto.SendLetterToMailScheduler, letterRepo letter_repository.LetterRepository) {
	var recipient []string

	for _, user := range scheduler.Recipients {
		recipient = append(recipient, user.Email)
	}

	subject := "Reminder Kegiatan " + scheduler.About

	zone, _ := scheduler.Datetime.Zone()

	message := fmt.Sprintf("Halo Bapak/Ibu\n\nJangan lupa untuk hadir pada kegiatan %s pada tanggal %d %s %d pukul %d:%d %s. \n\nTerima kasih", scheduler.About, scheduler.Datetime.Day(), scheduler.Datetime.Month(), scheduler.Datetime.Year(), scheduler.Datetime.Hour(), scheduler.Datetime.Minute(), zone)

	mailTime := scheduler.Datetime.Add(-24 * time.Hour)

	go schedule(scheduler.LetterID, recipient, subject, message, mailTime, letterRepo)
}
