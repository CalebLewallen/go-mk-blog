---
Page Name: You Will Always Have More Problems Than You Can Solve
Published: true
Published Date: 2023-02-20
Author: Caleb Lewallen
RobotsAllowed: true
---

![Feature Image](/static/images/photo-1505664194779-8beaceb93744)

I’m going to get philosophical.

We humans always long for something we can’t have. We’re constantly trying to tinker with and experiment with the way we operate and do things to find this thing that alludes us. We’ve experimented with different political and economics systems in the search of utopia. We constantly innovate trying to build a better mousetrap. New influencers pop up every day to convince us that their dietary model is the answer to better health.

But the thing that alludes them all — **perfection**.

It’s such a human thing, to constantly try to achieve the impossibility of perfection. Yet, humanity has constantly fought internally over making changes to try and achieve this.

Software is no different. Countless methods and systems, check and balances, different coding languages and documentation styles, automated and manual testing processes — all created with the same goal in mind — to eliminate bugs.

### You Can’t Have Perfection

There’s a fundamental law of the universe that will always prevent us from reaching perfection. It’s simple really: **it’s far easier to break things that to build them**. Anyone who’s tried to assemble anything from IKEA knows this is true. I think building software makes this more blatantly obvious than in other things.

Software platforms are best described as “[Complex Systems](https://en.wikipedia.org/wiki/Complex_system)”. Simply put, a complex system in one that cannot be understood in it’s whole (within the limitations of the human brain), but by studying the behaviors of its parts.

There are several characteristics that are symptomatic of complex systems. The first, and most obvious from it’s name, is **Complexity**. That is, a system’s behaviors cannot be easily inferred from its properties. The can exhibit **Non-Linearity**, which is a system responding in different ways to the same inputs. Parents of small children will know this effect by another name: *Chaos*. Lastly, and probably most importantly, is the concept of **Emergence**. This are traits of a system that are not apparent from its components in isolation, but result from the relationships and interactions between parts of a system. The most frustrating bugs in software come from this emergence trait.

### Bugs Live In Code

My CTO is fond of a saying that has entered a regular part of my vocabulary. It’s a simple phrase that holds a lot of lessons to unpack:

> *Where do bugs live? Bugs live in code.*

This is a hard truth to learn, especially for young guys like me that think we can solve anything if we’re given enough time. But there’s wisdom to this way of thinking.

To quote Mark Twain¹, “I didn’t have time to write a short letter, so I wrote a long one instead.” In writing, it’s often said that it’s much harder to write shorter than it is to write longer (clearly evident by the length of this post — *does this guy know when to stop typing???*). The same is true with app development.

The best way to write good code is to write less code. Fewer lines are easier to read, there’s fewer dependencies, and just fewer places to make mistakes. Fewer requirements means the engineers can write less.

This is also taken into consideration when we’re faced with a *build vs buy* decision. When you BUILD, you have to write a LOT of code. That’s opportunity for bugs, and only builds onto your complex system. When you BUY, the code you write is de minimis, just enough to integrate with the other service. There’s fewer opportunities for bugs, and therefore usually easier to maintain.

### Problems Increase Exponentially

It’s way easier to write code the first time through than it is to go back through and debug it. Both from a time and number of lines of code written, it’s the debugging process that kills forward momentum.

In practice, what this means is that the number of bugs you add to your software increases exponentially by the number of features you add to it.

Take adding a simple form for example:

![](https://170.187.153.50/content/images/max/800/0-_2td8krifmr_688j.png)

*From Github:* [AdamBrodzinski/reason-simple-form-example](https://github.com/AdamBrodzinski/reason-simple-form-example?ref=caleblewallen.com)

The function of what a form does is to take an input and record them to a database. It’s pretty simple, it only adds a single new piece of functionality. However, by creating this form, I create an exponentially larger number of *risks*. Thankfully, computer science has progressed to a point where these risks are minimal, but there’s a number of things that could go wrong that can become bugs that need to be fixed:

- *Multiple UI elements that may not render correctly*
- *Buttons that need to be wired to API calls*
- *Multiple API call that may not function as expected*
- *SQL queries to add the inputs to the database*

Obviously, any skilled developer could probably build a simple form like this with no fanfare, but imagine all of the potential pitfalls from a simple form being multiplied across an increasingly complex system. In order to gain one new piece of functionality, there are dozens or hundreds of potential points of failure that would need to be debugged.

So what’s the implication? That software isn’t to be trusted??? No, of course not. All of these risks have have run through by thousands of engineers before us, so we know how to handle these issues reliably. BUT… the more features you add to a platform, the more that need to be supported and maintained.

In other words, the number of real world problems you can solve with your application is limited by the size of your technology team, which in turn is limited by the amount of revenue you generate from your customers, which is in turn limited by the number and severity of problems you’re solving for your customers. There comes an equilibrium of all of those forces you can’t push beyond.

### Filters Instead of Prioritization

One of the most annoying practices in agile culture (in my humble opinion) is prioritization ad nauseum. Everything in agile software development is about prioritization of solutions. This leads to a constant state of meeting after meeting prioritizing the latest set of problems within the always growing list of existing problems.

A better method is filtering. What problems need to be solved and which don’t matter if they’re never solved? Empowering your people to filter these problems will yield incredible results. Your engineers and product people don’t need to know the order of prioritization for 20 different items in order to begin solving problems, that just slows them down. However, they will know which buckets of items need to be fixed asap and which don’t matter if they’re never fixed.

### Progress Is Incremental

Think about your daily life as a busy adult. You wake up and get ready for the day. That takes time. If you’re a parent, you also have some little minions to get ready for a daily routine. More time. Get kids to school, commute to work. Daily hygiene routines, eating meals, cleaning and maintaining your home and car. Mundane tasks at work. Dozens of little activities that are needed of you just to keep going. How much of that time do you have left to improve yourself?

Do you think building software is any different? There’s a lot of work that has to be done just to keep things humming. Software updates have breaking changes that need to be fixed. Security patches need to be applied and tested. Changes to real-world processes and workflows need to be applied to existing software functions (do you really think it took 4+ years to transition to SOFR because it took that long to say, “We want that one”?).

This leaves only a small part of your available capacity available to be committed towards forward momentum. Anything less, and you have a fragile system.

### To Err Is Human

Perfection isn’t our default state. Don’t get too caught up trying to achieve it, or thinking that it’s somehow attainable. That sort of reasoning will only end in failure, *if we just work hard enough, plan hard enough, everything will come out perfectly. If we don’t, it’s because someone screwed up somewhere.*

This just isn’t reality. Chaos is our default state. We build things to make our situation better, but anything we build is ultimately built by people, which are only human. You can’t change that. Instead of trying to fight it, instead build processes with the knowledge that things will never be perfect. Your software will be better, your business will be better.

Make your software better, not perfect.

¹*Okay, like a lot of quotes we attribute to famous people, he didn’t actually say this, but it sounds cool.*

* * *

If you liked this article, consider following me and [subscribing to email updates](#/portal/signup) whenever I post an article. You can also follow me on [Twitter](https://twitter.com/LewallenCaleb) or connect with me on [LinkedIn](https://www.linkedin.com/in/caleb-lewallen-b8699365/).
