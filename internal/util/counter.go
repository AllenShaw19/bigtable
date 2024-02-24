package util

import "sync/atomic"

type Counter struct {
	val int64
}

func NewCount(val int64) Counter {
	return Counter{val: val}
}

func (c *Counter) Add(v int64) int64 {
	return atomic.AddInt64(&c.val, v)
}

func (c *Counter) Sub(v int64) int64 {
	return atomic.AddInt64(&c.val, -v)
}

func (c *Counter) Inc() int64 {
	return atomic.AddInt64(&c.val, 1)
}

func (c *Counter) Dec() int64 {
	return atomic.AddInt64(&c.val, -1)
}

func (c *Counter) Set(v int64) bool {
	return atomic.CompareAndSwapInt64(&c.val, c.val, v)
}

func (c *Counter) Clear() bool {
	return atomic.CompareAndSwapInt64(&c.val, c.val, 0)
}
