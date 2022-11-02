FROM golang:1.17.0-alpine
RUN apk add --no-cache build-base=0.5-r3 \
    openssh=8.6_p1-r3
WORKDIR /src
COPY . ./
COPY run.yml .
ENV CONFIG_FILE_PATH="."
RUN go build -o indexer-binary ./*.go
CMD ["./indexer-binary"]