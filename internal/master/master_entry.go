package master

import "sync/atomic"

type MasterEntry struct {
	started atomic.Bool

	master *Master
}

func NewMasterEntry() *MasterEntry {

	ms := &MasterEntry{}
	return ms
}

func (s *MasterEntry) StartServer() bool {

	return true
}

func (s *MasterEntry) Run() bool {

	return true
}

func (s *MasterEntry) ShutdownServer() {

}
