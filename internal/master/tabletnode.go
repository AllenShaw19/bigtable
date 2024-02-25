package master

import "sync"

type NodeState int

type TabletNode struct {
	mutex     sync.Mutex
	addr      string
	uuid      string
	state     NodeState
	timestamp int64
}
