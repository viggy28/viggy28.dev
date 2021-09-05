---
date: 2021-09-05T21:58:00UTC
description: "What is DHCP and DHCPCD"
featured_image: ""
tags: ["system"]
title: "DHCP"
---

I have seen this term here and there. I remember fiddling with this in [pain of setting up port forwarding in xfinity/](https://viggy28.dev/article/pain-of-setting-up-port-forwarding-in-xfinity/). 

DHCP - Dynamic Host Configuration Protocol (DHCP) *is a network management protocol used on Internet Protocol (IP) networks for automatically assigning IP addresses and other communication parameters to devices connected to the network using a client–server architecture* (from [wiki](https://en.wikipedia.org/wiki/Dynamic_Host_Configuration_Protocol) )

client-server architecture?
In a home network, router is the server and client is my PCs, laptops, Raspberry Pis etc.

From wiki, "Many routers and residential gateways have DHCP server capability."

Continuing "Most residential network routers receive a unique IP address within the ISP network. Within a local network, a DHCP server assigns a local IP address to each device."

So the server aspect of DHCP is taken care by the router provided by ISP.

In the client side, for example in Raspberry Pi I see [dhcpcd](https://roy.marples.name/projects/dhcpcd/). It stands for DHCP Client Daemon.

> "It’s also an IPv4LL (aka ZeroConf) client"

So i need to know what is IPv4LL first. It stands for IPv4 [Link Local](https://en.wikipedia.org/wiki/Link-local_address) address.

> "In IPv4, link-local addresses are normally only used when no external, stateful mechanism of address configuration exists, such as the Dynamic Host Configuration Protocol (DHCP), or when another primary configuration method has failed"

From What I understand, link-local addresses are used by DHCP is not available. However dhcpcd is also a link local client !!

```
Sep 05 14:29:58 raspberrypi dhcpcd[21460]: vetha80e16d: using IPv4LL address 169.254.218.50
Sep 05 14:29:58 raspberrypi dhcpcd[21460]: vetha80e16d: adding route to 169.254.0.0/16
```