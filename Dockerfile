FROM golang:1.16-alpine as builder

WORKDIR /go/src/github.com/test_github_actions/

COPY go.mod go.sum ./
RUN go mod download

RUN echo hello

COPY . .

RUN go build -o out

FROM alpine

COPY --from=builder /go/src/github.com/test_github_actions/out /app/out

RUN echo world

CMD ["/app/out"]
