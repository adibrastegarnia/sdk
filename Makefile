# SPDX-FileCopyrightText: 2022-present Intel Corporation
#
# SPDX-License-Identifier: Apache-2.0

.PHONY: build api

build:
	goreleaser release --snapshot --rm-dist

release:
	goreleaser release

api: build
	@cd api && (rm -r **/*.pb.go **/*.md || true) && cd ..
	docker run -it -v `pwd`:/build \
		--entrypoint build/bin/compile-protos.sh \
		atomix/proto-build:latest

reuse-tool: # @HELP install reuse if not present
	command -v reuse || python3 -m pip install reuse

license: reuse-tool # @HELP run license checks
	reuse lint
