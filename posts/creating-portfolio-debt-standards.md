---
Page Name: Creating Portfolio Debt Standards
Published: true
Published Date: 2022-12-12
Author: Caleb Lewallen
RobotsAllowed: true
---

![Feature Image](/static/images/photo-1546250001-6b7b60086f17)

*This article is part of the Debt Policy Series. The first part of the series can be found* [*here*](__GHOST_URL__/why-your-middle-market-real-estate-fund-needs-a-debt-policy/)*.*

So far we’ve covered why we need to change the way we approach debt management and quickly outlined a plan for fixing these issues. This next post will discuss specifics on how to implement these policy standards. The first thing you have to be able to do is establish your portfolio risk metric levels for total capitalization and debt service. These will levels will depend on your risk tolerance and business plans.

### Capitalization Rate

This is simply the amount of debt you’re willing to take on for the value of your assets, essentially a portfolio level LTV. This level needs to be set based on your risk tolerance and business plans, but shouldn’t be determined based on how well the economy is doing or the what the Fed is doing, especially when times are good. Market conditions can change way faster than you can de-lever.

In this instance, I’m an advocate for prudence. I’ve known a company to set a portfolio-wide capitalization rate as low as 45%, which may be a bit extreme for many companies, so it depends entirely on your risk appetite.

*Why is this important?* We want financing flexibility. We want dry powder when things take a down turn. I’ve seen a lot of borrowers with loans maturing this year, and those with maturities coming in the next several months, get stuck in a tight place. Lenders are still willing to lend, but at higher cap rates and lower LTVs.

### Example

Let’s put some numbers to this and narrow in on a single asset. We’re going to assume that in 2019 we bought a property with a $1,000,000 NOI at a 4% Cap Rate. We were able to secure a loan at 75% LTV.

Fast forward to 2022, and it’s time for a refi, since our old loan is maturing. For simplification, we’re going to assume the NOI is still $1,000,000. I have no problem getting quotes from lenders, they’re still willing to lend. Aside from an increase in borrowing costs, the biggest change is a drop in valuations and lower leverage being offered. The best loan you can find is a 65% LTV at a 5% Cap Rate.

So if we plug these values in we get.

Good Times Loan = ($1M NOI / 4% Cap Rate) * 75% LTV = **$18.75M Loan**

Bad Times Refi = ($1M NOI / 5% Cap Rate) * 65% LTV = **$13M Loan**

Even though nothing’s really changed about our property, I still have to come to the table with more than $5M just to refi this loan! That $13M loan is the equivalent of a **52% LTV** loan back in 2019.

If we were smarter about this, we would have modelled out these risk factors back in 2019 to prepare for this. We can simplify the work we did above to determine our maximum LTV. Instead of underwriting our property to the good times, we’ll underwrite what we believe to be a bad times scenario, like a 65% LTV at a 5% Cap Rate.

The math is then: (65% Bad Times LTV * 4% Good Times Cap Rate) / 5% Bad Times Cap Rate = **52% Good Times LTV**. If this had been my Capitalization Rate standard in 2019, I would have only borrowed to a maximum of 52%.

### Creating a Standard

The way to create a standard is now to simply expand this to your entire portfolio. Using this methodology, you can calculate the maximum portfolio capitalization rate, as well as the maximum LTV each individual asset can take on to meet the standard at both an individual and portfolio level.

I’ve created a spreadsheet that you can use to manage this on your own. You can grab a copy of it and use it as is, or as a basis for creating something of your own.

