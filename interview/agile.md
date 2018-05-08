# Agile Notes

> Collective notes on Agile and software development practice.

## Agile Software

### Agile tools comparison

| Name  | Price | Rating |Environment| Review and notes |
|:------|------:|:-------|:----------|------------------|
|[AcitveCollab](https://activecollab.com/pricing)|$25/5u,5G||Android/iOS/web||
|[Asana](https://asana.com/)|Free/15u; $10/u/Premium|3.8|Android/iOS/web||
|[Apa](https://apascrum.com/#pricing)|Free|2|web||
|[Kanban Flow](https://kanbanflow.com/board/RrAn3opg)|Free; $5/u/mo/Premium|web||
|[Kanban Tool](https://kanbantool.com/pricing)|Free/2u; $5/u/mo|3|web||
|[PivotalTracker](https://www.pivotaltracker.com/help/articles/quick_start/)|Free/3u,2p,2G; $12.5/5u,5p,5G|3.3|Android/iOS/web|https://www.pivotaltracker.com/help/|
|[Podio](https://podio.com/dockerian/project-management/apps/projects)|Free/9u|3.5|Android/iOS/web||
|[SprintGround](https://www.sprintground.com/)|Free/3u,2p,50MB; $29/mo,8u,1G|2|web|con: empty Taskboard|
|[ScrumTotal](http://www.scrumtotal.com/)|Free/15u; $7/mo/user|3|Android/iOS/web||

  Reference:
  - https://blog.capterra.com/agile-project-management-software/
  - https://www.lifewire.com/kanban-board-tools-for-project-collaboration-771630
  - https://opensource.com/business/16/3/top-project-management-tools-2016
  - https://www.pcmag.com/roundup/356732/the-best-kanban-apps
  - https://www.techworld.com/tutorial/apps-wearables/6-top-free-mobile-project-management-tools-3597574/
  - https://teamweek.com/blog/2017/11/best-project-management-mobile-apps-2018/




## Agile Methodology

### General

  * sdlc process: agile vs waterfall
    - requirements collection (including doc/draft/mock, and a.c.)
    - design / architecture
    - documentation/testing against requirements (iteration)
    - prototype / coding
    - unit testing (iteration back to coding)
    - behavioral, functional / integration test
    - documentation (optional)
    - deploy / release

  * selecting methodology based on
    - Budget
    - Team size
    - Project criticality
    - Technology used
    - Documentation
    - Training
    - Best practices/lessons learned
    - Tools and techniques
    - Existing processes
    - Software

  * lightweight (vs heavyweight)
    - accommodate change well.
    - people oriented rather than process oriented; tend to work with people rather than against them.
    - complemented by the use of dynamic checklists.
    - project teams are smaller.
    - rely on working in a team environment.
    - foster knowledge sharing.
    - feedback is almost instantaneous.
    - project manager doesn’t need to develop "heavy" project documentation but instead is able to focus on the absolute necessary documentation (e.g., the project schedule).
    - learning methodologies. After each build or iteration, the team learns to correct issues on the project, forming an improvement cycle throughout the project.
    - see
      - [techrepublic](https://www.techrepublic.com/article/heavyweight-vs-lightweight-methodologies-key-strategies-for-development/)
      - [medium](https://medium.com/pminsider/design-sprints-vs-agile-dev-sprints-eb5a11a997a8)

  * each story
    - acceptance criteria
    - design/documentation/review/test
    - check gates to close

  * code quality
    - readable
    - unit-testable
    - design pattern (interface and composition)
    - DRY, SOLID (single responsibility, open-closed, Listkov/subtyping, interface segregation, dependency inversion) principles
    - linter/pep8/styling police

### Reference: Manifesto for Agile Software Development

  * More value on the left items
    - Individuals and interactions over processes and tools
    - Working software over comprehensive documentation
    - Customer collaboration over contract negotiation
    - Responding to change over following a plan

  * Twelve principles:
    - Highest priority is to satisfy the customer thru early & continuous delivery of valuable software.
    - Welcome changing requirements, even late in development. Agile processes harness change for the customer's competitive advantage.
    - Deliver working software frequently, from a couple of weeks to a couple of months, with a preference to the shorter timescale.
    - Business people and developers must work together daily throughout the project.
    - Build projects around motivated individuals. Give them the environment and support they need, and trust them to get the job done.
    - The most efficient and effective method of  conveying information to and within a development  team is face-to-face conversation.
    - Working software is the primary measure of progress.
    - Agile processes promote sustainable development.  The sponsors, developers, and users should be able  to maintain a constant pace indefinitely.
    - Continuous attention to technical excellence  and good design enhances agility.
    - Simplicity--the art of maximizing the amount  of work not done--is essential.
    - The best architectures, requirements, and designs  emerge from self-organizing teams.
    - At regular intervals, the team reflects on how  to become more effective, then tunes and adjusts  its behavior accordingly.


## Software Development Principles

### S.O.L.I.D

  * Single responsibility principle
    i.e. changes to only one part of the software's specification should be able to affect the specification of the class.
    <br/>
  * Open/closed principle
    software entities should be **open for extension, but closed for modification**.
    <br/>
  * Liskov substitution principle
    derived types must be completely substitutable for their base types - objects in a program should be replaceable with instances of their subtypes without altering the correctness of that program.
    <br/>
  * Interface segregation principle
    many client-specific interfaces are better than one general-purpose interface.
    <br/>
  * Dependency inversion principle
    depend upon abstractions, not concretions.
    <br/>

### Java/C#/C++ (additional/alternative to SOLID)
  * Encapsulate what varies
  * Programming for Interface not implementation
  * Favor Composition over Inheritance
  * Strives for loosely coupled designs between objects that interact (Loosely coupled high cohesion)

### Python

  * DRY (Don't Repeat Yourself)
  * KISS (Keep it simple, St\*pid)


### Design Patterns

  * **Abstract Factory**, Factory Method, Builder, Prototype, **Singleton** (- creational patterns)
  * **Adapter**, Bridge, **Composite**, **Decorator**, Facade, Flyweight, Proxy (- structural patterns)
  * Chain of Responsibility, **Command**, Intrepreter, Iterator, Mediator, Memento, **Observer**, **State**, **Strategy**, Template Method, Visitor (- behavioural patterns)
