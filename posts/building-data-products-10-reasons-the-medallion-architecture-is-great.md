---
Page Name: Building Data Products: 10 Reasons the Medallion Architecture is Great
Published: false
Published Date: 
Author: Caleb Lewallen
RobotsAllowed: true
---

![Feature Image](/static/images/b5ae9597-29d3-44a8-acaa-de20c699941f.png)

If you spend any time in your career building data products, you will inevitably work with some data folks who want to utilize an architecture that's almost ubiquitous with the very concept of the data lake itself. As the title has already given away, this is the medallion architecture.

I, and many other data professionals, have plenty of criticisms of this approach to building data products. I'll spend some time diving into those at some point, but for now I want to talk about why this architecture just works. Despite some of its flaws, its robust and flexible enough to power almost every data lake and warehouse out there.

There are 10 reasons why I think this works so well. Some of these reasons are technical, but many are more practical in nature – why they're easier to **build**, and why they can make for a good data **product**.

# The Medallion Architecture

First, let's talk about what this architecture is. The beauty of this design is it can easily fit onto a single slide. Data comes in a raw state on the left and comes out ready to use on the right.

![](/static/images/image-13.png)

Even my kids can understand how this works

### Bronze Layer

This is just a landing zone for data coming from its source. This might come in the form of direct SQL queries, data files, and/or event streams. Anything that you need to pick up gets dumped here. There are some optional steps you can do here, like converting everything into a standard data format or doing some light organizing (split flat files into logical structures, merging data sets, etc).

### Silver Layer

This layer is where we want to unify our sources into a common ontology. Really spend some time here to get this right, you'll save yourself so much time and energy later. Because you're pulling data from multiple source systems, data will be structured and named differently for each source. You will typically normalize your data into a single ontology here.

### Gold Layer

The last step is kind of like taking a step backwards. We just spent a lot of time working up into a highly normalized structure, and here we're going to de-normalize it all over again. It's for a good reason though, many of the data sets we produce will be for a specific use case. It could be de-normalizing things into a flat table for query performance, or creating a very specific data set that requires a lot of complex business logic to produce.

# 10 Reasons This Architecture Rocks

### 01 - It's easy to understand

Seriously, this is so simple it can fit on a single slide. Don't underestimate the benefit of being able to communicate your data strategy effectively, especially to folks who don't have a background in data. This includes both your technical and non-technical stakeholders. Business folks don't really care, and most technical folks think in terms of "stateless" functions – that is they think about logical structures first, and then try to backfill the data later.

### 02 - It's a one-way pipeline

This moves all of your data in the same direction and doesn't branch off. Every dataset you work with comes along for the ride. This is really useful for coordinating efforts amongst multiple teams. There's a checklist of steps data has to go through, so you don't need to handhold everyone along the way.

### 03 - It enforces incremental changes

Since you can only move through one step at a time, each change to the data is incremental. This prevents you pushing too much complexity into any individual step. One step is creating common data formats, another step is performing quality checks, and another is creating logical schemas. It's easy to introduce defects you can't easily debug if you tried to do it all at once.

### 04 - Its flexibility gives cheaper storage options

Storage is cheap, but when you're dealing the volumes of data that made you think building a data lake was a good idea, then we're dealing with real dollars. Given the separate steps, you can use cheaper object storage for lower layers and more purpose-built database options for higher layers.

### 05 - It gives multiple data assets for different users

Your silver and gold layers are all accessible data assets. Data scientists might be interested in using your sanitized data from source to run direct queries, while business users might be interested in higher-level data sets that are purpose-built for them.

### 06 - It's simple to govern

I know, I know, data governance is rarely "simple". However, this architecture makes it straightforward to catalog, track lineage, and understand the ownership of different attributes in different layers of the stack.

### 07 - It's resilient against source changes

Since the bronze layer preserves the original data (often in a schema-less architecture), allowing you to tweak downstream layers while maintaining the original data, releasing the source data once the changes have been made.

### 08 - It's horizontally scalable

The system isn't tied to any one particular monolithic process (although it can be, if that's what works for you), so you can scale different layers depending on how much data you move from one layer to the next.

### 09 - It gives flexibility with compute

Okay, this might be another take on #8, but it's worth calling out here. Some workflows, especially as you move into our silver and gold layers, might not need massively concurrent, horizontally scalable systems like Spark. Because each layer is separate, you can utilize something like a simple Python/Pandas/Polars script for smaller datasets, saving Spark for bigger workloads

### 10 - It's easier to coordinate across multiple teams

You can assign different layers of work to different groups. You can separate work out by layer, by transformer, or by whatever differentiating factor you prefer. There's enough commonality that different teams are able to coordinate, but also enough differentiation between different types of workloads that different specialties can be assigned.

# Conclusion

The medallion architecture is a time-tested architecture adopted across many organizations. There's a lot of resources out there to help you build your own data platform using this approach.

Next, I'll take some time to talk about some of the shortcomings of this approach and how to solve for them.
