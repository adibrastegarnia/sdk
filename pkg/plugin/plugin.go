// SPDX-FileCopyrightText: 2022-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package plugin

import (
	"fmt"
	"github.com/gofrs/flock"
	"io"
	"os"
)

func newPlugin(name, version, path string) *Plugin {
	return &Plugin{
		Name:    name,
		Version: version,
		Path:    path,
		lock:    flock.New(fmt.Sprintf("%s.lock", path)),
	}
}

type Plugin struct {
	Name    string
	Version string
	Path    string
	lock    *flock.Flock
}

func (p *Plugin) Create() (io.WriteCloser, error) {
	if err := p.lock.Lock(); err != nil {
		return nil, err
	}
	file, err := os.Create(p.Path)
	if err != nil {
		return nil, err
	}
	return &Writer{
		plugin: p,
		writer: file,
	}, nil
}

func (p *Plugin) Open() (io.ReadCloser, error) {
	if err := p.lock.RLock(); err != nil {
		return nil, err
	}
	file, err := os.Open(p.Path)
	if err != nil {
		return nil, err
	}
	return &Reader{
		plugin: p,
		reader: file,
	}, nil
}

func (p *Plugin) String() string {
	return getPluginName(p.Name, p.Version)
}

type Writer struct {
	plugin *Plugin
	writer io.WriteCloser
}

func (w *Writer) Write(bytes []byte) (n int, err error) {
	return w.writer.Write(bytes)
}

func (w *Writer) Close() error {
	defer w.plugin.lock.Unlock()
	return w.writer.Close()
}

var _ io.WriteCloser = (*Writer)(nil)

type Reader struct {
	plugin *Plugin
	reader io.ReadCloser
}

func (r *Reader) Read(bytes []byte) (n int, err error) {
	return r.reader.Read(bytes)
}

func (r *Reader) Close() error {
	defer r.plugin.lock.Unlock()
	return r.reader.Close()
}

var _ io.ReadCloser = (*Reader)(nil)

func getPluginName(name, version string) string {
	return fmt.Sprintf("%s/%s", name, version)
}
