// Copyright (c) 2018 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package yarpcpeer

import (
	"context"

	yarpc "go.uber.org/yarpc/v2"
)

// Single implements the Chooser interface for a single peer
type Single struct {
	dialer yarpc.Dialer
	pid    yarpc.Identifier
}

// NewSingle creates a static Chooser with a single Peer
func NewSingle(pid yarpc.Identifier, dialer yarpc.Dialer) *Single {
	single := &Single{
		pid:    pid,
		dialer: dialer,
	}
	return single
}

// Choose returns the single peer
func (single *Single) Choose(ctx context.Context, _ *yarpc.Request) (yarpc.Peer, func(error), error) {
	peer, err := single.dialer.RetainPeer(single.pid, _nopSubscriber)
	if err != nil {
		return nil, nil, err
	}
	err = single.dialer.ReleasePeer(single.pid, _nopSubscriber)
	if err != nil {
		return nil, nil, err
	}
	peer.StartRequest()
	return peer, func(err error) {
		peer.EndRequest()
	}, nil
}

type nopSubscriber struct{}

func (nopSubscriber) NotifyStatusChanged(_ yarpc.Identifier) {}

var _nopSubscriber nopSubscriber
