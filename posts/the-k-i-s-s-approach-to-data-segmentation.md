---
Page Name: The K.I.S.S. Approach to Data Segmentation
Published: false
Published Date: 
Author: Caleb Lewallen
RobotsAllowed: true
---

![Feature Image](/static/images/photo-1586892477901-f70e288a7318)

Stop me if this sounds familiar. You're asked to build some new feature on top of your application data: "We want to improve customer engagement," your manager says, "so we want to send customers a follow-up email 3 days after they sign up."

No problem, I can have that ready to go by lunchtime.

```SQL
-- We're using Postgres, obviously...
SELECT
    customer_name,
    customer_email,
    signup_date,
    (CURRENT_DATE - signup_date) as days_since_signup
FROM users
WHERE (CURRENT_DATE - signup_date) = 3
```

Problem solved. Well, until that starts getting results and suddenly we're being asked to send more emails, maybe even some SMS messages here and there. Eventually, you've got a flying spaghetti monster and rules start conflicting. Customers are getting multiple emails a day and are getting annoyed.

How do we manage this mess? How can we make any of this transparent and make sure that we're not creating conflicting rules?

# Data Segmentation

Segmentation is a tried and true approach. Conceptually, all we're really doing is separating data objects into some sort of "bucket" that we'll then take an action on. In our case, we can make sure that I'm only sending a specific email to a subset of customers.

Let's design.

## Segmentation Rules

First things first, I need to decide whether segments should be exclusive. In our case, we want to limit the number of times we contact a customer, or they get (understandably) irritated with us. So we want to build in a "blocking" service within our segmentation, where a customer can only qualify for one contact event at a time.

## Rule Conflicts

Since humans are coming up with our segmentation rules, and we all know "to err is human", then we need to be able to detect if we have rule conflicts. In our case, a conflict is a situation where a customer could fall into more than one segment.

## Rule Simulations
