package service

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/Aracelimartinez/email-platform-challenge/server/internal/model"
	"github.com/Aracelimartinez/email-platform-challenge/server/internal/service/zincsearch"
)

// Get the names of the users' folders
func getUsers() (*[]string, error) {
	var users []string

	path := filepath.Join(model.EmailDataSetRoot)

	entries, err := os.ReadDir(path)
	if err != nil {
		return nil, fmt.Errorf("failed to obtain the name of the user's folder: %w", err)
	}

	for _, e := range entries {
		users = append(users, e.Name())
	}

	return &users, nil
}

// Walk through the user's directory to map every email
func extractEmailsByUser(user string) ([]*model.Email, error) {
	var emails []*model.Email
	var err error
	var mtx sync.Mutex
	path := filepath.Join(model.EmailDataSetRoot, user)

	err = filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return fmt.Errorf("failed to go through the files and directories: %w", err)
		}

		if !info.IsDir() {
			// Ejecutar la función ProcessEmail para cada archivo
			mtx.Lock()
			email, err := processEmail(path)
			if err != nil {
				return fmt.Errorf("failed to process the email in the path '%s': %v", path, err)
			} else {
				// Agregar el email al slice de emails
				emails = append(emails, email)
			}
			mtx.Unlock()
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return emails, nil
}

// Read the email files and process it into an Email struct
func processEmail(emailPath string) (*model.Email, error) {

	// Lee el contenido del archivo
	content, err := os.ReadFile(emailPath)
	if err != nil {
		return nil, fmt.Errorf("failed to reading the file: %w", err)
	}

	// Convierte el contenido a string y separa el email en 2 partes
	parts := strings.SplitN(string(content), "\r\n\r\n", 2)

	email, err := mapEmail(parts[0], parts[1])
	if err != nil {
		return nil, err
	}

	return email, nil
}

// Map the content of the file into a email struct
func mapEmail(headers, body string) (*model.Email, error) {
	email := model.Email{}
	email.Body = body
	processedHeader := preprocessHeaders(headers)
	lines := strings.Split(processedHeader, "\r\n")

	for _, line := range lines {

		if colonIndex := strings.Index(line, ":"); colonIndex != -1 {
			key := line[:colonIndex]
			value := strings.TrimSpace(line[colonIndex+1:])

			switch key {
			case "Message-ID":
				email.MessageID = value
			case "Date":
				var err error
				email.Date, err = parseEmailDate(value)
				if err != nil {
					return nil, fmt.Errorf("invalid date format: %v", err)
				}
			case "From":
				email.From = value
			case "To":
				email.To = processEmailAddress(value)
			case "Cc":
				email.Cc = processEmailAddress(value)
			case "Bcc":
				email.Bcc = processEmailAddress(value)
			case "Subject":
				email.Subject = value
			case "Content-Type":
				email.ContentType = value
			}
		}
	}
	return &email, nil
}

// Pre-process the headers to avoid data losses
func preprocessHeaders(rawHeaders string) string {
	lines := strings.Split(rawHeaders, "\r\n")
	var processedLines []string

	for _, line := range lines {

		if len(line) > 0 && (line[0] == ' ' || line[0] == '\t') {
			if len(processedLines) > 0 {
				processedLines[len(processedLines)-1] += " " + strings.TrimSpace(line)
			}
		} else {
			processedLines = append(processedLines, line)
		}
	}
	return strings.Join(processedLines, "\r\n")
}

func processEmailAddress(emails string) []string {
	var processedEmails []string
	splitEmails := strings.Split(emails, ", ")
	for _, emailAddress := range splitEmails {
		emailAddress = strings.TrimSpace(emailAddress)
		processedEmails = append(processedEmails, emailAddress)
	}
	return processedEmails
}

// Parse the emails date
func parseEmailDate(dateStr string) (time.Time, error) {
	if idx := strings.LastIndex(dateStr, " ("); idx != -1 {
		dateStr = dateStr[:idx]
	}
	const layout = "Mon, 2 Jan 2006 15:04:05 -0700"
	return time.Parse(layout, dateStr)
}

// MapZincSearchEmails response into an email structure
func MapZincSearchEmails(zincSearchResponse *zincsearch.SearchDocumentsRsponse) ([]*model.Email, error) {
	emails := make([]*model.Email, 0, len(zincSearchResponse.Hits.Hits))

	for _, hit := range zincSearchResponse.Hits.Hits {
		var email model.Email

		sourceBytes, err := json.Marshal(hit.Source)
		if err != nil {
			fmt.Println("Error in mapping the emails: ", err)
			continue
		}

		err = json.Unmarshal(sourceBytes, &email)
		if err != nil {
			fmt.Println("Error in mapping the emails: ", err)
			continue
		}

		emails = append(emails, &email)
	}

	return emails, nil
}
