FROM golang:alpine AS builder
RUN apk update && apk add --no-cache git
WORKDIR /go/src/main
ENV GO111MODULE=on
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /app .


FROM scratch
COPY --from=builder /app /app
ENTRYPOINT ["/app"]
