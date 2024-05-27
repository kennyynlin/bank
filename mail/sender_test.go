package mail

import (
	"github.com/kennyynlin/bank/util"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSendEmailGmail(t *testing.T) {
	if testing.Short() {
		t.Skip("skip Gmail test in short mode")
	}
	config, err := util.LoadConfig("..")
	require.NoError(t, err)

	sender := NewGmailSender(config.EmailSenderName, config.EmailSenderAddress, config.EmailSenderPassword)
	subject := "A test email"
	content := `
	<h1>Hello World! </h1>
	<p> A test email </p>
	`
	to := []string{"ynlin1996@icloud.com"}
	attachFiles := []string{"../README.md"}

	err = sender.SendEmail(subject, content, to, nil, nil, attachFiles)
	require.NoError(t, err)
}
