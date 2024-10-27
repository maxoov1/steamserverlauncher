FROM golang:1.22.2 AS builder

WORKDIR /usr/src/steamserverlauncher

COPY go.mod go.sum ./
RUN go mod download

ENV CGO_ENABLED=0

COPY ./ ./
RUN go build -o steamserverlauncher ./cmd/main.go

FROM scratch AS runner

COPY --from=builder \
  /usr/src/steamserverlauncher/templates templates

COPY --from=builder \
  /usr/src/steamserverlauncher/static static

COPY --from=builder \
  /usr/src/steamserverlauncher/steamserverlauncher ./

ENV LAUNCHER_ADDRESS=0.0.0.0:80
EXPOSE 80

CMD [ "./steamserverlauncher" ]
