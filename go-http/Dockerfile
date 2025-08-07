FROM golang:1.24 AS build

WORKDIR /app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -v -o server ./...

FROM scratch

COPY --from=build /app/server /server

ENTRYPOINT ["/server"]
