FROM golang:1.19

WORKDIR /usr/app
RUN go install github.com/cespare/reflex@latest
RUN go install github.com/go-task/task/v3/cmd/task@v3.10.0
ENTRYPOINT ["task", "run"]