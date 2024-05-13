package service

import (
	"archive/tar"
	"compress/gzip"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/Aracelimartinez/email-platform-challenge/server/internal/model"
)

func IndexerProcessor() (chan []*model.Email, chan error) {
	//Start email processing
	users, err := getUsers()
	if err != nil {
		errChan := make(chan error, 1)
		errChan <- fmt.Errorf("failed to obtain usernames: %w", err)
		close(errChan)
		return nil, errChan
	}

	emailChan := make(chan []*model.Email, len(*users))
	errorChan := make(chan error)
	go func() {
		var wg sync.WaitGroup

		for _, user := range *users {
			wg.Add(1)
			go func(user string) {
				defer wg.Done()
				log.Printf("Processing emails from %s... \n", user)

				userEmails, err := extractEmailsByUser(user)
				if err != nil {
					errorChan <- fmt.Errorf("failed to extract emails for user %s: %s", user, err)
					return
				}
				emailChan <- userEmails
			}(user)
		}
		wg.Wait()
		close(emailChan)
		close(errorChan)
	}()

	return emailChan, errorChan
}

func DownloadProcessor() error {
	//Verify if the dataset folder exists
	if checkIfNotExists(model.EmailDataSetRoot) {

		//Download dataset for indexing
		log.Println("The dataset doesn't exists")
		log.Println("Downloading dataset...")

		fileURL := "https://www.cs.cmu.edu/~enron/enron_mail_20150507.tar.gz"
		destFolder := "/tmp"
		filePath, err := downloadData(fileURL, destFolder)
		if err != nil {
			return fmt.Errorf("failed to download file: %w", err)
		}

		//Decompress file
		log.Println("Decompressing files...")
		if err = decompressTarGz(filePath, destFolder); err != nil {
			return fmt.Errorf("failed to descompress the files: %w", err)
		}
	}
	return nil
}

// checkIfNotExists checks if a folder exists and returns a boolean.
func checkIfNotExists(path string) bool {
	_, err := os.Stat(path)
	return errors.Is(err, os.ErrNotExist)
}

// downloads a file from a URL.
func downloadData(url, destFolder string) (string, error) {
	splitUrl := strings.Split(url, "/")
	filename := splitUrl[len(splitUrl)-1]

	downloadPath := filepath.Join(destFolder, filename)

	out, err := os.Create(downloadPath)
	if err != nil {
		return "", err
	}
	defer out.Close()

	response, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	_, err = io.Copy(out, response.Body)
	if err != nil {
		return "", err
	}

	return downloadPath, nil
}

// decompresses a .tgz file
func decompressTarGz(src, dest string) error {
	file, err := os.Open(src)
	if err != nil {
		return err
	}
	defer file.Close()

	gzipReader, err := gzip.NewReader(file)
	if err != nil {
		return err
	}
	defer gzipReader.Close()

	tarReader := tar.NewReader(gzipReader)

	for {
		header, err := tarReader.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		targetPath := filepath.Join(dest, header.Name)
		switch header.Typeflag {
		case tar.TypeDir:
			if err := os.MkdirAll(targetPath, 0755); err != nil {
				return err
			}
		case tar.TypeReg:
			outFile, err := os.OpenFile(targetPath, os.O_CREATE|os.O_RDWR, os.FileMode(header.Mode))
			if err != nil {
				return err
			}
			if _, err := io.Copy(outFile, tarReader); err != nil {
				outFile.Close()
				return err
			}
			outFile.Close()
		}
	}
	return nil
}

// deletes the enron dataset
func RemoveAllData(dir string) error {
	return filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if path != dir {
			err := os.RemoveAll(path)
			if err != nil {
				return err
			}
		}
		return nil
	})
}
