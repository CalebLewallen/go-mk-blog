---
Page Name: FDIC Takeover: An SVB Primer
Published: true
Published Date: 2023-03-13
Author: Caleb Lewallen
RobotsAllowed: true
---

![Feature Image](/static/images/svbcollapse1.jpg)

By now we've all heard about the implosion of Silicon Valley Bank. If you haven't, then this is a really odd place to be hearing about it for the first time.

This bank failure represents the second-largest in US history, but the way it failed is largely unprecedented (to my knowledge). This post is intended to be a primer on the long road to its insolvency, and whether this is an indicator of things to come.

# How Banking Works

Before we can get to how and why SVB failed, it's important to understand how the business of banking functions. Banks have a really interesting business model, unlike any other business in the world. Most businesses make money by *selling* you something. A product, a service, etc. It might be something you produced directly, or it might be as a market-maker.

For example, Samsung designs and manufactures TV's. They make money by selling TV's for more money than it cost to produce. The economies of scale they operate at doesn't allow them to sell directly to consumers though, so they sell their product to wholesalers, which eventually makes its way down to a retailer. You, the consumer, can then buy your TV from Walmart, who also sells the TV for more than they bought it for.

While banks do have divisions that sell financial products, the way the business as a whole is stitched together is pretty unique.

## Depositors

If you decide to open your own bank, the first thing you're going to have to do is raise money from **depositors**. This can be individuals, or from other businesses. Depositors will place funds in accounts at your bank, and in exchange you agree to give them access to those funds anywhere in the globally connected financial world.

In return for offering that service, you don't charge a monthly fee, in fact you will often PAY depositors to keep their money in that account.

So, with that kind of a business model, how do you make money?

## Collecting Assets

Banks make money in a variety of ways, but one of the primary engines of that income is by issuing debt, either on the primary market is the issuance of loans, or on the secondary market purchasing debt.

The idea here is that you will make more money on the products you lend than the interest you have to pay to depositors to keep money at your bank.

Pretty straightforward right? Except...

## Fractional Reserve Banking

Ok, so this is where things get weird, and where things can get fragile. **Banks lend out more money than they have cash on their balance sheet**. How does this work?

Meet Joe. Joe is local businessman who owns a shop on main street. He decides to open a new account at your bank. He's been doing well, and has decided to move all of his bank business to you. He opens up the accounts he needs for his business and deposits $100,000.

Here's our current situation:

| Depositor | Deposits | Required Reserves | Amount Available to Lend |
|-----------|----------|-------------------|--------------------------|
| Joe       | $100,000 | $10,000           | $90,000                  |

Ok, so we have something new here: **Required Reserves**. Even though we want to loan out Joe's money, we have to keep some of it on hand in case Joe needs it (you know, to make payroll). The Federal Reserve has regulations on the reserves that banks must have on hand for different financial products. For now, let's call it 10%.

Now that we have Joe's deposit, we can lend out up to $90,000 (Joe's deposit, less 10%). The next day, Emily walks in and wants to take out a loan so she can start up a business down the street. She's got a great business plan, and a track record of building successful businesses. You make the loan on the condition she does her banking with you. She agrees, and deposits her newly acquired $90,000 of cash with you.

Let's review where we are now:

| Depositor | Deposits     | Required Reserves | Amount Available to Lend | Amount Lent |
|-----------|--------------|-------------------|--------------------------|-------------|
| Joe       | $100,000     | $10,000           | $0                       | $0          |
| Emily     | $90,000      | $9,000            | $81,000                  | $90,000     |
| **Total** | **$190,000** | **$19,000**       | **$81,000**              | **$90,000** |

We have $190,000 deposited at our institution now. We are obligated to provide that cash to our depositors *on demand*. However, I've only got the original $100,000 in cash to cover those demands. If you carry this out to it's logical conclusion, I could actually lend out as much as $900,000 on the original deposit!

I'm sure the intelligent among you can see the problem. I could have as much as $900,000 in deposits I'm required to produce on demand, but I've only got $100,000 in cash to cover the demands.

