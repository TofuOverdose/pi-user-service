# dev container with hot reloading
FROM golang:1.16-alpine

RUN go get github.com/cespare/reflex

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY reflex.conf /

CMD ["reflex", "-c", "/reflex.conf"]