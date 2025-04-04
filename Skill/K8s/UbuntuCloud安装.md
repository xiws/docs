# 安装Ubuntu cloud

1. 环境 Ubuntu
2. 工具KVM/QEMU

安装步骤
1. 编排init信息
2. 创建isoxinx
3. 创建虚拟机

## 准备

cloud_init.yaml

```yml
#cloud-config
hostname: xiwpc
manage_etc_hosts: true
timezone: Asia/Shanghai  # 设置时区

users:
  - name: xiw
    sudo: ALL=(ALL) NOPASSWD:ALL
    groups: users, admin
    home: /home/xiw
    shell: /bin/bash
    lock_passwd: false
    passwd: "$6$XeDSXqVqEzu0i3JT$oELvDbN13W421wrY7hKOMIuK26mo.xmGAOGwfguLZGrmx7vblHnlV/l3UundguDzdvwmVk8qPEOPK5Ogxu48Q/" # root 密码生成规则：openssl passwd -6
    ssh_authorized_keys:
      - ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQCusOwS7KZPu0eTZC8X/n7CS7AfnQjMiK6UZtWHfKomuxTK+trp/KzHTAzEghvfFMMMKi+zlG1K5tgolErYtZL+k2+BuyseANBlubCL/WkZ2zBRSYAkmtnqu5ALn6Nt9BbER51Py6AzrzPbPd4xSA4EOA4R51zDKHe3/dnmEy6tqxVCxZKEJ6T4CKfJQQteXaQ/S7WTYgOY6AoC4Rwq2lwAf6ZDMmV9+optR9V96JoGSjyz+fKadfZfeSERONRjmi+QdzvmzVCDUQ6zmylYAcJBSaL4hrw/CJTvT1J6OEKEXGoFCAhehMS6k1huVeNfWaGGsE3RdOOG0LM2RH2Wrula1fnEztLdLGYoI0EHXQUQ5Km31mrcf7JBiU3czy+6xGAHIBAqKo9s3B9Qxv914aeVa9TIEevbnFXwXT9owpMVblePk07HPS3od1TYZA2Ct9S8wkGvRI3Fq6oNJkyVvbbz6ICCejo438kG8n0xCPQyxbPPk5vvWeDCKOTYQI3u5JQWsAid/H1YY1RjR2QbpVgQgNiU9m+RW7xMArZiBPq3r9uFpczmWiVIaxXaOb1/XHZr/Guzul2bRuOuiqgAhJNrbjk4lXguYfowCD+RRZZUz2TJie5g/vZni97NmY5F9CSLIvW32Sa/fztN1zDr4nqsSCz8+invRztkhm9WzP5cJQ== zhongxiwang@live.com
      - ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQDdYsg/rgXg/7tqmwlS6tqVzH4VBh3hnqDHpospVO9S78bGNoXwJv9Io1Xmtl0IsVZrrCUdrk0EcTMDkeFW7ebjVc9syd9pYpm7ZXKgD9IDJ4T5h/OwUCQY4oavhhX24QXKCH4BgiGTMvLdnPPLJ1tJ53rpZZnvRKMjs3+mkBa74Uvrg38S4xOCeEUw0TrOE5GbjSyrNuTfyxfWXW85hpteNGBvLFRdlJyG+CiTc6JugZ8AYk4DI1FGgED2S+HIdmuB1Xcv81qgAYj/MYgp4mpE1Mv5+gDkuJVWSHwK926Au0Te/OD5xJy2qUa5w99XCPR6x65CLV+byfyf6EQ26NjjushvUg5WL8FaQn2w/K+fK/QxZWfG9DNjUmxiouTAmKR46Abp5SylNUXBwQ+9OCoMlKLqbhokaxyaD/HE2O2G+8fo1MJpcEOsK3muaH702CLrovyp32B/V/2ptpNtg6MaRzd8CrKwXGZDQiJoQie6aDHYkg7xbhLdzXIJnMfn1WYE7rN3CSQ9k6i/BqAGHBMKyDfWyqrCTe7w0SxA8LBiTwh9ZQIHnyV70b4rG+VTZDhBRZvn5RsJG5gaq+1RWE7+bOLlbploCXGy9L3NUAI0ZiZ5TgWlxrZRbb0EraP/qjC5MDteO4QzXtoHxOr9XMGD70ZISF8rs/YJCd79FIFDnQ== zhongxiwnag@foxmail.com

ssh_pwauth: true
chpasswd:
  list: |
    ubuntu:ubuntu  # 默认用户名 ubuntu，密码 ubuntu  只能修改已经存在的账户密码
  expire: false # false表示登录后不需要改密码

runcmd:
  - echo "Cloud-init is running!" > /root/cloud-init.log

```

network-config

```yaml

version: 2
ethernets:
  ens3:
    dhcp4: no
    addresses: [192.168.1.100/24]
    gateway4: 192.168.1.1
    nameservers:
      addresses:
        - 8.8.8.8
        - 8.8.4.4
```


## 打包

```shell
mkdir cloud-init
cp user-data cloud-init/
cp meta-data cloud-init/
cp network-config cloud-init/  # 这里手动加入 network-config

genisoimage -output cloud-init.iso -volid cidata -joliet -rock cloud-init/

# cloud-localds cloud-init.iso user-data meta-data network-config 
# 自带的不支持网络

```

## 安装命令

```shell
virt-install --name=nodeOnde \
--vcpus=4 --memory=4048 \
--disk path=/var/lib/libvirt/images/ubuntu-22.10-server-cloudimg-amd64.img,size=10 \
--disk path=/var/lib/libvirt/images/cloud-init.iso,device=cdrom \
--os-variant=ubuntu22.04 \
--import --network bridge=virbr0,model=virtio \
--graphics none --console pty,target_type=serial

```

## 登录虚拟机

```sshell
virsh console ubuntu-cloud

```

## 删除虚拟机

```shell
sudo virsh list --all  # 先查看虚拟机名称
sudo virsh shutdown <虚拟机名称>

sudo virsh destroy ubuntu-vm
sudo virsh undefine --remove-all-storage ubuntu-vm

```