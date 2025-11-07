# 1) build stage — use Go 1.25+
FROM golang:1.25 AS build
WORKDIR /app

# Cache deps early
COPY go.mod go.sum ./
RUN go env -w GOTOOLCHAIN=auto \
 && go mod download

# Now copy the rest
COPY . .

# Produce a static Linux binary
RUN CGO_ENABLED=0 GOOS=linux go build -o /go-api

# 2) run stage — minimal image
FROM gcr.io/distroless/base-debian12
WORKDIR /app
COPY --from=build /go-api /go-api
EXPOSE 8080
USER 65532:65532
ENTRYPOINT ["/go-api"]

FROM python:3.12-slim
ENV PYTHONUNBUFFERED=1
WORKDIR /app
COPY requirements.txt .
RUN pip install --no-cache-dir -r requirements.txt
COPY consumer.py .
CMD ["python", "-u", "consumer.py"]

