package common

type ProcessDocumentRequest struct {
	DocumentId  int    `json:"document_id"`
	StoragePath string `json:"storage_path"`
}

type UpdateDocumentRequest struct {
	DocumentId int    `json:"document_id"`
	Status     string `json:"status"`
}
