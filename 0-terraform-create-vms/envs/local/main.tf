terraform {
  required_version = "1.9.5"
}

locals {
  vargant_image = "../../boxes/rocky8.box"
}

module "master" {
  source = "../../modules/vm"

  vm_name  = "master"
  vm_image = local.vargant_image
  vm_cpu   = 2
  vm_mem   = "4.0 gib"
}

module "workers" {
  source = "../../modules/vm"

  vm_name  = "worker"
  vm_count = 2
  vm_image = local.vargant_image
  vm_cpu   = 4
  vm_mem   = "4.0 gib"
}
