provider "smartos" {
    "host" = "10.99.50.60"
}
resource "smartos_machine" "test" {
    "alias" = "provider-test"
    "brand" = "joyent"
    "image_uuid" = "c193a558-1d63-11e9-97cf-97bb3ee5c14f"
}
