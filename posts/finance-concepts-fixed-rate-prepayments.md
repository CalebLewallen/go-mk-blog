---
Page Name: Finance Concepts: Fixed Rate Prepayments
Published: true
Published Date: 2022-09-28
Author: Caleb Lewallen
RobotsAllowed: true
---

![Feature Image](/static/images/1-ejtjgcx6n2xlkwuy50xvrw.png)

As all CRE borrowers know, if you get a fixed rate loan, you’re more than likely going to be saddled with some sort of prepayment premium. There’s several common calculation methods, all going by different names (sometimes the wrong one), but they all follow the same basic concept.

The **Make Whole Prepayment** is straightforward in concept. When the lender agrees to make the loan, they expect to earn some rate of interest for a specific period of time. This rate of return is modeled into an ROE, and income is projected for the lender. If you want to prepay the loan early, the lender then has to redeploy those funds into an unknown interest rate and lending environment. This presents a series of risks to the lender:

- The lender has already earmarked those funds for deployment through the maturity of the loan. In order for you to prepay, they have to find an investment vehicle to redeploy those funds through the maturity of the paid off loan. This means the lender has to lend those funds out for a shorter amount of time, generally at a lower yield, an effect known as **rolling down the yield curve**.
- New loans are typically priced as a spread over Treasurys, meaning the rate they’re able to lend at is dependent on rates at the time of the prepayment. If rates have fallen since the loan originally closed, the lender will get a lower rate of return, exposing them to **interest rate risk**.
- Lenders are in the business of, well, lending. In order for the lender to redeploy those funds, there has to be 1) borrowers willing and qualified to take on debt and 2) sufficient appetite from the risk management department for new debt. This exposes lenders to various types of **lending risk**.

Ultimately, the lender is trying to solve for the difference between the amount of interest income they thought they would receive and what you’ve paid so far, like so:

