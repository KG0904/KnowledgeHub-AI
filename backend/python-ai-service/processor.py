import os
from pdf_reader import extract_text
from chunker import create_chunks
from embedding import create_embeddings
from qdrant_collection import store_embeddings


def process_document(document_id: int, storage_path: str) -> bool:

    try:
        print("Processing started")
        print("Document ID:", document_id)
        print("File path:", storage_path)

        # Check whether file exists
        if not os.path.exists(storage_path):
            print("File does not exist")
            return False

        print("File found")

        # TODO:
        # 1. Read PDF
        # 2. Extract text
        # 3. Split into chunks
        # 4. Generate embeddings
        # 5. Store embeddings in Qdrant
        print("Reading PDF...")

        text = extract_text(storage_path)

        print("PDF read successfully\n\nInitial 500 words\n")

        #Print first 500 words
        print(text[:500])

        chunks = create_chunks(text)

        print("Total chunks:", len(chunks))

        #To print chunks
        # for i, chunk in enumerate(chunks[:5]):
        #     print(f"Chunk{i}: {chunk}")
        #     print(f"--------------------------------------------------------------------------\n")

        embeddings = create_embeddings(chunks)

        print("Total embeddings:", len(embeddings))

        print("Dimension:", len(embeddings[0]))

        print("Vector embedding first chunck",embeddings[0])

        store_embeddings(document_id, chunks, embeddings)

        print("Document processing completed")

        return True

    except Exception as e:
        print("Error while processing document:", e)
        return False