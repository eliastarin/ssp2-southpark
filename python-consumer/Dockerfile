# Use a lightweight Python image
FROM python:3.12-slim

# Ensure logs are shown in real time (no buffering)
ENV PYTHONUNBUFFERED=1

# Set working directory inside container
WORKDIR /app

# Install dependencies first (for build cache)
COPY requirements.txt .
RUN pip install --no-cache-dir -r requirements.txt

# Copy the rest of the consumer files
COPY consumer.py .

# Run the consumer
CMD ["python", "-u", "consumer.py"]
