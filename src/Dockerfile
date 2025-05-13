FROM golang:latest
WORKDIR /src
COPY . /src
RUN go mod download
EXPOSE 8000

ENTRYPOINT [ "go", "run", "main.go" ]
# ENTRYPOINT [ "./gopeople" ]