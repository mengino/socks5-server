FROM golang

WORKDIR /src
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-s' -o ./server

FROM scratch

COPY --from=0 /src/server /
ENTRYPOINT ["/server"]
