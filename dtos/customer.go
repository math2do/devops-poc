package dtos

// ProvisionVMRequest ...
type ProvisionVMRequest struct {
	VMName        string `json:"vm_name"`
	VMNetwork     string `json:"vm_network"`
	VMMemory      string `json:"vm_memory"`
	VMVCPUs       string `json:"vm_vcpus"`
	VMZone        string `json:"vm_zone"`
	CorrelationID string `json:"corelation_id"`
}

// ProvisionVMResponse ...
type ProvisionVMResponse struct {
	CorrelationID string `json:"corelation_id"`
}
