VAGRANTFILE_API_VERSION = "2"

Vagrant.configure(VAGRANTFILE_API_VERSION) do |config|
  config.vm.box = "ubuntu/trusty64"

  config.vm.network :private_network, ip: "192.168.19.20"
  config.vm.network "forwarded_port", guest: 80, host: 8090

  config.ssh.forward_agent = true

  config.vm.synced_folder "~/Workspace/go-project", "/home/vagrant/go/project"

  config.vm.provision "ansible" do |ansible|
    ansible.playbook = "provisioning/playbook.yml"
    ansible.verbose = "vvv"
  end
end
