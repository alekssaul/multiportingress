
FROM golang:latest as builder
WORKDIR /go/src/github.com/alekssaul/multiportingress/
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o server ./


FROM alpine:latest
RUN mkdir /app
WORKDIR /app
COPY --from=builder /go/src/github.com/alekssaul/multiportingress/server .
CMD /app/server
EXPOSE 8080
EXPOSE 8081
EXPOSE 8082
