package main

import (
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceMachine() *schema.Resource {
	return &schema.Resource{
		Create: resourceMachineCreate,
		Read:   resourceMachineRead,
		Update: resourceMachineUpdate,
		Delete: resourceMachineDelete,

		Schema: map[string]*schema.Schema{
			"alias": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			/*
				"archive_on_delete": &schema.Schema{
					Type:     schema.TypeBool,
					Optional: true,
				},
			*/
			"autoboot": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			/*
				"billing_id": &schema.Schema{
					Type:     schema.TypeString,
					Optional: true,
				},
				"bhyve_extra_opts": &schema.Schema{
					Type:     schema.TypeString,
					Optional: true,
					ForceNew: true,
				},
				"boot": &schema.Schema{
					Type:     schema.TypeString,
					Optional: true,
				},
				"bootrom": &schema.Schema{
					Type:     schema.TypeString,
					Optional: true,
					ForceNew: true,
				},
			*/
			"brand": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"cpu_cap": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			/*
				"cpu_shares": &schema.Schema{
					Type:     schema.TypeInt,
					Optional: true,
				},
					"cpu_type": &schema.Schema{
						Type:     schema.TypeString,
						Optional: true,
						ForceNew: true,
					},
			*/
			"customer_metadata": &schema.Schema{
				Type:     schema.TypeMap,
				Optional: true,
			},
			/*
				"delegate_dataset": &schema.Schema{
					Type:     schema.TypeBool,
					Optional: true,
				},
				// 'disks.*' - disk object array

				"disk_driver": &schema.Schema{
					Type:     schema.TypeString,
					Optional: true,
					ForceNew: true,
				},
				"do_not_inventory": &schema.Schema{
					Type:     schema.TypeBool,
					Optional: true,
				},
				"dns_domain": &schema.Schema{
					Type:     schema.TypeString,
					Optional: true,
				},
				// "filesystems.*"
				"firewall_enabled": &schema.Schema{
					Type:     schema.TypeBool,
					Optional: true,
				},
				"flexible_disk_size": &schema.Schema{
					Type:     schema.TypeInt,
					Optional: true,
				},
				"fs_allowed": &schema.Schema{
					Type:     schema.TypeString,
					Optional: true,
				},
				"hostname": &schema.Schema{
					Type:     schema.TypeString,
					Optional: true,
				},
			*/
			"image_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			/*
				"internal_metadata": &schema.Schema{
					Type:     schema.TypeList,
					Optional: true,
					Elem: &schema.Schema{
						Type: schema.TypeString,
					},
				},
				"internal_metadata_namespaces": &schema.Schema{
					Type:     schema.TypeList,
					Optional: true,
					Elem: &schema.Schema{
						Type: schema.TypeString,
					},
				},
				"indestructible_delegated": &schema.Schema{
					Type:     schema.TypeBool,
					Optional: true,
				},
				"indestructible_zoneroot": &schema.Schema{
					Type:     schema.TypeBool,
					Optional: true,
				},
				"kernel_version": &schema.Schema{
					Type:     schema.TypeString,
					Optional: true,
					ForceNew: true,
				},
				"limit_priv": &schema.Schema{
					Type:     schema.TypeString,
					Optional: true,
				},
				"maintain_resolvers": &schema.Schema{
					Type:     schema.TypeBool,
					Optional: true,
				},
				"max_locked_memory": &schema.Schema{
					Type:     schema.TypeInt,
					Optional: true,
				},
				"max_lwps": &schema.Schema{
					Type:     schema.TypeInt,
					Optional: true,
				},
			*/
			"max_physical_memory": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			/*
				"max_swap": &schema.Schema{
					Type:     schema.TypeInt,
					Optional: true,
				},
				"mdata_exec_timeout": &schema.Schema{
					Type:     schema.TypeInt,
					Optional: true,
				},
			*/
			"nics": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"gateways": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"interface": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
						"ips": &schema.Schema{
							Type:     schema.TypeList,
							Required: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"nic_tag": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},
			/*
				"nic_driver": &schema.Schema{
					Type:     schema.TypeString,
					Optional: true,
				},
				"nowait": &schema.Schema{
					Type:     schema.TypeBool,
					Optional: true,
				},
				"owner_uuid": &schema.Schema{
					Type:     schema.TypeString,
					Optional: true,
				},
				"qemu_opts": &schema.Schema{
					Type:     schema.TypeString,
					Optional: true,
				},
				"qemu_extra_opts": &schema.Schema{
					Type:     schema.TypeString,
					Optional: true,
				},
				"quota": &schema.Schema{
					Type:     schema.TypeInt,
					Optional: true,
				},
				"ram": &schema.Schema{
					Type:     schema.TypeInt,
					Optional: true,
					ForceNew: true,
				},
			*/
			"resolvers": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			// "routes.*" - object
			/*
				"spice_opts": &schema.Schema{
					Type:     schema.TypeString,
					Optional: true,
				},
				"spice_password": &schema.Schema{
					Type:     schema.TypeString,
					Optional: true,
				},
				"spice_port": &schema.Schema{
					Type:     schema.TypeString,
					Optional: true,
				},
				"tmpfs": &schema.Schema{
					Type:     schema.TypeInt,
					Optional: true,
				},
				"uuid": &schema.Schema{
					Type:     schema.TypeString,
					Optional: true,
				},
			*/
			"vcpus": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				ForceNew: true,
			},
			/*
				"vga": &schema.Schema{
					Type:     schema.TypeString,
					Optional: true,
				},
				"virtio_txburst": &schema.Schema{
					Type:     schema.TypeInt,
					Optional: true,
				},
				"virtio_txtimer": &schema.Schema{
					Type:     schema.TypeInt,
					Optional: true,
				},
				"vnc_password": &schema.Schema{
					Type:     schema.TypeString,
					Optional: true,
				},
				"vnc_port": &schema.Schema{
					Type:     schema.TypeInt,
					Optional: true,
				},
				"zfs_data_compression": &schema.Schema{
					Type:     schema.TypeString,
					Optional: true,
				},
				"zfs_data_recsize": &schema.Schema{
					Type:     schema.TypeInt,
					Optional: true,
				},
				"zfs_filesystem_limit": &schema.Schema{
					Type:     schema.TypeInt,
					Optional: true,
				},
				"zfs_io_priority": &schema.Schema{
					Type:     schema.TypeInt,
					Optional: true,
				},
				"zfs_root_compression": &schema.Schema{
					Type:     schema.TypeString,
					Optional: true,
				},
				"zfs_root_recsize": &schema.Schema{
					Type:     schema.TypeInt,
					Optional: true,
				},
				"zfs_snapshot_limit": &schema.Schema{
					Type:     schema.TypeInt,
					Optional: true,
				},
				"zlog_max_size": &schema.Schema{
					Type:     schema.TypeInt,
					Optional: true,
				},
				"zpool": &schema.Schema{
					Type:     schema.TypeString,
					Optional: true,
				},
			*/
		},
	}
}

