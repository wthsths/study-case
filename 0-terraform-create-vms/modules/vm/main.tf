terraform {
  required_providers {
    virtualbox = {
      source  = "terra-farm/virtualbox"
      version = "0.2.2-alpha.1"
    }
  }
}

resource "virtualbox_vm" "vm" {
  count  = var.vm_count
  name   = format("${var.vm_name}-%02d", count.index + 1)
  image  = var.vm_image
  cpus   = var.vm_cpu
  memory = var.vm_mem

  network_adapter {
    type           = var.vm_network_type
    host_interface = var.vm_network_host_interface
  }
}
