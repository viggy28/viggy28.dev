---
date: 2019-03-25T18:00:08-04:00
description: "How to secure your remote applications"
featured_image: "/images/vpn1.png"
tags: ["firebase", "hugo", "hosting"]
title: "How to secure your remote applications"
---

I am writing this around the time where we are at the beginning of a
[**global
pandemic**](https://www.forbes.com/sites/brucelee/2020/03/11/the-covid-19-coronavirus-is-now-a-pandemic-what-does-that-mean/#d9bcd0277699).
Everyone is trying to figure out how to keep their business up and
running during this challenging period. Most of the employers have
mandated work from remote policy. One of the main challenges is how to
connect to your applications remotely. Traditionally the answer is
[VPN](https://en.wikipedia.org/wiki/Virtual_private_network). There are
so many challenges wrt scalability, availability, and performance of
these traditional VPN software. Don’t worry, there are alternatives.

Here I am going to explain the setup that I have used to secure my
[application](https://viggy28.dev/book) using [Cloudflare
Access](https://teams.cloudflare.com/access/index.html), which provides
Simple, secure access for internal apps. Check out the link which
provides a great explanation on how Access works. Also, another good
news is that this product is completely **free** (at least for the next
6 months due to the outbreak of COVID-19). Yes, you read it right.
Honestly, this is a massive news. Here is an excerpt from the
[announcement](https://blog.cloudflare.com/cloudflare-during-the-coronavirus-emergency/).

***If you are not yet using Cloudflare for Teams, and if you or your
employer are struggling with limits on the capacity of your existing VPN
or Firewall, we stand ready to help and have removed the limits on the
free trials of our Access and Gateway products for at least the next six
months.***

I always wanted to try how Access works and there can’t be a better time
than now to try. Here is the simple 4 steps which I did to get up and
running:

1.  Select an Identify provider: Identity provider is nothing but an
    authentication method that your end user is going to be
    authenticated. It can be your existing identity provider such as
    Gmail, Facebook, Github. Of course, you can also use Okta, AzureAD
    but thats available for enterprise plan. I have chosen One-Time pin,
    and Github.

2.  Get a Login Page Domain: Here is what I created for this demo.
    [covid19-how-to-secure-remote-application.cloudflareaccess.com](https://covid19-how-to-secure-remote-application.cloudflareaccess.com/).
    Feel free to click, I assure there is no catch. ![set up with
    identity provider](/images/identity-provider.png)

3.  Set up an Access Control Policy: [Access-control
    List](https://en.wikipedia.org/wiki/Access-control_list) defines who
    can do what. Similar to authorization in the databases.

    -   Provide a suitable Application Name.
    -   Specify the domain where your application is hosted on. For
        example, mine runs on
        [viggy28.dev/book](https://viggy28.dev/book)
    -   I have added Github organization, emails from a specific domain
        and individual emails as a different authentication method to
        reach my application. There are other ways like IP ranges. This
        shows how granular you can get into when it comes to policies.
        ![Access Policies](/images/access-policies.png)

4.  Access App Launch: We are almost there. I have set it up as
    **Everyone** to access the App. Since I have secured it using ACL in
    the above step I am not concerned about anyone trying to reach my
    application. If you want you can still use the same authentication
    methods used as your Identity Provider. ![App
    Policies](/images/app-policies.png)

That’s it !!! Check out your app access URL. Mine is
[https://covid19-how-to-secure-remote-application.cloudflareaccess.com/](https://covid19-how-to-secure-remote-application.cloudflareaccess.com/)

I am using the One-Time pin Identity provider method to authenticate.
![authenticate](/images/authentication.png)

That email address is already added to our access policy. So I did get a
6 digit pin within a minute to that email!
[email](/images/one-time-pin.png).

Note: Not only HTTP based application, you can also secure SSH, RDP
based applications. Access with [Gateway]
([https://teams.cloudflare.com/gateway/](https://teams.cloudflare.com/gateway/))
one can build a fully fledged VPN.

-   [networking](/tags/networking)
-   [Cloudflare](/tags/cloudflare)
