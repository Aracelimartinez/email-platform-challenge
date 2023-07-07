package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/mail"
	"os"
	"path/filepath"

	"github.com/Aracelimartinez/email-platform-challenge/server/models"
	"github.com/Aracelimartinez/email-platform-challenge/server/services/zincsearch"
)

// Get the names of the users' folders
func GetUsers() (*[]string, error) {
	var users []string

	path := filepath.Join(models.EmailDataSetRoot)

	entries, err := os.ReadDir(path)
	if err != nil {
		return nil, fmt.Errorf("failed to obtain the name of the user's folder: %w\n", err)
	}

	for _, e := range entries {
		users = append(users, e.Name())
	}

	return &users, nil
}

// Walk through the user's directory to map every email
func ExtractEmailsByUser(user string) ([]*models.Email, error) {
	var emails []*models.Email
	var err error
	path := filepath.Join(models.EmailDataSetRoot, user)

	err = filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return fmt.Errorf("failed to go through the files and directories: %w\n", err)
		}

		// Verificar si no es una carpeta
		if !info.IsDir() {
			// Ejecutar la función ProcessEmail para cada archivo
			email, err := processEmail(&path)
			if err != nil {
				return fmt.Errorf("failed to process the email in the path '%s': %v\n", path, err)
			} else {
				// Agregar el email al slice de emails
				emails = append(emails, email)
			}
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return emails, nil
}

// Read the email files and process it into an Email struct
func processEmail(emailPath *string) (*models.Email, error) {
	email := models.Email{}

	// Lee el contenido del archivo
	content, err := os.ReadFile(*emailPath)
	if err != nil {
		return nil, fmt.Errorf("failed to reading the file: %w\n", err)
	}

	r := bytes.NewReader(content)
	m, err := mail.ReadMessage(r)
	if err != nil {
		return nil, fmt.Errorf("failed to reading the email file: %w\n", err)
	}
	email.MessageID = m.Header.Get("Message-ID:")
	email.Date = m.Header.Get("Date:")
	email.From = m.Header.Get("From:")
	email.To = m.Header.Get("To:")
	email.Subject = m.Header.Get("Subject:")
	email.ContentType = m.Header.Get("Content-Type:")
	body, err := io.ReadAll(m.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to reading the body: %w\n", err)
	}
	email.Body = string(body)

	return &email, nil
}

//MapZincSearchEmails response into an email structure
func MapZincSearchEmails(zincSearchResponse *zincsearch.SearchDocumentsRsponse) ([]*models.Email, error) {
	emails := make([]*models.Email, 0, len(zincSearchResponse.Hits.Hits))

	for _, hit := range zincSearchResponse.Hits.Hits {
		var email models.Email

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
