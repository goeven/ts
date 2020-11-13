FROM golang:1.15-alpine as build

WORKDIR /src
ADD . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags="-w -s" -o /usr/local/bin/ts .

FROM scratch

COPY --from=build /usr/local/bin/ts /usr/local/bin/ts

ENTRYPOINT [ "/usr/local/bin/ts" ]
