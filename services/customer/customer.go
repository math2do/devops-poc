package customer

import (
	"log"

	"github.com/veltris/devops-poc/dtos"
)

// Service ...
type Service struct {
	log *log.Logger
}

// NewService ...
func NewService(log *log.Logger) *Service {
	return &Service{
		log: log,
	}
}

// ProvisionVMs ...
func (s *Service) ProvisionVMs(req *dtos.ProvisionVMRequest) (*dtos.ProvisionVMResponse, error) {
	return &dtos.ProvisionVMResponse{VMID: "sample"}, nil
}
