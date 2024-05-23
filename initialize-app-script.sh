#!/bin/bash

#Docker installation
sudo yum install docker -y
sudo service docker start
sudo chkconfig docker on # make docker  autostart


# docker-compose (latest version)
sudo curl -L https://github.com/docker/compose/releases/latest/download/docker-compose-$(uname -s)-$(uname -m) -o /usr/local/bin/docker-compose
sudo chmod +x /usr/local/bin/docker-compose # Fix permissions after download

# Git installation
sudo yum install -y git

# Clone repository
git clone https://github.com/Aracelimartinez/email-platform-challenge.git
cd /email-platform-challenge

# Set environment variables
export ZINC_FIRST_ADMIN_USER="${ZINC_FIRST_ADMIN_USER}"
export ZINC_FIRST_ADMIN_PASSWORD="${ZINC_FIRST_ADMIN_PASSWORD}"
export API_PORT="${API_PORT}"
export ZINCSEARCH_USERNAME="${ZINCSEARCH_USERNAME}"
export ZINCSEARCH_PASSWORD="${ZINCSEARCH_PASSWORD}"
export ZINCSEARCH_HOST="${ZINCSEARCH_HOST}"

# Run the containers
docker-compose -f docker-compose.production.yml up -d
