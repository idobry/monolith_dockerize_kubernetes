FROM golang:1.8.5-jessie
# create a working directory
WORKDIR /go/src

RUN go get github.com/gorilla/mux
RUN go get github.com/lib/pq

# add source code
ADD src .
# run main.go
CMD ["go", "run", "/go/src/insertVote.go"]
