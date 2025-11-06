# 1) build
FROM golang:1.22 AS build
WORKDIR /app
COPY . .
RUN go build -o /go-api

# 2) run
FROM gcr.io/distroless/base-debian12
WORKDIR /app
COPY --from=build /go-api /go-api
EXPOSE 8080
ENTRYPOINT ["/go-api"]
