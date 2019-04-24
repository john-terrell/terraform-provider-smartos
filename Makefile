all:
	go build -o terraform-provider-smartos
	terraform init
