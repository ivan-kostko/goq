#!/bin/sh

GO_OUTDIR=${OUTDIR}/go
JAVA_OUTDIR=${OUTDIR}/java
JS_OUTDIR=${OUTDIR}/js

rm -rf ${GO_OUTDIR}
mkdir -pv ${GO_OUTDIR}
rm -rf ${JAVA_OUTDIR}
mkdir -pv ${JAVA_OUTDIR}
rm -rf ${JS_OUTDIR}
mkdir -pv ${JS_OUTDIR}


for filename in ${PROTO_DIR}/*.proto; do
    filename=${filename##*/} # remove path
    filename=${filename%.*} # remove extension

    rm -rf ${GO_OUTDIR}/${filename}
    mkdir -pv ${GO_OUTDIR}/${filename}
    rm -rf ${JAVA_OUTDIR}/${filename}
    mkdir -pv ${JAVA_OUTDIR}/${filename}
    rm -rf ${JS_OUTDIR}/${filename}
    mkdir -pv ${JS_OUTDIR}/${filename}

    ${PROTOC} \
    -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
    -I${PROTO_DIR} \
    --grpc-gateway_out=logtostderr=true,grpc_api_configuration=${PROTO_DIR}/rest_api_config.yaml:${GO_OUTDIR}/${filename} \
	--go_out=plugins=grpc:${GO_OUTDIR}/${filename} \
    --java_out=${JAVA_OUTDIR} \
    --js_out=${JS_OUTDIR}/${filename} \
    ${filename}.proto

done
