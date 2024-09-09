variable "vm_name" {
  description = "The name of VM"
  type        = string
}

variable "vm_count" {
  description = "The count of VM"
  type        = number
  default     = 1
}

variable "vm_cpu" {
  description = "The CPU of VM"
  type        = number
  default     = 1
}

variable "vm_mem" {
  description = "The mem of VM"
  type        = string
  default     = "512 mib"
}

variable "vm_image" {
  description = "The image of VM"
  type        = string
}

variable "vm_network_type" {
  description = "The network type of VM"
  type        = string
  default     = "bridged"
}

variable "vm_network_host_interface" {
  description = "The network host interface of VM"
  type        = string
  default     = "enp8s0"
}