If this is the first time you're hearing this, just know that the banking system of the entire world runs this way. This system runs smoothly so long as *everyone* doesn't want *all* of their money *at the same time*.

We'll get back to this soon...

# Enter, SVB

Silicon Valley Bank was founded on October 17, 1983 by Bill Biggerstaff and Robert Medearis. The idea was simple – they were going to cater to startups funded by venture capital. Once banking customers, they could add additional services to continue catering to these clients as they scaled their business.

The bank has had its ups and downs over the years, largely influenced by investment cycles. During the dot-com boom, their stock price soared and fell again as soon as it burst. In 2000, the bank gained a new CEO, Ken Wilcox, who chose to continue to concentrate in servicing tech companies instead of diversifying into a more traditional bank servicing model.

By 2015, the [bank claimed to serve 65% of all US Tech startups](https://web.archive.org/web/20200401051042/https://www.nytimes.com/2015/04/02/business/dealbook/silicon-valley-bank-strengthens-its-roots.html), and many prominent VC firms.

## The Pandemic

As many of you will recall, financial markets went through an interesting period during the pandemic. Instead of collapsing in the face of an economic slowdown, the stock market boomed. Armed with stimulus checks, a lot of people stuck at home began doing what every bored person with some extra cash does – they start gambling. This pushed valuations for public companies through the roof, and as a result, the VC business model become even more profitable.

With public valuations getting pushed ever higher, private valuations weren't far behind. Plus, without the votes of the masses influencing stock prices on Robinhood, private valuations can sometimes be even more insane than GameStop circa 2021.

WFH meant more people were spending money online than ever before, and a new generation of entrepreneurs were ready to build new products for them to spend that money on. Those new products usually needed funding, and money was cheap (and often easy to get). It almost reminded you of the party scenes from The Great Gatsby – the money flowed and everyone had a good time.

Venture Capitalists had a lot of money to place, and if you had a URL ending in ".ai", you got a check. A big one, and Silicon Valley Bank was standing by ready to serve your needs. Over the course of 2020 into 2022, SVB's total assets swelled from $71B to over $211B – The 16th largest bank in the nation.

## Treasury Portfolio

With all of these new deposits coming in, SVB decided that instead of rushing poor quality loans out of the door, they would try to do the responsible thing and invest their depositor's money into the safest type of investment out there – US Treasury Bonds.

At face value, it's pretty smart actually. But they didn't execute on it very well.

For starters, as we all know, bond prices are inversely related to yields. Yields up, price down. They got all of these new deposits in a zero interest rate policy environment, and subsequently bought some of the most expensive Treasurys in history. But what ultimately ended up causing issues, was that most of their investments were in long-dated stuff that wasn't going to mature anytime soon. This exposes them to interest rate risk.

The problem with being concentrated in the startup world is that startups tend to burn cash. So, as rates rose, two things happened: 1) the party stopped, the money stopped flowing, and depositors as a whole were pulling more money out than they were putting in, and 2) bond prices were in free-fall.

Now, all of this money they had tied into bonds had to get liquidated in order to fulfill their obligations to their depositors. They had to sell sub-2% coupon bonds in a ~3.90% rate environment, leading to a loss. They're obligated to make depositors whole, so they had to back-stop this loss by issuing stock. This spooked markets, and more depositors, who demanded more cash.

The cycle had begun...

## Hedging

Banks use hedging for the exact same reason you do – they don't want to get pinched when rates move the wrong way. In this case, the swaps they used for hedging were Fixed Payer swaps. This means that they paid a fixed rate, and received a floating rate.

In this position, the swap value moves in their favor as rates *rise*, and moves against them as rates *fall*. It's a handy instrument to have when you hold a bunch of long-dated Treasurys.

SVB did actually have some hedges in place at the end of 2021, about $10.7B worth. It was less than 10% of their total portfolio, but you don't really *need* to hedge your entire treasury portfolio, unless all of your depositors suddenly start demanding all of their money back at the same time 😬.

