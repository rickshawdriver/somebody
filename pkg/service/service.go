package service

type Service struct {
	ID       uint32 `json:"id"`
	EndPoint string `json:"endpoint"`
	MaxQps   uint32 `json:"maxQps"` // support max qps
}

func (d *Dispatcher) addService(service *Service) error {
	d.Services[service.ID] = service

	return nil
}

func (s *Service) GetID() uint32 {
	return s.ID
}
