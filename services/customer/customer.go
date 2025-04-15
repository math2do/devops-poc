package customer

import (
	"log"

	"github.com/google/uuid"
	"github.com/veltris/devops-poc/dtos"
	"github.com/veltris/devops-poc/services/kubernetes"
	v1 "k8s.io/api/core/v1"
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
	err := s.k8sClient.LaunchK8sJob("infra-provisioner-job-"+uuid, "sriramgopalvarma/infra-provisioner:v1",

		[]string{"/bin/bash", "-c"}, []string{
			`
			./run.sh
		`,
		},
		[]v1.EnvVar{
			{
				Name:  "TF_VAR_vm_name",
				Value: req.VMName,
			},
			{
				Name:  "TF_VAR_network_name",
				Value: req.VMNetwork,
			},
			{
				Name:  "TF_VAR_vm_memory",
				Value: req.VMMemory,
			},
			{
				Name:  "TF_VAR_vm_vcpus",
				Value: req.VMVCPUs,
			},
			{
				Name:  "TF_VAR_zone_name",
				Value: req.VMZone,
			},
			{
				Name:  "CORELATION_ID",
				Value: uuid,
			},
		},
	)

	if err != nil {
		return nil, err
	}

	return &dtos.ProvisionVMResponse{CorrelationID: uuid}, nil
}
