package master

import "sync"

type TabletNodeList map[string]*TabletNode

type TabletNodeManager struct {
	mutex              sync.Mutex
	master             *Master
	tabletNodeList     TabletNodeList
	reconnectingTsList map[string]int64
}
