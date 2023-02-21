package server

type ManagerServer struct {
}

func NewManagerServer() (*ManagerServer, error) {
	return &ManagerServer{}, nil
}

func (m *ManagerServer) Run() error {
	return nil
}
