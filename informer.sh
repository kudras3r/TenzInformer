#!/bin/bash

RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[0;33m'
BLUE='\033[0;36m'
NC='\033[0m' 

APP_NAME="tenzinformer"
SERVICE_FILE="/etc/systemd/system/${APP_NAME}.service"
LOG_DIR="/var/log/tenzir-example"
CONFIG_DIR="/etc/tenzir-example"
TMP_DIR="/var/tmp/tenzir-example"
BINARY_NAME="informer"

USER=$(whoami)
GROUP=$(id -gn)


spin() {
  local pid=$1
  local delay=0.1
  local spinchars="/-\\|"

  while ps -p $pid > /dev/null; do
      for i in $(seq 0 3); do
          echo -ne "\r${spinchars:i:1} "
          sleep $delay
      done
  done

  echo -ne "\r"  
}

# check root
if [ "$EUID" -ne 0 ]; then
  echo -e "${RED}Please, restart installation with sudo${NC}"
  exit 1
fi

create_config() { 
  cat <<EOF > $CONFIG_DIR/config.yml
os:
  family: "Linux"
  name: "Lubuntu"
  kernel: "5.4.0-42-generic"
  codename: "Focal Fossa"
  type: "Linux"
  platform: "x86_64"
  version: "20.04"

name: "examplePC"
mac: "00:1A:2B:3C:4D:5E"
EOF
}

create_service() {
  cat <<EOL > $SERVICE_FILE

[Unit]

Description=${APP_NAME}
After=network.target tenzir-node.service

[Service]

User =${USER}
Group=${GROUP}
ExecStart=/usr/bin/informer
Restart=on-failure

[Install]

WantedBy=multi-user.target

EOL
}

install() {
  # pkg update / upgrade
  echo -e "${BLUE}Updating packages...${NC}"
  (apt update && apt upgrade -y) &
  spin $!
  echo -e "\n"

  # install snap
  echo -e "${BLUE}Installing snap...${NC}"
  if ! command -v snap &> /dev/null; then
    (apt install snapd -y) &
    spin $!
  fi

  # install go
  echo -e "${BLUE}Installing Go with snap...${NC}"
  if ! command -v go &> /dev/null; then
    (snap install go --classic) &
    spin $!
  fi
  echo -e "\n"

  # mk dirs / log file / cfg file
  echo -e "${GREEN}Creating logs directory: ${LOG_DIR}${NC}"
  mkdir -p $LOG_DIR
  
  echo -e "${GREEN}Creating config directory: ${CONFIG_DIR}${NC}"
  mkdir -p $CONFIG_DIR
  
  echo -e "${GREEN}Creating log file...${NC}"
  touch $LOG_DIR/tenzinformer.log

  echo -e "${GREEN}Creating config file...${NC}"
  create_config
  
  echo -e "${GREEN}Creating tmp directory: ${TMP_DIR}${NC}"
  mkdir -p $TMP_DIR
  echo -e "\n"

  # compile source
  echo -e "${BLUE}Compiling...${NC}"
  (go build -o /usr/bin/$BINARY_NAME $(pwd)/cmd/informer/main.go) &
  spin $!

  # mk .service file 
  create_service  

  chmod 644 $SERVICE_FILE

  echo -e "${BLUE}Reloading systemd...${NC}"
  systemctl daemon-reload

  echo -e "\n"

  echo -e "${GREEN}Starting the service...${NC}"
  systemctl start $APP_NAME

  echo -e "${GREEN}Done!${NC}"
  echo -e "\n"
}

# Функция деинсталляции
uninstall() {
  echo -e "${YELLOW}Stopping service...${NC}"
  systemctl stop $APP_NAME

  echo -e "${YELLOW}Removing service...${NC}"
  systemctl disable $APP_NAME
  rm -f $SERVICE_FILE

  echo -e "${YELLOW}Removing binary file...${NC}"
  rm -f /usr/local/bin/$BINARY_NAME

  echo -e "${YELLOW}Removing directories...\n${NC}"
  rm -rf $LOG_DIR
  rm -rf $CONFIG_DIR
  rm -rf $TMP_DIR

  echo -e "${GREEN}Uninstallation completed!\n${NC}"
}


if [ "$1" == "install" ]; then
  install
elif [ "$1" == "uninstall" ]; then
  uninstall
else
  echo -e "${RED}Usage: sudo $0 {install|uninstall}${NC}"
  exit 1
fi