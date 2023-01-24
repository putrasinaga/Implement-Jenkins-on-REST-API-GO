FROM golang:1.16-buster AS builder 
WORKDIR /app
COPY go.* ./
RUN go mod download 
COPY *.go ./
RUN go build -o /web-simple-api


# Create a new release build stage
FROM gcr.io/distroless/base-debian10
# Set the working directory to the root directory path
WORKDIR /
# Copy over the binary built from the previous stage
COPY --from=builder /web-simple-api . 
EXPOSE 9090
ENTRYPOINT ["/web-simple-api"]