[Get the Google Sheet](https://docs.google.com/spreadsheets/d/1ln6fVFQm22CAtzC9o7B7S8sWF75Dzxaivyqq3PHdpC4/edit#gid=0)

### Keeping it simple

Sometimes the barrier to implementing something new is the initial lift. It’s a lot of work to go figure out the exact cap rate of a property at purchase, or what the cap rate of a property is now. If you know what your average cap rate is, then you can use the math above to figure out what your maximum capitalization rate should be.

Remember, the point is to set a standard for the MAX average LTV across all of your properties. This is the type of situation where fuzzy math is completely ok. The point is prevent (or at least minimize) the risk of having to ask investors for additional capital in order to refi, this doesn’t necessarily have to be really precise.

### Debt Service Coverage

The DSC is the amount of principal and interest your property support with its net operating income. Just like with the capitalization rate, this should be set based on your risk tolerance and business plans. We want to plan for the future — especially when money is cheap and easy to find. Rates have increased dramatically, the time to make sure we could cover it was a year ago, not today.

*Why is this important?* This policy standard serves dual purposes — to make sure that you can continue servicing floating rate loans if interest rates rise, but also to ensure that you don’t get DSCR constrained if you have to refi in a volatile, high interest rate environment. Because of this, we’re going to be approaching the DSC problem a little differently than your lender will. They don’t care how you perform once the loan is paid off, so we can’t use a credit facility example here. We have to refi this at some point!

### Example

We’re going to start with a single asset again to illustrate how this works. For our portfolio level work, we’re going to ignore what the actual loan definitions for debt service are — our goal is maintaining solvency, we’ll worry maintaining financial covenant requirements at the asset level.

In our example, we’re going to start with our earlier example — in 2019 I bought a property with a $1,000,000 NOI. The loan I got was sized at 4.5% over 30 years at a 1.25x DSCR, or roughly **$13.15M**.

Fast forward to 2022, it’s time to refi. We went back to the same lender who agreed to underwrite at the same terms of 30 year amort at a 1.25x DSCR, but now rates are up and we’re borrowing at around 7% (*Side Note: I think there’s something comical about the fact that this reasonable scenario would have sound ludicrous to all of us just a year ago*). At this rate, our loan gets sized back down to just $10M.

Just like with the capitalization rate scenario, nothing about our property has really changed, but I’ve still got to come up with over $3M just to refi my property. If we went back in time and borrowed money assuming that underwriting rates could jump to 7%, our DSCR would have been:

**NOI** = $1,000,000

**Monthly Debt Service** = ~$10M debt @ 4.50% over 30yrs = $50,772.42

**Annual Debt Service** = $50,772.42 * 12 = $609,269.09

**DSCR** = $1,000,000 / $609,269.09 = **1.64x**

Had we had a crystal ball, we would have known to borrow at a max of a 1.64x DSCR at those rates.

### Creating a Standard

To apply this to our portfolio, we need to understand what our current borrowing rates are and what we think a worst case scenario is. Like always, there’s no magic to this, where you set that portfolio DSCR is dependent on what you envision as a reasonable “Bad Times” scenario.

Before we can apply this to our portfolio, we need to simplify the math, and preferably remove any absolute numbers from our math (no need to deal with $ values if we don’t have to). We can simplify the above work into the following (using Excel Formulas so you can plug these into your own).

> *Amort Periods = 360*

> *DS1 = PMT(Bad Times Rate / 12, Amort Periods, -1) x 12 x Min DSCR*

> *DS2 = PMT(Good Times Rate / 12, Amort Periods, \[-1 / DS1])*

**Good Times DSCR** = 1 / DS2 * 12

What we’re solving for is the DSCR we should be sizing to given our assumptions about the bad times market environment we don’t want to be surprised by.

I’ve created a Google Sheet you can use as a starting point in designing your own portfolio standard.

[Get the Google Sheet](https://docs.google.com/spreadsheets/d/1sStdIwufp0jgyZPHMAMojv14aKLB0O0D3yGmKunYR3Q/edit?usp=sharing)

### Why Do This?

If you have a crystal ball and can tell when big market swings are going to happen, then this isn’t for you. For the rest of us, the point of this kind of exercise and rigor around our debt portfolio management is precisely because we don’t have a crystal ball.

Are you potentially leaving yourself with a lower ROE by purposefully limiting your debt exposure? Of course we are, but we’re also not finding ourselves completely underwater when we’re all surprised by the next black swan event.

You’ll notice that in our portfolio management practices we’ve purposefully left out absolute property values. This isn’t just to make the math simpler, it’s also to help remove the temptation of treating a ***debt management*** problem like it’s an ***asset management*** problem. If we only deal with debt variables, we can better manage our debt risk.

*This is Part 3 of the Debt Management Policy series.*

**Next:** [Debt Standard Checklists for Individual Loans](__GHOST_URL__/debt-standard-checklists-for-individual-loans/)

* * *

The Debt Management Policy is a short series of articles that provide the resources and materials for building and establishing Debt Management Policies for your own CRE firm.

1. [Why Your Middle-Market Real Estate Firm Needs a Debt Management Policy](__GHOST_URL__/why-your-middle-market-real-estate-fund-needs-a-debt-policy/)
2. [A Better Approach To Debt Management](__GHOST_URL__/a-better-approach-to-debt-management/)
3. [Creating Portfolio Debt Standards](__GHOST_URL__/creating-portfolio-debt-standards/)
4. [Debt Standard Checklists for Individual Loans](__GHOST_URL__/debt-standard-checklists-for-individual-loans/)
5. [Setting and Enforcing Debt Management Policies](__GHOST_URL__/setting-and-enforcing-debt-management-controls/)
6. [Challenges You Will Face Implementing Debt Management Policies](__GHOST_URL__/challenges-you-will-face-implementing-debt-management-strategies/)

* * *

If you liked this article, consider following me and [subscribing to email updates](#/portal/signup) whenever I post an article. You can also follow me on [Twitter](https://twitter.com/LewallenCaleb) or connect with me on [LinkedIn](https://www.linkedin.com/in/caleb-lewallen-b8699365/).
