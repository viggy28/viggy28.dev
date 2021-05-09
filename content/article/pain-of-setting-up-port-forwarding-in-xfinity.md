---
date: 2021-05-07T23:00:08-04:00
description: "Challenges and concepts learnt during port forwarding set up"
featured_image: "/images/rpi.jpeg"
tags: ["networking"]
title: "port forwarding in xfinity"
---

This article is part of my Today I Learnt. I have been trying to share whatever I learnt newly on the day. Some of them are basic, but still I want to write it down for two reasons. 
1. To help me understand the concepts better 
2. Potential help other people who running in to same issue.

Alright, It's all started when I bought a new domain (synergy.net) for my Dad. Who doesn't like yet another domain for their future project? He wanted a blog, so I thought it would be nice to host a [ghost](https://ghost.org/) [1] blog on my Raspberry Pi which has been sitting idle on my desk. As like every good netizen, I did a Google search and found this medium [article](https://medium.com/swlh/install-ghost-on-your-raspberry-pi-b7cdc8e7e37f) [2] which exactly solves what I want. Software installation went fine and I reached the step of forwarding traffic to my Raspberry Pi.

```
There are two steps to that:
Adding a DNS record that points to your public IP
Adding port forwarding from your home router to your Raspberry Pi.
```

This is when things got interesting. Xfinity is my Internet Service Provider and I use their device for using Internet. I didn't have a clear idea what that device is. Looking further, I learnt it is a xFi Gateway device which combines both modem + router. That was interesting. Also, at http://10.0.0.1/ I can reach the gateway. First thing on the Gateway site I noticed is **Bridge Mode** which is disabled in my case. I never heard of that term before. If I want to run modem and router separately then I need to enable it [3].

A primer on what's the main difference between modem and router [4]:

**modem** 

1. brings internet to your home
2. has public IP address. For eg. when you visit website like [whatismyipaddress.com](whatismyipaddress.com) whatever IP that shows is provided by your modem

**router** 

1. brings internet to your devices
2. assigns local IP address. For eg. when you type `hostname -I` on your raspberry Pi.

**?** One thing which I don't understand is why my IPv6 is different between what I see in [whatismyipaddress.com](whatismyipaddress.com) and what I see in gateway interface. IPv4 address is same.

Adding a DNS record was straightforward. In my Domain Registrar I added an A record which points to my public ipv4. 

Next I needed to set up portforwarding in my router. Here I learnt there are ways one can do port forwarding. Xfinity xFi users, can manage it here https://internet.xfinity.com/network/advanced-settings/portforwarding. [5]
For non xFi users they can set up in the admin interface which is 10.0.0.1

On the xFi gateway the options seem limited. No support for HTTP 

![port-forwarding-xfinity-site.png](/images/port-forwarding-xfinity-site.png)

From [5],

```
Advanced Users: Setting Up Port Forwarding With the Admin Tool
We recommend xFi for most subscribers, but more advanced users can alternately use the Gateway’s Admin Tool.
```

**?** I am not sure how to disable xFi and start using Xfinity gateway.  

**DHCP**

Moving on, I was wondering how the IP is assigned to my raspberry pi. Looks like Dynamic Host Configuration Protocal (**DHCP)** is responsible for assigning IPs to the devices connected to the router. On raspberry pi, I see a service called **dhcpcd** which has leased IP address (i guess from the router).

![port-forwarding](/images/port-forwarding.jpeg)

The critical part of port forwarding I couldn't get it done easliy. So, I decided to use an alternative method which I will write in an another post.

References:

1. https://ghost.org/
2. https://medium.com/swlh/install-ghost-on-your-raspberry-pi-b7cdc8e7e37f

3. https://www.xfinity.com/support/articles/wireless-gateway-enable-disable-bridge-mode
4. https://www.xfinity.com/hub/internet/modem-vs-router
5. https://www.xfinity.com/support/articles/port-forwarding-xfinity-wireless-gateway