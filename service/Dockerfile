
ARG GOLANG_IMAGE_TAG=1.10.4-alpine3.8
ARG RUNNER_ALPINE_IMAGE_TAG=3.8

FROM golang:${GOLANG_IMAGE_TAG} as base

ARG PROTOBUF_VERSION_ARG=3.3.5.2-r0
ARG GO_PROJECT_ARG=github.com/ivan-kostko/goq/service

ENV PROTOBUF_VERSION=${PROTOBUF_VERSION_ARG}
ENV GO_PROJECT=${GO_PROJECT_ARG}

RUN apk add git binutils protobuf	#=${PROTOBUF_VERSION}

RUN go tool cgo -V && \
    protoc --version

RUN go get -u google.golang.org/grpc
RUN go get -u github.com/golang/protobuf/protoc-gen-go
RUN go get -u github.com/grpc-ecosystem/grpc-gateway/...
# RUN go get -u github.com/davecheney/godoc2md


FROM base AS generator

ENV GO_PROJECT_PATH=${GOPATH}/src/${GO_PROJECT}

RUN env

RUN mkdir -p ${GO_PROJECT_PATH}

ADD . ${GO_PROJECT_PATH}

WORKDIR ${GO_PROJECT_PATH}

RUN env
RUN echo ${PWD}

RUN PROTOC="protoc" PROTO_DIR=./proto OUTDIR=./proto/src ./generate_proto.sh
RUN go get ./... && \
    go generate ./...

# Generates :
# * ./proto/src from proto definitions
# * all go generate
# * documentation on packages
CMD PROTOC="protoc" PROTO_DIR=./proto OUTDIR=./proto/src ./generate_proto.sh && \
    echo "Generating..." && go generate ./... && \
    cd $GOPATH/src &&  \
    find "./$GO_PROJECT" \( ! -name .git -o -prune \) -type d | grep -v "\.git" | sed 's:^\.\/::' | while read PACKAGE; \
    do \
        echo "Generating README for ${PACKAGE}"; \
        # $(${GOPATH}/bin/godoc2md ${PACKAGE} > ./${PACKAGE}/README.MD) ; \
    done;

