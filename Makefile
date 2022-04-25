# SPDX-FileCopyrightText: 2019-present Open Networking Foundation <info@opennetworking.org>
#
# SPDX-License-Identifier: Apache-2.0

.PHONY: build api

VERSION := latest

BUILD_VERSION := latest

api-go:
	@cd api && (rm -r **/*.pb.go **/*.md || true) && cd ..
	docker run -it -v `pwd`:/build atomix/atomix-build:latest \
		api go \
		--input api \
		--output api \
		--package github.com/atomix/atomix-runtime/api \
		.

api-docs:
	@cd api && (rm -r **/*.md || true) && cd ..
	docker run -it -v `pwd`:/build atomix/atomix-build:latest \
		api docs \
		--input api \
		--output api \
		.

api: api-go api-docs

atomix-plugin-registry-docker:
	docker build \
		-f build/docker/registry.Dockerfile \
		-t atomix/atomix-plugin-registry:$(VERSION) \
		--build-arg BUILD_VERSION=$(BUILD_VERSION) .

atomix-plugin-compiler-docker:
	docker build \
		-f build/docker/compiler.Dockerfile \
		-t atomix/atomix-plugin-compiler:$(VERSION) \
		--build-arg BUILD_VERSION=$(BUILD_VERSION) .

images: atomix-plugin-registry-docker atomix-plugin-compiler-docker

reuse-tool: # @HELP install reuse if not present
	command -v reuse || python3 -m pip install reuse

license: reuse-tool # @HELP run license checks
	reuse lint
