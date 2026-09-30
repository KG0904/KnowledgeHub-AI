from fastapi import FastAPI, HTTPException
from pydantic import BaseModel
from processor import process_document

app = FastAPI()

@app.get("/test")
def test():
    print("Inside test")
    return {
        "message": "status OK"
    }


#Inheritence of BaseModel class by ProcessDocumentRequest class
class ProcessDocumentRequest(BaseModel):
    document_id: int
    storage_path: str

@app.post("/v1/process-document")
def process_document_api(request: ProcessDocumentRequest):

    print("Received document ID:", request.document_id)

    success = process_document(
        request.document_id,
        request.storage_path,
    )

    if success:
        return {
            "document_id": request.document_id,
            "status": "READY",
        }

    raise HTTPException(
        status_code=500,
        detail={
            "document_id": request.document_id,
            "status": "FAILED",
        }
    )