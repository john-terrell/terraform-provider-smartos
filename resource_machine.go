package main

import (
	"log"

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
			"metadata": &schema.Schema{
				Type:     schema.TypeMap,
				Computed: true,
			},
			/*
				"delegate_dataset": &schema.Schema{
					Type:     schema.TypeBool,
					Optional: true,
				},
			*/
			"disks": {
				Type:     schema.TypeList,
				Optional: true,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"boot": &schema.Schema{
							Type:     schema.TypeBool,
							Optional: true,
							ForceNew: true,
						},
						"compression": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
						},
						"image_uuid": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
						},
						"image_size": &schema.Schema{ // in MiB
							Type:     schema.TypeInt,
							Optional: true,
							ForceNew: true,
						},
						"model": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
						},
						"size": &schema.Schema{ // in MiB
							Type:     schema.TypeInt,
							Optional: true,
							ForceNew: true,
						},
					},
				},
			},
			/*
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
				Optional: true,
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
			*/
			"kernel_version": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			/*
				"limit_priv": &schema.Schema{
					Type:     schema.TypeString,
					Optional: true,
				},
			*/
			"maintain_resolvers": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			/*
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
			"max_swap": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			/*
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
						"allow_restricted_traffic": &schema.Schema{
							Type:     schema.TypeBool,
							Optional: true,
							ForceNew: true,
						},
						"allow_ip_spoofing": &schema.Schema{
							Type:     schema.TypeBool,
							Optional: true,
							ForceNew: true,
						},
						"allow_mac_spoofing": &schema.Schema{
							Type:     schema.TypeBool,
							Optional: true,
							ForceNew: true,
						},
						"gateways": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							ForceNew: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"interface": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"ips": &schema.Schema{
							Type:     schema.TypeList,
							Required: true,
							ForceNew: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"nic_tag": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"model": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
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
			*/
			"primary_ip": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
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
	log.Printf("---------------- MachineCreate")
	d.SetId("")

	client := m.(*SmartOSClient)
	machine := Machine{}
	err := machine.LoadFromSchema(d)
	if err != nil {
		return err
	}

	uuid, err := client.CreateMachine(&machine)
	if err != nil {
		return err
	}

	d.SetId(uuid.String())

	err = resourceMachineRead(d, m)
	log.Printf("---------------- MachineCreate (COMPLETE)")
	return err
}

func resourceMachineRead(d *schema.ResourceData, m interface{}) error {
	log.Printf("---------------- MachineRead")
	client := m.(*SmartOSClient)
	uuid, err := uuid.Parse(d.Id())
	if err != nil {
		log.Printf("Failed to parse incoming ID: %s", err)
		return err
	}

	machine, err := client.GetMachine(uuid)
	if err != nil {
		log.Printf("Failed to retrieve machine with ID %s.  Error: %s", d.Id(), err)
		return err
	}

	err = machine.SaveToSchema(d)
	log.Printf("---------------- MachineRead (COMPLETE)")
	return err
}

func resourceMachineUpdate(d *schema.ResourceData, m interface{}) error {
	log.Printf("---------------- MachineUpdate")
	machineId, err := uuid.Parse(d.Id())
	if err != nil {
		return err
	}

	d.Partial(true)

	machineUpdate := Machine{
		ID: &machineId,
	}

	updatesRequired := false

	if d.HasChange("alias") && !d.IsNewResource() {
		_, newValue := d.GetChange("alias")

		machineUpdate.Alias = newValue.(string)
		updatesRequired = true
	}

	if d.HasChange("autoboot") && !d.IsNewResource() {
		_, newValue := d.GetChange("autoboot")

		machineUpdate.Autoboot = newBool(newValue.(bool))
		updatesRequired = true
	}

	if d.HasChange("cpu_cap") && !d.IsNewResource() {
		_, newValue := d.GetChange("cpu_cap")

		machineUpdate.CPUCap = newUint32(uint32(newValue.(int)))
		updatesRequired = true
	}

	if d.HasChange("customer_metadata") && !d.IsNewResource() {
		oldSchemaValue, newSchemaValue := d.GetChange("customer_metadata")
		oldMap := oldSchemaValue.(map[string]interface{})
		newMap := newSchemaValue.(map[string]interface{})

		var addItem func(key string, value interface{}) = machineUpdate.setCustomerMetadata
		var removeItem func(key string) = machineUpdate.removeCustomerMetadata

		if ReconcileMaps(oldMap, newMap, addItem, addItem, removeItem, stringsAreEqual) {
			updatesRequired = true
		}
	}

	if d.HasChange("maintain_resolvers") && !d.IsNewResource() {
		_, newValue := d.GetChange("maintain_resolvers")

		machineUpdate.MaintainResolvers = newBool(newValue.(bool))
		updatesRequired = true
	}

	if d.HasChange("max_physical_memory") && !d.IsNewResource() {
		_, newValue := d.GetChange("max_physical_memory")

		machineUpdate.MaxPhysicalMemory = newUint32(uint32(newValue.(int)))
		updatesRequired = true
	}

	if d.HasChange("max_swap") && !d.IsNewResource() {
		_, newValue := d.GetChange("max_swap")

		machineUpdate.MaxSwap = newUint32(uint32(newValue.(int)))
		updatesRequired = true
	}

	if d.HasChange("quota") && !d.IsNewResource() {
		_, newValue := d.GetChange("quota")

		machineUpdate.Quota = newUint32(uint32(newValue.(int)))
		updatesRequired = true
	}

	if d.HasChange("resolvers") && !d.IsNewResource() {
		_, newSchemaValue := d.GetChange("resolvers")

		var resolvers []string
		for _, resolver := range newSchemaValue.([]interface{}) {
			resolvers = append(resolvers, resolver.(string))
		}
		machineUpdate.Resolvers = resolvers
		updatesRequired = true
	}

	if d.HasChange("nics") && !d.IsNewResource() {
		_, newSchemaValue := d.GetChange("nics")

		var nics []NetworkInterface
		for _, nic := range newSchemaValue.([]interface{}) {
			nics = append(nics, nic.(NetworkInterface))
		}
		machineUpdate.NetworkInterfaces = nics
		updatesRequired = true
	}

	if updatesRequired {
		client := m.(*SmartOSClient)

		err = client.UpdateMachine(&machineUpdate)
		if err != nil {
			return err
		}
	}

	d.Partial(false)
	err = resourceMachineRead(d, m)
	log.Printf("---------------- MachineUpdate (COMPLETE)")
	return err
}

func resourceMachineDelete(d *schema.ResourceData, m interface{}) error {
	log.Printf("Request to delete machine with ID: %s\n", d.Id())

	client := m.(*SmartOSClient)
	machineId, err := uuid.Parse(d.Id())
	if err != nil {
		return err
	}

	return client.DeleteMachine(machineId)
}
