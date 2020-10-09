#!/bin/bash
# This is the entry point for configuring the system.
#####################################################

#install basic tools

sudo apt-get install -y mysql-server-5.7
sudo apt install -y mysql-client-core-5.7

sudo apt-mark hold mysql-server-5.7
