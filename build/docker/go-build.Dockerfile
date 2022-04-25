# SPDX-FileCopyrightText: 2022-present Intel Corporation
#
# SPDX-License-Identifier: Apache-2.0

FROM golang:1.18

RUN apt-get update && \
    apt-get install -y unzip git

RUN mkdir -p /build/src /build/tgt

COPY go.mod /build/tgt/go.mod

WORKDIR /build
