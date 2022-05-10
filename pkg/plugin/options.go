// SPDX-FileCopyrightText: 2022-present Intel Corporation
//
// SPDX-License-Identifier: Apache-2.0

package plugin

const (
	defaultPath   = "."
	defaultSymbol = "Plugin"
)

type RepoOptions struct {
	Symbol     string
	Downloader DownloadFunc
}

func (o RepoOptions) apply(opts ...RepoOption) {
	o.Symbol = defaultSymbol
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

func WithSymbol(symbol string) RepoOption {
	return func(options *RepoOptions) {
		options.Symbol = symbol
	}
}

func WithDownloader(f DownloadFunc) RepoOption {
	return func(options *RepoOptions) {
		options.Downloader = f
	}
}

type CacheOptions struct {
	Path string
}

func (o CacheOptions) apply(opts ...CacheOption) {
	o.Path = defaultPath
	for _, opt := range opts {
		opt(&o)
	}
}

type CacheOption func(*CacheOptions)

func WithCacheOptions(options CacheOptions) CacheOption {
	return func(opts *CacheOptions) {
		*opts = options
	}
}

func WithPath(path string) CacheOption {
	return func(options *CacheOptions) {
		options.Path = path
	}
}
