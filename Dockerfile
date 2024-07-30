# Use a newer version of Go that includes the net/netip package
FROM golang:1.21-alpine AS builder

LABEL maintainer="Khalid Alhabibie <Khalidalhabibie07@gmail.com>"

# Move to working directory (/build).
WORKDIR /build

# Copy and download dependency using go mod.
COPY go.mod go.sum ./
RUN go mod download

# Copy the code into the container.
COPY . .

# Set necessary environment variables needed for our image and build the API server.
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
RUN go build -ldflags="-s -w" -o goarticle .

FROM scratch

# Copy binary and config files from /build to root folder of scratch container.
COPY --from=builder ["/build/goarticle", "/build/.env", "/"]

# Command to run when starting the container.
CMD ["./goarticle"]
