package handler

import (
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func TestHandler(c *gin.Context) {
	c.JSON(200, gin.H{"message": "status OK"})
}

func UploadDocument(c *gin.Context) {
	log.Println("Inside UploadDocument")

	file, err := c.FormFile("file")
	if err != nil {
		log.Println("Error getting file:", err)

		c.JSON(http.StatusBadRequest, gin.H{
			"error": "file is required",
		})
		return
	}

	uploadDir := "/app/uploads"
	log.Println("upload directory: ", uploadDir)
	err = os.MkdirAll(uploadDir, os.ModePerm)
	if err != nil {
		log.Println("Error creating upload directory:", err)

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "unable to create upload directory",
		})
		return
	}
	log.Println("/app/uploads created")
	filePath := filepath.Join(uploadDir, file.Filename)

	err = c.SaveUploadedFile(file, filePath)
	if err != nil {
		log.Println("Error saving file:", err)

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "unable to save file",
		})
		return
	}

	log.Println("File uploaded successfully:", filePath)

	documentID, err := SaveMetaDataInPostgres(file, filePath)

	if err != nil {
		log.Println("Error saving Metadata in postgres:", err)

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "unable to update postgres",
		})
		return
	}

	statusCodePython, err := NotifyPythonService(documentID, filePath)
	if err != nil {
		log.Println("Error Notifying Python AI Service:", err.Error())

		c.JSON(statusCodePython, gin.H{
			"error": "Unable to notify python service",
		})
		return
	}

	statusCode, err := UpdateDocumentStatusInPostgres(documentID, statusCodePython)
	if err != nil {
		log.Println("Error in updating document status:", err.Error())

		c.JSON(statusCode, gin.H{
			"message": "Document upload successfully however unable to update postgres",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message":     "file uploaded and processed successfully",
		"filename":    file.Filename,
		"document_id": documentID,
	})
}

// func UpdateDocumentStatus(c *gin.Context) {

// 	documentID := c.Param("id")

// 	var req common.UpdateDocumentRequest

// 	if err := c.ShouldBindJSON(&req); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"error": "invalid request body",
// 		})
// 		return
// 	}

// 	statusCode, err := UpdateDocumentStatusInPostgres(documentID, req)
// 	if err != nil {
// 		log.Println("Error in updating document status:", err.Error())

// 		c.JSON(statusCode, gin.H{
// 			"error": "unable to update postgres",
// 		})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{
// 		"message": "document updated successfully",
// 	})
// }
