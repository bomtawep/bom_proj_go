#!/bin/bash

# Update the package repository
sudo yum update -y

# Install Docker using Amazon Linux Extras
sudo amazon-linux-extras install docker -y

# Start the Docker service
sudo service docker start

# Add the user to the docker group
sudo usermod -a -G docker ec2-user

# Enable Docker to start on boot
sudo chkconfig docker on

# Install Docker Compose
sudo curl -L https://github.com/docker/compose/releases/latest/download/docker-compose-$(uname -s)-$(uname -m) -o /usr/local/bin/docker-compose
sudo chmod +x /usr/local/bin/docker-compose

# Display Docker and Docker Compose versions for verification
docker --version
docker-compose --version