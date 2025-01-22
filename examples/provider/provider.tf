# Copyright (c) Digitec Galaxus AG
# SPDX-License-Identifier: MIT

terraform {
  required_providers {
    hetznerrobot = {
      source = "DigitecGalaxus/hetznerrobot"
    }
  }
}

provider "hetznerrobot" {
  username = ""
  password = ""
}
