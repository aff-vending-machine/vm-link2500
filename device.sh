#!/bin/bash

# ------------------
# lsusb
## Bus 001 Device 004: ID 2341:0043 .......
#
## vendor ID = 2341
## product ID = 0043
#
# -----------------

# Set the vendor and product IDs for your USB device
VENDOR_ID="0000"
PRODUCT_ID="0000"
# Set the symlink name for your device
SYMLINK_NAME="link2500"

# Determine the device path for the USB device
DEVICE_PATH="$(udevadm info -q path -n /dev/ttyACM0)"

# Create a udev rule for the device
# RULE="SUBSYSTEM==\"tty\", ATTRS{idVendor}==\"${VENDOR_ID}\", ATTRS{idProduct}==\"${PRODUCT_ID}\", SYMLINK+=\"${SYMLINK_NAME}\""
RULE="KERNEL==\"ttyACM[0-9]*\", ATTRS{idVendor}==\"${VENDOR_ID}\", ATTRS{idProduct}==\"${PRODUCT_ID}\", MODE=\"0666\", SYMLINK+=\"${SYMLINK_NAME}\""

# Write the rule to a new file in /etc/udev/rules.d/
echo "$RULE" | sudo tee /etc/udev/rules.d/99-usb-serial.rules > /dev/null

# Reload the udev rules
sudo udevadm control --reload-rules

# Print the device path and symlink name
echo "Device path: $DEVICE_PATH"
echo "Symlink name: /dev/${SYMLINK_NAME}"

# ------------------
# docker-compose.yaml
#   ...
#     devices:
#       - /dev/link2500:/dev/ttyACM0
#     user: root:dialout
#   ...
# -----------------