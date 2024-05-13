package model

import "time"

const (
	// EmailDataSetRoot = "/tmp/maildir"
	EmailDataSetRoot = "enron_mail_20150507/maildir"
	EmailIndexName   = "emails"
)

// var emailHeaders = [15]string{"Message-ID", "Date", "From", "To", "Subject", "Mime-Version", "Content-Type", "Content-Transfer-Encoding", "X-From", "X-To", "X-cc", "X-bcc", "X-Folder", "X-Origin", "X-FileName"}

type Email struct {
	MessageID   string    `json:"message_id"`
	Date        time.Time `json:"date"`
	From        string    `json:"from"`
	To          []string  `json:"to"`
	Cc          []string  `json:"cc"`
	Bcc         []string  `json:"bcc"`
	Subject     string    `json:"subject"`
	ContentType string    `json:"content_type"`
	Body        string    `json:"body"`
}

//Example of the email
// Message-ID: <18782981.1075855378110.JavaMail.evans@thyme>
// Date: Mon, 14 May 2001 16:39:00 -0700 (PDT)
// From: phillip.allen@enron.com
// To: tim.belden@enron.com
// Subject:
// Mime-Version: 1.0
// Content-Type: text/plain; charset=us-ascii
// Content-Transfer-Encoding: 7bit
// X-From: Phillip K Allen
// X-To: Tim Belden <Tim Belden/Enron@EnronXGate>
// X-cc:
// X-bcc:
// X-Folder: \Phillip_Allen_Jan2002_1\Allen, Phillip K.\'Sent Mail
// X-Origin: Allen-P
// X-FileName: pallen (Non-Privileged).pst

// Here is our forecast
