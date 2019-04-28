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
NOTE: Currently, only native OS virtual machines are supported (LX, KVM and Bhyve machines will be supported next)

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

- smartos_image (Data source - the images must already have been imported using 'imgadm import')
- smartos_machine (Resource)

### Example ###

The following example shows you how to configure a simple zone running Illumos (base-64-lts from Joyent)

```hcl
provider "smartos" {
    "host" = "10.99.50.60:22"
    "user" = "root"
}

data "smartos_image" "illumos" {
    "name" = "base-64-lts"
    "version"  = "18.4.0"
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

```
