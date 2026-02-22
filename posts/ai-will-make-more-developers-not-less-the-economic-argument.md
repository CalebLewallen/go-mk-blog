---
Page Name: AI Will Make More Developers, Not Less: The Economic Argument
Published: true
Published Date: 2023-05-17
Author: Caleb Lewallen
RobotsAllowed: true
---

![Feature Image](/static/images/photo-1522252234503-e356532cafd5)

I think the growing consensus among the tech world right now is that AI is going to eventually take over most white collar jobs. I've written before about how [AI is going to take over (half) your job](__GHOST_URL__/ai-will-take-your-job-and-its-totally-fine/), but I'm starting to think I was wrong, at least when it comes to software developers. To be clear, I don't think I'm wrong about it taking over half your job, but I think I was wrong about the impact it's going to have on developers.

Since ChatGPT came out, and we've gotten a taste of what a well-trained, commercial LLM product can do, I've kind of viewed this as a net-neutral impact on the developer community. After all, I haven't heard from a single company yet that they're short of things to go on their roadmap – they're just going to accelerate their plans.

Sure, we've seen [massive tech layoffs](https://techcrunch.com/2023/05/12/tech-industry-layoffs/), but those were all economic layoffs from Big Tech, not productivity problems. Those were problems of **focus** at the top, not a problem with output from the bottom. The layoffs from Big Tech have created [massive hiring opportunities for startups](https://www.forbes.com/sites/forbesbusinesscouncil/2023/03/03/big-tech-layoffs-a-hiring-opportunity/?sh=5e13aefa78c9), and led to a [surge in new startups](https://www.wired.com/story/tech-layoffs-are-feeding-a-new-startup-surge/). This is the impact of economic forces, not AI, but it's indicative of what's to come.

# The AI Revolution

Ok, punch line first. The AI revolution is going to **increase** the number of developers, not decrease them.

When we evaluate this hypothesis, I want to be clear that this is in reference to the total number of people employed as developers, or have a hand in development work, NOT that individual companies will hire more developers. In fact, I would argue that the number of developers per company will stay the same or drop, but due to economic reasons, not technical ones.

This is going to come about because of several reasons:

1. Decreased Barriers to Entry
2. The Rise of Low-Code and No-Code
3. Product Niches

# Decreased Barriers to Entry

The prevailing argument towards AI decreasing the number of developer roles is that AI will prove itself an example of **Disruptive Innovation**, sometimes referred to as destructive innovation. This scary naming comes from the fact that certain jobs are destroyed whenever a new type of technology is developed.

For example, the tractor destroyed farming as a way of life for millions. It transformed our way of life positively, moving us from an agrarian society into an industrial one, with better incomes and a better quality of life. However, it temporarily displaced a lot of workers who suddenly found themselves not needed to tend to crops.

There's countless other examples, but I want to stick to the tractor for this evaluation. The tractor has made farming so efficient that you can't make a living farming without one. However, the tractor presents a massive barrier to entry. Planting most crops involves the use of a class of equipment called row tractors. These are the machines that can plant, well, rows of different crops.

![](/static/images/photo-1614977645540-7abd88ba8e56)

