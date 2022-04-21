FROM golang:1.17-alpine

# Install git + SSL ca certificates.
# Git is required for fetching the dependencies.
# Ca-certificates is required to call HTTPS endpoints.
# Else you will get error => local error: tls: bad record MAC 
RUN apk update && apk add --no-cache git ca-certificates tzdata && update-ca-certificates

ENV GIN_MODE=release
ENV PORT=8080 

WORKDIR /myapp/
COPY . .

# Fix the error cannot find package "golang.org/x/net/html" in any of:
# default: #10 0.286  /usr/local/go/src/golang.org/x/net/html (from $GOROOT)
# default: #10 0.286  /go/src/golang.org/x/net/html (from $GOPATH)
# RUN go mod vendor

RUN go get ./...

# Build the binary
# RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o /go/bin/blockexplorer -mod vendor .
# -ldflags="-w" to disable debug -> smaller image. -s make the linked C part also static into the binary, reducing compatibility risk
RUN go build -ldflags="-w -s" -o ./blockexplorer . 

### Cross compile, CGO_ENABLED not work on cross compiling -> disable it
## Compiling Linux and windows 64 bit executable program under mac
# CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build main.go
# CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build main.go

## Compiling 64 bit executable program of MAC and Linux under Windows
# CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build main.go
# CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build main.go

ENTRYPOINT ./blockexplorer -port=${8080}]

EXPOSE $PORT