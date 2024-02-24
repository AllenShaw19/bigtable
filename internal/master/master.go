package master

import (
	pb "github.com/allen-shaw/bigtable/internal/proto"
	"github.com/allen-shaw/bigtable/internal/util"
)

type Master struct {
	stateMachine MasterStateMachine

	tabletManager     *TabletManager
	tabletNodeManager *TabletNodeManager
	userManager       *UserManager
	sizeScheduler     Scheduler
	loadScheduler     Scheduler

	// election

	//
	sequenceID util.Counter

	// access entry

}

func NewMaster() *Master {
	m := &Master{}

	m.stateMachine = NewMasterStateMachine(IsSecondary)

	// m.tabletManager = n

	return m
}

func (m *Master) Init() error {

	return nil
}

func (m *Master) CreateTable(req *pb.CreateTableReq) (*pb.CreateTableResp, error) {

	resp := &pb.CreateTableResp{}
	return resp, nil
}
