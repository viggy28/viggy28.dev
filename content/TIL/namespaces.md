---
date: 2021-09-15T16:20:00UTC
description: "What is namespaces in Linux"
featured_image: ""
tags: ["system"]
title: "namespaces"
---

namespaces is a Linux Kernel feature. There are different types of namespaces
- Partitions Kernel resources such that one set of Processes see a different resources while others different
- Apartment complex analogy - 7 different namespaces
- PID namespace - Watching TV show 
- Net namespace - each namespace gets an unique IP address s list of port. - apartment address analogy 
- Uts namespace - host names for Ip address - Telling taxi driver apartment name instead of address
- User namespace - Files are associated with VID Euser Identification
- Mailbox associated to apartment unit # not with the name of the person - Mit namespace o to isolate mount points such that processes in different namespaces cannot view each others files.
- Similar to chroot

**Commands:**

* lsns -> To list namespaces
* pgrep --ns $pid -a  -> To list all the processes running on the namespace
* unshare -> To create user namespace 

references:
* https://www.redhat.com/sysadmin/7-linux-namespaces
* https://www.redhat.com/sysadmin/building-container-namespaces

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