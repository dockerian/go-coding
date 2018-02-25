# System Design


<a name="contents"><br /></a>
## Contents

  * [Thought Process](#thought-process)
  * [Design a Cache System](#cache-system)
  * [Design a TinyURL System](#tiny-url)
  * [Design a Chat Messenger](#chat-messenger)
  * [Links](#links)


<a name="thought-process"><br /></a>
## Thought Process

  * Abstraction.
    - understand the systems you're building upon.
    - know roughly how an OS, file system, and database work.
    - know about the various levels of caching in a modern OS.
  * Concurrency.
    - threads, deadlock, and starvation
    - parallelize algorithms
    - consistency and coherence?
  * Networking.
    - IPC and TCP/IP
    - Difference between throughput and latency, and when each is the relevant factor
    - BIND and DNS lookup
  * Real-World Performance.
    - be familiar with the speed of everything your computer can do, including the relative performance of RAM, disk, SSD and your network.
  * Estimation.
    - especially in the form of a back-of-the-envelope calculation
    - narrow down the list of possible solutions to only the ones that are feasible
    - have only a few prototypes or micro-benchmarks to write
  * Availability and Reliability.
    - thinking about how things can fail, especially in a distributed environment
    - design a system to cope with network failures
    - understand durability
  * Machine Learning (optional)
  * Pros and cons


<a name="cache-system"><br /></a>
## Design a Cache System

  * Examples: DNS lookup, web server
  * LRU (least recently used) design
    - hash table for fast lookup by identifier as key
    - data in queue/stack/sorted array (can be implemented by double-linked list)
    - keep least recently used entry at the end of the queue
  * Eviction policy
    * Random Replacement (RR) - just randomly delete an entry.
    * Least frequently used (LFU)
      - keep the count of how frequent each item is requested and delete the one least frequently used.
      - problem: an item is only used frequently in the past, but LFU will still keep this item for a long while.
    * W-TinyLFU - by calculating frequency within a time window.
  * Concurrency (read/write conflict)
    - clients may compete for the same cache slot and the last one wins.
    - common solution of using a lock has its downside to affect the performance.
    - optimization is to split the cache into multiple shards and have a lock for each.
    - alternative is to use commit logs (commonly in database design):
      - store all the mutations into logs rather than update immediately.
      - some background processes execute all the logs asynchronously.
  * Scalability
    - distributed cache on multiple machines.
    - hash table maps each resource to the corresponding machine.
    - redirect request to resource machine.


<a name="tiny-url"><br /></a>
## Design a TinyURL System


<a name="chat-messenger"><br /></a>
## Design a Chat Messenger



<a name="links"><br /></a>
## Links

  - [Gainlo](http://www.gainlo.co)
  - [GeekyPrep.com](https://www.geekyprep.com/)
  - [Grokking the System Design Interview](https://www.educative.io/collection/5668639101419520/5649050225344512)
  - [High scalability](http://highscalability.com/)
  - [HiredInTech's System Design Course](https://www.hiredintech.com/system-design/)
  - [How to ace system design](https://www.palantir.com/how-to-ace-a-systems-design-interview/)
  - [Interviewbit.com](http://interviewbit.com/)
  - [Problem Solving in Data Structures & Algorithms Using Java: The Ultimate Guide to Programming Interview](https://www.amazon.com/Problem-Solving-Structures-Algorithms-Using/dp/1539724123)
  - [Refdash - Real time interview feedback from senior engineers](https://refdash.com/)
  - [Structured, Rigorous bootcamp: Large Scale Systems Design Interview Preparation Bootcamp](http://interviewkickstart.com/)
  - [System design questions and approaches](https://www.youtube.com/watch?v=0s1aVoeF0Gs)
  - [Source](http://blog.gainlo.co/index.php/category/system-design-interview-questions/)
