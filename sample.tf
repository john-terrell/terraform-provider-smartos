provider "smartos" {
    "host" = "10.99.50.60:22"
    "user" = "root"
}

data "smartos_image" "test" {
    "name" = "base-64-lts"
    "version"  = "18.4.0"
}

resource "smartos_machine" "test" {
    "alias" = "provider-test-512"
    "brand" = "joyent"
    "cpu_cap" = 100

    "customer_metadata" = {
        "root_authorized_keys" = "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDYIHv3DoXnAMn+dggUup1a+jjSqpZiIU5ThgljXHG9KM+iy1W3zo9qshUE7vBj/l7l5aHzRKyXsmWb6EdmtlVBnYl7SH5IMGaEFlB6n7T+yoMRl7VczZZxvP+VSAac2HeLPvdrCDeJCckfkHeTg9E3rt2PcAz0REKDCm34lpsedgM4QrVh8D54NgqLCdpT+QpidEBwE1T5wGMId4OwBB+r1VJZyn+lstJreQ0mu67qn3TFKu5AxZoTdDj6BSDqqHEos5KirS4pz3zt3r5IbC3mv8vDm9+o6O5M2f7R6RRNfD9IPJANmO0k2Ajf529I0bGgAGgIpIXb8OaI6G+L48dR john@Johns-MacBook-Pro.local"
        "user-script" = "/usr/sbin/mdata-get root_authorized_keys > ~root/.ssh/authorized_keys"
    }

    "image_uuid" = "${data.smartos_image.test.id}"
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
