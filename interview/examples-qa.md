# Interview Session Examples

<a name="session1"><br /></a>
## Session 1

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


<a name="session2"><br /></a>
## Session 2

- Q: When should you use NOLOCK for SQL statements ?

- Q: What is a CTE and what is a common use for them ?

**Answer:**

```sql
;WITH cte AS
(SELECT ROW_NUMBER()
        OVER (PARTITION BY col1, col2, col3 ORDER BY rowId) RowNum
   FROM SomeTable)
DELETE FROM cte WHERE RowNum > 1
```

otherwise, the query would be

```sql
DELETE FROM SomeTable
  LEFT OUTER JOIN (
   SELECT MIN(rowId) as rowId, *
     FROM SomeTable
    GROUP BY col1, col2, col3 -- any columns to identify unique row
 ) as KeptRows ON SomeTable.rowId = KeptRows.rowId
 WHERE KeepRows.rowId IS NULL
```

- Q: What is the null-conditional operator ("Elvis") operator (?.) e.g. myObj?.member?.submember

- Q: What is JSON? Why is it important? How do you read and write it ?

- Q: What are lambda expressions ?

- Q: When are the lambda expressions executed in LINQ ?

- Q: If you had a list of lists of strings how would you write that type in C#? How would you flatten it into a list of strings using LINQ ?

**Answer:**

```csharp
ListOfListOfString.SelectMany(ll => ll).Distinct().ToList()
```

- Q: What’s the difference between a concrete method, a virtual method and an abstract method ?

- Q: What is SOLID code? What is DRY code ?

- Part of a 3-tier application is very slow. How do you fix it ?

- Q: In 3-4 member-team scrum model, how do you handle a situation when your current tasks will likely to exceed the original estimates and overflow into next sprint ?

- Q: Have you ever had a situation when you discover missing scenarios/requirements after you started coding? What would you do ?


<a name="session3"><br /></a>
## Session 3

- Q: Explain Polymorphism.

**Answer:**

```java
Animal a = new Cat();
a.Run();

IMailService s; // or MailService s;
s = new MyMailService();
s.SendMail();
```

- Q: (1) Explain Abstract class. (2) Comparing to Interface.

- Q: In the context of Java or .NET generics, what is type erasure?

**Answer:** Type erasure refers to the compile-time process by which explicit type annotations are removed from a program, before it is executed at run-time.
  - replace all type parameters in generic types with their bounds or Object if the type parameters are unbounded. The produced bytecode, therefore, contains only ordinary classes, interfaces, and methods.
  - insert type casts if necessary to preserve type safety.
  - generate bridge methods to preserve polymorphism in extended generic types.
  - ensures that no new classes are created for parameterized types; consequently, generics incur no runtime overhead.

- Q: What is the difference between an overloaded and an overridden method?

- Q: Are you familiar with any Design Patterns? Give examples and describe.

- Q: Have you used any separation of concerns patterns, e.g. MV* ?

- Q: Compare MVC with MVVM

- Q: Explain Hashtable, and its Big-O time notation on read and write. How is it implemented internally?

**Answer:** TBD

- Q: Can you have a memory leak in Java/C# ? Give Example.


- Q: Are you familiar with the atoi function? Write a function that takes a string numeric value and returns the corresponding integer value, e.g. "1234" => `1234`

**Answer:**

```javascript
function atoi(s) {
  if (!s) return 0;
  var result = 0;
  for(var i=0; i< s.length; i++) {
      v = s[i] - '0'; // no valid digit check here
      result = result*10 + v; // no overflow check
  }
  return result;
}
```
See also in golang (with more checking):

```go
func Atoi(s string) (int, error) {
  var result int
  // checking null and/or empty input
  if len(s) return result
  for _, ch := range s {
      r := result
      v := int(ch) - int('0')
      // checking valid digit
      if 0 <= v && v <= 9 {
        r = r*10 + v
        if r < result {
          // checking overflow
          return result, fmt.Errorf("Overflow on Atoi: %s", s)
        }
        continue;
      }
      return result, fmt.Errorf("Parsing error on Atoi %s", s)
  }
  return result, nil
}
```

- Q: Assume we have 50,000 HTML files in a directory tree, under a directory called "/website". We have 2 days to get a list of file paths that contain telephone numbers. This is a one-time, non-repetitive task. The phone numbers are in the following two formats: (xxx) xxx-xxxx and xxx-xxx-xxxx. How would you solve this problem? Keep in mind our team is on a short (2-day) timeline.

**Answer:**

```bash
grep -e '\(?\d{3}\)? \d{3}-\d{3}'
```

- Q: Given an array of integers where all the values occur an odd number of times except one, how do you efficiently determine which value is the one that occurs an even number of times.

```javascript
function findTheOnlyEvenCountIntegerArray(a) {
    var counts = {};
    for(var i=0; i < a.length; i++) {
        key = a[i]
        counts[key] = 1 + (counts[key] ? counts[key] : 0)
    }
    for (var key in counts) {
        if (counts[key] % 2 == 0) return key;
    }
}
// see the function to find only odd count
function findTheOnlyOddCountInIntegerArray(a) {
    return a.reduce(function(x, y) {
        return x ^ y
    });
}
```

- Q: Write a function that counts the number of set bits in an integer value.

  ```
  6 -> 110 -> 2
  3 -> 011 -> 2
  4 -> 100 -> 1
  ```
  **Answer:**

```javascript
function countBits(n) {
    var count = 0;
    while (n > 0) {
      if (n & 1) count++;
        n = n >> 1;
   }
   return count;   
}
```

- Q: Write a function to find the smallest 1 million numbers in a large store of 1 billion numbers.

**Answer:** TBD

- Q: Give 5 mins summary of your current work and things you are passionate about?


<a name="session4"><br /></a>
## Session 4

- Q: Calculate degrees between clock hands (with or without precise to second).

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

  **Answer:**

  ```
  #include <cstdlib>
  #include <cmath>
  #include <iostream>
  using namespace std;

  // Prints combinations of upper/lowercase letters
  int PrintAllCases(const char *s, int size) {
    if (!s || size < 1 || size > 31) return 0;
    int max = pow(2.0, size); // int max = 1 << size;

    for (int i = 0; i < max; i++) {
      int flag = i;
      for (int j = 0; j < size; j++) {
        if (flag & 0x1) {
          putchar(toupper(s[j]));
        } else {
          putchar(tolower(s[j]));
        }
        flag >>= 1; // divided by 2
      }
      printf("\n");
    }
  }
  ```

- Q: Reverse a linked list at each pair and write test cases.

- Q: Spiral printing of 2D array.

- Q:

<br />
