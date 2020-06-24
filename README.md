### Introduction/Outline

Probe is a lightweight TCP/IP suite build with Go which tries to adhere two basic principles. Simplicity
and speed.

Usual protocol suite design tends to segregate each protocol differently and are handled
as different processes and communication happens through these processes.
While, Probe segregates each  'handling of packet' to run as a separate process for better
speed, as it reduces context switching.

Basically this is just a code design thinking which tries to follow an
object orientation into the stack structure as well.

The idea is to bypass the kernel and use network hardware directly from userspace process.

##### Short note on TUN TAP

TUN/TAP provides packet reception and transmission for user space programs.
It can be seen as a simple Point-to-Point or Ethernet device, which,
instead of receiving packets from physical media, receives them from
user space program and instead of sending packets via physical media
writes them to the user space program.

- How to test the interface once probe initilizes it?

Just run the `run_interface.sh` script once you execute the probe binary.
