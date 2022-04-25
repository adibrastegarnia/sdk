# SPDX-FileCopyrightText: 2019-present Open Networking Foundation <info@opennetworking.org>
#
# SPDX-License-Identifier: Apache-2.0

.PHONY: build api

VERSION := latest

BUILD_VERSION := latest

all: build proxy registry compiler

api:
	@cd api && (rm -r **/*.pb.go **/*.md || true) && cd ..
	docker run -it -v `pwd`:/build \
		--entrypoint build/bin/compile-protos.sh \
		`docker build -q build/docker/api`

atomix-build:

atomix-proxy:

atomix-registry:
	docker build \
		-f build/docker/registry/Dockerfile \
		-t atomix/atomix-plugin-registry:$(VERSION) \
		--build-arg BUILD_VERSION=$(BUILD_VERSION) .

atomix-compiler:
	docker build \
		-f build/docker/compiler/Dockerfile \
		-t atomix/atomix-plugin-compiler:$(VERSION) \
		--build-arg BUILD_VERSION=$(BUILD_VERSION) .

images: atomix-build atomix-proxy atomix-registry atomix-compiler

reuse-tool: # @HELP install reuse if not present
	command -v reuse || python3 -m pip install reuse

license: reuse-tool # @HELP run license checks
	reuse lint
