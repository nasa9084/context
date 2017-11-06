// Package context wraps stdlib context package and
// add some functional context like context.withValueContext.
package context

import "context"

// stdlib context aliases

var (
	// Canceled is the error returned by Context.Err when the context is canceled.
	Canceled = context.Canceled
	// DeadlineExceeded is the error returned by Context.Err when the context's deadline passes.
	DeadlineExceeded = context.DeadlineExceeded
)

type (
	// A CancelFunc tells an operation to abandon its work. A CancelFunc does not wait for the work to stop.
	// After the first call, subsequent calls to a CancelFunc do nothing.
	CancelFunc = context.CancelFunc
	// A Context carries a deadline, a cancelation signal, and other values across API boundaries.
	// Context's methods may be called by multiple goroutines simultaneously.
	Context = context.Context
)

var (
	// Background returns a non-nil, empty Context.
	Background = context.Background
	// TODO returns a non-nil, empty Context.
	TODO = context.TODO
	// WithCancel returns a copy of parent with a new Done channel.
	WithCancel = context.WithCancel
	// WithDeadline returns a copy of the parent context with the deadline adjusted
	// to be no later than d.
	WithDeadline = context.WithDeadline
	// WithTimeout returns WithDeadline(parent, time.Now().Add(timeout)).
	WithTimeout = context.WithTimeout
	// WithValue returns a copy of parent in which the value associated with key is
	// val.
	WithValue = context.WithValue
)
