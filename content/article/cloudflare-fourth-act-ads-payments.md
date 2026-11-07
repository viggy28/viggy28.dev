---
date: 2026-11-06T00:00:00-04:00
description: "Cloudflare's fourth act is the economic layer for the agentic Internet: payments first, and potentially advertising next."
featured_image: "/images/cloudflare-four-acts.svg"
tags: ["cloudflare", "ai", "agents", "payments", "advertising"]
title: "Cloudflare Is Going to Be an Ads and Payments Company"
toc: true
draft: false
---

*I previously worked at Cloudflare. This is my personal analysis based entirely on public information. It does not represent Cloudflare, and it does not use confidential information.*

Cloudflare likes to describe its strategy in **acts**.

Its first act put Cloudflare in front of websites. Its second act extended the network to people and private networks. Its third act let developers run applications on that network.

For a long time, the fourth act was not obvious. Would Cloudflare become an ISP? Would it sell routers for offices? Would it build something else on top of its global network?

I think we now have a clear answer.

**Cloudflare's fourth act is the economic layer for the agentic Internet. Payments are already visible. Advertising is the logical next step.**

![Cloudflare's four acts](/images/cloudflare-four-acts.svg)

## Act 1: control traffic going to websites

Cloudflare launched publicly in 2010 as a service that made websites faster and safer. Its early advantage was a large [Anycast](https://www.cloudflare.com/learning/cdn/glossary/anycast-network/) network: traffic could enter Cloudflare close to the user and then be filtered, cached, or routed before reaching the website's origin server.

Calling this only a CDN undersells what Cloudflare built. The important product was the **reverse proxy position** between a website and the rest of the Internet.

Once Cloudflare occupied that position, it could add services such as:

- authoritative DNS
- caching and content delivery
- DDoS protection
- a web application firewall
- bot management
- smart routing through Argo
- free TLS through [Universal SSL](https://blog.cloudflare.com/introducing-universal-ssl/)

Cloudflare's network was the foundation. Each new service made that network more valuable, while each new customer gave Cloudflare more traffic and better threat intelligence.

## Act 2: control enterprise and network traffic

Cloudflare's official shorthand for Act 2 is the **forward proxy**: products that protect users as they connect to the Internet and to corporate applications.

My broader mental model is that Cloudflare moved one level down and one level inward. It was no longer protecting only public websites. It was helping companies control employees, offices, private applications, and entire IP networks.

Products in this era include:

- [Cloudflare Access](https://blog.cloudflare.com/introducing-cloudflare-access/) for identity-aware access to internal applications
- Gateway and WARP for filtering and securing user traffic
- Cloudflare Tunnel for connecting origins without exposing inbound ports
- Spectrum for TCP and UDP applications beyond HTTP
- [Magic Transit](https://blog.cloudflare.com/magic-transit-network-functions/) for protecting and routing IP networks
- Magic WAN and Magic Firewall for connecting and securing enterprise networks

The acts are not perfectly separated product folders. Magic Transit, for example, is not literally a forward proxy. The useful abstraction is the progression: **Cloudflare went from controlling requests to websites to applying identity and policy across networks.**

## Act 3: run applications on the network

Once Cloudflare had a global network carrying and filtering traffic, the next question was obvious: why stop at moving requests? Why not run the application there too?

That became the developer platform:

- [Workers](https://blog.cloudflare.com/introducing-cloudflare-workers/) for compute
- Workers KV for globally distributed key-value storage
- Durable Objects for state and coordination
- R2 for object storage
- D1 for relational data
- Queues, Workflows, and other application primitives
- Workers AI and Vectorize for AI inference and retrieval

Act 3 made Cloudflare programmable. The network was no longer only a pipe or a security boundary. Developers could build complete applications on it.

Each act expanded what Cloudflare's network could mediate: first requests to websites, then access across company networks, and finally the applications themselves. As some of those applications become agents that fetch information, call tools, and make purchases, the next missing layer is economic: how those machines pay and get paid.

That is the bridge from Act 3 to Act 4. Cloudflare does not automatically *own* either side of an agent transaction. But a publisher may already use Cloudflare to protect its website, while the agent may run on Workers or use Cloudflare's agent tooling. Cloudflare can provide infrastructure to both and often sits in the middle when they communicate.

## Agents break the web's economic loop

The human web has a simple economic loop.

A publisher pays to create and operate a website. People visit it. The publisher recovers its cost through subscriptions, commerce, or advertising. Even when a visitor does not buy anything, an ad impression can have value.

Agents break that loop.

Imagine asking an agent to find a restaurant in Manhattan. A person might open three restaurant websites. An agent can inspect a thousand menus, compare prices and reviews, and return one answer.

All thousand restaurants paid some cost to serve those requests. But 999 of them received no visit, no ad impression, and possibly no attribution. The agent extracted the useful information while removing the economic mechanism that funded it.

At human scale, that mismatch was tolerable. At machine scale, it can overwhelm the economics that paid for the content.

## Payments come first

In 2024, Cloudflare introduced [dedicated tools](https://blog.cloudflare.com/cloudflare-ai-audit-control-ai-content-crawlers/) for website owners to identify, audit, and block AI crawlers. In July 2025 it introduced [Pay Per Crawl](https://blog.cloudflare.com/introducing-pay-per-crawl/), allowing participating publishers to choose among three responses to an AI crawler:

1. allow it,
2. charge it, or
3. block it.

Pay Per Crawl uses HTTP `402 Payment Required`, but it does not require an on-chain payment for every page. A crawler receives a price and retries with payment intent. Cloudflare records the authenticated requests, aggregates the charges, bills the crawler, and distributes the publisher's earnings.

Cloudflare later partnered with Coinbase around [x402](https://blog.cloudflare.com/x402/), an open protocol for machine-to-machine payments over HTTP.

In September 2026, Cloudflare moved beyond crawl access with two beta products. [Monetization Gateway](https://blog.cloudflare.com/monetization-gateway-beta/) lets sellers charge per request for websites, APIs, datasets, or MCP tools. [Pay Per Use](https://blog.cloudflare.com/pay-per-use/) lets an AI company define a paid use—such as a citation or excerpt—and publishers decide whether to accept the offer. The buyer reports each use; Cloudflare handles the records, billing, and payment.

Together, these products turn the identity and policy checks Cloudflare already performs into metering, billing, and settlement.

## Pay Per Crawl is not advertising

This distinction is important.

In a normal ad market, an **advertiser** pays to reach a potential customer. The **publisher** supplies the place where the ad appears. Demand-side platforms, supply-side platforms, and exchanges match the two sides and run an auction.

Pay Per Crawl moves money in the other direction. The agent pays the publisher for access. Economically, it looks more like a toll, API charge, or content license than an advertisement.

So why do I think advertising is next?

Because agents do not only bring traffic. They can bring **structured intent**.

A browser request says, roughly, “someone opened this page.” An agent may be able to say, with the user's permission, “I am looking for a five-night hotel stay in Hawaii for a family of four, within this budget.”

That is extremely valuable commercial intent. A hotel might happily waive the access fee to receive that request. It may even pay to submit a relevant offer.

The advertisement is no longer a banner competing for a person's attention. It is a machine-readable offer competing for a place in the agent's candidate set.

**Old advertising sells human attention. Agentic advertising may sell machine consideration.**

![Payments and ads inside an agent request](/images/cloudflare-agentic-market.svg)

## Why Cloudflare could become an exchange

Google became powerful because it indexed the web and owned the search box. Meta owns attention and identity inside its applications. Amazon owns shopping intent close to the purchase.

Cloudflare owns none of those things directly. Its advantage is different: **it operates the gate.**

For a meaningful portion of the web, Cloudflare can already:

- recognize the incoming bot or agent,
- apply the publisher's allow, block, or charge policy,
- observe which resource is being requested,
- protect the transaction from abuse,
- serve code close to both parties, and
- through its beta products, handle billing and settlement.

## Advertising shows the scale, not the market size

Advertising is a useful comparison because it shows how valuable aggregated intent can become. In 2025, Alphabet, Meta, and Amazon generated roughly **$560 billion in advertising revenue** combined: $294.7 billion, $196.2 billion, and $68.6 billion, respectively, according to their annual reports.

That figure shows the scale of existing advertising businesses; it does not predict Cloudflare's opportunity. No one yet knows how often agents will accept paid offers, which interactions will become billable, or what role Cloudflare will play. The narrower point is that Act 4 gives Cloudflare a new way to monetize its network position—as an economic intermediary rather than only an infrastructure provider.

## The fourth act

Cloudflare's first act made the web faster and safer. Its second act applied the network to companies and their users. Its third act made the network programmable.

The fourth act makes the network **economic**.

Payments are the clearest part: agents pay for access to content, data, APIs, and services. Advertising is the larger prediction: merchants pay for transparent, machine-readable consideration when an agent expresses commercial intent.

In its Q3 2026 earnings call, Cloudflare said it had recorded its first Act 4 revenue. During the Q&A, Matthew Prince said he expected “some version of micropayments,” described advertising opportunities around an agent's declared intent, and said that, looking ten years out, he imagined this becoming “the largest part of Cloudflare's business.”

That is a remarkable destination for a company that began by caching websites and blocking attacks.

Cloudflare started by moving packets. Its fourth act may be about moving money.

## References

- [Cloudflare Q2 2025 earnings call: the four acts](https://cloudflare.net/files/doc_financials/2025/q2/CORRECTED-TRANSCRIPT_-Cloudflare-Inc-NET-US-Q2-2025-Earnin.pdf)
- [Cloudflare Q3 2026 earnings-call transcript](https://stockanalysis.com/stocks/net/transcripts/726816-q3-2026/)
- [Introducing Pay Per Crawl](https://blog.cloudflare.com/introducing-pay-per-crawl/)
- [Cloudflare and Coinbase launch the x402 Foundation](https://blog.cloudflare.com/x402/)
- [Monetization Gateway beta](https://blog.cloudflare.com/monetization-gateway-beta/)
- [Pay Per Use beta](https://blog.cloudflare.com/pay-per-use/)
- [Introducing Universal SSL](https://blog.cloudflare.com/introducing-universal-ssl/)
- [Introducing Cloudflare Access](https://blog.cloudflare.com/introducing-cloudflare-access/)
- [Introducing Magic Transit](https://blog.cloudflare.com/magic-transit-network-functions/)
- [Introducing Cloudflare Workers](https://blog.cloudflare.com/introducing-cloudflare-workers/)
- [IAB advertising ecosystem overview](https://data.iab.com/ecosystem.html)
- [Alphabet 2025 Form 10-K](https://www.sec.gov/Archives/edgar/data/1652044/000165204426000018/goog-20251231.htm)
- [Meta 2025 Form 10-K](https://www.sec.gov/Archives/edgar/data/1326801/000162828026003942/meta-20251231.htm)
- [Amazon 2025 Form 10-K](https://www.sec.gov/Archives/edgar/data/1018724/000101872426000004/amzn-20251231.htm)
