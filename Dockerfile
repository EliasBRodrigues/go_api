FROM golang:1.22-alpine as stage

WORKDIR /app

COPY  go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o goApi ./main

FROM scratch

COPY  --from=stage /app/goApi /

ENTRYPOINT [ "/goApi" ]