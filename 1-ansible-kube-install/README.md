## ansible-kube-spray
kubespray ile kubernetes kurulumu projesidir.

## Kurulum
- ```make install``` downloads and prepares kubespray
- ```cd kubespray && ansible-playbook -i inventory/mycluster/inventory.ini -u vagrant --ask-pass --become --ask-become-pass cluster.yml``` run playbook
