FROM golang:1.18-alpine
LABEL maintainer="SpectreH"
LABEL org.label-schema.schema-version="1.0"
LABEL org.label-schema.description="Social-Network frontend"

RUN apk add build-base
RUN mkdir -p /usr/src/back/
WORKDIR /usr/src/back/
COPY . /usr/src/back/
RUN rm /usr/src/back/db/network.db
RUN touch /usr/src/back/db/network.db

RUN go mod download
RUN go build -o /server
EXPOSE 4000

CMD [ "/server" ]
