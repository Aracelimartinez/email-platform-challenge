package handler

import (
	"log"
	"net/http"
	"runtime"
	"sync"

	"github.com/Aracelimartinez/email-platform-challenge/server/internal/model"
	"github.com/Aracelimartinez/email-platform-challenge/server/internal/response"
	"github.com/Aracelimartinez/email-platform-challenge/server/internal/service"
	"github.com/Aracelimartinez/email-platform-challenge/server/internal/service/zincsearch"
)

func IndexEmails(w http.ResponseWriter, r *http.Request) {
	//Download processor
	err := service.DownloadProcessor()
	if err != nil {
		response.Err(w, http.StatusBadRequest, err)
		return
	}

	// indexer processor
	log.Println("Starting file processing...")
	extractedEmails, errChan := service.IndexerProcessor()

	// zincsearch service indexer
	log.Println("Indexing users emails...")
	var wg sync.WaitGroup
	maxWorkers := runtime.NumCPU()
	log.Println(maxWorkers)
	workerSem := make(chan struct{}, maxWorkers)
	indexingErrChan := make(chan error)
	go func() {
		for userEmails := range extractedEmails {
			wg.Add(1)
			workerSem <- struct{}{}
			go func(userEmails []*model.Email) {
				defer func() {
					<-workerSem
					wg.Done()
				}()
				res, err := zincsearch.CreateDocument(model.EmailIndexName, userEmails)
				if err != nil {
					indexingErrChan <- err
					return
				}
				log.Printf("Indexed %d documents\n", res.RecordCount)
			}(userEmails)

		}
		wg.Wait()
		close(indexingErrChan)
	}()

	// Handle errors from the IndexerProcessor
	if err := <-errChan; err != nil {
		response.Err(w, http.StatusBadRequest, err)
		return
	}

	// Handle any indexing errors
	if err := <-indexingErrChan; err != nil {
		response.Err(w, http.StatusInternalServerError, err)
		return
	}

	// indexer processor finish by deleting unnecessary files
	log.Println("Removing unnecesary files..")
	err = service.RemoveAllData("/tmp/")
	if err != nil {
		log.Println("unnable to delete the directory data: ", err)
	}

	response.JSON(w, http.StatusCreated, "Emails indexed succesfully")
}
