# ---------- BUILD STAGE ----------
# Use the latest Go toolchain (you need Go ≥ 1.25 for GOTOOLCHAIN=auto)
FROM golang:1.25 AS build

# Set the working directory inside the builder container
WORKDIR /app

# Copy go.mod and go.sum first (for caching)
COPY go.mod go.sum ./

# Enable automatic toolchain updates and download dependencies
RUN go env -w GOTOOLCHAIN=auto && go mod download

# Copy the rest of the project source files
COPY . .

# Build the Go binary with static linking
RUN CGO_ENABLED=0 GOOS=linux go build -o /go-api

# ---------- RUNTIME STAGE ----------
# Use a minimal runtime image (no shell, no package manager)
FROM gcr.io/distroless/base-debian12

# Working directory inside the final container
WORKDIR /app

# Copy only the compiled binary from the builder
COPY --from=build /go-api /go-api

# Expose port 8080 (the Go API + embedded web UI)
EXPOSE 8080

# Run the binary as a non-root user for security
USER 65532:65532

# Start the app
ENTRYPOINT ["/go-api"]
