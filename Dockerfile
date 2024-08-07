FROM golang:1.22 AS build

ARG SERVICE

WORKDIR /go/src

COPY go.mod go.sum .
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /app ./cmd/${SERVICE}

FROM gcr.io/distroless/static:nonroot
COPY --from=build /app /app
ENTRYPOINT ["/app"]