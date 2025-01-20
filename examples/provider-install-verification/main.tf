terraform {
  required_providers {
    hetznerrobot = {
      source = "DigitecGalaxus/hetznerrobot"
    }
  }
}

provider "hetznerrobot" {}

data "hetznerrobot_server" "example" {}
