package master

import "sync/atomic"

type MasterServer struct {
	started atomic.Bool

	master *Master
}

func NewMasterServer() *MasterServer {

	ms := &MasterServer{}
	return ms
}

func (s *MasterServer) StartServer() bool {

	return true
}

func (s *MasterServer) Run() bool {

	return true
}

func (s *MasterServer) ShutdownServer() {

}
