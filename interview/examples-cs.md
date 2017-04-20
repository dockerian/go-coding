# C&#35; Interview Questions

<a name="general"><br/></a>
## General Questions

- Q: Does C# support multiple-inheritance?
<br/>A: No.

- Q: Who is a protected class-level variable available to?
<br/>A: It is available to any sub-class (a class inheriting this class).

- Q: Are private class-level variables inherited?
<br/>A: Yes, but they are not accessible.  Although they are not visible or accessible via the class interface, they are inherited.

- Q: Describe the accessibility modifier “protected internal”.
<br/>A: It is available to classes that are within the same assembly and derived from the specified base class.

- Q: What’s the top .NET class that everything is derived from?
<br/>A: System.Object.

- Q: What does the term immutable mean?
<br/>A: The data value may not be changed.  Note: The variable value may be changed, but the original immutable data value was discarded and a new data value was created in memory.

- Q: What’s the difference between System.String and System.Text.StringBuilder classes?
<br/>A: System.String is immutable.  System.StringBuilder was designed with the purpose of having a mutable string where a variety of operations can be performed.

- Q: What’s the advantage of using System.Text.StringBuilder over System.String?
<br/>A: StringBuilder is more efficient in cases where there is a large amount of string manipulation.  Strings are immutable, so each time a string is changed, a new instance in memory is created.

- Q: Can you store multiple data types in System.Array?
<br/>A: No.

- Q: What’s the difference between the System.Array.CopyTo() and System.Array.Clone()?
<br/>A: The Clone() method returns a new array (a shallow copy) object containing all the elements in the original array.  The CopyTo() method copies the elements into another existing array.  Both perform a shallow copy.  A shallow copy means the contents (each array element) contains references to the same object as the elements in the original array.  A deep copy (which neither of these methods performs) would create a new instance of each element's object, resulting in a different, yet identacle object.

- Q: How can you sort the elements of the array in descending order?
<br/>A: By calling Sort() and then Reverse() methods.

- Q: How can you sort list of a particular type objects?
<br/>A: Implement `IComparable<T>`, using sort delegate, or `OrderBy`

```csharp
List<Foo> unsorted;
unsorted.Sort((x, y) => x.Id.CompareTo(y.Id));
List<Foo> sortedList = unsorted.OrderBy(o => o.Id).ThenBy(o => o.Name).ToList();
```

- Q: How to use `SortedSet<T>` ?
<br/>A: Use `SortedSet<T>` as a heap, or `SortedDictionary<T>` as priority queue

```csharp
SortedSet<Foo> sortedById = new SortedSet<Foo>(new ByFooId());

public class ByFooId : IComparer<Foo> {
  public int Compare(x, y) {
    return x.Id.CompareTo(y.Id)
  }
}
```

- Q: What’s the .NET collection class that allows an element to be accessed using a unique key?
<br/>A: HashTable.

- Q: What class is underneath the SortedList class?
<br/>A: A sorted HashTable.

- Q: Will the finally block get executed if an exception has not occurred?­
Yes.

- Q: What’s the C# syntax to catch any possible exception?
<br/>A: A catch block that catches the exception of type System.Exception.  You can also omit the parameter data type in this case and just write catch {}.

- Q: Can multiple catch blocks be executed for a single try statement?
<br/>A: No.  Once the proper catch block processed, control is transferred to the finally block (if there are any).

- Q: Explain the three services model commonly know as a three-tier application.
<br/>A: Presentation (UI), Business (logic and underlying code) and Data (from storage or other sources).


<a name="class"><br/></a>
## Class Questions

- Q: What is the syntax to inherit from a class in C#?
<br/>A: Place a colon and then the name of the base class. Example: class MyNewClass : MyBaseClass

- Q: Can you prevent your class from being inherited by another class?
<br/>A: Yes.  The keyword “sealed” will prevent the class from being inherited.

- Q: Can you allow a class to be inherited, but prevent the method from being over-ridden?
<br/>A: Yes.  Just leave the class public and make the method sealed.

- Q: What’s an abstract class?
<br/>A: A class that cannot be instantiated.  An abstract class is a class that must be inherited and have the methods overridden.  An abstract class is essentially a blueprint for a class without any implementation.

