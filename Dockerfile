FROM alpine

ARG SOURCE_REPOSITORY=github.com/Project-Prismatica/prismatica-core

ENV GOPATH=/go \
    CORE_AMBASSADOR_SOURCE_CONFIG_DIR=/source-configs/ambassador/config \
    CORE_AMBASSADOR_CONFIG_DIR=/external-configs/ambassador/config

ADD . /go/src/${SOURCE_REPOSITORY}

RUN apk add --no-cache dumb-init git go musl-dev && \
    go get $SOURCE_REPOSITORY && \
    apk del --no-cache git go musl-dev && \
    mkdir -p \
        $CORE_AMBASSADOR_SOURCE_CONFIG_DIR \
        $CORE_AMBASSADOR_CONFIG_DIR && \
    adduser -S -D -H prismatica-core

USER prismatica-core
ENTRYPOINT [ "/usr/bin/dumb-init", "--" ]
CMD [ "/go/bin/prismatica-core" ]
