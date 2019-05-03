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
    "alias" = "provider-test-illumos"
    "brand" = "joyent"
    "cpu_cap" = 100

    "customer_metadata" = {
        # Note: this is my public SSH key...use your own.  :-)
        "root_authorized_keys" = "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDYIHv3DoXnAMn+dggUup1a+jjSqpZiIU5ThgljXHG9KM+iy1W3zo9qshUE7vBj/l7l5aHzRKyXsmWb6EdmtlVBnYl7SH5IMGaEFlB6n7T+yoMRl7VczZZxvP+VSAac2HeLPvdrCDeJCckfkHeTg9E3rt2PcAz0REKDCm34lpsedgM4QrVh8D54NgqLCdpT+QpidEBwE1T5wGMId4OwBB+r1VJZyn+lstJreQ0mu67qn3TFKu5AxZoTdDj6BSDqqHEos5KirS4pz3zt3r5IbC3mv8vDm9+o6O5M2f7R6RRNfD9IPJANmO0k2Ajf529I0bGgAGgIpIXb8OaI6G+L48dR john@Johns-MacBook-Pro.local"
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
        "root_authorized_keys" = "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDYIHv3DoXnAMn+dggUup1a+jjSqpZiIU5ThgljXHG9KM+iy1W3zo9qshUE7vBj/l7l5aHzRKyXsmWb6EdmtlVBnYl7SH5IMGaEFlB6n7T+yoMRl7VczZZxvP+VSAac2HeLPvdrCDeJCckfkHeTg9E3rt2PcAz0REKDCm34lpsedgM4QrVh8D54NgqLCdpT+QpidEBwE1T5wGMId4OwBB+r1VJZyn+lstJreQ0mu67qn3TFKu5AxZoTdDj6BSDqqHEos5KirS4pz3zt3r5IbC3mv8vDm9+o6O5M2f7R6RRNfD9IPJANmO0k2Ajf529I0bGgAGgIpIXb8OaI6G+L48dR john@Johns-MacBook-Pro.local"
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
    "brand" = "lx"
    "kernel_version" = "3.16.0"
    "vcpus" = 2

    "customer_metadata" = {
        # Note: this is my public SSH key...use your own.  :-)
        "root_authorized_keys" = "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDYIHv3DoXnAMn+dggUup1a+jjSqpZiIU5ThgljXHG9KM+iy1W3zo9qshUE7vBj/l7l5aHzRKyXsmWb6EdmtlVBnYl7SH5IMGaEFlB6n7T+yoMRl7VczZZxvP+VSAac2HeLPvdrCDeJCckfkHeTg9E3rt2PcAz0REKDCm34lpsedgM4QrVh8D54NgqLCdpT+QpidEBwE1T5wGMId4OwBB+r1VJZyn+lstJreQ0mu67qn3TFKu5AxZoTdDj6BSDqqHEos5KirS4pz3zt3r5IbC3mv8vDm9+o6O5M2f7R6RRNfD9IPJANmO0k2Ajf529I0bGgAGgIpIXb8OaI6G+L48dR john@Johns-MacBook-Pro.local"
        "user-script" = "/usr/sbin/mdata-get root_authorized_keys > ~root/.ssh/authorized_keys"
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