- Q: When do you absolutely have to declare a class as abstract?
1. When the class itself is inherited from an abstract class, but not all base abstract methods have been overridden.
2. When at least one of the methods in the class is abstract.

- Q: What is an interface class?
<br/>A: Interfaces, like classes, define a set of properties, methods, and events. But unlike classes, interfaces do not provide implementation. They are implemented by classes, and defined as separate entities from classes.

- Q: Why can’t you specify the accessibility modifier for methods inside the interface?
<br/>A: They all must be public, and are therefore public by default.

- Q: Can you inherit multiple interfaces?
<br/>A: Yes.  .NET does support multiple interfaces.

- Q: What happens if you inherit multiple interfaces and they have conflicting method names?
<br/>A: It’s up to you to implement the method inside your own class, so implementation is left entirely up to you. This might cause a problem on a higher-level scale if similarly named methods from different interfaces expect different data, but as far as compiler cares you’re okay.
To Do: Investigate

- Q: What’s the difference between an interface and abstract class?
<br/>A: In an interface class, all methods are abstract - there is no implementation.  In an abstract class some methods can be concrete.  In an interface class, no accessibility modifiers are allowed.  An abstract class may have accessibility modifiers.

- Q: What is the difference between a Struct and a Class?
<br/>A: Structs are value-type variables and are thus saved on the stack, additional overhead but faster retrieval.  Another difference is that structs cannot inherit.


<a name="method"><br/></a>
## Method and Property Questions

- Q: What’s the implicit name of the parameter that gets passed into the set method/property of a class?
<br/>A: Value.  The data type of the value parameter is defined by whatever data type the property is declared as.

- Q: What does the keyword “virtual” declare for a method or property?
<br/>A: The method or property can be overridden.

- Q: How is method overriding different from method overloading?
<br/>A: When overriding a method, you change the behavior of the method for the derived class.  Overloading a method simply involves having another method with the same name within the class.

- Q: Can you declare an override method to be static if the original method is not static?
<br/>A: No.  The signature of the virtual method must remain the same.  (Note: Only the keyword virtual is changed to keyword override)

- Q: What are the different ways a method can be overloaded?
<br/>A: Different parameter data types, different number of parameters, different order of parameters.

- Q: If a base class has a number of overloaded constructors, and an inheriting class has a number of overloaded constructors; can you enforce a call from an inherited constructor to a specific base constructor?
<br/>A: Yes, just place a colon, and then keyword base (parameter list to invoke the appropriate constructor) in the overloaded constructor definition inside the inherited class.


<a name="event"><br/></a>
## Events and Delegates

- Q: What’s a delegate?
<br/>A: A delegate object encapsulates a reference to a method.

- Q: What’s a multicast delegate?
<br/>A: A delegate that has multiple handlers assigned to it.  Each assigned handler (method) is called.


<a name="xml"><br/></a>
## XML Documentation Questions

- Q: Is XML case-sensitive?
<br/>A: Yes.

- Q: What’s the difference between `//` comments, `/* */` comments and `///` comments?
<br/>A: Single-line comments, multi-line comments, and XML documentation comments.

- Q: How do you generate documentation from the C# file commented properly with a command-line compiler?
<br/>A: Compile it with the /doc switch.


<a name="debug"><br/></a>
## Debugging and Testing Questions

- Q: What debugging tools come with the .NET SDK?
1. CorDBG – command-line debugger.  To use CorDbg, you must compile the original C# file using the /debug switch.
2. DbgCLR – graphic debugger.  Visual Studio .NET uses the DbgCLR.

- Q: What does assert() method do?
<br/>A: In debug compilation, assert takes in a Boolean condition as a parameter, and shows the error dialog if the condition is false.  The program proceeds without any interruption if the condition is true.

- Q: What’s the difference between the Debug class and Trace class?
<br/>A: Documentation looks the same.  Use Debug class for debug builds, use Trace class for both debug and release builds.

- Q: Why are there five tracing levels in System.Diagnostics.TraceSwitcher?
<br/>A: The tracing dumps can be quite verbose.  For applications that are constantly running you run the risk of overloading the machine and the hard drive.  Five levels range from None to Verbose, allowing you to fine-tune the tracing activities.

