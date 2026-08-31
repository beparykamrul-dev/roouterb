FROM golang:1.24-alpine AS build
WORKDIR /src
COPY go.mod ./
COPY . .
RUN CGO_ENABLED=0 go build -trimpath -ldflags='-s -w' -o /out/routerb ./cmd/routerb

FROM alpine:3.22
RUN adduser -D -H -u 10001 routerb
COPY --from=build /out/routerb /usr/local/bin/routerb
USER 10001
EXPOSE 8080
ENTRYPOINT ["/usr/local/bin/routerb"]
