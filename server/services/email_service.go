package services

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

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
			email, err := processEmail(path)
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
func processEmail(emailPath string) (*models.Email, error) {

	// Lee el contenido del archivo
	content, err := os.ReadFile(emailPath)
	if err != nil {
		return nil, fmt.Errorf("failed to reading the file: %w\n", err)
	}

	// Convierte el contenido a string y separa el email en 2 partes
	lines := strings.SplitN(string(content), "\r\n\r\n", 2)

	return mapEmail(lines), nil
}

// Map the content of the file into a email struct
func mapEmail(lines []string) *models.Email {
	email := models.Email{}
	detailsLines := strings.SplitAfter(string(lines[0]), "\n")

	for _, line := range detailsLines {

		if strings.HasPrefix(line, "Message-ID:") {
			email.MessageID = strings.TrimPrefix(line, "Message-ID: ")
		} else if strings.HasPrefix(line, "Date:") {
			email.Date = strings.TrimPrefix(line, "Date: ")
		} else if strings.HasPrefix(line, "From:") {
			email.From = strings.TrimPrefix(line, "From: ")
		} else if strings.HasPrefix(line, "To:") {
			email.To = strings.TrimPrefix(line, "To: ")
		} else if strings.HasPrefix(line, "Subject:") {
			email.Subject = strings.TrimPrefix(line, "Subject: ")
		} else if strings.HasPrefix(line, "Content-Type:") {
			email.ContentType = strings.TrimPrefix(line, "Content-Type: ")
		} else {
			continue
		}
	}
	email.Body = lines[1]

	return &email
}

// MapZincSearchEmails response into an email structure
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
