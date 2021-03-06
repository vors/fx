// Copyright (c) 2016 Uber Technologies, Inc.
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

package service

import (
	"go.uber.org/fx/core/config"
	"go.uber.org/fx/core/ulog"

	"github.com/uber-go/tally"
)

// A Option configures a service host
type Option func(Host) error

// WithConfiguration adds configuration to a service host
func WithConfiguration(config config.ConfigurationProvider) Option {
	return func(svc Host) error {
		// TODO(ai) verify type assertion is correct
		svc2 := svc.(*host)
		svc2.configProvider = config
		return nil
	}
}

// WithLogger adds ulog to a service host
func WithLogger(log ulog.Log) Option {
	return func(svc Host) error {
		// TODO(ai) verify type assertion is correct
		svc2 := svc.(*host)
		svc2.log = log
		return nil
	}
}

// WithMetricsRootScope configures a service host with metrics
func WithMetricsRootScope(scope tally.RootScope) Option {
	return func(svc Host) error {
		svc2 := svc.(*host)
		svc2.scope = scope
		return nil
	}
}

// WithObserver configures a service with an instance lifecycle observer
func WithObserver(observer Observer) Option {
	return func(svc Host) error {
		service := svc.(*host)
		service.observer = observer
		service.serviceCore.observer = observer
		return nil
	}
}
