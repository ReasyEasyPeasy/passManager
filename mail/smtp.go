package mail

import (
	"github.com/spf13/viper"
	"log"
	"crypto/tls"
	"net/mail"
)
var Ready = true



var smtpHost string


func init()  {
	if !viper.IsSet("SMTPHost") {
		log.Println("No smtp-server set.")
		Ready = false
	}
	if !viper.IsSet("SMTPsPort") {
		log.Println("No smtp-port set. Using default 465")
		viper.Set("SMTPsPort",465)
	}

	smtpHost = viper.GetString("SMTPHost")

}

func SendMail() {


	from := mail.Address{"", "username@example.tld"}
	to   := mail.Address{"", "username@anotherexample.tld"}
	
	subj := "This is the email subject"
	body := "This is an example body.\n With two lines."

	// Setup headers
	headers := make(map[string]string)
	headers["From"] = from.String()
	headers["To"] = to.String()
	headers["Subject"] = subj

	tlsconfig := &tls.Config {
		InsecureSkipVerify: false,
		ServerName: smtpHost,
	}
}