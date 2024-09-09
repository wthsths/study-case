## Terraform Vms
VMleri oluşturmak için kullandığımız terraform projesidir.

### Gereklilikler
- Virtualbox
- terraform

### Adımlar
- ```make install``` komutu ile Vagrant box'ı indirin
- ```terraform plan```
- ```terraform apply```

### Sorunlar
Eğer destroy edildikten sonra tekrar create işlemi sırasında hata alıyorsanız. Lütfen ~/.config/Virtualbox/Virtualbox.xml dosyası içerisindeki VM disklerin destroy edildiğinden emin olunuz. Eğer diskler silinmemiş ise aşağıdaki komutla diskleri manuel olarak silebilirsiniz.

```
vboxmanage closemedium disk UUID --delete
```
