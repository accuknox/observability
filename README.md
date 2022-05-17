# Observability


## Step to Follow


###  Install Cilium 
    curl -L --remote-name-all https://github.com/cilium/cilium-cli/releases/latest/download/cilium-linux-amd64.tar.gz{,.sha256sum}
    
    sha256sum --check cilium-linux-amd64.tar.gz.sha256sum
    
    sudo tar xzvfC cilium-linux-amd64.tar.gz /usr/local/bin
    
    rm cilium-linux-amd64.tar.gz{,.sha256sum}

    cilium install 

    cilium hubble enable

### Install Kubearmor

    curl -sfL https://raw.githubusercontent.com/kubearmor/kubearmor-client/main/install.sh | sudo sh -s -- -b /usr/local/bin

    karmor install --image kubearmor/kubearmor:latest