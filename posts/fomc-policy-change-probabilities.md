---
Page Name: Calculating the odds of a change in Fed Funds
Published: true
Published Date: 2024-06-26
Author: Caleb Lewallen
RobotsAllowed: true
---

![Feature Image](/static/images/DALL-E-2024-06-25-21.06.31---A-professional--flat-2D-financial-graphic-designed-in-a-style-similar-to-Figma.-Jerome-Powell-is-in-the-center-of-the-image--seen-from-the-waist-up-wi.webp)

**An interactive version of this article can be found here:** [https://ratehike.caleblewallen.com/](https://ratehike.caleblewallen.com/)

Here's a question for you – as of late June 2024, how many rate cuts, if any, do you expect the FOMC to make by the end of this year? Outside of the talking heads on CNBC and the Fed dots in the quarterly SEP, do you have any data points to back up your point of view?

Thankfully, financial markets give us a view into the consensus view from billions of dollars of bets being made at any given time. We can use Fed Funds Futures contracts to calculate our own probability of a policy change.

## Fed Funds Futures Contracts

Before we get into how to calculate the **probability** of a policy move, first we need to understand how Fed Funds and their corresponding futures contracts function.

Fed Funds Futures (FFF) are quoted as a price, not a rate. The rate is derived from the price by subtracting 100 from the price. For example, if the contract you're viewing is priced at 94.88, then the implied rate in that contract is:

\\\[100−94.88=5.12\\]

### Contract Data[](https://ratehike.caleblewallen.com/Fed_Funds_Futures#contract-data)

Contracts for Fed Funds Futures are quote monthly, and have an expiry on the last day of every single month. When you get quoted a series of these contracts, they'll look something like this:

Quote Code Expiry Month Price ZQM3 06/01/2023 5.1200 ZQN3 07/01/2023 5.1750 ZQQ3 08/01/2023 5.2850 ZQU3 09/01/2023 5.2750 ZQV3 10/01/2023 5.2500 ZQX3 11/01/2023 5.1450 ZQZ3 12/01/2023 5.0550

I know what you're thinking, "But isn't fed funds an overnight rate? How is the contract monthly?" I'm glad you asked.

### Contract Pricing[](https://ratehike.caleblewallen.com/Fed_Funds_Futures#contract-pricing)

Fed Funds futures are quoted as the ***average*** fed funds settlement over a given month. So if FF is 5.125% for the first half of the month, and 5.375% for the last half, then the contract will settle at 5.25%.

![](/static/images/newplot--1-.png)

### FOMC Meetings[](https://ratehike.caleblewallen.com/Fed_Funds_Futures#fomc-meetings)

How does this help us calculate the odds of a policy move? The FOMC has an iron grip over the front end of the yield curve. They decide what the overnight rate is going to be and uses a trading desk in downtown Manhattan to keep settlement prices within a given range. With this in mind, knowing that only a Fed action can influence those contract rates up or down by more than a bp or two, we can start piecing together the expected path of hike.

### Step 1: Collect Futures Data[](https://ratehike.caleblewallen.com/Calculation_Walkthrough#step-1-collect-futures-data)

Below is a table of futures data. You can get this data from the CME Group. This data is publicly available on a delayed basis. That's totally fine, we're not trading on this information, we're just trying to get a pulse on how likely the market thinks a rate hike is.

[CME Group - Fed Funds Futures Data](https://www.cmegroup.com/markets/interest-rates/stirs/30-day-federal-fund.quotes.html)

priorSettle volume quoteCode expirationDate productName 94.67 0 ZQM4 6/1/2024 30 Day Federal Funds Futures 94.67 15 ZQN4 7/1/2024 30 Day Federal Funds Futures 94.7 978 ZQQ4 8/1/2024 30 Day Federal Funds Futures 94.76 616 ZQU4 9/1/2024 30 Day Federal Funds Futures 94.855 84 ZQV4 10/1/2024 30 Day Federal Funds Futures 94.925 15 ZQX4 11/1/2024 30 Day Federal Funds Futures 95.03 111 ZQZ4 12/1/2024 30 Day Federal Funds Futures

### Step 2: Get FOMC Meeting Dates[](https://ratehike.caleblewallen.com/Calculation_Walkthrough#step-2-get-fomc-meeting-dates)

Policy Decisions are typically going to happen on predetermined meeting dates (outside of black swan events that require immediate action), so we'll need to cross-reference our futures data with our FOMC meeting dates to figure out when we can expect a policy change to occur.

For the remainder of 2024, these are the meetings remaining. The 2025 meetings won't get announced until closer to the end of the year.

FOMC Meeting Dates 2024-07-31 2024-09-18 2024-11-07 2024-12-18

You can get this data from the [FOMC's website](https://www.federalreserve.gov/monetarypolicy/fomccalendars.htm). Note that decisions are announced on the second day of the meeting, therefore you only need to collect the second date of each meeting.

### Step 3: Build Projected Rate Movements[](https://ratehike.caleblewallen.com/Calculation_Walkthrough#step-3-build-projected-rate-movements)

This step involves taking our futures pricing and calculating the implied rates and rate movements that occur in a given month. Our goal is to calculate Fed Funds at the start of the month vs the end, and then seeing the difference between the two. We're going to make a few assumptions when we work with this data:

- The Beginning Rate for a month is simply the Ending Rate for the prior month. The exception to this is when there are no Fed Meetings.
- Rate movements only occur in months with Fed Meetings. Since we know the rate at the beginning of the month and the implied rate per the futures data, we can calculate the end rate with:

\\\[r\_{ending} = ((r\_{implied} / days\_{total}) - (r\_{begin} * days\_{priortochange})) / days\_{afterchange}\\]

- We can only calculate probabilities as far out as we have planned FOMC Meetings. This means that as we approach the latter half of the year, we can only project out a couple of meetings until the FOMC publishes next year's meeting schedule.

![](/static/images/image-2.png)

Projections as of end of June 2024

📉

Check out live [projected movements](https://ratehike.caleblewallen.com/Calculation_Walkthrough#step-3-build-projected-rate-movements)

## Step 4: Probability Tree[](https://ratehike.caleblewallen.com/Calculation_Walkthrough#step-4-probability-tree)

Calculating the probability tree occurs in two stages. The first is in the 'Front Meeting', which is just the very next meeting after today. This done by dividing the expected rate change by the expected policy move. Historically, this is typically 25 bps, although in recent memory it has been by as much as 75bps at a time.

\\\[Policy Change Probability = Expected Rate Change / Expected Policy Move\\]

As an example, if the front meeting has an expected rate change of 12.5 bps, and we expect that the FOMC's policy change is 25 bps change, then we would say there's a 50% likelihood of a hike.

As we move further down the probability tree, things get a little little more complex, but not all that much more. In concept, what we're trying to figure out is the probability that we land on one of the nodes in the chart below. For example, in the n2 nodes below (representing 2 meetings ahead of today), what's the likelihood that we will have experienced 2 hikes (n2(++))? What about no hike and then a hike (n2(+))?

![](/static/images/image-1.png)

The way we accomplish this is by first calculating the isolated probabilities of each meeting. In other words, what's the likelihood of a rate movement for a given meeting, based solely on the expected rate movement?

Next, we'll use those isolated probabilities to figure out if the anticipated policy change is a hike or a cut, based on whether the isolated probabilities are positive or negative.

The formula we'll use to calculate the probability of one of our nodes is: \\\[(p\_{hikeNow}\*p\_{noHikePreviousMonth}) + (\[1 - p\_{hikeNow}] * p\_{hikePreviousMonth})\\] where 'p' is probability. If futures data indicates a cut, then you would simply reverse the direction of the formula.

For example, here's how we would lay out calculating the probability of a cut as of late June 2024.

![](/static/images/image-3.png)

📐

Check out the [interactive probability tree](https://ratehike.caleblewallen.com/Calculation_Walkthrough#step-4-probability-tree)

### Policy Decision Prediction[](https://ratehike.caleblewallen.com/Calculation_Walkthrough#policy-decision-prediction)

This is probably the most straightforward of the concepts. Now that we have our probability tree built, we can start with deducing the project policy decisions.

We'll begin by marking the FF range we're currently in. We'll add up all of the probabilities above and below that range, excluding the value from the current range. Once the summed probabilities exceeds a pre-defined threshold (in my case, 50%), then we'll mark a policy decision and move our mark up or down into a new FF range. From that, we'll repeat the process. You can tweak this process according to your risk thresholds by adjusting the move threshold up or down.

In our example above, the July meeting shows 90% probability of no policy change. However, in September there's a 58% chance that rates will be 25bps lower and a 6% chance that rates will be 50bps lower, for a overall probability of 64% that there will be a policy move down by the September meeting.

📊

Check out the live [Policy Decision Prediction](https://ratehike.caleblewallen.com/Calculation_Walkthrough#policy-decision-prediction)

## That's it?

That's it! With some statistics and a little patience, we can calculate our own probabilities. While the Fed is projecting only one rate cut by the end of this year, our calculations show that the market is betting on two cuts.

![](/static/images/newplot--2-.png)

As of late June 2024, markets project two rate cuts by the FOMC

* * *

If you liked this article, consider following me and [subscribing to email updates](__GHOST_URL__/ai-will-make-more-developers-not-less-the-economic-argument/#/portal/signup) whenever I post an article. You can also follow me on [Twitter](https://twitter.com/LewallenCaleb?ref=caleblewallen.com) or connect with me on [LinkedIn](https://www.linkedin.com/in/caleb-lewallen-b8699365/?ref=caleblewallen.com).
