# Interview Session Examples

<a name="session1-answers"><br/></a>
## Session 1 - [Answers](examples-qa.md#session1)

### Coding/Language

- Q: In C++ what is a typical (good) way to allocate a pointer to memory in an Exception Safe way?
- Q: In object-oriented speak – what is the difference between "Is-A" and "Has-A"? Or put another way `inheritance` and `ownership`?
- Q: What does `thread safe` mean?
- Q: What does `exception safe` mean?
- Q: What are common sorting algorithms?


### Data Structures

- Q: What is a `dequeue`?
- Q: What is a hash table?
- Q: What is a priority queue?
- Q: What is a tree data structure? Design a data model for a family tree problem.
- Q: What is a well-known data structure for saving/retrieving/deleting a value from persistent storage (disk) given a key?
- Q: Use "Big-O" notation to describe the performance of inserting to the front of a singly linked list.


### Multi-threading

- Q: What is a "critical section"?
- Q: What is a "thread pool" and when might one be used?
- Q: What is a "thread safe producer consumer queue"?


### Design and Design Patterns

- Q: Define a solution to find all pages that contain US phone numbers from a directory that contain 50,000 pages where each phone number can be of format `+1 xxx-xxx-xxxx` OR `(xxx) xxx-xxxx`.
- Q: Have you ever heard of a `pimpl` or "private (or pointer to an) implementation" design pattern? When is it used?
- Q: What is a Singleton?


### Coding/Algorithms

- Q: Code and articulate a binary search implementation on the phone, ability to analyze complexity. Implication between recursion vs. non-recursion.
- Q: Design an algorithm to multiply two really large numbers where each number is big enough to not fit into usual data types and is represented as string.
- Q: Design an OO class to shuffle a deck of cards (OO).


### Assignments

#### Assignment 1

- Part 1: Imagine the following compare_and_swap (CAS) primitive function is available to you. Using this function, write an integer counter class providing increment, decrement, set and get APIs in an atomic and thread-safe way.

  ```cpp
  /* In a critical section, if (*dest) == old_val, then
   * updates dest to new_val and returns true; otherwise
   * leaves dest unchanged and returns false.
   */
  bool compare_and_swap(int* dest, int old_val, int new_val)
  ```

- Part 2: Imagine the code you wrote in Part 1 were being used in a context subject to periodic highly intensive concurrent activity (i.e. at peak, hundreds of threads looking to update a given counter in a short period of time.) Could you imagine any problems with the code as written, if any? If you identify problems, what might be done to address them?

**Note**: we do not expect you to implement compare_and_swap function. As mentioned above, it is available to you, so that you can simply use it to implement an integer counter class. In addition, we look for a solution that can avoid using a mutex.


#### Assignment 2

Please write a method to normalize a string which represents a file path. For the purposes of this question, normalizing means:

- all single dot components of the path must be removed. For example, `foo/./bar` should be normalized to `foo/bar`.
- all double dots components of the path must be removed, along with their parent directory. For example, `foo/bar/../baz` should be normalized to `foo/baz`.

Normally, a path normalization algorithm would do a lot of other stuff, but for this question, do not try any other kind of normalization or transformation of the path. As an example, `foo//bar` should be normalized to `foo//bar` (i.e. should be a no-op).

While you can use any language you feel comfortable in, we prefer Java. Please use the following interface:

```java
public interface PathNormalizer {
  // Take a string and return a string representing the normalized path.  
  public String normalizePath(String path);
}
```


<a name="session2-answers"><br/></a>
## Session 2 - [Answers](examples-qa.md#session2)

- Q: When should you use NOLOCK for SQL statements ?

- Q: What is a CTE and what is a common use for them ?

- Q: What is the null-conditional operator ("Elvis") operator (?.) e.g. myObj?.member?.submember

- Q: What is JSON? Why is it important? How do you read and write it ?

- Q: What are lambda expressions ?

- Q: When are the lambda expressions executed in LINQ ?

- Q: If you had a list of lists of strings how would you write that type in C#? How would you flatten it into a list of strings using LINQ ?

- Q: What’s the difference between a concrete method, a virtual method and an abstract method ?

- Q: What is SOLID code? What is DRY code ?

- Q: Part of a 3-tier application is very slow. How do you fix it ?

- Q: In 3-4 member-team scrum model, how do you handle a situation when your current tasks will likely to exceed the original estimates and overflow into next sprint ?

- Q: Have you ever had a situation when you discover missing scenarios/requirements after you started coding? What would you do ?


<a name="session3-answers"><br/></a>
## Session 3 - [Answers](examples-qa.md#session3)

- Q: Explain Polymorphism.

- Q: Explain Abstract class, comparing to Interface.

- Q: In the context of Java or .NET generics, what is type erasure?

- Q: What is the difference between an overloaded and an overridden method?

- Q: Are you familiar with any Design Patterns? Give examples and describe.

- Q: Have you used any separation of concerns patterns, e.g. MV* ?

- Q: Compare MVC with MVVM

- Q: Explain Hashtable, and its Big-O time notation on read and write. How is it implemented internally?

- Q: Can you have a memory leak in Java/C# ?


- Q: Are you familiar with the atoi function? Write a function that takes a string numeric value and returns the corresponding integer value, e.g. "1234" => `1234`

- Q: Assume we have 50,000 HTML files in a directory tree, under a directory called "/website". We have 2 days to get a list of file paths that contain telephone numbers. This is a one-time, non-repetitive task. The phone numbers are in the following two formats: (xxx) xxx-xxxx and xxx-xxx-xxxx. How would you solve this problem? Keep in mind our team is on a short (2-day) timeline.

- Q: Given an array of integers where all the values occur an odd number of times except one, how do you efficiently determine which value is the one that occurs an even number of times.

- Q: Write a function that counts the number of set bits in an integer value.

  ```
  6 -> 110 -> 2
  3 -> 011 -> 2
  4 -> 100 -> 1
  ```

- Q: Write a function to find the smallest 1 million numbers in a large store of 1 billion numbers.

- Q: Give 5 mins summary of your current work and things you are passionate about?


<a name="session4-answers"><br/></a>
## Session 4 - [Answers](examples-qa.md#session4)

- Q: Describe when you know the code quality is good enough.

- Q: Design a card game, or a scrabble, or a board game.

- Q: Find first K closest points from a point.

- Q: Find the number of occurrences of the most frequent substring of length L (L>=2) in a string of length N (N<10000).

- Q: Function to return the nth Fibonacci number in the Fibonacci Sequence.
  - Fibonacci Sequence: 1, 1, 2, 3, 5, 8, 13, 21, 34
  - Example: N = 6
  - Return: 8

- Q: Function to reverse the order of words in a string, with reversing the order of the letters in the words.
  - Input: "The quick brown fox jumped over the lazy dog."
  - Return: "dog. lazy the over jumped fox brown quick The"

- Q: Function passes a file as an argument that checks a file and excludes duplicative words.

- Q: Given a string, print out all possible variations of upper and lower case.

- Q: Reverse a linked list at each pair and write test cases.

- Q: Spiral printing of 2D array.

- Q:

<br />
