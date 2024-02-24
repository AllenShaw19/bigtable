package master

import (
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
