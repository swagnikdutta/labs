doing this in linux was easier, straight forward. 

in linux
- you can leave the fields empty - total len, id, checksum, src address
- refer man page rules, checksum gets filled in always
- 

in mac 
- certain fields are in host byte order like, total len and offset (although I didn't require to deal with offsets)
- about checksum, there's no need to calculate the checksum, it gets calculated by the NIC (on systems that support 
  checksum offloading), irrespective of whether checksum offloading is set to true or false.
- mac os zeros your computed checksum, in tcpdump output and wireshark output. This happens if checksum offloading is set to true.
- What happens with checksum offloading = 0, is, your calculated checksum is not zeroed out. It is respected. Wireshark would not show that message "due to checksum offloading"
  - however, your checksums won't match, because while you have sent out checksum using host byte order, in tcpdump, it appreas reversed, hence checksum does not match.
- but you do get the echo reply back.
  - the rules of raw sockets in mac os apply - total length does not include ip header length, and it is in host byte order.


But you get the response back.

What a journey it has been!

# disable/enable checksum offloading

Since checksum bits were being zeroed, I am trying to find out why that's happening.
Till now, checksum offload on the loopback interface lo0 was set (1)

I am unsetting it.
```code
➜ ~ ifconfig lo0
lo0: flags=8049<UP,LOOPBACK,RUNNING,MULTICAST> mtu 16384
	options=1203<RXCSUM,TXCSUM,TXSTATUS,SW_TIMESTAMP>
	inet 127.0.0.1 netmask 0xff000000
	inet6 ::1 prefixlen 128
	inet6 fe80::1%lo0 prefixlen 64 scopeid 0x1
	nd6 options=201<PERFORMNUD,DAD>
➜ ~
➜ ~ sudo sysctl net.link.generic.system.hwcksum_tx
net.link.generic.system.hwcksum_tx: 1
➜ ~ sudo sysctl net.link.generic.system.hwcksum_rx
net.link.generic.system.hwcksum_rx: 1

➜ ~ sudo sysctl net.link.generic.system.hwcksum_rx=0
net.link.generic.system.hwcksum_rx: 1 -> 0
➜ ~ sudo sysctl net.link.generic.system.hwcksum_tx=0
net.link.generic.system.hwcksum_tx: 1 -> 0

➜ ~ sudo sysctl net.link.generic.system.hwcksum_rx net.link.generic.system.hwcksum_tx
net.link.generic.system.hwcksum_rx: 0
net.link.generic.system.hwcksum_tx: 0
➜ ~
```

Let's try again and see if the checksum bits are still being zeroed.

Now checksum is not being zeroed at all. Yay!
While sending, 
my calculated checksum was `60 f0`, but it was `7c d4` in wireshark and tcpdump.
The mismatch is probably(until I verify) due to total length field being represented in little-endian,
i.e 1c 00 creates checksum 60 f0
and 00 1c creates checksum 7c d4

# Update: Yes, above is true verified using chatgpt

While receiving, 
tcpdump/wireshark first receives the packet, there checksum is `c9 dd`, 
and it's the same in the application program.

# Conclusion come observation
So what we take away from here is that, the "zeroing" of checksum has indeed stopped
after checksum offloading was disabled or set to 0.


