---
Page Name: Composable Data Platforms are the Future
Published: true
Published Date: 2025-02-22
Author: Caleb Lewallen
RobotsAllowed: true
---

![Feature Image](/static/images/composable-data.png)

*This post is inspired by* [*this amazing article*](https://jack-vanlightly.com/blog/2025/2/17/towards-composable-data-platforms) *by Jack Vanlightly. I've taken a lot of his ideas and shaped them around some of the enterprise workflows I've been a part of recently, and tweaked them in a way that I think works for a couple of real-world use cases.*

Anyone who's worked with me over the last year or so knows I'm generally not a fan of the microservices pattern. Don't get me wrong, sometimes it's a necessary pattern, still doesn't mean I have to like it 🤗. However, I think I've stumbled across something that might make them tolerable.

While most of the tech news you see is around AI and who's the latest tech giant to spend a couple billion to train a marginally better LLM, there's been a quiet revolution in the data space that is undoubtedly more impactful to how people actually build software. Especially systems at an enterprise scale. This "new" development is in the form of Open Table Formats (OTF).

OTF's aren't new, but are emerging as a more popular standard in recent years for data stores as the need for managing larger datasets begins to exceed traditional database paradigms. I want to spend a little time exploring how these new formats can be used to make a microservices architecture tolerable.

# Databases & Microservices

First, let's very quickly define microservices for the people in the back (not you, of course). Microservice architectures are an approach for working on large scale systems that breaks different areas of functionality into independent services that can be developed and deployed without affecting other services. Proponents will also tout some (IMHO made up) benefits of improved reliability – in practice I think this is rarely the case.

Each service is effectively its own independent software application, complete with its own data store (read: data silo). These services typically perform a specific business function. You don't have to be really creative to guess what these might look like. A payment processing service, a communications service, a user management service, etc. The payments service holds the data it needs to process payments, communications holds the data it needs to send emails/texts/whatever, etc.

![](/static/images/image-10.png)

Each service must maintain its own data store

The databases you would use to make all of this work are the usual suspects: Postgres, SQL Server, Oracle, MongoDB, DynamoDB, etc. What all of these have in common is that they can only access data within their respective stores. In other words, the compute is **coupled** to the data storage. While the payments service only contains data it needs to process payments, it helps to know the user who made the payment. To do this, it needs user information. If we want to send an email on a successful payment, then the communications service needs both customer information and payment information. In order to do this, data must be duplicated and synchronized across services.

Think about this like writing in a secret cipher that only you can read. If someone else needs that information to be able to use it later, then you have to read it to them first so they can write it down in their own secret cipher.

Sometimes, that's a helpful pattern. In order to get the data for this blog post, you needed to ask the database engine powering this blog to go fetch it for you – you couldn't just go read the database directly. However, in a microservices architecture, we want all of the internal services to be able to manage and read the data like it's all one system. This is a **technical barrier** that we usually overcome by utilizing some sort of messaging service. This might be API requests, web hooks, or more often than not, messaging queues. Whenever an action is taken, that message is published to a queue and consumed by all the services who are subscribed to the event.

![](/static/images/image-11.png)

So now I have three copies of my user data and multiple copies of payment information. What if the user service or payment service needs to know if an email was sent? In practice, a lot of data ends up being duplicated over and over again. And keeping all of this data in sync is a fool's errand. You need to periodically resync all of your data stores just in case something fails and events aren't pick up. This is the tradeoff against having to maintain and deploy changes against a monolithic application.

What would make this better, is if each service could just read the data it needs from everyone's data stores when its needed – too bad traditional databases just don't allow for that. What if I could just have a single copy of the data and read from it wherever I need to?

# Open Table Formats

Enter the Open Table Format (OTF). The major benefit that we care about for our use case is that an OTF allows us to separate the physical data that we want to read from the compute needed to read it. This separation of concerns is called **virtualization**. This virtualization allows different resources to be shared more efficiently.

The OTF does two key things:

1. It separates **data** from its **metadata**, and
2. separates **storage** from **compute**.

Table metadata is information about the table itself – where the data is, how it's organized by table and column names, column indexes, and the like. It's the table of contents for a database engine. By separating this away from the raw data beneath it, I'm able to send this metadata to distributed systems without creating multiple copies of the data itself.

This creates a number of benefits:

- **Single source of truth.** One common challenge in copying data is synchronization. Errors happen all the time – drives get corrupted, connections are interrupted, edge cases cause data to not copy as expected, etc. Ask anyone who's worked with Kafka events how flawless it can be. A single source of data means it's easier to prevent replication errors from happening and I'm able to operate from the same source of truth.
- **Utilizing Object Storage.** Believe it or not, the physical data in your database is just a file (often multiple files, but a file nonetheless). If we separate this from the compute and metadata, I can just stick the files in cheap object storage and not have to worry about complex storage management.
- **Federated access.** Data catalogs are pure metadata. They can easily be distributed across a federated system to grant access. The owner of the data can easily present the data to another system without copying data over.
- **Standardization.** Having a common format means that data can be collaborated on by multiple parties. Since the physical data is in a common format, multiple engines can process on the same data for different purposes.

Thinking about these virtualization benefits yields some powerful implications. Being able to surface the *same table* in *different platforms* as if they were both native to the same platform is insane. With this capability, there's no need to move data between platforms, you just need to expose the tables needed to feed other platforms. Data access to peer services is truly *real time*.

![](/static/images/image-8-1771789986.png)

# Data Governance & Access

Okay, so two systems having access to the same table sounds cool in concept, but what are the practical implications? I can already see someone raising their hand to tell me two systems shouldn't be able to access the same table. The chaos! It's anarchy!

Remember, these are just tools. How you use them is the key.

Tables come in two flavors. The first is a **native** table. This is exactly how you think about tables now. It's both the table metadata as well as the physical data itself. The second flavor is an **external** table. This is just the table metadata. It gives us the information needed to read the table without a copy of the data itself.

We expose this data to multiple systems by granting access to native and external tables. One system will have access to the native table, while one or more other systems can have access to the external table.

## Data Sharing

As a service or application, data sharing is two-way. You can share out, meaning you own the native table and expose table metadata to other applications. You can also share in, where you add external tables from other applications. In other words, the native table owner publishes metadata changes that the external table owner must subscribe to.

Once you have the metadata, then it becomes a matter of permissions for your typical CRUD operations. At this moment in time, I haven't been able to come up with a use case where I would allow an external table owner being able to do anything to the physical data except read it. Allowing external platforms to write to your table creates a host of support, data quality, and data reliability issues.

The use cases I've experienced so far revolve around a common pattern. There is a System of Record owner that publishes data changes that other systems need to react to. Right now, that's done via a batch file import or some event queue. In our composable data pattern, the SoR would write to a native table that other systems can read using an external table.

![](/static/images/image-5.png)

Data that needs to be written back to the SoR should be handled in one of two ways. The first is to utilize typical messaging interfaces like API calls or message queues to let the SoR know it needs to write the data back to a native table. The second is that if we think the outside service should be writing that data, then we should evaluate whether that data truly belongs to the SoR.

![](/static/images/image-6.png)

Do you see where this is going? Take an example where our customer platform from earlier needs to send a text message to a user, for which they use a communications microservice. If the text message is unable to reach the customer, then the bounce back data needs to be written and recorded back to the user. In our composable model, since the customer application isn't the creator of the bounce back data, the communications service can own the native table with the bounce back data, giving the SoR direct access via an external table.

# Operational Data and Analytics

So far, we've really only looked at this from a SoR/transactional lens. What about as we get into data-specific uses like Operational and Analytical platforms?

## Operational Data

A common data sharing primitive in the operational space is one of data streaming – directing messages to different microservices to take different actions, as well as consuming response messages from the microservices. This pattern can be conceptualized for our composable data platform in a producer/consumer pattern like so:

![](/static/images/image-12.png)

Using virtual tables to produce and consume event data

This approach differs slightly from the traditional approach to event-based microservices architectures in that the microservice isn't necessarily producing events, instead centralizing that responsibility to the operational data platform. This approach allows services within a given domain to simply read/write to database tables without needing to implement messaging protocols. Any events that need to be broadcast back to the enterprise is delegated to an operational platform.

## Analytical Data

I won't pretend to be an expert in the analytics space, I haven't spent nearly as much time in that area as operational data. However, given I know just enough to be confident in my ignorance, I'll let you tell me where this falls short.

The role of the analytical data function is in aggregations of long-term business datasets, pre-calculating expensive queries into something that can be easily consumable by end users. Usually this begins in the form of copying large quantities of raw data away from SoR systems into a place where we can clean and process data.

![](/static/images/image-2-1771789986.png)

The foundation of analytics is copying large amounts of data

Like before, one of the most expensive parts of this operation is copying large amounts of data to a separate store before we start to clean and conform the data into something useful. This results in a mountain of data that is really nothing more than a 1:1 copy of the original data. We haven't added any value yet and cost a ton in additional storage costs. Virtual tables allow the analytical space to incrementally read the raw data directly where it sits and store refined/conformed data in native tables.

## Reverse ETL

There's a fuzzy line between analytical and operational data. In theory we like to consider these things like separate ideas, but in practice they often overlap. Like a report created from SoR data that contains a new derived dataset, that's consumed by the original SoR to perform new workflows.

![](/static/images/image-7.png)

What's the common theme we find ourselves with once again? The physical copying of data. Two copies of the same data that are needed by different applications in the same system. More replication, more points of failure.

## Data Fabric

What we want to start shifting towards is a data "fabric", where we begin to think about tables in a logical ownership model rather than a physical one. With a data fabric as the model for our composable data platform, a native table can be exposed to any other system that requires access, so that the same set of data can be read and used by any system that needs it.

![](/static/images/image-4.png)

In our old model, our data access is restricted by physical and technological barriers. All operations have to be performed by the same application. In our new model, the physical location of the table is no longer important (save for latency issues with edge compute, but that's a different discussion). In theory, you could have the physical data for dozens of services all in the same object storage layer, with ownership distributed in a logical model.

To prevent this from becoming too chaotic and maintaining good data governance, we still want to interface with other domain areas of the enterprise through some sort of central data platform. An example for a financial services company might be separating domains around loan origination and servicing. The origination domain largely focuses on marketing, sales, and underwriting functions. Servicing shifts into daily operations. Each one of those business domains likely operates independently from the other, with a customer handoff between functions. In this case, I wouldn't want to mix domain data. I'd still consider an enterprise data interface to shift relevant data attributes between domain spaces. Once inside the domain, we can move back into our composable data structures. Your mileage may vary here, it will be up to you to determine how wide you want your data fabric to stretch.

# A New Ownership Model

With this way of looking at ownership and governance, our new focus is on how to present the data. It becomes less and less about how data should move from A to B, scheduling batch updates, what events should be published, etc. We no longer have to concern ourselves with different platforms with different capabilities and needs. There's no more siloed application data within a business domain – instead we define data ownership using a logical model. Our applications start behaving like a hive instead of as separate entities, and we can focus more on the logical constructs of our applications instead of how we're going to communicate between them.

Those of you that are starting to grasp the vision should see why this is such an appealing idea. I don't see any real barriers to being able to run complex systems this way. The most difficult thing to map out is who is the owner of a specific native table (i.e. who is the producer and sole authorized modifier of the data) and ensuring the external table is exposed to the proper users.

# Fitting Into Your Data Strategy

To be clear, this isn't a simple way to manage data for singular platforms. For startups and those with monolithic applications, my take is still to "just use Postgres" and be done with it. However, if you're in a large organization and need to federate data access across dozens of services, then a composable data platform approach using OTF's can radically simplify your tech stack and the way you operate.

To me, this type of data strategy fits into the "just use Postgres" mantra, but designed for the enterprise. We want to skip out on setting up and using messaging queues for most services and just query data. We want to allow our operational and analytical data platforms to read the data they need without having to invest in massive ETL compute and storage costs to create 1:1 copies of the same data. Services should read the data where it sits and move on with what they need to do.

* * *

If you liked this article, leave a comment and let me know! You can also follow me on [Twitter](https://twitter.com/LewallenCaleb?ref=caleblewallen.com) or connect with me on [LinkedIn](https://www.linkedin.com/in/caleb-lewallen-b8699365/?ref=caleblewallen.com).
