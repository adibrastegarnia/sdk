# SPDX-FileCopyrightText: 2022-present Intel Corporation
#
# SPDX-License-Identifier: Apache-2.0

.PHONY: build api

build:
	go build ./...

release:
	goreleaser release

api: build
	@cd api && (rm -r **/*.pb.go **/*.md || true) && cd ..
	docker run -it -v `pwd`:/build \
		--entrypoint build/bin/compile-protos.sh \
		`docker build -f build/docker/proto-build.Dockerfile -q .`

reuse-tool: # @HELP install reuse if not present
	command -v reuse || python3 -m pip install reuse

license: reuse-tool # @HELP run license checks
	reuse lint
