package customer

import (
	"log"

	"github.com/google/uuid"
	"github.com/veltris/devops-poc/dtos"
	"github.com/veltris/devops-poc/services/kubernetes"
)

// Service ...
type Service struct {
	k8sClient *kubernetes.K8sClient
	log       *log.Logger
}

// NewService ...
func NewService(log *log.Logger, k8sClient *kubernetes.K8sClient) *Service {
	return &Service{
		k8sClient: k8sClient,
		log:       log,
	}
}

// ProvisionVMs ...
func (s *Service) ProvisionVMs(req *dtos.ProvisionVMRequest) (*dtos.ProvisionVMResponse, error) {

	// create unique job per request
	uuid := uuid.New().String()
	err := s.k8sClient.LaunchK8sJob("hello-"+uuid, "bash:5.1",

		[]string{"/bin/bash", "-c"}, []string{
			`
		echo "Hello from inside the job!";
		echo "Listing files:";
		ls -la;
		`,
		})

	if err != nil {
		return nil, err
	}

	return &dtos.ProvisionVMResponse{VMID: "sample"}, nil
}
