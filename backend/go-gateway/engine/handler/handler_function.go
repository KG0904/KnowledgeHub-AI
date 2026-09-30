package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"go-gateway/engine/common"
	"go-gateway/engine/db/psqldb"
	"log"
	"mime/multipart"
	"net/http"
	"os"
)

func SaveMetaDataInPostgres(file *multipart.FileHeader, filePath string) (int, error) {

	fileInfo, err := os.Stat(filePath)
	if err != nil {
		log.Println(err)
		return 0, err
	}

	query := `
	INSERT INTO documents (
		owner_type,
		owner_id,
		filename,
		storage_path,
		file_size,
		mime_type,
		status
	)
	VALUES ($1,$2,$3,$4,$5,$6,$7)
	RETURNING id; 
	`

	var documentID int
	err = psqldb.Pool.QueryRow(
		context.Background(),
		query,
		"DEMO",
		"demo-session-id",
		file.Filename,
		filePath,
		fileInfo.Size(),
		file.Header.Get("Content-Type"),
		"UPLOADED",
	).Scan(&documentID)

	if err != nil {
		return 0, err
	}

	return documentID, nil
}

func NotifyPythonService(documentID int, filePath string) (int, error) {

	request := common.ProcessDocumentRequest{
		DocumentId:  documentID,
		StoragePath: filePath,
	}

	jsonData, err := json.Marshal(request)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	response, err := http.Post(
		"http://python-ai-service:8000/v1/process-document",
		"application/json",
		bytes.NewBuffer(jsonData),
	)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return response.StatusCode, fmt.Errorf("python service returned status %d", response.StatusCode)
	}

	var updateReq common.UpdateDocumentRequest

	err = json.NewDecoder(response.Body).Decode(&updateReq)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	log.Println("Response recieved from python service:", updateReq)

	return http.StatusOK, nil

}

func UpdateDocumentStatusInPostgres(documentID int, statusCodePython int) (int, error) {
	var status string
	/* status can have UPLOAD, READY, PROCESS FAILED*/
	if statusCodePython == http.StatusOK {
		status = "READY"
	} else {
		status = "PROCESS FAILED"
	}

	query := `
		UPDATE documents
		SET
			status = $1,
			updated_at = NOW()
		WHERE id = $2
	`

	commandTag, err := psqldb.Pool.Exec(
		context.Background(),
		query,
		status,
		documentID,
	)

	if err != nil {
		return http.StatusInternalServerError, err
	}

	if commandTag.RowsAffected() == 0 {
		return http.StatusNotFound, nil
	}

	return http.StatusOK, nil
}
