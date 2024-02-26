package master

import (
	"sync"

	pb "github.com/allen-shaw/bigtable/internal/proto"
)

type TabletManager struct {
	allTables  TabletList
	mutex      sync.Mutex
	master     *Master
	metaTablet *MetaTablet
}

func NewTabletManager(master *Master) *TabletManager {
	tm := &TabletManager{
		allTables: make(TabletList),
		master:    master,
	}
	return tm
}

func (tm *TabletManager) Init() error {
	return nil
}

func (tm *TabletManager) Stop() {}

func (tm *TabletManager) Destory() {
	tm.ClearTableList()
}

func (tm *TabletManager) CreateTable(meta *pb.TabletMeta) *Table {
	t := &Table{}

	return t
}

func (tm *TabletManager) ClearTableList() {

}
