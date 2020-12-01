FROM golang

WORKDIR /src
COPY . .
RUN export GOARCH=amd64 GOOS=linux && go build -o server

FROM scratch

COPY --from=0 /src/server /
ENTRYPOINT ["/server"]
