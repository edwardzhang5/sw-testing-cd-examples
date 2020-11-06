#!/bin/bash
# This is the entry point for configuring the system.
#####################################################

#install basic tools
sudo pacman -Syu --noconfirm
sudo pacman -S git vim neovim devtools base-devel --needed --noconfirm

# install google cloud sdk
# git clone https://aur.archlinux.org/google-cloud-sdk.git

#get golang 1.14.10
curl -O https://storage.googleapis.com/golang/go1.14.10.linux-amd64.tar.gz

#unzip the archive
tar -xvf go1.14.10.linux-amd64.tar.gz

#move the go lib to local folder
mv go /usr/local

#delete the source file
rm go1.14.10.linux-amd64.tar.gz

#only full path will work
touch /home/vagrant/.bash_profile

echo "export PATH=$PATH:/usr/local/go/bin" >> /home/vagrant/.bash_profile

echo "export GOPATH=/home/vagrant/go:$PATH" >> /home/vagrant/.bash_profile

export GOPATH=/home/vagrant/go

mkdir -p "$GOPATH/bin"

