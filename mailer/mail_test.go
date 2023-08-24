package mailer

import "testing"

func TestMail_SendSMTPMessage(t *testing.T) {
	msg := Message{
		From: "me@here.com",
		FromName: "Joe",
		To: "you@there.com",
		Subject: "Test",
		Template: "test",
		Attachments: []string{"./testdata/mail/test.html.tmpl"},
		Data: nil,
	}

	err := mailer.SendSMTPMessage(msg)
	if err != nil {
		t.Errorf("SendSMTPMessage() error = %v", err)
	}
}
