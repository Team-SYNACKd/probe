### Introduction/Outline

Probe is a lightweight TCP/IP suite build with Go which tries to adhere two basic principles. Simplicity
and speed.

Usual protocol suite design tends to segregate each protocol differently and are handled
as different processes and communication happens through these processes.
While, Probe segregates each  'handling of packet' to run as a separate process for better
speed, as it reduces context switching.

Basically this is just a code design thinking which tries to follow an
object orientation into the stack structure as well.
