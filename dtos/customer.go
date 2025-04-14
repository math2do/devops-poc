package dtos

// ProvisionVMRequest ...
type ProvisionVMRequest struct {
	VMName  string `json:"vm_name"`
	VMRAM   string `json:"vm_ram"`
	VMVcpus string `json:"vm_vcpus"`
	OsType  string `json:"os_type"`
}

// ProvisionVMResponse ...
type ProvisionVMResponse struct {
	VMID string `json:"vm_id"`
}