![](https://170.187.153.50/content/images/max/800/1-iasv7cvodgcma5ndvk9c1g.png)

This leaves really only one investment vehicle for lenders with the necessary liquidity to facilitate any sort of rate of return: **US Treasurys**

“Wait,” you may be thinking, “didn’t I lock this as a spread *over* Treasurys? Isn’t that going to leave the lender short?” If so, you would be correct. That means that you have to make up the difference.

So then the problem we have to solve is this — what lump sum of money, invested into a risk-free rate of interest, would yield the same dollar amount (FV) as the amount of interest still owed on the loan? There are three primary ways this is calculated.

1. Yield Maintenance
2. Make Whole
3. Defeasance

### Financial Concept

All fixed rate prepayments share the same financial concept, which is to replace or discount the future fixed rate loan payments with some risk-free rate of interest that the creditor could redeploy their funds to. For calculation purposes, USD based loans are typically discounted using US Treasury securities.

To calculate a back-of-the-hand estimation of a fixed rate prepayment, one would use this formula:

![](https://170.187.153.50/content/images/max/800/1-dwdq809nep1ictruxrsgbw.png)

> L\_avg *= Average loan balance through end of prepayment period*

> r\_note *= Interest Rate on the Note*

> r\_risk-free *= Yield on a US Treasury security equal to the term from the prepayment date through the end of the prepayment period*

> T\_n — t\_0 *= Time in years from prepayment date to end of prepayment period*

For example, assume we have a ***10 year, non-amortizing loan of $25M priced at 5%***. To make the math easy, let’s also assume a 30/360 daycount. Five years have passed since the loan close, and five years are remaining until maturity. The current five year treasury yield is 2.25%.

$25,000,000 x (5.00% — 2.25%) * 5\_yrs = **$3,437,500**

Keep in mind that this is just an estimate. This back-of-the-hand math has not PV’d the result and will generally be higher than your actual penalty.

### Yield Maintenance

First thing to note, this term is very often used a a catch-all term for prepayments in loan agreements. Just because your docs say “Yield Maintenance”, doesn’t mean your prepayment will be calculated as a yield maintenance.

The yield maintenance calculation method follows the concept above by calculating a hypothetical rate of interest that is the difference between the Note Rate and the Risk Free rate of interest. That is then discounted back to the prepayment date.

![](https://170.187.153.50/content/images/max/800/1-kkskhpveisonqvobnblfyg.png)

The first couple of prepayment periods in the loan above would look like this:

![](https://170.187.153.50/content/images/max/800/1-woh_9sqrhy86oyo5g5kqta.png)

This process is repeated for the entirety of the remaining set of cashflows. The sum of the Yield Maintenance Interest column is taken, which represents the **Prepayment Penalty** of the loan.

In our example loan, the prepayment penalty is $3,248,315.86, which is pretty close to what we estimated in our back-of-the-hand method above.

> Most Yield Maintenance terms include a minimum penalty of 1%, which needs to be accounted for in your final result.

### Bonus: Agency Style Yield Maintenance

The big multifamily agency lenders, like Freddie Mac and Fannie Mae, have their own flavor of yield maintenance that they employ. This method assumes that all cashflows after the prepayment date are a hypothetical interest-only cashflow, discounted by an interpolated US Treasury Rate. This disadvantages the borrower if their loan amortizes.

![](https://170.187.153.50/content/images/max/800/1-lruadp_hmfmtld4ioa37xq.png)

*Where OPB = Outstanding Principal Balance | r = Treasury Rate | z = time in years to end of prepayment period | R = Note Rate*

This will also almost always have some minimum penalty, usually 1%.

### Make Whole

The Make Whole method, often done under the moniker of “Yield Maintenance”, is probably the most common form of fixed rate prepayments in bank loans. While it will produce a penalty materially similar to that of Yield Maintenance, it is a distinct calculation with it’s own method.

Make whole calculations are essentially discounted cashflows. Mostly commonly, the language will require that you discount the remaining payments of the loan over through maturity, including the balloon payment. The discount rate is almost always Treasurys, often an interpolated rate from the daily H.15 publication from the Federal Reserve.

![](https://170.187.153.50/content/images/max/800/1-hqydogotxhcz4r0cw9t-yg.png)

Where r = discount rate | T-t = time to today in years

The goal of this method is to find the **total amount due to lender** at prepayment. This, subtracted from the loan balance at the time of prepayment, is your penalty. A sampling of those cashflows would look something like this.

![](https://170.187.153.50/content/images/max/800/1-myfolkdaxshmcpeftlcvbw.png)

The last period of the PV’d cashflows would include any principal balloon payment

Since this loan is interest only, the Make Whole penalty is identical to Yield Maintenance, at **$3,248,315.86.** The values will differ slightly if the loan amortizes.

Like our previous methods, some minimum penalty will usually apply.

### Defeasance

Ah, Defeasance. A word known to cause dread in many an operator. If you haven’t been a part of a defeasance before, it’s just a process by which real estate cashflows are replaced with a series of maturing Treasury Bonds/Bills (i.e. a “substitution of collateral”). While the first several methods we covered try to answer the replacement question in a theoretical sense, the madman that created defeasance decided to actually carry it out.

This replacement is achieved by structuring a portfolio of securities, that by the virtue of their maturation, exactly replicates the cashflows that would have been delivered by the real estate. The goal of the **securities broker** and **defeasance consultant** is to find the most efficient portfolio (read: cheapest) portfolio of securities that can meet the cashflow requirements of the loan.

The goal of defeasance is to provide a means for a borrower to unencumber an assets from a securitized loan, while protecting the returns of the investors in the CMBS securitization. Your portfolio cost will be materially similar to that of a Make Whole, but will also incur $50K — $75K of additional transaction fees.

### Debt Management — How to Manage Prepayments

Since this is a publication about debt management, let’s take a look at some characteristics of these prepayments and how to manage them to minimize risk.

### Prepayment Burndown

Due to the calculation principles behind fixed rate prepayments, these loans typically have a characteristic of “burning down” over time. Each month a prepayment is delayed, the penalty will burndown by roughly the amount of the monthly loan payment, *ceteris paribus*.

The chart below shows the effect of the burndown over the 60 remaining months of our imaginary loan, assuming the risk-free rate of interest remains constant for the remaining term of the loan, calculated using the Yield Maintenance method. Note that the prepayment penalty decreases consistently until the minimum penalty is reached.

You can use this information strategically in timing a sale of an asset. Even delaying a sale a few months can have a measurable impact on return metrics

![](https://170.187.153.50/content/images/max/800/1-wyqilvls5yh3wxuzfdezcw.png)

### Prepayment Language

According to Freddie Mac loan data, roughly half of all loans will be prepaid early. Since you’ve got a 50/50 chance of having to deal with a prepayment at some point, we need to plan for it. The best way to minimize the risk of a prepayment blowing up a deal is to de-risk the loan before it ever closes. Below are a few prepayment provisions to include in your negotiation with the lender.

#### Yield Maintenance and Make Whole Prepayments

**Par Prepayment Periods**

The par prepayment period is the time leading up to the loan maturity when you can prepay the loan without penalty. This is usually 3 months, but we want to make this long as possible. I’ve seen periods as long as 2 years successfully negotiated, but you’re more likely to land on somewhere around 6 months.

**Prepayment Calculation Periods**  
This is how long we have to run the calculation for. We never, ever want to run the calculation to maturity. Whenever possible, run the calculation to the start of the Par Prepayment Period. If you can prepay without penalty 6 months prior to maturity, why would we calculate the prepayment through maturity? If our par prepayment window is 3 months long, we would have saved *$153,891.41* in our above fictitious loan, or *$308,650.08* if it was 6 months.

**Discount/Replacement Rate Spread**  
Try to get a spread added to the replacement/discount rate. I’ve seen many prepays with language allowing for Treasurys + 50bps, but obviously get that spread as high as you can. The higher the spread, the cheaper the prepayment. A 50 bps spread would have brought our 2.25% discount rate to 2.75%, saving us *$623,413.97.*

Keep in mind these changes are cumulative — including both this and the prepayment calculation periods would save nearly $900K.

**Waive Penalty for Refinancing with Lender**  
Let’s say you locked a fixed rate loan in the next couple of months. Your rate is going to be pretty high. If we enter a recession in the next 12 months and the Fed is forced to cut rates, you’re likely going to want to refinance. Unfortunately, there’s an inverse relationship between a prepayment penalty and Treasury rates. As those rates fall, your prepayment penalty will spike, and likely make it prohibitively expensive for you to refinance. Some lenders, especially local/regional lenders, will agree to waive the prepayment penalty in the event that the refinancing is done with them.

#### Defeasance

**Portfolio Length & Par Prepayments**  
Similar to the above, make sure to negotiate as long of a par prepayment period as possible. You also only want to structure the security portfolio through the end of the defeasance period.

**Successor Borrower Designation**  
This is arguably the most important term to negotiate, depending on how much your lender will agree to with portfolio length and par prepayments. Because the defeasance is just a substitution of collateral, someone has to step in as a replacement borrower to the loan after you step out. This entity is the Successor Borrower (SB). The security portfolio generates revenue for the SB over its life, which is a value that can be sold for an upfront fee or revenue sharing agreement, offsetting some of the prepayment cost.

**Security Types**  
Since we actually have to build a portfolio to compensate the security investors, we can’t artificially inflate spreads like with YM/Make Whole calculations. Instead, we want the basket of securities that we can pick from to be diverse as possible. These usually include different agency securities, like Freddie Mac, Fannie Mae, or FHLB securities. The more your broker has to pick from, the more options they have to decrease portfolio costs.

### Summary

Fixed rate prepayment methods all follow the same core principles, understanding the concepts behind how they are calculated can help you improve your debt management practices and improve your deal profitability.

* * *

If you liked this article, consider following me and [subscribing to email updates](#/portal/signup) whenever I post an article. You can also follow me on [Twitter](https://twitter.com/LewallenCaleb) or connect with me on [LinkedIn](https://www.linkedin.com/in/caleb-lewallen-b8699365/).
