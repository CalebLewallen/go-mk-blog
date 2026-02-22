---
Page Name: Systems of Record vs Reference
Published: false
Published Date: 
Author: Caleb Lewallen
RobotsAllowed: true
---

One thing I'm not sure I completely appreciated until working for an S&P500 company was the complexity and interconnectivity of systems at scale. This interplay of systems introduces the need for a new data paradigm that I think is useful in smaller scale operations, albeit applied differently.

This concept is understanding the different between a System of Record and a System of Reference. If you're in charge of managing the systems or data for your org, having a concept of what your systems of record vs reference are will make your job a lot easier.

# Systems of Record

This is what most people think of when they think about their systems and how their business processes work. My customer data is in Hubspot. My accounting data is in Quickbooks. My support tickets are in Zendesk.

A System of Record is really just a *application specific data store*. They're always tied to a specific system, and it's the data generated and stored by that specific system. If you enter all of your accounting data into Quickbooks, then that is the system of record for your accounting data.

# System of Reference

This gets a little more nuanced. A system of record *can be* a system of reference, but not the othe other way around. A system of reference is a domain specific system that enables a particular business context to operate.

Think about a sales department. They will have a CRM application, an application to generate and send invoices, maybe an applicaiton for generating contracts, and many other applications. However that data isn't always accessed (or *referenced*, get it?)
