# SPDX-FileCopyrightText: 2022-present Intel Corporation
#
# SPDX-License-Identifier: Apache-2.0

FROM golang:1.18

RUN apt-get update && \
    apt-get install -y unzip git

RUN mkdir -p /runtime /build

COPY ./go.mod /runtime/go.mod
COPY ./go.sum /runtime/go.sum
COPY ./api/ /runtime/api/
COPY ./cmd/ /runtime/cmd/
COPY ./pkg/ /runtime/pkg/

RUN cd /runtime && go build -o /usr/local/bin/atomix-build github.com/atomix/atomix-runtime/cmd/atomix-build

WORKDIR /build

ENTRYPOINT ["atomix-build"]