- Q: Where is the output of TextWriterTraceListener redirected?
<br/>A: To the Console or a text file depending on the parameter passed to the constructor.

- Q: How do you debug an ASP.NET Web application?
<br/>A: Attach the aspnet_wp.exe process to the DbgClr debugger.

- Q: What are three test cases you should go through in unit testing?
1. Positive test cases (correct data, correct output).
2. Negative test cases (broken or missing data, proper handling).
3. Exception test cases (exceptions are thrown and caught properly).

- Q: Can you change the value of a variable while debugging a C# application?
<br/>A: Yes.  If you are debugging via Visual Studio.NET, just go to Immediate window.


<a name="ado"><br/></a>
## ADO.NET and Database Questions

- Q: What is the role of the DataReader class in ADO.NET connections?
<br/>A: It returns a read-only, forward-only rowset from the data source.  A DataReader provides fast access when a forward-only sequential read is
needed.

- Q: What are advantages and disadvantages of Microsoft-provided data provider classes in ADO.NET?
<br/>A: SQLServer.NET data provider is high-speed and robust, but requires SQL Server license purchased from Microsoft. OLE-DB.NET is universal for accessing other sources, like Oracle, DB2, Microsoft Access and Informix.  OLE-DB.NET is a .NET layer on top of the OLE layer, so it’s not as fastest and efficient as SqlServer.NET.

- Q: What is the wildcard character in SQL?
<br/>A: Let’s say you want to query database with LIKE for all employees whose name starts with La. The wildcard character is %, the proper query with LIKE would involve ‘La%’.

- Q: Explain ACID rule of thumb for transactions.
<br/>A: A transaction must be:
1. Atomic - it is one unit of work and does not dependent on previous and following transactions.
2. Consistent - data is either committed or roll back, no “in-between” case where something has been updated and something hasn’t.
3. Isolated - no transaction sees the intermediate results of the current transaction).
4. Durable - the values persist if the data had been committed even ifthe system crashes right after.

- Q: What connections does Microsoft SQL Server support?
<br/>A: Windows Authentication (via Active Directory) and SQL Server authentication (via Microsoft SQL Server username and password).

- Q: Between Windows Authentication and SQL Server Authentication, which one is trusted and which one is untrusted?
<br/>A: Windows Authentication is trusted because the username and password are checked with the Active Directory, the SQL Server authentication is untrusted, since SQL Server is the only verifier participating in the transaction.

- Q: What does the Initial Catalog parameter define in the connection string?
<br/>A: The database name to connect to.

- Q: What does the Dispose method do with the connection object?
<br/>A: Deletes it from the memory.
To Do: answer better.  The current answer is not entirely correct.

- Q: What is a pre-requisite for connection pooling?
<br/>A: Multiple processes must agree that they will share the same connection, where every parameter is the same, including the security settings.  The connection string must be identical.


<a name="assembly"><br/></a>
## Assembly Questions

- Q: How is the DLL Hell problem solved in .NET?
<br/>A: Assembly versioning allows the application to specify not only the library it needs to run (which was available under Win32), but also the version of the assembly.

- Q: What are the ways to deploy an assembly?
<br/>A: An MSI installer, a CAB archive, and XCOPY command.

- Q: What is a satellite assembly?
<br/>A: When you write a multilingual or multi-cultural application in .NET, and want to distribute the core application separately from the localized modules, the localized assemblies that modify the core application are called satellite assemblies.

- Q: What namespaces are necessary to create a localized application?
<br/>A: System.Globalization and System.Resources.

- Q: What is the smallest unit of execution in .NET?
<br/>A: an Assembly.

- Q: When should you call the garbage collector in .NET?
<br/>A: As a good rule, you should not call the garbage collector.  However, you could call the garbage collector when you are done using a large object (or set of objects) to force the garbage collector to dispose of those very large objects from memory.  However, this is usually not a good practice.

- Q: How do you convert a value-type to a reference-type?
<br/>A: Use Boxing.

- Q: What happens in memory when you Box and Unbox a value-type?
<br/>A: Boxing converts a value-type to a reference-type, thus storing the object on the heap.  Unboxing converts a reference-type to a value-type, thus storing the value on the stack.
