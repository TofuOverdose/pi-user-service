# dev container with hot reloading
FROM golang:1.16-alpine

RUN go get github.com/cespare/reflex

COPY reflex.conf /

CMD ["reflex", "-c", "/reflex.conf"]
