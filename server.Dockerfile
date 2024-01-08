FROM golang:1.21.5

# Set the working directory inside the container
WORKDIR /app

COPY go.mod go.sum ./

# Download and install Go module dependencies
RUN go mod download

# Copy the current directory contents into the container at /app
COPY . .

# Build the Golang application
RUN go build -o main .

# Run the application when the container starts
CMD ["./main"]