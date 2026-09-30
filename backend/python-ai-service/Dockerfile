FROM python:3.12-slim

WORKDIR /app

COPY requirements.txt .

# Install CPU-only PyTorch before sentence-transformers (avoid default GPU version, complex using requirements.txt).
RUN pip install torch --index-url https://download.pytorch.org/whl/cpu
# sentence-transformers has lot of other dependencies so we install it using RUN
RUN pip install sentence-transformers

RUN pip install --no-cache-dir -r requirements.txt

COPY . .

EXPOSE 8000

CMD ["uvicorn", "app:app", "--host", "0.0.0.0", "--port", "8000"]
