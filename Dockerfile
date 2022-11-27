FROM golang:1.19-alpine AS build

ENV SERVER_PORT=9000
ENV DB_HOST=db
ENV DB_USER=postgres
ENV DB_PASS=secret
ENV DB_PORT=5432
ENV DB_NAME=api_todo

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 go build -o go-chi-api

FROM scratch

COPY --from=build /app/go-chi-api /go-chi-api

ENTRYPOINT [ "/go-chi-api" ]
