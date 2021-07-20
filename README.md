NOTE: This provider has been superceded by:

https://github.com/CoolpeopleNetworks/terraform-provider-smartos.git

This repository will no longer be updated.


SmartOS Terraform Provider
=========================

<img src="https://cdn.rawgit.com/hashicorp/terraform-website/master/content/source/assets/images/logo-hashicorp.svg" width="600px">

Requirements
------------

-	[Terraform](https://www.terraform.io/downloads.html) 0.10.x
-	[Go](https://golang.org/doc/install) 1.9 (to build the provider plugin)

Using the provider
------------------

This provider can be used to provision machines with a SmartOS host via SSH.  SSH public keys are expected to already be installed on the SmartOS host in order for this provider to work.   

NOTE: Currently, this provider only supports a subset of properties for SmartOS virtual machines.

### Setup ###

```hcl
provider "smartos" {
    "host" = "10.99.50.60:22"
    "user" = "root"
}
```

The following arguments are supported.

- `host` - (Required) This is the address of the global zone on the SmartOS host.
- `user` - (Required) This is the authenticated SSH user which will run provisioning commands.   Normally this is 'root'.

### Resources and Data Providers ###

Currently, the following data and resources are provided:

- smartos_image (Data source - the images will be imported on first use by a smartos_machine stanza.)
- smartos_machine (Resource)

NOTE: The property names supported by this provider match (as much as possible) those defined by Joyent for use with their 'vmadm' utility.   See the man page (specifically the PROPERTIES section) for that utility for more info:

https://smartos.org/man/1m/vmadm

Many of the properties defined in the man page are not yet supported by the provider.

### Example ###

The following example shows you how to configure two simple zones - one running Illumos (base-64-lts from Joyent) and the other running Ubuntu 16.04.

(See the included sample.tf)

```hcl
provider "smartos" {
    "host" = "10.99.50.60:22"
    "user" = "root"
}

data "smartos_image" "illumos" {
    "name" = "base-64-lts"
    "version"  = "18.4.0"
}

data "smartos_image" "linux" {
    "name" = "ubuntu-16.04"
    "version"  = "20170403"
}

data "smartos_image" "linux_kvm" {
    "name" = "ubuntu-certified-16.04"
    "version" = "20190212"
}

resource "smartos_machine" "illumos" {
    "alias" = "illumos"
    "brand" = "joyent"
    "cpu_cap" = 100

    # These fields are required in order for provisioning (below) to function.
    "customer_metadata" = {
        "root_authorized_keys" = "... copy this from your ~/.ssh/id_rsa.pub ..."
        "user-script" = "/usr/sbin/mdata-get root_authorized_keys > ~root/.ssh/authorized_keys"
    }

    "image_uuid" = "${data.smartos_image.illumos.id}"
    "maintain_resolvers" = true
    "max_physical_memory" = 512
    "nics" = [
        {
            "nic_tag" = "external"
            "ips" = ["10.0.222.222/16"]
            "gateways" = ["10.0.0.1"]
            "interface" = "net4"
        }
    ]
    "quota" = 25

    "resolvers" = ["1.1.1.1", "1.0.0.1"]

    provisioner "remote-exec" {
        inline = [
            "pkgin -y update",
            "pkgin -y in htop",
        ]
    }
}

resource "smartos_machine" "linux" {
    "alias" = "provider-test-linux"
    "brand" = "lx"
    "kernel_version" = "3.16.0"
    "cpu_cap" = 100

    "customer_metadata" = {
        # Note: this is my public SSH key...use your own.  :-)
        "root_authorized_keys" = "... copy this from your ~/.ssh/id_rsa.pub ..."
        "user-script" = "/usr/sbin/mdata-get root_authorized_keys > ~root/.ssh/authorized_keys"
    }

    "image_uuid" = "${data.smartos_image.linux.id}"
    "maintain_resolvers" = true
    "max_physical_memory" = 512
    "nics" = [
        {
            "nic_tag" = "external"
            "ips" = ["10.0.222.223/16"]
            "gateways" = ["10.0.0.1"]
            "interface" = "net5"
        }
    ]
    "quota" = 25

    "resolvers" = ["1.1.1.1", "1.0.0.1"]

    provisioner "remote-exec" {
        inline = [
            "apt-get update",
            "apt-get -y install htop",
        ]
    }
}

resource "smartos_machine" "linux-kvm" {
    "alias" = "provider-test-linux-kvm"
    "brand" = "kvm"
    "kernel_version" = "3.16.0"
    "vcpus" = 2

    "customer_metadata" = {
        # Note: this is my public SSH key...use your own.  :-)
        "root_authorized_keys" = "... copy this from your ~/.ssh/id_rsa.pub ..."
    }

    "maintain_resolvers" = true
    "ram" = 512
    "nics" = [
        {
            "nic_tag" = "external"
            "ips" = ["10.0.222.224/16"]
            "gateways" = ["10.0.0.1"]
            "interface" = "net0"
            "model" = "virtio"
        }
    ]
    "quota" = 25

    "resolvers" = ["1.1.1.1", "1.0.0.1"]

    "disks" = [
        {
            "boot" = true
            "image_uuid" = "${data.smartos_image.linux_kvm.id}"
            "compression" = "lz4"
            "model" = "virtio"
        }
    ]

    provisioner "remote-exec" {
        inline = [
            "apt-get update",
            "apt-get -y install htop",
        ]
    }
}

```
