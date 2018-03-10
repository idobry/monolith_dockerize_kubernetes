FROM golang:1.8.5-jessie
# create a working directory
WORKDIR /go/src
# add source code
ADD src .
# run main.go
CMD ["go", "run", "/go/src/main.go"]