Here's a summary of SVB's hedge positions for FY 2021 and FY 2022. You can see that most of their hedges had been unwound over the course of the year

![](https://170.187.153.50/content/images/size/w1000/2023/03/image-2.png)

*This is pure speculation on my part*, but I think this is one of the most concrete signs that SVB knew things were bad way back in 2022. They had a LOT of fixed-payer swaps in a steeply rising interest rate environment, they would have been worth a lot in SVB's favor. I'm sure they made a killing unwinding those positions. If you needed a way to raise some cash without selling any assets, this would have been a good way to do it.

# Bank Runs

What's the surest way to kill a bank? Convince enough depositors to all demand their cash at the same time. That's exactly what happened with SVB.

To be clear, it wasn't the bad hedging, it wasn't the losses from their treasury portfolio, it was the good old fashioned rumor mill that collapsed SVB. Last Wednesday, SVB surprised investors with the announcement that they would need to raise $2.25B to cover their realized bond losses. On Thursday, SVB seemingly botched an investor call that sent their stock price plummeting, and by Friday they were taken over by the FDIC. So what triggered this?

## The Dominos

First things first – nobody actually pays attention to what's going on at their bank. This includes the VC's and founders that used SVB for their banking. That is, until on February 23rd, Byrne Hobart's newsletter, The Diff, called out SVB for their less than stellar risk management situation.

![](https://170.187.153.50/content/images/size/w1000/2023/03/image-3.png)

Every VC on the planet reads Byrne Hobart. Every VC on the planet is also in the business of evaluating the viability of businesses, so after reading this they do what VC's do, and they started looking into this situation themselves. It wouldn't be too much longer until the next earnings call, so they started pay attention.

### Strategic Actions/Q1'23 Mid-Quarter Update

This was the title of the investor presentation sent out on March 8th. It was a pretty standard looking bank presentation. You can take a look at it yourself [here](https://s201.q4cdn.com/589201576/files/doc_downloads/2023/03/Q1-2023-Mid-Quarter-Update-vFINAL3-030823.pdf).

First, if your bank is sending out a voluntary mid-quarter update, it's probably not good news. The first couple of pages they rip the band aid off and say that they're taking a big realized loss and then spend the rest of the presentation trying to convince you it's not a big deal.

It's a little bit like when one of my kids tells me they got an F on an exam, but it's okay because as long as they get a 55 on the final they'll still pass the class.

This, of course, spooked the hell out of not only their investors, but their depositors (and more importantly, the VC's that backed them).

### The VC Scare

Peter Thiel, USV, and Coatue send out messages to their portfolio companies to pull out funds from SVB. Cue the first couple waves of bank runs, but nothing the bank can't dig up cash for.

You can't keep this sort of thing a secret for long though (measured in hours, not days). Pretty soon, tech twitter catches word that all of their buddies have been told to get out, making them wonder if they should be pulling out as well. That's when the real bank run starts.

In the space of 48 hours, a staggering $42B in deposits are pulled out of SVB, roughly 25% of their base. I'm not sure even the most well capitalized banks are in a position for 25% of their deposits to evaporate overnight.

### Back to Fractional Reserve Banking

Recall our explanation of fractional reserve banking above. The bank is fine so long as everyone doesn't want all of their money at the same time. What's interesting about SVB though, is their assets didn't really outstrip their deposits by all that much. Most of it was just stashed in Treasurys.

The problem was still the same though. They hedged badly (or rather, not at all), and when depositors came for their cash, the bank was forced to take a loss. That loss is sizeable for a bank their size.

# A Sign of Things to Come?

So let's get this out of the way, is this a sign of things to come? Yes and no. Yes, there are going to be banks that hedged badly and got caught by the changing interest rate environment. No (probably), the largest banks that represent a systemic risk if they collapsed probably won't face any trouble. They're well capitalized, stress tested annually, have way better hedging strategies in place, and aren't diversified into a single, fickle group of depositors.

However, we can never say never. This could just be the first domino to fall. But let's also try to see this in context. SVB got pinched because they were too heavily concentrated in one type of clientele, and stuck all of their cash into an instrument that would expose them to interest rate risk without hedging the exposure.

### How Did This Happen? Aren't Banks Regulated?

Yes, banks are regulated, but many of the regulations, hedging requirements, and reserve requirements only apply to banks with more than $250B in assets under management. This includes the annual stress tests that the largest banks in the country undergo each year. SVB fell shy of that, with around $212B at its peak.

I think the irony in all of this, is that their asset portfolio was ~50% Treasurys – they tried to do the responsible thing! Well, responsible from the perspective of a risk committee with little no risk management experience. What they didn't prepare for was the potential of having to realize paper losses.

### It took more than 48 hours

SVB knew about this long before last week. They did operate without a Chief Risk Officer for most of the last 18 months, with Kim Olsen not joining until late 2022 – too late to do anything about SVB's precarious position.

The folks at Forbes [released some disclosures](https://www.forbes.com/sites/noahbarsky/2023/03/12/silicon-valley-bank-proxy-shows-boards-secret-yearlong-risk-panic/?sh=6a1976f91e7b) they believe point to some signs that over the last year, SVB was acutely aware of the position they were in.

- It's 2023 proxy statement called for **seven** of its 11 board members to serve on its risk committee, while no other committee has more than 5 members. Interestingly, their most qualified director, Thomas King, a former IB CEO, wasn't on the board. His experience would have been far more useful than that of a Napa vineyard owner, a retired healthcare CIO, a former US Treasury undersecretary, VC partners, and consultants.
- The risk committee met about twice as frequently in 2022 as it did in 2021 – 18 times vs 7 times. It's not clear whether these were evenly spaced, or whether the frequency accelerated towards the end of the year.
- Most troubling was that SVB operated without a CRO for most of 2022. Their former CRO was moved into a non-executive role in April 2022, and a replacement wasn't announced until January 2023.

## Takeaways

There's a couple of things that make SVB's failure (and most recently, Signature Bank) unique from the banking sector as a whole:

- Each bank was concentrated in a very niche market. SVB serviced startups, and Signature was focused heavily in crypto. Both segments are coming out of recent boom cycles, leaving the banks exposed to a rapidly shrinking deposit base.
- SVB tried to do the conservative thing and place deposits into Treasurys. A smart move, if their exposure had been hedged, or if they had rolled short-term investments to ensure they could always be no more than a month away from getting cash they needed.
- The rumor mill killed SVB. Their collapse is putting every other bank under scrutiny, so there is absolutely some non-zero chance that some banks will face scrutiny from their depositors.
- First Republic was suggested as an alternative by the same VC's that told their portfolio companies to pull their funds out of SVB (the same VC's that told them to put their money there to begin with). Not the kind of recommendation you want, sending their stock plummeting.
- The Federal Reserve and the US Treasury is going to be going full-blast trying to reassure American citizens that the financial sector is stable. They've already back-stopped all SVB deposits, beyond the normal FDIC insurance of $250,000 per depositor. If you work for a startup that uses SVB, your employer will get access to their funds to make payroll. The Fed's primary concern, before inflation and unemployment, is ensuring the stability of the US financial system.
- Interest Rates are DOWN. Two-year Treasurys had dropped 100bps since Thursday, and are down by ~50bps as I'm typing this (Monday morning, 3/13/23). This represents the largest single drop in US history.
- Emergency rate cuts are on the table to ensure the stability of the financial system. Not guaranteed, just on the table.

# The Story Isn't Over Yet

This saga is still unfolding. Take everything above with a grain of salt. This is only based on the information that was available in the days following the collapse, so everything is subject to change in the face of new information.

* * *

If you liked this article, let's connect! You can follow me on [Twitter](https://twitter.com/LewallenCaleb) or connect with me on [LinkedIn](https://www.linkedin.com/in/caleb-lewallen-b8699365/).
