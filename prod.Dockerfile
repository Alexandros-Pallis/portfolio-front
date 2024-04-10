ARG ALPINE_VERSION=latest
FROM node:21.7.2-alpine as assets_builder
ARG artifacts_folder

WORKDIR /app

COPY $artifacts_folder ./
RUN npm install
RUN npm run build

FROM golang:1.21-alpine as build

WORKDIR /usr/src/app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY --from=assets_builder /app/go.mod /app/go.sum ./
RUN go mod download && go mod verify

COPY $artifacts_folder ./
COPY --from=assets_builder /app ./
RUN go install github.com/a-h/templ/cmd/templ@latest
RUN templ generate
RUN go build -v -o /usr/local/bin/app .

FROM alpine:${ALPINE_VERSION}
RUN echo "ALPINE_VERSION=${ALPINE_VERSION}"
WORKDIR /usr/src/app
COPY --from=build /usr/src/app/assets ./assets
COPY --from=build /usr/src/app/dist ./dist
COPY --from=build /usr/local/bin/app /usr/local/bin/app

EXPOSE 80

CMD ["app"]
