# SPDX-FileCopyrightText: 2022-present Intel Corporation
#
# SPDX-License-Identifier: Apache-2.0

.PHONY: build api

.PHONY: api build docs

build:
	goreleaser release --snapshot --rm-dist

release:
	goreleaser release --rm-dist

api-go:
	@cd api && (rm -r **/*.pb.go || true) && cd ..
	docker run -it -v `pwd`:/build \
		atomix/proto-build:latest \
		go --package github.com/atomix/runtime-api/api ./api

api-docs:
	@cd api && (rm -r **/*.md || true) && cd ..
	docker run -it -v `pwd`:/build \
		atomix/proto-build:latest \
		docs ./api

api: api-go api-docs

reuse-tool: # @HELP install reuse if not present
	command -v reuse || python3 -m pip install reuse

license: reuse-tool # @HELP run license checks
	reuse lint