Photo by [Julia Koblitz](https://unsplash.com/fr/@jkoblitz?utm_source=ghost&utm_medium=referral&utm_campaign=api-credit) / [Unsplash](https://unsplash.com/?utm_source=ghost&utm_medium=referral&utm_campaign=api-credit)

The upfront cost of one of these machines is in excess of $100K, even for used machines. When financed, the cost of the machine has to be amortized over 5 years. This represents a huge barrier to entry for anyone who wants to try and farm at scale.

## AI Lowers That Barrier

These GPT models are creating (at least initially) a much lower barrier to entry to create new software products. PHP/MySQL apps are going to always make money, and AI is going to make it easier than ever for non-developers to make these products!

Here's a simple example. I'm not developer by trade. I've worked a lot with data over my career, so I'm familiar with SQL, and can sometimes scrape enough Python together to do a quick task, but I'm not someone who could make any form of functional software. At least until this afternoon...

![Webpage of ChatGPT, a prototype AI chatbot, is seen on the website of OpenAI, on a smartphone. Examples, capabilities, and limitations are shown.](/static/images/photo-1681460590033-67b0d1413550)

Photo by [Sanket Mishra](https://unsplash.com/@sanketgraphy?utm_source=ghost&utm_medium=referral&utm_campaign=api-credit) / [Unsplash](https://unsplash.com/?utm_source=ghost&utm_medium=referral&utm_campaign=api-credit)

Today I decided to create a simple HTML app, ChartGPT. It's a pretty simple app – it takes input from a user to create an org chart. Here's the thing, I didn't write a single line of code.

I asked ChatGPT (GPT4) the following question:

```
Create for me a javascript script that will create an org chart. The user should be able to enter in the name, title, and person they report to into a form to add a node. Each node should have a connector that connects the nodes together. Generate the HTML, CSS, and Javascript code in separate scripts.
```

It then created for me some code that I was able to test using CodePen. It took some tweaking (again, not from me – ChatGPT made all the code changes when I requested them). Eventually, it created a workable demo that I was able to use. This isn't perfect, but it's getting better. Tools like AutoGPT make it possible to connect different models together and follow multi-step process independently. Work is being done on these code generation models to allow models to recursively test and update their own code so that the desired outcome is produced.

Here's a demo of the code ChatGPT produced to give you an idea of what's possible, just a few months into the AI revolution. I didn't write a single character of this code (not that I could if I wanted to), it's just a copy/paste of the code the GPT4 model produced. I had to ask it to fix a couple of things, and if I spent a little more time on it I'm sure I could fix a few other issues with it. But it's still pretty impressive that in 5 min, someone like me could take a generic LLM and produce functioning code.

Given a few months (or weeks, given the pace of AI innovation these days), more specialized tools are going to come out that put the power of *good* code generation into the hands of idiots like me who can just tell it step by step what I want and what tweaks to make and create a functioning app.

> This doesn't just lower the barrier to entry, it almost eliminates any technical barriers to entry.

There's countless examples of this sort of practice occurring. One recent example that comes to mind is the easy access to financial markets, made possible by free trading from user-friendly apps like Robinhood and Webull. Only a few years prior, if you wanted to get involved in financial markets you would need to use a cumbersome brokerage platform that would take a long time to learn, and THEN pay $8 for the privilege of entering or exiting a position. Lockdown and stimulus in 2020 definitely contributed to the mayhem, but the easy access to markets was the catalyst that made it all possible.

The AI revolution will be no different. It's all building up to become the perfect storm. Big tech laying off thousands of skilled developers, who are either being scooped up by startups, or making companies of their own. Rising interest rates seem to be pushing us towards a recession, which will put others in the non-tech world out of work. Do you think unemployment and free access to ChatGPT or Bard is going to create more people trying to create software products, or less?

# The Rise of Low-Code and No-Code

![Papercraft toy in front of PHP echo statement](/static/images/photo-1535551951406-a19828b0a76b)

Photo by [KOBU Agency](https://unsplash.com/@kobuagency?utm_source=ghost&utm_medium=referral&utm_campaign=api-credit) / [Unsplash](https://unsplash.com/?utm_source=ghost&utm_medium=referral&utm_campaign=api-credit)

I know, it takes more than just a chat bot pushing out code snippets to build a product. Have no fear, there's plenty of low code and no code products out there to bridge the gap between where you are and where you need to be.

It's no longer necessary to have any ability to code to create useful SaaS apps (although some conceptual knowledge is really useful). These platforms, like Bubble, WeWeb, WebFlow, Xano, and Backendless remove almost all of the technical barriers, like user authentication, security, and infrastructure maintenance from the hands of the app creators and allows them to focus on making cool products.

## Economies of Scale

Once you have the code written, most SaaS apps are expensive to run. All of the EC2 compute time, database and S3 storage, serverless infrastructure, etc. cost a lot of money to run. If you've ever built a SaaS app you also know that most of the time those resources sit around underutilized. You're still paying for that dedicated EC2 instance, even when it's not running. This makes an expensive barrier to entry even if you already know how to code.

Most no-code platforms are built on shared resources, meaning you're sharing a pool of compute and storage resources with a bunch of other applications. This is completely and totally fine for 99% of applications. From a cost efficiency perspective, this is even preferred.

This makes the cost of building these tools on low or no code platforms dirt cheap. All of these services do offer dedicated resource packages, but you only need to migrate to these more expensive tiers once you have the user load to demand it (and the revenue to pay for it). You can get an app up and running for FREE in most cases, only paying when you think you have something ready to be published and to show to the world.

We've very quickly knocked down both the technical barriers and economic barriers to creating new products. I can build a fairly sophisticated application in my pajamas, orders of magnitude faster than would have been possible just a couple of years ago. That's *insane*.

# Product Niches

This is one of my absolute favorite charts ever.

![](https://170.187.153.50/content/images/size/w1000/2023/05/image-8.png)

It's called the product adoption curve. It's been around forever and is used to describe the types of customers you're going to encounter as you reach market saturation.

Moving a product up the curve is pretty formulaic in strategy, albeit not in execution. The strategic approach to unseating an incumbent is to find a *customer niche* that your competitor can't serve very well (the customer niche is too small to matter for your competitor, their product is too generic to pivot towards that customer, etc). This niche is your foothold to storm the proverbial beach and take market share.

The AI revolution and the rise of no/low code tools is going to create waves of niche products. Wish your trading app was focused on only trading tech stocks? You can make a niche product for that. Wish your fitness app was focused on making you a better powerlifter? Strongman? Bodybuilder? Long distance runner? Sprinter? You can make a niche product for that.

The economic forces and incentives towards getting lots of niche products out the door that serve smaller groups of users are here. I can't help but wonder if the loud pundits crying out for regulating AI to prevent "the loss of jobs" are doing so out of compassion or self-interest?

# More Devs, Not Less

All signs seem to point towards more people being employed to create software products, not less. I don't mean that there will be more people learning how to use an IDE to code the old fashioned way – those folks will still be needed, but gap between what you have to code the old fashioned way and what you can do with a new generation of tools is closing fast, being chipped away almost daily.

Who knows, maybe one day your full time job, or side hustle, might just be building and maintaining a specialized piece of software.

* * *

If you liked this article, consider following me and [subscribing to email updates](#/portal/signup) whenever I post an article. You can also follow me on [Twitter](https://twitter.com/LewallenCaleb?ref=caleblewallen.com) or connect with me on [LinkedIn](https://www.linkedin.com/in/caleb-lewallen-b8699365/?ref=caleblewallen.com).
