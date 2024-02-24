package master

import (
	"sync"

	pb "github.com/allen-shaw/bigtable/internal/proto"
)

type Tablet struct {
	mutex        sync.Mutex
	meta         pb.TabletMeta
	stateMachine TabletStateMachine

	node   *TabletNode
	tablet *Table

	updateTime      int64
	lastMoveTimeUs  int64
	dataSizeOnFlash int64
	serverID        string
}
