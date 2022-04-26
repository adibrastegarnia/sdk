# SPDX-FileCopyrightText: 2022-present Intel Corporation
#
# SPDX-License-Identifier: Apache-2.0

FROM golang:1.18 AS build

RUN mkdir /build

COPY ./go.mod /build/go.mod
COPY ./go.sum /build/go.sum
COPY ./api/ /build/api/
COPY ./cmd/ /build/cmd/
COPY ./pkg/ /build/pkg/

WORKDIR /build

RUN go build -o /build/atomix-plugin-build ./cmd/atomix-plugin-build

FROM golang:1.18

RUN apt-get update && \
    apt-get install -y unzip git

RUN mkdir -p /root/.atomix/build

WORKDIR /root

COPY --from=build /build/go.mod /root/.atomix/build/go.mod
COPY --from=build /build/atomix-plugin-build /usr/local/bin/atomix-plugin-build

ENTRYPOINT ["atomix-plugin-build"]
