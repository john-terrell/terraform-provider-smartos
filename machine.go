package main

import "github.com/google/uuid"

type Machine struct {
	ID                         *uuid.UUID `json:"uuid,omitempty"`
	Alias                      string     `json:"alias"`
	Autoboot                   bool       `json:"autoboot,omitempty"`
	Brand                      string     `json:"brand"`
	CPUCap                     uint32     `json:"cpu_cap,omitempty"`
	CPUShares                  uint32     `json:"cpu_shares,omitempty"`
	CustomerMetadata           []string   `json:"customer_metadata,omitempty"`
	DelegateDataset            bool       `json:"delegate_dataset,omitempty"`
	DNSDomain                  string     `json:"dns_domain,omitempty"`
	FirewallEnabled            bool       `json:"firewall_enabled,omitempty"`
	Hostname                   string     `json:"hostname,omitempty"`
	ImageUUID                  uuid.UUID  `json:"image_uuid"`
	InternalMetadata           []string   `json:"internal_metadat,omitempty"`
	InternalMetadataNamespaces []string   `json:"internal_metadata_namespaces,omitempty"`
	IndestructableDelegated    bool       `json:"indestructible_delegated,omitempty"`
	IndestructableZoneRoot     bool       `json:"indestructible_zoneroot,omitempty"`
	KernelVersion              string     `json:"kernel_version,omitempty"`
	MaxPhysicalMemory          uint32     `json:"max_physical_memory,omitempty"`
	MaxSwap                    uint32     `json:"max_swap,omitempty"`
	Quota                      uint32     `json:"quota,omitempty"`
	RAM                        uint32     `json:"ram,omitempty"`
	Resolvers                  []string   `json:"resolvers,omitempty"`
	VirtualCPUCount            uint16     `json:"vcpus,omitempty"`
}
