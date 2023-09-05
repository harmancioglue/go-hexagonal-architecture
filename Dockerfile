FROM golang:1.17.0-alpine3.13

WORKDIR /app

# add some necessary packages
RUN apk update && \
    apk add libc-dev && \
    apk add gcc && \
    apk add make

# Install Delve for debugging
RUN go install github.com/go-delve/delve/cmd/dlv@latest

COPY ./go.mod  ./
COPY ./go.sum  ./

RUN go mod download && go mod verify


#ınstall compileDeamon

RUN go install -mod=mod github.com/githubnemo/CompileDaemon

COPY . .
COPY ./entrypoint.sh /entrypoint.sh


ADD https://raw.githubusercontent.com/eficode/wait-for/v2.1.0/wait-for /usr/local/bin/wait-for
RUN chmod +rx /usr/local/bin/wait-for /entrypoint.sh


ENTRYPOINT ["sh", "/entrypoint.sh"]