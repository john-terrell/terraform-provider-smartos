NOTE: This provider has been superceded by:

https://github.com/CoolpeopleNetworks/terraform-provider-smartos.git

This repository will no longer be updated.


SmartOS Terraform Provider
=========================

<img src="https://cdn.rawgit.com/hashicorp/terraform-website/master/content/source/assets/images/logo-hashicorp.svg" width="600px">

Requirements
------------

-	[Terraform](https://www.terraform.io/downloads.html) 0.12.x
-	[Go](https://golang.org/doc/install) >=1.9 (to build the provider plugin)


Building The Provider
---------------------
Make sure you have Go installed and your `$GOPATH` setup.

Clone the repo:

```sh
$ mkdir -p $GOPATH/src/github.com/john-terrell
$ cd $GOPATH/src/github.com/john-terrell
$ git clone https://github.com/john-terrell/terraform-provider-smartos.git
```

Build the provider

```sh
$ cd $GOPATH/src/github.com/john-terrell/terraform-provider-smartos
$ go get
$ make build
```

In your Terraform project, initialize it by passing the directory that contains the built provider binary.

```sh
$ cd ~/git/my_terra_project
$ terraform init --plugin-dir=$GOPATH/bin
```

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
  host = "10.99.50.60:22"
  user = "root"
}

data "smartos_image" "illumos" {
  name    = "base-64-lts"
  version = "18.4.0"
}

data "smartos_image" "linux" {
  name    = "ubuntu-16.04"
  version = "20170403"
}

data "smartos_image" "linux_kvm" {
  name    = "ubuntu-certified-16.04"
  version = "20190212"
}

resource "smartos_machine" "illumos" {
  alias   = "provider-test-illumos"
  brand   = "joyent"
  cpu_cap = 100

  customer_metadata = {
    # Note: this is my public SSH key...use your own.  :-)
    "root_authorized_keys" = "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDYIHv3DoXnAMn+dggUup1a+jjSqpZiIU5ThgljXHG9KM+iy1W3zo9qshUE7vBj/l7l5aHzRKyXsmWb6EdmtlVBnYl7SH5IMGaEFlB6n7T+yoMRl7VczZZxvP+VSAac2HeLPvdrCDeJCckfkHeTg9E3rt2PcAz0REKDCm34lpsedgM4QrVh8D54NgqLCdpT+QpidEBwE1T5wGMId4OwBB+r1VJZyn+lstJreQ0mu67qn3TFKu5AxZoTdDj6BSDqqHEos5KirS4pz3zt3r5IbC3mv8vDm9+o6O5M2f7R6RRNfD9IPJANmO0k2Ajf529I0bGgAGgIpIXb8OaI6G+L48dR john@Johns-MacBook-Pro.local"
    "user-script"          = "/usr/sbin/mdata-get root_authorized_keys > ~root/.ssh/authorized_keys"
  }

  image_uuid          = data.smartos_image.illumos.id
  maintain_resolvers  = true
  max_physical_memory = 512
  nics {
    nic_tag   = "external"
    ips       = ["10.0.222.222/16"]
    gateways  = ["10.0.0.1"]
    interface = "net4"
  }
  quota = 25

  resolvers = ["1.1.1.1", "1.0.0.1"]

  filesystems {
    source = "/zdata/videos"
    target = "/data/videos"
    type   = "lofs"
  }
  filesystems {
    source = "/zdata/music"
    target = "/data/music"
    type   = "lofs"
  }

  provisioner "remote-exec" {
    inline = [
      "pkgin -y update",
      "pkgin -y in htop",
    ]
    connection {
      type     = "ssh"
      host     = self.primary_ip
      user     = "root"
      private_key = file("~/.ssh/id_rsa")
    }
  }
}

resource "smartos_machine" "linux" {
  alias          = "provider-test-linux"
  brand          = "lx"
  kernel_version = "3.16.0"
  cpu_cap        = 100

  customer_metadata = {
    # Note: this is my public SSH key...use your own.  :-)
    "root_authorized_keys" = "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDYIHv3DoXnAMn+dggUup1a+jjSqpZiIU5ThgljXHG9KM+iy1W3zo9qshUE7vBj/l7l5aHzRKyXsmWb6EdmtlVBnYl7SH5IMGaEFlB6n7T+yoMRl7VczZZxvP+VSAac2HeLPvdrCDeJCckfkHeTg9E3rt2PcAz0REKDCm34lpsedgM4QrVh8D54NgqLCdpT+QpidEBwE1T5wGMId4OwBB+r1VJZyn+lstJreQ0mu67qn3TFKu5AxZoTdDj6BSDqqHEos5KirS4pz3zt3r5IbC3mv8vDm9+o6O5M2f7R6RRNfD9IPJANmO0k2Ajf529I0bGgAGgIpIXb8OaI6G+L48dR john@Johns-MacBook-Pro.local"
    "user-script"          = "/usr/sbin/mdata-get root_authorized_keys > ~root/.ssh/authorized_keys"
  }

  image_uuid          = data.smartos_image.linux.id
  maintain_resolvers  = true
  max_physical_memory = 512
  nics {
    nic_tag   = "external"
    ips       = ["10.0.222.223/16"]
    gateways  = ["10.0.0.1"]
    interface = "net5"
  }
  quota = 25

  resolvers = ["1.1.1.1", "1.0.0.1"]

  filesystems {
    source = "/zdata/videos"
    target = "/data/videos"
    type   = "lofs"
  }
  filesystems {
    source = "/zdata/music"
    target = "/data/music"
    type   = "lofs"
  }

  provisioner "remote-exec" {
    inline = [
      "apt-get update",
      "apt-get -y install htop",
    ]
  }
}

resource "smartos_machine" "linux-kvm" {
  alias          = "provider-test-linux-kvm"
  brand          = "kvm"
  kernel_version = "3.16.0"
  vcpus          = 2

  customer_metadata = {
    # Note: this is my public SSH key...use your own.  :-)
    "root_authorized_keys" = "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDYIHv3DoXnAMn+dggUup1a+jjSqpZiIU5ThgljXHG9KM+iy1W3zo9qshUE7vBj/l7l5aHzRKyXsmWb6EdmtlVBnYl7SH5IMGaEFlB6n7T+yoMRl7VczZZxvP+VSAac2HeLPvdrCDeJCckfkHeTg9E3rt2PcAz0REKDCm34lpsedgM4QrVh8D54NgqLCdpT+QpidEBwE1T5wGMId4OwBB+r1VJZyn+lstJreQ0mu67qn3TFKu5AxZoTdDj6BSDqqHEos5KirS4pz3zt3r5IbC3mv8vDm9+o6O5M2f7R6RRNfD9IPJANmO0k2Ajf529I0bGgAGgIpIXb8OaI6G+L48dR john@Johns-MacBook-Pro.local"
  }

  maintain_resolvers = true
  ram                = 512
  nics {
    nic_tag   = "external"
    ips       = ["10.0.222.224/16"]
    gateways  = ["10.0.0.1"]
    interface = "net0"
    model     = "virtio"
  }
  quota = 25

  resolvers = ["1.1.1.1", "1.0.0.1"]

  disks {
    boot        = true
    image_uuid  = data.smartos_image.linux_kvm.id
    compression = "lz4"
    model       = "virtio"
  }

  provisioner "remote-exec" {
    inline = [
      "apt-get update",
      "apt-get -y install htop",
    ]
    connection {
      type     = "ssh"
      host     = self.primary_ip
      user     = "root"
      private_key = file("~/.ssh/id_rsa")
    }
  }
}
```
