package master

import (
	pb "github.com/allen-shaw/bigtable/internal/proto"
)

type MasterStatus pb.StatusCode

const (
	IsSecondary MasterStatus = MasterStatus(pb.StatusCode_MasterIsSecondary)
	IsRunning   MasterStatus = MasterStatus(pb.StatusCode_MasterIsRunning)
	OnRestore   MasterStatus = MasterStatus(pb.StatusCode_MasterOnRestore)
	OnWait      MasterStatus = MasterStatus(pb.StatusCode_MasterOnWait)
	IsReadonly  MasterStatus = MasterStatus(pb.StatusCode_MasterIsReadonly)
)

type MasterStateMachine struct {
	currentStatus MasterStatus
}

func NewMasterStateMachine(initStatus MasterStatus) MasterStateMachine {
	return MasterStateMachine{
		currentStatus: initStatus,
	}
}

func (sm *MasterStateMachine) GetStatus() MasterStatus {
	return sm.currentStatus
}
