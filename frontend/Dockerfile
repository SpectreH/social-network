FROM golang:1.18-alpine
LABEL maintainer="SpectreH"
LABEL org.label-schema.schema-version="1.0"
LABEL org.label-schema.description="Social-Network frontend"

RUN mkdir -p /usr/src/front/
RUN mkdir -p /usr/src/front/dist/
WORKDIR /usr/src/front/
COPY app.go /usr/src/front/
COPY go.mod /usr/src/front/
COPY dist /usr/src/front/dist/

RUN go build -o /app
EXPOSE 3000

CMD [ "/app" ]
