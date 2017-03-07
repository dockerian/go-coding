# Interview Questions

<a name="contents"><br /></a>
## Contents

- [Books](#books)
- [Resources](#resources) | My [Archives](#archives)
- [Interview tips](#tips)
- [General questions](#general)
- [Data Structure](#data-structure)
- [Design and scrum](#design-and-scrum)
- [C#](#csharp)
- [Go (Golang)](#golang) | [Notes](../demo/golang-notes.md)
- [Java](#java)
- [JavaScript](#javascript)
- [Database](#database)
- [Client/Server and Network](#client-server-network)
- [Math](#math)
- [Linked List](#linked-list)
- [Map](#map)
- [String](#string)
- [Tree](#tree)


<a name="books"><br /></a>
## Books

- Aziz, [Elements of Programming Interviews](https://www.amazon.com/Elements-Programming-Interviews-Insiders-Guide/dp/1479274836): 300 Questions and Solutions by Aziz, Adnan, Prakash, Amit, Lee, Tsung-Hsien 1st (first) Edition (10/11/2012), 2012, 481 pages, 6 x 9, $25, 300 problems (mostly C++, concurrency in Java, discrete math in formulas and English)
- Aziz, [Elements of Programming Interviews in Java: The Insiders' Guide](https://www.amazon.com/Elements-Programming-Interviews-Java-Insiders/dp/1517435803/): 2nd Edition (9/19/2015), 542 pages · ISBN 1517435803 $27, 250 Problems and solutions in Java
- McDowell, [Cracking the Coding Interview](https://www.amazon.com/Cracking-Coding-Interview-Programming-Questions/dp/098478280X): 150 Programming Questions and Solutions, 2011 (5th edition), 500 pages, 6 x 9, $23, 150 problems, (mostly all Java except of course the C, C++ question sections!)
- Guiness, [Ace the Programming Interview](https://www.amazon.com/Ace-Programming-Interview-Questions-Answers/dp/111851856X): 160 Questions and Answers for Success, 2013, 419 pages, 6 x 9, $20, 160 problems, (mostly Java and C# but some unusual JavaScript, SQL, Ruby and Perl examples too)
- Mongan, [Programming Interviews Exposed](https://www.amazon.com/Programming-Interviews-Exposed-Secrets-Landing/dp/1118261364): Secrets to Landing Your Next Job, 2013 (ed. 3), 301 pages, 7.4 x 9, $18, 150+ problems (C, C++, C#, Java)


<a name="resources"><br /></a>
## Online Resources

- [Big-O Cheatsheet](http://bigocheatsheet.com) (review data structure time complexity)
- [CareerCup](http://careercup.com) (actual interview questions)
- [CodeChef](http://codechef.com) (judge code by other engineers)
- [CodePad](https://codepad.remoteinterview.io) (remote interview supports golang)
- [CodeShare](https://codeshare.io/) (sharing code in real time with others)
- [CollabEdit](http://collabedit.com/) (online coding interview, free version of https://codinghire.com)
- [EdRepublic](edrepublic.com) (problems by companies and focus area)
- [Educative](https://www.educative.io) (online lessons, and free courses)
- [Gainlo](http://www.gainlo.co/#!/faq) (mock interviews with professionals)
- [Geeks for geeks](http://www.geeksforgeeks.org/) (a computer science portal for geeks)
- [Glassdoor](http://www.glassdoor.com/) (company rates/reviews, salary comparisons, and interviews)
- [HackerEarth](https://www.hackerearth.com/@jason_zhuyx) (programming challenges & coding competitions)
- [HackerRank](http://hackerrank.com) (rank programmer on coding skill)
- [Hiredintech](http://hiredintech.com) (tips/tricks on algorithm and systems design)
- [LeetCode](http://leetcode.com) (interview code online judge)
- [Project Euler](https://projecteuler.net/archives) (a series of computational problems)
- [SPOJ](http://www.spoj.com/tutorials/) (problemset archive, online judge and contest)
- [Topcoder](https://www.topcoder.com/members/jason_zhuyx/) (online computer programming competitions)
- [Case study](http://www.careerprofiles.info/case-study-interview-examples.html)
- [Courses](https://www.coursera.org/courses?categories=cs-theory&languages=en)
  - [Algorithms by Pinceton](https://www.coursera.org/learn/algorithms-part1)
  - [Algorithms: Design and Analysis, Part 1 (Stanford)](https://www.coursera.org/course/algo)
  - [Algorithms: Design and Analysis, Part 2 (Stanford)](https://www.coursera.org/course/algo2)
  - [Data Science Tutorials](https://www.topcoder.com/community/data-science/data-science-tutorials/)


<a name="archives"><br /></a>
## Archives
  - coding [notes](coding.md)
  - hard/bigger [problems](problems.md)
  - puzzle/quiz/teaser [questions](puzzles.md)
  - some [examples](examples.md) and [answers](examples-qa.md) | [c# qa](examples-cs.md)
  - go [quiz/puzzle solutions](../puzzle)


<a name="tips"><br /></a>
## Interview Tips

- Start writing on the board.
- Loud thinking, talk it through.
- Think data structures.
  - Array
  - Hashset / Hashmap / Hashtable / Dictionary
  - Heap
  - Stack / Queue
  - Tree / binary tree
  - Graph
- Think algorithms.
  - Sorting (plus searching / binary search)
  - Divide-and-conquer
  - Dynamic programming / memoization
  - Greediness
  - Recursion
- Think about similar problems and solutions
  - dynamic programming
  - divide and conquer
  - composition
- Breaking down to smaller problems
- See blogs
  - [here](https://www.palantir.com/2011/09/how-to-ace-an-algorithms-interview/)
  - [ace the coding interview](https://www.linkedin.com/pulse/20141120061048-6976444-ace-the-coding-interview-every-time)
  - [Facebook interview](https://www.facebook.com/notes/facebook-engineering/get-that-job-at-facebook/10150964382448920)
  - [Google interview](http://steve-yegge.blogspot.com/2008/03/get-that-job-at-google.html)
  - [Big-O](http://bigocheatsheet.com/)

  | Data Structure | Read (Avg/Worst) | Write       | Space    |
  | -------------- |:----------------:|:-----------:|:--------:|
  | Array          | 1 / n            | n           | n        |
  | Stack / Queue  | n                | 1           | n        |
  | Linked List    | n                | 1           | n        |
  | Skip List      | log(n) / n       | log(n) / n  | n log(n) |
  | Hash Table     | 1 / n            | 1 / n       | n        |
  | BST / Tree     | log(n) / n       | log(n) / n  | n        |

  | Algorithm      | Best     | Average  | Worst    | Space    |
  | -------------- |:--------:|:--------:|:--------:|:--------:|
  | Bubble-sort    | n        | n^2      | n^2      | 1        |
  | Bucket-sort    | n+k      | n+k      | n^2      | n        |
  | Cube-sort      | n        | n log(n) | n log(n) | n        |
  | Heap-sort      | n log(n) | n log(n) | n log(n) | 1        |
  | Insertion-sort | n        | n^2      | n^2      | 1        |
  | Merge-sort     | n log(n) | n log(n) | n log(n) | n        |
  | Quick-sort     | n log(n) | n log(n) | n^2      | log(n)   |
  | Radix-sort     | n\*k     | n\*k     | n\*k     | n+k      |
  | Selection-sort | n^2      | n^2      | n^2      | 1        |
  | Shell-sort     | n log(n) |n log(n)^2|n log(n)^2| 1        |
  | Tree-sort      | n log(n) | n log(n) | n^2      | n        |


<a name="general"><br /></a>
## General

- data team: DSA and data science questions; algorithms and data structures questions with attention to complexity and possible optimizations.
- describe the principles of object-oriented programming
- design an API for an elevator.
- design a cache
- design question for some hypothetical Task object. write the algorithm and design all the classes required to complete all the tasks, subtasks, and their dependent tasks.
- describe a particularly difficult concurrency problem you have faced and you solved it?
- difference between cloud (computing) and virtualization. (virtualization is a virtualized simulation of a device or resource, e.g. storage, network, memory; cloud is shared computing resources, software, or data are delivered as a service and on-demand through the Internet.)
- explain server caching/sessions
- given a log of pages clicked on website by users (sorted by timestamp), find the top ten most clicked 3 page sequence. A 3 page sequence, is a sequence of 3 pages clicked by the same user in successive order.
- how does garbage collection work in .NET?
- implement a LRU cache
- implement an infinite/large sized Tic Tac Toe game? how to check for win conditions?
- optimize a memory situation, e.g. millions lines of data needs to be read into a server.
- parse json and store the result in a csv file
- string and array manipulations, minesweeper game, LRU Cache, Blackjack game, trees
- what are the four pillars of Object Oriented Programming?
- what is different between null and undefined ?
- what is JSON ? how to read and write ? (key-value pair; use stream/text reader/writer to deserialize/serialize)
- write an elevator controller


<a name="data-structure"><br /></a>
## Data Structure

- big-O questions, hash map vs binary tree, heap data structure.
- compare min-heap, max-heap, and priority queue
- describe a hash table in as much detail as possible
- describe different types of sorting methods, eval time complexity from best to worst.
- design and implement a hash map


<a name="design-and-scrum"><br /></a>
## Design and Scrum

- design:
  - a card game, with Card, Hand, Deck, and interfaces
  - a chess game, or a borad game
  - a database model for movie ticketing system
  - a parking lot, e.g. with FindBestSpot()
  - a scalable twitter feed filter system which builds a public sentiment every minute.
  - an elevator system
- design patterns
  - creational patterns
    - abstract factory
    - builder
    - factory method
    - prototype
  - behaviorial patterns
    - command
    - strategy
    - state
  - structural patterns
    - adapter
    - bridge
    - decorator
    - facade
  - concurrency patterns
- explain SOLID, KISS, and DRY (principles of software development)
  * DRY=Don't Repeat Yourself;
  * KISS=Keep It Simple, S*****;
  * SOLID
    - **Single responsibility**: only one potential change to affect the spec;
    - **Open-closed**: a well-encapsulated and highly-cohesive system open for extension but closed for modification;
    - **Liskov substitution**: every subclass/derived class should be substitutable for their base/parent class
    - **Interface segregation**: many client-specific interfaces are better than one general-purpose interface;
    - **Dependency inversion**: entities must depend on abstractions not on concretions;
    - see http://www.codemag.com/article/1001061
- dependency inversion (IoC)
  - a special form of decoupling
  - high-level modules should not depend on low-level modules. Both should depend on abstractions.
  - abstractions should not depend upon details. Details should depend upon abstractions.
  - implementations:
    - factory pattern
    - service locator pattern
    - dependency injection:
      - interface (type 1)
      - constructor (type 3)
      - setter (type 3)
- compare mvp vs mvc vs mvvm
  - MVP: view <=> presenter <=> model (inputs begin with view)
  - MVC: views <= controller => model \[=> view] (starts with controller)
  - MVVM: view <= viewmodel <=> model (inputs begin with view)

- compare mutex vs semaphore

- describe properties of RESTful
  - Representational State Transfer (REST) is an architectural style
  - **Uniform interface**: a fixed set of CRUD (create, read, update, delete) operations: POST, GET, PUT, and DELETE
  - **Resource identification**: data and functionality are considered resources and accessible through URI
  - **Self-descriptive messages**: Resources are decoupled from representation so that the content can be accessed in a variety of formats, e.g. HTML/text, XML, JSON, JPEG, PDF
  - **Stateless to stateful interactions**: use stateless communication protocol, e.g. HTTP, (with stateless resource and self-contained request messages) to transfer states (e.g. embedded in response message)

- explain Responsive Web Design (RWD). how it comares to adaptive design ?
  - a design approach which prioritizes on giving the user
    an optimal viewing/reading and navigational experience
    across multiple devices and screen resolutions
    by utilizing many design concepts.
  - the goal is to have one content base with multiple 'disconnected' views.
    The word responsive signifies that your content responds to
    the user's current view (i.e. resolution, capabilities etc.) and
    the design process is all about optimizing and creating views.
  - Bootstrap is the most popular CSS, HTML and JS framework
    used for developing responsive web design
  - see http://www.alistapart.com/articles/responsive-web-design/
  - adaptive web design essentially utilizes many of the components of
    progressive enhancement (PE) as a way to define the set of
    design methods that focus on the user and not the browser.
    Using a predefined set of layout sizes based on device screen size
    along with CSS and JavaScript, the AWD approach adapts to the detected device.
    The three layers of Progressive Enhancement:
      - Content layer = rich semantic HTML markup
      - Presentation layer = CSS and styling
      - Client-side scripting layer = JavaScript or jQuery behaviors

- discuss [NP-complete problems](https://en.wikipedia.org/wiki/List_of_NP-complete_problems)

- how to balance conflicting, urgent priorities from different teams ?


<a name="csharp"><br /></a>
## C&#35;

- books
  - [C# in Depth](http://csharpindepth.com/Articles.aspx)
  - [Threading in C#](http://www.albahari.com/threading/)
  - [more ...](http://codecall.net/2014/05/26/16-free-ebooks-for-c-learner/)
- boxing vs unboxing
- class vs struct (reference vs value type, on heap vs stack, inheritance vs no)
- compare `abstract`, `virtual`, and concrete methods; `override` vs overloading
- compare `const` vs `readonly` vs `sealed`
- compare `ref` vs `out` parameters
- compare `Dispose()` and `Finalize()` (`protected` for GC only, unmanaged resource vs `IDesposable`)
- compare `System.Array.CopyTo()` and `System.Array.Clone()`: which shallow?
- compare `System.String` vs `System.Text.StringBuilder`
- compare default value of `String` vs `DateTime` (class vs value type)
- compare .Any and .Where in LINQ extension methods
- explain the purpose of LINQ (Language Integrated Query) and Enterprise Framework
- LINQ and lambda

  ```csharp
  // public delegate TResult Func<T1, T2, TResult>(T1 a1, T2 a2);
  Func<int, int, bool> compare = (a, b) => a > b
  // bool compare(int a, int b) { return a > b; }
  u = users.Where(u => u.Age > 20).First();  
  ```

- how to use `using` statement ?
- how to configure a service or web site for SSL/TLS or https? (using ABC, address-binding-contract, configuration in web.config)
- null-conditional operator (for `nullable`)

  ```csharp
  int? v = objInstance?.Property // null if objInstance is null
  ```

- what is `delegate`, `Delegate`, `MulticastDelege`, handler, event ?

  ```
  public delegate FooTypeDelegateOrHandler Func<int, bool>;
  public event FooTypeDelegateOrHandler FooEvent; // class property
  ```

- what is good and bad about `try catch` ?
- what is nullable type ?


<a name="golang"><br /></a>
## Go (Golang)

- advantages and disadvantages of golang ?
  - simpler (as interpreted lang), strongly typed, fast (compiling), portable
  - no generics (yet)
- compare type conversion vs type assertion
- compare `v.(type)` vs `v.(string)` (`v.(SomeTypeName)`), where `v` is `interface{}`
- use [reflection](https://blog.golang.org/laws-of-reflection)


<a name="java"><br /></a>
## Java

- Does Java support multiple inheritance?
- Given 2 arrays of integers, how do you find out their intersection in O(n) efficiency?
- How does Java implement polymorphism?
- How many bits do you need to hold a number that is between 0 and 100,000?
- Is Java pass by reference or pass by value?
- What are the 2 kind of exceptions in java?
- What does it mean to have memory leak in Java? How do you detect and resolve memory leaks?
- What is multi-threading? How does it work? Explain the "synchronized" keyword. Think of a situation where multiple threads would cause a “deadlock”?
- What is polymorphism?
- What is the difference between == and "equals"?
- What's a hash table? how do you implement it?
- What's a singleton? How do you implement one?
- What's the difference between interface and abstract class?
- What's the pros and cons between SQL and NoSQL?
- What's your experience with multi-threading?
- What's hashcode, hastable, equals methods, checked exception, volatile, finalize, GC etc ?
- Why string is immutable?



<a name="javascript"><br /></a>
## JavaScript

- async/promise (Node.js)
- explain about the various modules you used in the Node.js app
- explain box model (html/css)
- explain closure
- how prototype works, comparing to e.g. inheritance in C++/C#/Java
- javascript method overloading - see [here](http://ejohn.org/blog/javascript-method-overloading/)
- LESS/Bootstrap


<a name="database"><br /></a>
## Database

- collation is defined to specify the sort order in a table. 3 types (case in/sensitive, binary)
- common table expression (CTE), as a temp query/view

  ```sql
  WITH foo_cte AS (
    SELECT a, b FROM [foo_table]
  ) --, bar_cte AS (
    --  SELECT c, d FROM [bar_table]
  ),
  SELECT * FROM [foo_cte] -- JOIN [bar_cte] ON ...
  ```

- find duplicate rows

  ```sql
  SELECT col_name, COUNT(*) count FROM table GROUP BY col_name HAVING count > 1
  SELECT DISTINCT col_name, count(col_name) FROM table GROUP BY col_name
  ```

- delete duplicate rows

  ```sql
  WITH cte_rows AS (
      SELECT ROW_NUMBER()
      OVER (PARTITION BY Col1, Col2 ORDER BY (SELECT 0)) RN
      FROM #MyTable)
  DELETE FROM cte_rows
  WHERE  RN > 1;
  ```
  or on MySQL (without `CTE` or `PARTITION BY` clause)

  ```sql
  -- to keep highest id
  DELETE t1 FROM table t1, table t2
   WHERE t1.id < t2.id AND t1.col_name == t2.col_name
  ```
  or

  ```sql
  -- to keep lowest id
  DELETE t1 FROM table t1, table t2
   WHERE t1.id < t2.id AND t1.col_name == t2.col_name
  ```
  or

  ```sql
  CREATE TABLE tmp_table LIKE src_table;
  INSERT INTO tmp_table(id) SELECT MAX(id) as id
    FROM src_table
    GROUP BY field1, field2
    HAVING COUNT(*) > 1
  DELETE FROM src_table
    WHERE id IN (SELECT id FROM tmp_table)
  DROP TABLE tmp_table
  ```
  or on other system allows deleting from the same table

  ```sql
  DELETE FROM src_table src
    WHERE src.id IN (
      SELECT MAX(dup.id)
      FROM src_table as dup
      GROUP BY dup.field1, dup.field2
      HAVING COUNT(*) > 1 )
  ```

- difference between clustered and a non-clustered index? (clustered index reorder the row as physically stored, the leaf nodes contain the data pages)
- difference between primary key and unique key? (PR constraint is a unique identifier for each row, it creates clustered index and does not allow NULL)
- difference between `DELETE` and `TRUNCATE` commands?
- difference between `HAVING` clause and `WHERE` clause? (`HAVING` used only with the GROUP BY)
- explain ACID: 4 properties to qualify a transaction (which is a sequence of operations performed as a single logical unit of work)
  - Atomicity: as atomic unit of work, either all or none of modifications are perform.
  - Consistency: all data in a consistent state; all rules must be applied to maintain all data integrity.
  - Isolation: modifications must be isolated from modifications made by any other concurrent transaction.
  - Durability: modifications persist permanently in the system, even in the event of system failure.
  - Lock modes: Intent shared (IS), Intent exclusive (IX), Shared with intent exclusive (SIX), Intent update (IU), Shared intent update (SIU), Update intent exclusive (UIX)
  - see https://technet.microsoft.com/en-us/library/jj856598(v=sql.110).aspx

  | Isolation level  | Dirty read | Nonrepeatable read | Phantom |
  | ---------------- |:----------:|:------------------:|:-------:|
  | Read uncommitted | Yes        | Yes                | Yes     |
  | Read committed   | No         | Yes                | Yes     |
  | Repeatable read  | No         | No                 | Yes     |
  | Snapshot         | No         | No                 | No      |
  | Serializable     | No         | No                 | No      |

- explain inner vs outer joins
- explain normalization vs. denomarlization
- explain NOLOCK hint, when to or not to use. (READ UNCOMMITTED/dirty data by engine faster, applicable to SELECT only)
- how to check locks in database ? (`sp_lock`)
- how to get a list of triggers ? types of triggers ? (`INSERT`, `DELETE`, `UPDATE`, `INSTEAD OF`)

  ```
  select * from sys.objects where type=’tr’
  ```

- use nested SQL select statements to determine a value in table A from information in table B
- use SQL Server `ROW_NUMBER()`

  ```
  SELECT ROW_NUMBER() OVER(PARTITION BY city ORDER BY age) AS rank, city, age
  FROM People
  ```
- use COALESCE to return first non-null expression within the arguments.
- use methods to protect against SQL injection attack:
  - Use Parameters for Stored Procedures
  - Use Parameter collection with Dynamic SQL
  - In like clause, user escape characters
  - Filtering input parameters


<a name="client-server-network"><br /></a>
## Client/Server and Network

- describe the OSI layers and give an example of each
- describe internet protocol suite
  - application (dhcp, ftp, http, imap, socks, ssh),
  - transport (tcp, ucp, rip),
  - internet (ip, icmp, ipsec),
  - link (arp, ndp, ppp, )
- how to fix a slow 3-tier application ?
- troubleshoot a network connection from your workstation to a server inside the company? and from your workstation to a client in another time zone?
- what is idempotent ? apply for which HTTP method ?

  | HTTP Method | Idempotent | Safe |
  | ----------- |:----------:| ----:|
  | HEAD        | yes        | yes  |
  | OPTIONS     | yes        | yes  |
  | GET         | yes        | yes  |
  | PUT         | yes        | no   |
  | DELETE      | yes        | no   |
  | PATCH       | no         | no   |
  | POST        | no         | no   |


<a name="math"><br /></a>
## Math

- cumulative sum of fibonacci series. fib(n) = addition of all the fibonacci numbers up to n-1
- design a recursive method that calculates a fibonacci sequence, rewrite it to remove the recursion and reduce its time complexity.
- find if an array of numbers that satisfy the fibonacci sequence.
- find prime numbers within 0..N
- find the sum of the fibonacci numbers up to the n-th one
- how to figure out if an integer is out of range
- giving a 2D grid with pixels valued 0 or 1. How to check if pixels A and B are connected through a path of 0-valued pixels
- give a deck of card, calculate total number of point that is closes to 21. e.g. A,A,J = 12; J,J,A,2 = 23; A,2 = 13
  - game of blackjack, implement a function for blackjack that returns the score of your hand - write code for `getScore()`:

  ```
  class Hand {
    List<Card> cards;
    int getScore( ) { }
  }
  ```
- print out Fibonacci series in MATLAB
- reverse a number


<a name="linked-list"><br /></a>
## Linked List

- find the loop starting node (if any) in a linked list
- find the middle of a linked list with only one pass and 2 pointers (references)
- merge two sorted linked lists. Questions on time and space complexity.
- reverse a linked list


<a name="map"><br /></a>
## Map

- traverse the map


<a name="string"><br /></a>
## String

- check if an expression has proper openings and closings of `()`, `[]`, and `{}`.
- find all anagrams of a given word. A array including all English words is provided.
- find first unique char in a string, find least greatest value with given target in a BST.
- find the length of string recursively.
- detects the first non-repeating character in a char array, and do so with only a single pass over the array.
- given a dollar amount, output a textual representation (i.e. $123.43 -> "One hundred twenty three dollars and forty three cents”).
- inverts the case of each character in a string.
- parse a string into a double, where the string could be anything like "1 1/2" or "2/5" or "-3”.


<a name="tree"><br /></a>
## Tree

- adding and removing nodes from a ternary tree
- binary search tree vs balanced BST vs heap (binary heap, min-/max-heap)
- check if a binary tree a BST (binary search tree)
- compare the node in an unsorted d-tree
- create a tree, e.g. insert and delete in a trinary tree
- find the common ancestor of two nodes
- find the Least Common Ancestor given two nodes of a binary tree. The nodes each have a reference to their parent node and you do not have the root node of the tree.
- find LCA of a Binary Tree and making it efficient.
- given a binary tree and a sum, determine if the tree has a root-to-leaf path such that adding up all the values along the path equals the given sum. Follow-up: find all the paths
- lowest common ancestor in a BT
- questions about trees with unconventional structures
- search through a binary search tree, what is the worst-case big-O complexity?
- traversal methods
  - breadth-first

  ```javascript
  //        A(root)
  //       /      \
  //      B        C
  //     /  \       \
  //    D    E       F
  //        / \     /
  //       G   H   I
  //
  breadthFirst(root) {
    var node = root;
    var nodeList = []; // node list
    var queue = [];
    queue.push(node); // push root into queue
    while (queue.length > 0) {
      node = queue.shift(); // dequeue the node
      nodeList.push(node);  // append node to the list
      if (node.left) queue.push(node.left);
      if (node.right) queue.push(node.right);
    }
    return nodeList;
  }
  ```

  - depth-first (pre-order):
      - top-down order
      - e.g. reading hierarchical document in natural oder (chapters/sections/... in a book)
      - e.g. prefix expression tree `"(a+b)*c" => * + a b c` (Polish Notation) in arithmetic parser

  ```javascript
  //        A(root)
  //       /      \
  //      B        G
  //     /  \       \
  //    C    D       H
  //        / \     /
  //       E   F   I
  //
  depthFirstPreOrder(root) {
    var node = root;
    var nodeList = [];
    var stack = [];
    stack.push(node);
    while (stack.length > 0) {
      node = stack.pop(); // pop the node from stack
      if (node.right) stack.push(node.right);
      if (node.left) stack.push(node.left);
    }
    return nodeList;
  }

  depthFirstPreOrderRecursive(node, nodeList) {
    if (node) {
      nodeList.push(node);
      depthFirstInOrder(node.left, nodeList);
      depthFirstInOrder(node.right, nodeList);
    }
  }
  ```

  - depth-first (in-order)
    - e.g. binary search tree, or infix `(a+b)*c`

  ```javascript
  //        F(root)
  //       /      \
  //      B        G
  //     /  \       \
  //    A    D       I
  //        / \     /
  //       C   E   H
  //
  depthFirstInOrder(node, nodeList) {
    if (node) {
      depthFirstInOrder(node.left, nodeList);
      nodeList.push(node);
      depthFirstInOrder(node.right, nodeList);
    }
  }
  ```
  - depth-first (post-order)
    - bottom-up order, visiting all leaves before their parent
    - e.g. postfix expression tree `"(a+b)*c" => a b + c *` (Reverse Polish Notation) in arithmetic parser

  ```javascript
  //        I(root)
  //       /      \
  //      E        H
  //     /  \       \
  //    A    D       G
  //        / \     /
  //       B   C   F
  //
  depthFirstPostOrder(node, nodeList) {
    if (node) {
      depthFirstInOrder(node.left, nodeList);
      depthFirstInOrder(node.right, nodeList);
      nodeList.push(node);
    }
  }
  ```

- tree construction
  - create from pre-order

  ```javascript
  fromPreOrder(inputs) {
     var root = {};
     var size = inputs.length;
     var half = size / 2 + 1;
     if (size > 0) {
       root.value = inputs[0]; // first item is the root
       if (size > 1) {
         root.left = fromPreOrder(inputs.slice(1, half));
       } else if (size > half) {
         root.left = fromPreOrder(inputs.slice(half, size));
       }
     }
     return null;
  }
  ```

  - create from post-order

  ```javascript
  fromPostOrder(inputs) {
     var root = {};
     var size = inputs.length;
     var half = size / 2;
     if (size > 0) {
       root.value = inputs[size-1]; // last item is the root
       if (size > 1) {
         root.left = fromPreOrder(inputs.slice(0, half));
       } else if (size > half) {
         root.left = fromPreOrder(inputs.slice(half, size-1));
       }
     }
     return null;
  }
  ```

  - create from in-order

  ```javascript
  fromInOrder(inputs) {
     var root = {};
     var size = inputs.length;
     var half = size / 2;
     if (size > 0) {
       root.value = inputs[half]; // middle item is the root
       if (size > 1) {
         root.left = fromInOrder(inputs.slice(0, half));
       } else if (size > half) {
         root.left = fromInOrder(inputs.slice(half+1, size));
       }
     }
     return null;
  }
  ```
