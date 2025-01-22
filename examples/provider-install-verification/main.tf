# Copyright (c) Digitec Galaxus AG
# SPDX-License-Identifier: MIT



terraform {
  required_providers {
    hetznerrobot = {
      source = "DigitecGalaxus/hetznerrobot"
    }
  }
}

/*
To test this with working call create a secrets.tfvars file and add it to your terraform plan command
Filecontents of secrets.tfvars
hetzner_username = ""
hetzner_password = ""

terraform plan -var-file secrets.tfvars
*/

provider "hetznerrobot" {
  username = var.hetzner_username
  password = var.hetzner_password
}

data "hetznerrobot_servers" "servers" {}
