FROM registry.semaphoreci.com/golang:1.19 as builder

ENV APP_HOME /booksapi

WORKDIR "$APP_HOME"
COPY . .

RUN go mod download
RUN go mod verify
RUN go build

FROM registry.semaphoreci.com/golang:1.19

ENV APP_HOME /booksapi
RUN mkdir -p "$APP_HOME"
WORKDIR "$APP_HOME"

COPY --from=builder "$APP_HOME"/booksapi $APP_HOME

CMD ["./booksapi"]
