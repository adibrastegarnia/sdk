# SPDX-FileCopyrightText: 2022-present Intel Corporation
#
# SPDX-License-Identifier: Apache-2.0

FROM golang:1.18 AS build

RUN mkdir -p /build

COPY ./go.mod /build
COPY ./go.sum /build
COPY ./cmd/ /build/cmd/
COPY ./internal/ /build/internal/

WORKDIR /build

RUN go build -o atomix-build ./cmd/atomix-build

FROM golang:1.18

RUN apt-get update && \
    apt-get install -y unzip git

RUN mkdir -p /build

COPY --from=build /build/atomix-build /usr/local/bin/atomix-build

WORKDIR /build

ENTRYPOINT ["atomix-build"]
