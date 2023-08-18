FROM golang:1.20-buster AS builder

WORKDIR /app

# copy dependencies
COPY go.* ./
RUN go mod download

# copy source
COPY . .

# build binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o /main ./cmd

# execute binary
FROM scratch

# copy binary
COPY --from=builder /main /main

# copy resources
COPY --from=builder /app/resources /app/resources

# run binary
CMD ["/main"]
