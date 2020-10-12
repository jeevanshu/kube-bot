FROM golang:1.14.4-alpine3.12 AS build

RUN apk update \
    && apk add --no-cache git libc6-compat

WORKDIR $GOPATH/src
COPY . .
RUN go build -o $GOPATH/bin/kube-bot
RUN chmod +x $GOPATH/bin/kube-bot


FROM golang:1.14.4-alpine3.12
RUN apk update \
    && apk add --no-cache git libc6-compat

COPY --from=build $GOPATH/bin/kube-bot $GOPATH/bin
COPY start.sh /start.sh
CMD [ "/start.sh" ]