func resourceMachineCreate(d *schema.ResourceData, m interface{}) error {
	d.SetId("")

	client := m.(*SmartOSClient)
	if client == nil {
		return fmt.Errorf("Client is NULL")
	}

	id, err := uuid.Parse(d.Get("image_uuid").(string))
	if err == nil {
		var uuid *uuid.UUID
		{
			alias := d.Get("alias").(string)
			brand := d.Get("brand").(string)

			machine := Machine{
				// Set the required items
				Alias:     alias,
				Brand:     brand,
				ImageUUID: id,
			}

			// Set the optional items
			if autoboot, ok := d.GetOk("autoboot"); ok {
				machine.Autoboot = autoboot.(bool)
			}

			if cpuCap, ok := d.GetOk("cpu_cap"); ok {
				machine.CPUCap = cpuCap.(uint32)
			}

			customerMetaData := map[string]string{}
			for k, v := range d.Get("customer_metadata").(map[string]interface{}) {
				customerMetaData[k] = v.(string)
			}
			machine.CustomerMetadata = customerMetaData

			if maxPhysicalMemory, ok := d.GetOk("max_physical_memory"); ok {
				machine.MaxPhysicalMemory = uint32(maxPhysicalMemory.(int))
			}

			if nics, ok := d.GetOk("nics"); ok {
				machine.NetworkInterfaces, err = getNetworkInterfaces(nics)
			}

			for _, resolver := range d.Get("resolvers").([]interface{}) {
				machine.Resolvers = append(machine.Resolvers, resolver.(string))
			}

			var err error
			uuid, err = client.CreateMachine(&machine)
			if err != nil {
				return err
			}
		}

		// Now, loop until provisioning is complete
		var createdMachine *Machine
		for {
			createdMachine, err = client.GetMachine(*uuid)
			if err != nil {
				_ = client.DeleteMachine(*uuid)
				return err
			}

			if createdMachine.State == "running" {
				break
			}

			log.Printf("Waiting for machine to enter running state. Current state: %s\n", createdMachine.State)
			time.Sleep(1 * time.Second)
		}

		createdMachine.UpdatePrimaryIP()
		log.Printf("Primary IP address updated to %s", createdMachine.PrimaryIP)

		// Set some important properties
		d.Set("primary_ip", createdMachine.PrimaryIP)
		d.Set("id", uuid.String())
		d.SetId(uuid.String())

		if createdMachine.PrimaryIP != "" {
			d.SetConnInfo(map[string]string{
				"type": "ssh",
				"host": createdMachine.PrimaryIP,
			})
		}
	}

	return err
}

func resourceMachineRead(d *schema.ResourceData, m interface{}) error {
	/*
		log.Printf("Request to read machine with ID: %s\n", d.Id())

		client := m.(*SmartOSClient)
		if client == nil {
			return fmt.Errorf("Client is NULL")
		}

		machineId, err := uuid.Parse(d.Id())
		if err != nil {
			return err
		}

		machine, err := client.GetMachine(machineId)
		if err != nil {
			return err
		}
	*/
	return nil
}

func resourceMachineUpdate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceMachineDelete(d *schema.ResourceData, m interface{}) error {
	log.Printf("Request to delete machine with ID: %s\n", d.Id())

	client := m.(*SmartOSClient)
	if client == nil {
		return fmt.Errorf("Client is NULL")
	}

	machineId, err := uuid.Parse(d.Id())
	if err != nil {
		return err
	}

	return client.DeleteMachine(machineId)
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
			// Required tags
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
