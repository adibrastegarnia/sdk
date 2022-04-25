# SPDX-FileCopyrightText: 2022-present Intel Corporation
#
# SPDX-License-Identifier: Apache-2.0

.PHONY: build api

VERSION := latest

BUILD_VERSION := latest

all: atomix atomix-build atomix-proxy atomix-registry atomix-compiler

api: proto-build
	@cd api && (rm -r **/*.pb.go **/*.md || true) && cd ..
	docker run -it -v `pwd`:/build \
		--entrypoint build/bin/compile-protos.sh \
		atomix/proto-build:$(VERSION)

atomix:
	goreleaser build --single-target

proto-build:
	docker build \
		-f build/docker/proto-build.Dockerfile \
		-t atomix/proto-build:$(VERSION) .

go-build:
	docker build \
		-f build/docker/go-build.Dockerfile \
		-t atomix/go-build:$(VERSION) .

test-release:
	goreleaser release --snapshot --rm-dist

release:
	goreleaser release

reuse-tool: # @HELP install reuse if not present
	command -v reuse || python3 -m pip install reuse

license: reuse-tool # @HELP run license checks
	reuse lint
