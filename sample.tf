provider "smartos" {
    "host" = "10.99.50.60:22"
    "user" = "root"
}
resource "smartos_machine" "test" {
    "count" = 1
    "alias" = "provider-test"
    "brand" = "joyent"
    "image_uuid" = "c193a558-1d63-11e9-97cf-97bb3ee5c14f"
}
