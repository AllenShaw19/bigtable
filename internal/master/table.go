package master

import (
	"sync"

	"github.com/allen-shaw/bigtable/internal/metric"
	pb "github.com/allen-shaw/bigtable/internal/proto"
)

type TableMetric struct {
	tableName string
	tabletNum metric.MetricCounter
}

type TabletList map[string]*Tablet

type Table struct {
	mutex            sync.Mutex
	name             string
	tabletList       TabletList
	schema           *pb.TableSchema
	snapshotList     []uint64
	rollbackNames    []string
	deletedTabletNum uint32
	maxTabletNo      uint64
	createTime       int64
	counter          *pb.TableCounter
	metric           *TableMetric
	schemaIsSyncing  bool
	oldSchema        *pb.TableSchema

	stateMachine TableStateMachine
}
