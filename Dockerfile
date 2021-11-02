FROM golang:alpine
RUN go version

ADD . /go/src/app
WORKDIR /go/src/app

# Gin will use the PORT env var
#ENV PORT 5000
EXPOSE 8080

# Install git
RUN apk add --no-cache git
# Fetch deps
COPY go.mod .
COPY go.sum .

RUN go mod download


# Remove git
# Compile app
RUN go build -o main .
# Run app
CMD ["./main"]