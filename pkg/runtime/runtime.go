// SPDX-FileCopyrightText: 2022-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package runtime

import (
	"context"
	"github.com/atomix/runtime-api/pkg/logging"
)

var log = logging.GetLogger()

type Runtime interface {
	Connect(ctx context.Context, store string) (Conn, error)
}
