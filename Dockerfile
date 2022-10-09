# Specify the base image for the go app.
FROM golang:1.16.0-alpine3.13
# Specify that we now need to execute any commands in this directory.
WORKDIR /go/src/github.com/yemrealtanay/sms_templates
# Copy everything from this project into the filesystem of the container.
COPY . .
# Obtain the package needed to run code. Alternatively use GO Modules.
RUN go get -u github.com/lib/pq

RUN go get -d -v ./...

# Install the package
RUN go install -v ./...

#Setup hot-reload for dev stage
RUN go get github.com/githubnemo/CompileDaemon
RUN go get -v golang.org/x/tools/gopls

RUN apk add build-base

ENTRYPOINT CompileDaemon --build="go build -a -installsuffix cgo -o main ." --command=./main