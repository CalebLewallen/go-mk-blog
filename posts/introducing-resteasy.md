---
Page Name: A New Chapter
Published: true
Published Date: 2023-11-20
Author: Caleb Lewallen
RobotsAllowed: true
---

![Feature Image](/static/images/photo-1698434156086-918aa526b531)

I know it's been a while since I've actually published anything. There's been a lot going on in my personal life, which meant that consistent writing has just fallen by the wayside a bit.

Over the summer of 2023, I decided to leave my role at LoanBoss, a company I helped to build since it's initial ideation in 2017. It was bittersweet for me – LoanBoss had started its life as a spreadsheet on my laptop, and it was hard to walk away from. I worked with a lot of amazing people at LoanBoss, and it will always be a special place to me. For anyone reading this that's currently a user of LoanBoss, the platform is in great hands and they have some amazing plans in store for the future.

With my summer off, I got to spend a lot of time with my wife and kids, and spent my time tinkering and building something new. I wanted to build an API service to deliver hard-to-find holiday information to subscribers. However, I pretty quickly discovered that the API service itself was by far the easiest part of building and launching this product.

# Building Something New

I love building things. It was my favorite part of working at LoanBoss. I loved the challenge of finding a problem and then having to work my way towards creating a solution for it. This new API service I was trying to build in my time off, was really a continuation of solving a particularly difficult problem I had face sourcing holiday information for financial markets. I had a great time figuring out how to source this data, calculate the holiday information, and serve it up in such a way that it was easily digestible by an end user. Then the real problems began.

First, I want to *sell* access to this API. It costs money to run a service like this, and I certainly didn't want to foot the bill on my own. So I needed a way to get customers onboarded and paying for this service. No problem, I thought, I'll just make a user portal.

Turns out it's easier said than done. First, I need a way for users to create an account. They also need a way to recover their password if it's forgotten, so that means integrating an email service. Once that's done and they're logged in, I need a way to collect payment information so they can subscribe to my service. No one wants to punch in their credit card information just anywhere, so I need a reputable processor, like Stripe, in order to create some legitimacy (side note: I don't know if you've ever integrated with a payments vendor, but it ain't easy).

After that, I need a way to secure my API services so I know only subscribed user are accessing them. That means issuing, encrypting, and securely storing API keys. Then I need to provide documentation so users can figure out how to get the data they want. What about when they want to cancel their subscription and delete their account? What if they lose their API key and need to generate a new one?

AGH!

[via GIPHY](https://giphy.com/gifs/reaction-8c9NInMpXixMjR6lTH)

# A Solution

After asking around online and with some developer friends, it seems like I'm not the only one facing this problem. I had lunch with an old friend of mine, who happens to be a software engineer by trade. I pitched him the idea of creating a service that handled all the boring, awful parts of offering an API-as-a-Service – builders should focus on building, not reinventing the wheel every time they make a new API service. He loved the idea too, and we have decided to make the idea a reality.

## Introducing: RestEasy

![](/static/images/Frame-33.png)

## [RestEasy](https://www.resteasyapi.com) is the front-end and middleware for you API service. The idea is simple:

- Deploy your API Service to your favorite cloud provider and set up your marketing site.
- Use RestEasy to set up your **custom, white-labelled** user portal and admin portal on your own **custom URL** (or ours, if you prefer).
- We will handle all of your user onboarding & management, payments & billing, API key management, logging, and auto-generating documentation.
- You will need to define your route -- simply add your endpoint and query parameters, along with any basic validation you want us to check for you. The user will make their API request through your RestEasy instance, and we will handle the auth, basic query validation, and logging.

When I was building my API product, I found that building the API itself was only 20% of the actual work that needed to be done. The other 80% is everything I've described above. In fact, I still haven't finished it (the user-servicing piece, the API itself has been done for a while), and my new plan is to run this product on RestEasy when we launch.

# Beta Program

If you'd like to join our beta program, we'd love to have you! We haven't opened up the platform yet, but you can join our waitlist here: [https://www.resteasyapi.com#waitlist/](https://www.resteasyapi.com/waitlist/).

We'll be selecting a few waitlisted members to join our beta platform, and in exchange for helping us get our product ready for prime-time, you will receive lifetime discounted access for the products you launch during the beta period.

I'm excited for this next chapter, thank you for coming along!
