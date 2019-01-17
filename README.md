# whoDB
It's a toy distributed data structure store inspired by Redis. I'm creating it to go deep and understand better real-world problems and solutions in distributed systems, database internals, and testing strategies for complex systems.

## Goals
I'll use this repo as a safety net to my experiments in building a distributed database, so here I'll have a lot of failed attempts and small achievements (I hope so).
Few free to suggest improvements, references, bug report via [Github issues](https://github.com/brunohenrique/whodb/issues/new).

I want to explore some concepts in this project, such as:

- **Distributed**
    - Why and how to “split” my database?
    - Which ways exist to do this?
    - Why choose this approach in preference to that?

- **Persistence**
    - What problems do I want to solve?
    - Do I need to persist my data?
    - Why not put only the data in-memory?
    - If my memory fulfill in data?
    - Can I have a hybrid approach?

- **Replication**
    - Why and how do I need to replicate my data?

- **Consistency**
    - What does it even mean?
    - Is it the same from `ACID`?
    - What does exist in the `consistency's` spectrum ?
    - Why eventual consistency?

- **Availability**
    - Can my system respond correctly?
    - How can I achieve high-availability?
    - What if my database fails?

- **Partition tolerance**
    - If a network partition occur, can my database keep working?
    - Will I lose data?
    - Can I prevent it?

_These are still “incomplete” questions and I'll update this document as new ones arise._

## Engineering Principles
Even it been an experimental project, I’ll not ignore (try not) the “best practices” of software engineering. And to guide this I’ll follow three basic principles: **Operability**, **Simplicity** and **Evolvability**.

Using Martin Kleppmann, words:

#### Operability
> “Make it easy for operations teams to keep the system running smoothly.”

#### Simplicity
> “Make it easy for new engineers to understand the system, by removing as much complexity as possible from the system. (Note this is not the same as simplicity of the user interface.)”

#### Evolvability
> “Make it easy for engineers in future to make changes to the system, adapting it for unanticipated use cases as requirements change. Also known as extensibility, modifiability or plasticity.”

_Excerpts From: [Martin Kleppmann. “Designing Data-Intensive Applications.”](http://shop.oreilly.com/product/0636920032175.do)_

_In the book, he explains in more details each one, so I really recommend the book._

Initially, I planned to write it in [Go](http://golang.org/). However, I started working full time as a Go developer, and I fulfilled my interest in writing more in this language. After a while doing researches, I realized that [Rust](https://www.rust-lang.org/) could be a good fit for it. Moreover, I was curious about it, so I decided to change the language's project, and now it'll be written and Rust.
