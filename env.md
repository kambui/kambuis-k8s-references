

# Using MobaXTerm as base 
I use MobaXTerm on windows as a base to interact with cloud and k8s clusters. The local terminal of mobaXterm is very powerful and gives you nix feeling on windows. I had customized moba so that these tools are avialable without too much set up time. 

You will need to use a persistant home directory for moba. Then setup the tools using bash_profile

Here is the snippet of .bash_profile
```
if [ ! -f /bin/kubectl ]; then
   ln -s /drives/c/ProgramData/chocolatey/bin/kubectl.exe /bin/kubectl
fi

if [ ! -f /bin/oc ]; then
   ln -s /drives/c/ProgramData/chocolatey/bin/oc.exe /bin/oc
fi

if [ ! -f /bin/minikube ]; then
   ln -s /drives/c/ProgramData/chocolatey/bin/minikube.exe /bin/minikube
fi

if [ ! -f /bin/vboxmanage ]; then
   ln -s /drives/d/Program\ Files/Oracle/VirtualBox/VBoxManage.exe /bin/vboxmanage
fi

if [ ! -f /bin/terraform ]; then
   ln -s /drives/c/ProgramData/chocolatey/bin/terraform.exe /bin/terraform
fi

```

# Minikube Environment
C:\>minikube start

# List virtual machines
C:\>vboxmanage list vms
"minikube" {fdf452f5-e866-4361-a0d8-eb8ccdc03e24}
"ubuntu" {aad283a9-edc0-4bcf-8501-1e9f09db5a3c}
"ubuntu-core" {d89ad4a6-8b79-489a-ac01-c525de258ed8}

-- I use Ubuntu and ubuntu-core for my docker builds. To know more about ubuntu core : https://ubuntu.com/core

# Start VM in headless mode
VBoxManage startvm ubuntu-core --type headless

# Show all running VMS
vboxmanage list runningvms

# Screenshot of the headless vm console output (This is useful to get the ipaddress due to DHCP)
vboxmanage controlvm ubuntu-core screenshotpng screen.png

# Power off VM
VBoxManage controlvm ubuntu-core poweroff