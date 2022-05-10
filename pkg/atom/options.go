// SPDX-FileCopyrightText: 2022-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package atom

import "github.com/atomix/sdk/pkg/plugin"

type RepoOptions struct {
	Path       string
	Downloader plugin.DownloadFunc
}

func (o RepoOptions) apply(opts ...RepoOption) {
	for _, opt := range opts {
		opt(&o)
	}
}

type RepoOption func(*RepoOptions)

func WithRepoOptions(options RepoOptions) RepoOption {
	return func(opts *RepoOptions) {
		*opts = options
	}
}

func WithPath(path string) RepoOption {
	return func(options *RepoOptions) {
		options.Path = path
	}
}

func WithDownloader(downloader plugin.DownloadFunc) RepoOption {
	return func(options *RepoOptions) {
		options.Downloader = downloader
	}
}
