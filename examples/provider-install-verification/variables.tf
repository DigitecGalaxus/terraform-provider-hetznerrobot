# Copyright (c) Digitec Galaxus AG
# SPDX-License-Identifier: MIT



variable "hetzner_username" {
  description = "Username for the hetzner robot webservice"
  type        = string
  sensitive   = true
}

variable "hetzner_password" {
  description = "Password for the hetzner robot webservice"
  type        = string
  sensitive   = true
}
