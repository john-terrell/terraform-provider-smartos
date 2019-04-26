package main

import (
	"github.com/google/uuid"
)

type Machine struct {
	ID                         *uuid.UUID         `json:"uuid,omitempty"`
	Alias                      string             `json:"alias"`
	Autoboot                   bool               `json:"autoboot,omitempty"`
	Brand                      string             `json:"brand"`
	CPUCap                     uint32             `json:"cpu_cap,omitempty"`
	CPUShares                  uint32             `json:"cpu_shares,omitempty"`
	CustomerMetadata           map[string]string  `json:"customer_metadata,omitempty"`
	DelegateDataset            bool               `json:"delegate_dataset,omitempty"`
	DNSDomain                  string             `json:"dns_domain,omitempty"`
	FirewallEnabled            bool               `json:"firewall_enabled,omitempty"`
	Hostname                   string             `json:"hostname,omitempty"`
	ImageUUID                  uuid.UUID          `json:"image_uuid"`
	InternalMetadata           map[string]string  `json:"internal_metadat,omitempty"`
	InternalMetadataNamespaces map[string]string  `json:"internal_metadata_namespaces,omitempty"`
	IndestructableDelegated    bool               `json:"indestructible_delegated,omitempty"`
	IndestructableZoneRoot     bool               `json:"indestructible_zoneroot,omitempty"`
	KernelVersion              string             `json:"kernel_version,omitempty"`
	MaxPhysicalMemory          uint32             `json:"max_physical_memory,omitempty"`
	MaxSwap                    uint32             `json:"max_swap,omitempty"`
	NetworkInterfaces          []NetworkInterface `json:"nics,omitempty"`
	Quota                      uint32             `json:"quota,omitempty"`
	RAM                        uint32             `json:"ram,omitempty"`
	Resolvers                  []string           `json:"resolvers,omitempty"`
	VirtualCPUCount            uint16             `json:"vcpus,omitempty"`
	State                      string             `json:"state,omitempty"`

	PrimaryIP string
}

func (m *Machine) UpdatePrimaryIP() {
	m.PrimaryIP = ""
	for _, networkInterface := range m.NetworkInterfaces {
		if networkInterface.IsPrimary {
			m.PrimaryIP = networkInterface.IPAddress
			break
		}
	}
}

type NetworkInterface struct {
	AllowDHCPSpoofing     bool     `json:"allow_dhcp_spoofing,omitempty"`
	AllowIPSpoofing       bool     `json:"allow_ip_spoofing,omitempty"`
	AllowMACSpoofing      bool     `json:"allow_mac_spoofing,omitempty"`
	AllowRestrictedTrafic bool     `json:"allow_restricted_traffic,omitempty"`
	BlockedOutgoingPorts  []uint16 `json:"blocked_outgoing_ports,omitempty"`
	Gateways              []string `json:"gateways,omitempty"`
	Interface             string   `json:"interface,omitempty"`
	IPAddresses           []string `json:"ips,omitempty"`
	IPAddress             string   `json:"ip,omitempty"`
	HardwareAddress       string   `json:"mac,omitempty"`
	Model                 string   `json:"model,omitempty"`
	Tag                   string   `json:"nic_tag,omitempty"`
	IsPrimary             bool     `json:"primary,omitempty"`
	VirtualLANID          uint16   `json:"vlan_id,omitempty"`
}
