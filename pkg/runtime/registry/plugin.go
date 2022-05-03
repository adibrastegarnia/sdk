// SPDX-FileCopyrightText: 2022-present Intel Corporation
//
// SPDX-License-Identifier: Apache-2.0

package registry

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"sync"
)

type Plugin struct {
	registry   *Registry
	Name       string
	Version    string
	APIVersion string
	Path       string
	mu         sync.RWMutex
}

func (p *Plugin) Create() (io.WriteCloser, error) {
	p.mu.Lock()
	writer, err := os.Create(p.Path)
	if err != nil {
		p.mu.Unlock()
		return nil, err
	}
	return &PluginWriter{
		plugin: p,
		writer: writer,
	}, nil
}

func (p *Plugin) Open() (io.ReadCloser, error) {
	p.mu.RLock()
	reader, err := os.Open(p.Path)
	if err != nil {
		p.mu.RUnlock()
		return nil, err
	}
	return &PluginReader{
		plugin: p,
		reader: reader,
	}, nil
}

func (p *Plugin) Download(owner, repo string) error {
	p.mu.Lock()
	defer p.mu.Unlock()

	url := fmt.Sprintf("https://github.com/%s/%s/releases/download/%s/%s", owner, repo, p.Version, getPluginFile(p.Name, p.Version, p.APIVersion))

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	file, err := os.Create(p.Path)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return err
	}
	return nil
}

type PluginWriter struct {
	plugin *Plugin
	writer io.WriteCloser
}

func (w *PluginWriter) Write(p []byte) (n int, err error) {
	return w.writer.Write(p)
}

func (w *PluginWriter) Close() error {
	defer w.plugin.mu.Unlock()
	return w.writer.Close()
}

type PluginReader struct {
	plugin *Plugin
	reader io.ReadCloser
}

func (r *PluginReader) Read(p []byte) (n int, err error) {
	return r.reader.Read(p)
}

func (r *PluginReader) Close() error {
	defer r.plugin.mu.RUnlock()
	return r.reader.Close()
}

func getPluginFile(name, version, apiVersion string) string {
	return fmt.Sprintf("%s-%s.%s.so", name, version, apiVersion)
}
