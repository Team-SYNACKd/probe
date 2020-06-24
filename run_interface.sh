sudo ip addr add 10.1.0.10/24 dev probe

sudo ip link set dev probe up

ping -c1 -b 10.1.0.255
