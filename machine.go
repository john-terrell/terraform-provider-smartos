package main

import (
	"github.com/google/uuid"
	"github.com/hashicorp/terraform/helper/schema"
)

type Machine struct {
	ID       *uuid.UUID `json:"uuid,omitempty"`
	Alias    string     `json:"alias,omitempty"`
	Autoboot *bool      `json:"autoboot,omitempty"`
	Brand    string     `json:"brand,omitempty"`
	CPUCap   *uint32    `json:"cpu_cap,omitempty"`
	/*
		CPUShares                  uint32             `json:"cpu_shares,omitempty"`
	*/
	CustomerMetadata       map[string]string `json:"customer_metadata,omitempty"`
	SetCustomerMetadata    map[string]string `json:"set_customer_metadata,omitempty"`    // for updates
	RemoveCustomerMetadata []string          `json:"remove_customer_metadata,omitempty"` // for updates

	/*
		DelegateDataset            bool               `json:"delegate_dataset,omitempty"`
		DNSDomain                  string             `json:"dns_domain,omitempty"`
		FirewallEnabled            bool               `json:"firewall_enabled,omitempty"`
		Hostname                   string             `json:"hostname,omitempty"`
	*/
	ImageUUID *uuid.UUID `json:"image_uuid,omitempty"`
	/*
		InternalMetadata           map[string]string  `json:"internal_metadat,omitempty"`
		InternalMetadataNamespaces map[string]string  `json:"internal_metadata_namespaces,omitempty"`
		IndestructableDelegated    bool               `json:"indestructible_delegated,omitempty"`
		IndestructableZoneRoot     bool               `json:"indestructible_zoneroot,omitempty"`
		KernelVersion              string             `json:"kernel_version,omitempty"`
	*/
	MaxPhysicalMemory *uint32 `json:"max_physical_memory,omitempty"`
	/*
		MaxSwap           uint32             `json:"max_swap,omitempty"`
	*/
	NetworkInterfaces []NetworkInterface `json:"nics,omitempty"`
	Quota             *uint32            `json:"quota,omitempty"`
	/*
		RAM               uint32             `json:"ram,omitempty"`
	*/
	Resolvers []string `json:"resolvers,omitempty"`
	/*
		VirtualCPUCount *uint16  `json:"vcpus,omitempty"`
	*/
	State     string `json:"state,omitempty"`
	PrimaryIP string `json:"-"`
}

func (m *Machine) UpdatePrimaryIP() {
	m.PrimaryIP = ""
	for _, networkInterface := range m.NetworkInterfaces {
		if networkInterface.IsPrimary != nil {
			m.PrimaryIP = networkInterface.IPAddress
			break
		}
	}
}

func newBool(value bool) *bool {
	n := value
	return &n
}

func newUint32(value uint32) *uint32 {
	n := value
	return &n
}

func newStringMap() *map[string]string {
	var n map[string]string
	return &n
}

func (m *Machine) LoadFromSchema(d *schema.ResourceData) error {

	m.Alias = d.Get("alias").(string)
	m.Brand = d.Get("brand").(string)

	imageUUID, err := uuid.Parse(d.Get("image_uuid").(string))
	if err != nil {
		return err
	}
	m.ImageUUID = &imageUUID

	if autoboot, ok := d.GetOk("autoboot"); ok {
		m.Autoboot = newBool(autoboot.(bool))
	}

	if cpuCap, ok := d.GetOk("cpu_cap"); ok {
		m.CPUCap = newUint32(uint32(cpuCap.(int)))
	}

	customerMetaData := map[string]string{}
	for k, v := range d.Get("customer_metadata").(map[string]interface{}) {
		customerMetaData[k] = v.(string)
	}
	m.CustomerMetadata = customerMetaData

	if maxPhysicalMemory, ok := d.GetOk("max_physical_memory"); ok {
		m.MaxPhysicalMemory = newUint32(uint32(maxPhysicalMemory.(int)))
	}

	if nics, ok := d.GetOk("nics"); ok {
		m.NetworkInterfaces, err = getNetworkInterfaces(nics)
	}

	if quota, ok := d.GetOk("quota"); ok {
		m.Quota = newUint32(uint32(quota.(int)))
	}

	for _, resolver := range d.Get("resolvers").([]interface{}) {
		m.Resolvers = append(m.Resolvers, resolver.(string))
	}

	return nil
}

func (m *Machine) SaveToSchema(d *schema.ResourceData) error {
	d.Set("primary_ip", m.PrimaryIP)
	d.Set("id", m.ID.String())

	if m.PrimaryIP != "" {
		d.SetConnInfo(map[string]string{
			"type": "ssh",
			"host": m.PrimaryIP,
		})
	}

	return nil
}

func (m *Machine) setCustomerMetadata(key string, value interface{}) {
	if m.SetCustomerMetadata == nil {
		m.SetCustomerMetadata = make(map[string]string)
	}

	m.SetCustomerMetadata[key] = value.(string)
}

func (m *Machine) removeCustomerMetadata(key string) {
	m.RemoveCustomerMetadata = append(m.RemoveCustomerMetadata, key)
}

func stringsAreEqual(a interface{}, b interface{}) bool {
	return a.(string) == b.(string)
}

type NetworkInterface struct {
	/*
		AllowDHCPSpoofing     bool     `json:"allow_dhcp_spoofing,omitempty"`
		AllowIPSpoofing       bool     `json:"allow_ip_spoofing,omitempty"`
		AllowMACSpoofing      bool     `json:"allow_mac_spoofing,omitempty"`
		AllowRestrictedTrafic bool     `json:"allow_restricted_traffic,omitempty"`
		BlockedOutgoingPorts  []uint16 `json:"blocked_outgoing_ports,omitempty"`
	*/
	Gateways    []string `json:"gateways,omitempty"`
	Interface   string   `json:"interface,omitempty"`
	IPAddresses []string `json:"ips,omitempty"`
	IPAddress   string   `json:"ip,omitempty"`
	/*
		HardwareAddress       string   `json:"mac,omitempty"`
		Model                 string   `json:"model,omitempty"`
	*/
	Tag          string `json:"nic_tag,omitempty"`
	IsPrimary    *bool  `json:"primary,omitempty"`
	VirtualLANID uint16 `json:"vlan_id,omitempty"`
}

func getNetworkInterfaces(d interface{}) ([]NetworkInterface, error) {
	networkInterfaceDefinitions := d.([]interface{})

	var networkInterfaces []NetworkInterface

	for _, nid := range networkInterfaceDefinitions {
		networkInterfaceDefinition := nid.(map[string]interface{})

		var gateways []string
		for _, gateway := range networkInterfaceDefinition["gateways"].([]interface{}) {
			gateways = append(gateways, gateway.(string))
		}

		interfaceName := networkInterfaceDefinition["interface"].(string)

		var ips []string
		for _, ip := range networkInterfaceDefinition["ips"].([]interface{}) {
			ips = append(ips, ip.(string))
		}

		nicTag := networkInterfaceDefinition["nic_tag"].(string)

		var vlanID uint16
		if vlanIDCheck, ok := networkInterfaceDefinition["vlan_id"].(int); ok {
			vlanID = uint16(vlanIDCheck)
		}

		networkInterface := NetworkInterface{
			Interface:    interfaceName,
			IPAddresses:  ips,
			Tag:          nicTag,
			Gateways:     gateways,
			VirtualLANID: vlanID,
		}

		networkInterfaces = append(networkInterfaces, networkInterface)
	}

	return networkInterfaces, nil
}
