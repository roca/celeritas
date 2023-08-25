package mailer

import "testing"

func TestMail_SendSMTPMessage(t *testing.T) {
	msg := Message{
		From:        "me@here.com",
		FromName:    "Joe",
		To:          "you@there.com",
		Subject:     "Test",
		Template:    "test",
		Attachments: []string{"./testdata/mail/test.html.tmpl"},
		Data:        nil,
	}

	err := mailer.SendSMTPMessage(msg)
	if err != nil {
		t.Errorf("SendSMTPMessage() error = %v", err)
	}
}

func TestMail_SendUsingChan(t *testing.T) {
	msg := Message{
		From:        "me@here.com",
		FromName:    "Joe",
		To:          "you@there.com",
		Subject:     "Test",
		Template:    "test",
		Attachments: []string{"./testdata/mail/test.html.tmpl"},
		Data:        nil,
	}

	mailer.Jobs <- msg
	res := <-mailer.Results
	if res.Error != nil {
		t.Errorf("SendUsingChan() error = %v", res.Error)
	}

	msg.To = "not_an_email_address"
	mailer.Jobs <- msg
	res = <-mailer.Results
	if res.Error == nil {
		t.Errorf("SendUsingChan() error = %v", res.Error)
	}
}

func TestMail_SendUsingAPI(t *testing.T) {
	msg := Message{
		To:          "you@there.com",
		Subject:     "Test",
		Template:    "test",
		Attachments: []string{"./testdata/mail/test.html.tmpl"},
		Data:        nil,
	}

	mailer.API = "unknown"
	mailer.APIKey = "abc123"
	mailer.APIUrl = "http://www.fake.com"

	err := mailer.SendUsingAPI(msg, mailer.API)
	if err == nil {
		t.Errorf("SendUsingAPI() error = %v", err)
	}
	mailer.API = ""
	mailer.APIKey = ""
	mailer.APIUrl = ""
}

func TestMail_buildHTMLMessage(t *testing.T) {
	msg := Message{
		From:        "me@here.com",
		FromName:    "Joe",
		To:          "you@there.com",
		Subject:     "Test",
		Template:    "test",
		Attachments: []string{"./testdata/mail/test.html.tmpl"},
		Data:        nil,
	}

	html, err := mailer.buildHTMLMessage(msg)
	if err != nil {
		t.Errorf("buildHTMLMessage() error = %v", err)
	}

	if html == "" {
		t.Errorf("buildHTMLMessage() html = %v", html)
	}
}

func TestMail_buildPlainTextMessage(t *testing.T) {
	msg := Message{
		From:        "me@here.com",
		FromName:    "Joe",
		To:          "you@there.com",
		Subject:     "Test",
		Template:    "test",
		Attachments: []string{"./testdata/mail/test.html.tmpl"},
		Data:        nil,
	}

	plainText, err := mailer.buildPlainTextMessage(msg)
	if err != nil {
		t.Errorf("buildHTMLMessage() error = %v", err)
	}

	if plainText == "" {
		t.Errorf("buildHTMLMessage() html = %v", plainText)
	}
}

func TestMail_Send(t *testing.T) {
	msg := Message{
		From:        "me@here.com",
		FromName:    "Joe",
		To:          "you@there.com",
		Subject:     "Test",
		Template:    "test",
		Attachments: []string{"./testdata/mail/test.html.tmpl"},
		Data:        nil,
	}

	err := mailer.Send(msg)
	if err != nil {
		t.Errorf("Send() error = %v", err)
	}

	mailer.API = "unknown"
	mailer.APIKey = "abc123"
	mailer.APIUrl = "http://www.fake.com"

	err = mailer.Send(msg)
	if err == nil {
		t.Errorf("Send() error = %v", err)
	}
	mailer.API = ""
	mailer.APIKey = ""
	mailer.APIUrl = ""
}

func TestMail_ChooseAPI(t *testing.T) {
	msg := Message{
		From:        "me@here.com",
		FromName:    "Joe",
		To:          "you@there.com",
		Subject:     "Test",
		Template:    "test",
		Attachments: []string{"./testdata/mail/test.html.tmpl"},
		Data:        nil,
	}

	mailer.API = "unknown"
	err := mailer.ChooseAPI(msg)
	if err == nil {
		t.Errorf("ChooseAPI() error = %v", err)
	}
	mailer.API = ""
}
