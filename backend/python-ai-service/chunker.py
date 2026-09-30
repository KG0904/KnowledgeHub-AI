from langchain_text_splitters import RecursiveCharacterTextSplitter

def create_chunks(text):
    splitter = RecursiveCharacterTextSplitter.from_tiktoken_encoder(
        chunk_size=500,
        chunk_overlap=100,
    )

    chunks = splitter.split_text(text)
    return chunks