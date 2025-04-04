# 定义变量
ISO_NAME = cloud-init.iso
IMG_NAME = ubuntu-22.10-server-cloudimg-amd64.img
IMG_PATH = /var/lib/libvirt/images
VM_NAME = node1

.PHONY: installKvm
installKvm:
	sudo apt update && sudo apt install -y qemu-kvm libvirt-daemon-system libvirt-clients bridge-utils cloud-utils genisoimage

.PHONY: all iso move install clean

# 生成 cloud-init ISO
iso:
	rm -rf $(ISO_NAME)
	genisoimage -output $(ISO_NAME) -volid cidata -joliet -rock cloud-init/

# 移动文件到 /var/lib/libvirt/images
move: iso
	cp $(IMG_NAME) $(IMG_PATH)/
	cp $(ISO_NAME) $(IMG_PATH)/

# 安装虚拟机
install: move
	virt-install --name=$(VM_NAME) \
		--vcpus=4 --memory=4048 \
		--disk path=$(IMG_PATH)/$(IMG_NAME),size=10 \
		--disk path=$(IMG_PATH)/$(ISO_NAME),device=cdrom \
		--os-variant=ubuntu22.04 \
		--import --network bridge=virbr0,model=virtio \
		--graphics none --console pty,target_type=serial

# 清理生成的 ISO 文件
clean:
	rm -f $(ISO_NAME)

.PHONY: shutdown
shutdown: 
	sudo virsh shutdown $(VM_NAME)

.PHONY: del
del: 
	sudo virsh undefine --remove-all-storage $(VM_NAME)

