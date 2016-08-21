# whoDB
It's a toy distributed data structure store inspired by Redis, wich I’m creating to go deep and understand better real world problems and solutions in distributed systems and database internals.

## Goals
I'll use this repo as a safety net to my experiments in building a distributed database, so here I'll have a lot of failed attempts and small achievements (I hope so).
This project will work as a normal open source project and you can suggest improvements, references, bug reports and even a good direction to the project using [Github issues](https://github.com/brunohenrique/whodb/issues/new).

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

I’ll use [Go](http://golang.org) in this project. First, because I want to write more software with this language and second because I think it’s a good choice for this project. The language is very simple - not easy -, has a good documentation, a very active community, others databases written in it where I can read the code and simple abstractions to concurrency (which does not eliminate previous knowledge about the subject.).
