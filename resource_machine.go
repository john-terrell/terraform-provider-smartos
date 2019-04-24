package main

import (
	"fmt"
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
				},
				"boot": &schema.Schema{
					Type:     schema.TypeString,
					Optional: true,
				},
				"bootrom": &schema.Schema{
					Type:     schema.TypeString,
					Optional: true,
				},
			*/
			"brand": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"cpu_cap": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"cpu_shares": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			/*
				"cpu_type": &schema.Schema{
					Type:     schema.TypeString,
					Optional: true,
				},
			*/
			"customer_metadata": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"delegate_dataset": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			// 'disks.*' - disk object array
			/*
				"disk_driver": &schema.Schema{
					Type:     schema.TypeString,
					Optional: true,
				},
				"do_not_inventory": &schema.Schema{
					Type:     schema.TypeBool,
					Optional: true,
				},
			*/
			"dns_domain": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			// "filesystems.*"
			"firewall_enabled": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			/*
				"flexible_disk_size": &schema.Schema{
					Type:     schema.TypeInt,
					Optional: true,
				},
				"fs_allowed": &schema.Schema{
					Type:     schema.TypeString,
					Optional: true,
				},
			*/
			"hostname": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"image_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
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
			},
			/*
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
			"max_swap": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			/*
				"mdata_exec_timeout": &schema.Schema{
					Type:     schema.TypeInt,
					Optional: true,
				},
				// "nics.*"
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
			"quota": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"ram": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
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
	id, err := uuid.Parse(d.Get("image_uuid").(string))
	if err == nil {
		machine := Machine{
			// Set the required items
			Alias:     d.Get("alias").(string),
			Brand:     d.Get("brand").(string),
			ImageUUID: id,
		}

		// Set the optional items
		if autoboot, ok := d.GetOk("autoboot"); ok {
			machine.Autoboot = autoboot.(bool)
		}

		if cpuCap, ok := d.GetOk("cpu_cap"); ok {
			machine.CPUCap = cpuCap.(uint32)
		}

		client := m.(*SmartOSClient)
		if client == nil {
			return fmt.Errorf("Client is NULL")
		}

		uuid, err := client.CreateMachine(&machine)
		if err == nil {
			d.SetId(uuid.String())
		}
	}

	return err
}

func resourceMachineRead(d *schema.ResourceData, m interface{}) error {
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
