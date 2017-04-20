# Golang Notes

## Golang Study Notes

### Golang Training

<a name="ccis"><br/></a>
#### CCIS material

  - Trainer: Cory LaNou
    - https://angel.co/corylanou
    - https://github.com/corylanou
    - interview https://medium.com/@IndianGuru/interview-gopher-cory-lanou-ded55150020b#.uobpqkpa2
  - wifi: dedicatedccis3/kirkland3; ccis1/washington
  - slack
	  - https://gophersinvite.herokuapp.com (invite)
	  - https://gophers.slack.com/messages/training (channel)
  - go playground: http://play.golang.org
  - main material: http://github.com/ardanlabs/gotrainings
  - awecome collection: [awesome-go](https://github.com/avelino/awesome-go)
  - books: [go books](https://github.com/dariubs/GoBooks)
  - Ultimate Go ([5-day material](https://github.com/ardanlabs/gotraining/blob/master/courses/ultimate/README.md))
  - package builtin https://golang.org/pkg/builtin/
  - https://forum.golangbridge.org/
  - https://sourcegraph.com/


<a name="concurreny"><br/></a>
#### Concurreny in Go

  - Channels

  ```go
  const MaxOutstanding = 5
  var semaphore = make(chan int, MaxOutstanding)
  func Serve(queue chan *Request) {
      for req := range queue {
          semaphore <- 1
          // goroutine with `req` parameter to create a closure
          go func(req *Request) {
              process(req)
              <-semaphore
          }(req)
      }
  }
  ```


<a name="testing"><br/></a>
#### Go testing

  - use `-timeout` to timeout the testing
  - failing one test vs stop all tests: `t.Errorf` vs `t.Fatalf`
  - mock http server https://golang.org/pkg/net/http/httptest/

```
func mockServer() *httptest.Server {
	f := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Header().Set("Content-Type", "application/xml")
		fmt.Fprintln(w, feed)
	}

	return httptest.NewServer(http.HandlerFunc(f))
}
```
  - http testing https://golang.org/pkg/net/http/httptest/
  - https://github.com/cespare/reflex
  - benchmark benchstat
  - vet


<a name="debugging"><br/></a>
#### Go debugging/gdb

  - https://github.com/wg/wrk (GUI http://www.graphviz.org/Download.php)
  - Debugging Go with VSCode: https://github.com/Microsoft/vscode-go
  - Using GDB is really not an option. Use Delve: https://github.com/derekparker/delve
  - Debugging worked after a few additional steps to these instructions.
    1. sudo chgrp procmod /usr/local/bin/gdb
    2. sudo chmod g+s /usr/local/bin/gdb
    3. add the legacy switch -p to taskgated by modifying the file /System/Library/LaunchAgents/com.apple.taskgated.plist
    4. Force kill taskgated process (it will restart) or reboot if necessary
    5. Try again
  - Solution comes from this stackoverflow thread: http://stackoverflow.com/questions/12050257/gdb-fails-on-mountain-lion
  - ALERT! Be sure to use standard .plist markup with the file modification. Otherwise osx won't start next time you reboot. This happened to me. Solution to this is to reboot with recovery option (Cmd-R) and modifying the file with vi to match the standard.
  - Error from debug console in Lite-IDE --> "Unable to find Mach task port for process-id 12383: (os/kern) protection failure (0x2). (please check gdb is codesigned - see taskgated(8)". This error came even though I had codesigned gdb 7.6 with the exact steps of your instructions. Now I can reproduce the error by unloading the /System/Library/LaunchDaemons/com.apple.taskgated.plist, deleting the -p switch and reloading taskgated. My Mac OS is same as yours. The problem has to do with codesigning, but I can't figure out in what way.
  - memory tracing https://github.com/ardanlabs/gotraining/tree/master/topics/memory_trace
  ```
  http.ListenAndServe(":6060", nil)
  ```
  - dump https://github.com/davecgh/go-spew


<a name="benchmarking"><br/></a>
#### Go benchmarking

  - example

```
var fa int // suppress go optimization
// BenchmarkFib provides performance numbers for the fibonacci function.
func BenchmarkFib(b *testing.B) {
var a int

for i := 0; i < b.N; i++ {
	a = fib(1e5)
}

fa = a
}
```
  - using `go test`

```
go test -run none -bench
go test -run none -bench . -cpuprofile cpu.out -benchtime 3s -benchmem
go tool pprof profiling.test cpu.out
```
  - live benchmark ?


<a name="logging"><br/></a>
#### Go logging

  - http://dave.cheney.net/2015/11/05/lets-talk-about-logging
  - https://github.com/op/go-logging - smaller than the other here
  - https://github.com/Sirupsen/logrus - used in many popular projects such as Docker
  - https://github.com/inconshreveable/log15
  - https://github.com/cloudfoundry/lager


<a name="tips"><br/></a>
#### Questions/tips

  - nested scope ?
  - fmt is pronounced as 'phumpt'
  - can struct be nested and self-reference ?
  - don' pass pointer unless needed; don't use pointer in struct ?
  - &mystruct.prop is the pointer to the prop field
  - println vs fmt.Println (using reflection)
  - gcflags ```go build -gcflags -m```
  - const cannot be struct
  - error is primitive type in go ? `error.New` or `fmt.Errorf`
  - no panic in lib; panic only for debugging ? which level ? recover ?
  - func return pointer or reference ? https://golang.org/doc/faq#methods_on_values_or_pointers
  - changing slice cap may create new heap
  - string is array of runes (not single byte); rune size 4-byte (int32) ?
  - use `[]byte()` to gain performance
  - variadic var must be the last one (caution ... operator)

	```
	func takeVariadic(foo int, users ...user) {}
	takeVariadic(users...)
	```

  - map order is not guaranteed
  - receiver: value vs pointer (attention to the signature)
  - should method defined by type pointer ?
  - cast from interface to concrete ?
  - embedding type is kind of inheritance ?
  - type assertion `v, ok := someVar.(someType)` vs type conversion `someType(someVar)`
  - init() is special, can be multiple in a package ?
  - custom err https://golang.org/doc/faq#nil_error
  - `defer fn()` - `fn` is called at the end of the `func` NOT {scope}
  - concurrency vs parallelism ```var wg sync.WaitGroup```
  - keyword ```defer``` vs ```keyword``` go
  - go playground or ```runtime.GOMAXPROCS(1)``` limits multi-threading
  - always use race detector (--race)
  - avoid racing condition: ```"atomic"``` and/or ```"mutex"```
  - channel has type: ```make(chan int)```
  - multiple go routines read/write one channel ? no order guaranteed
  - cannot write to a closed chan (panic)
  - use no memory: ```struct{}```
  - no order to pick from ```for { select { case <- chan1: ; case <- chan2 }}```
  - channel real examples:
    - https://github.com/influxdata/influxdb/blob/master/cmd/influxd/main.go
    - https://github.com/ardanlabs/gotraining/blob/master/topics/channels/exercises/template1/template1.go
    - https://github.com/ardanlabs/gotraining/blob/master/topics/channels/exercises/template1/template2.go
    - concurrency pattern https://github.com/ardanlabs/gotraining/tree/master/topics/concurrency_patterns
  - overloading http://changelog.ca/log/2015/01/30/golang
  - chan to use

```go
sigChan := make(chan os.Signal, 1)
signal.Notify(sigChan, os.Interrupt)
<-sigChan
```


<a name="others"><br/></a>
#### Other Topics

  - https://github.com/peterh/liner
  - https://github.com/spf13/cobra
  - https://github.com/ardanlabs/gotraining/blob/master/topics/writers_readers/README.md
  - Code review: https://github.com/golang/go/wiki/CodeReviewComments
  - LimitReader
  - TeeReader vs https://golang.org/pkg/bufio/#Reader.Peek
  - WebAPI http://go-talks.appspot.com/github.com/gSchool/go/Exercises/Web-Services-API/presentation.slide#1
  - WebAPI exercises https://github.com/gSchool/go/tree/master/Exercises/Web-Services-API
  - Restful JSON API http://thenewstack.io/make-a-restful-json-api-go/
  - Router https://github.com/julienschmidt/httprouter


<a name="see"><br/></a>
#### See also

  - https://github.com/dariubs/GoBooks
  - https://github.com/avelino/awesome-go
  - https://www.reddit.com/r/golang/


<a name="gophercon"><br/></a>
----------
----------
## Gophercon 2016 Session Notes
==========

- The entire conference was rather small with two days of conference sessions that broke down to a single track in the mornings and three (3) tracks in the afternoons and a third day of “hackathons” and small mini sessions
- The conference sessions were largely informative but most sessions felt more like mini-tutorials

----------
### TL;DR
----------
  * nil values are useful, so check for and use them rather than ignoring `nil`
  * goguru (golang.org/x/tools/cmd/guru) provides powerful tools/capabilities for working with go code in a variety of editors
  * go is very useful for data science applications where the relatively lower performance and variable behavior of python (the same code does not necessarily behave the same over time or across environments) can lead to undesirable results
  * handle errors by treating them as opaque (there is or is not an error, vs what is the specific error type) and annotate errors as they’re passed up the call stack rather than attempting to handle errors at every stage
  * nested vendoring can cause compile time issues in mismatched interface types, avoid vendoring multiple versions of a package and do not expose any vendored types via interfaces or functions
  * AtmanOS (https://github.com/atmanos/atmanos) – a golang unikernel
  * Go achieves cross platform support by providing abstractions at the assembler, which produces machine architecture specific assembly instructions from Go objects, by leveraging the fact that most architectures have assembly languages  with largely similar grammars based on early assembly language syntax/grammar
  * gomobile can be used to create android and iOS applications with Go, but it lacks UI support
  * when writing libraries, keep the library simple by doing the minimum possible/necessary and try to not rely on other libraries (avoid vendoring)

----------
### Usefulness of Nil
----------
  * https://speakerdeck.com/campoy/understanding-nil
  * Pointers
    - nil receivers are useful, when defining functions with pointer receivers, use the pointer receiver value of `nil` to implement some default behavior in unintialized cases (note: `nil` pointer is allowed to call a pointer receiver)
  * Slices
    - A `nil` slice can still be appended to, a slice does not have to be preallocated with `make([]type)` to start appending

	```
	var s int[] // len(s) == 0, cap(s) == 0, s == nil, s is "[]"
	s = append(s, 1)
	```
    - Reallocation to grow the size of a slice is usually fast enough so that preallocation is not always necessary
  * Maps
    - When a map is required for reading (not writing) nil is a valid empty map
    - When invoking functions that require a map as input, rather than passing an empty map as a zero input, use `nil`
  * Channels
    - When selecting on channel input, rather than signifying a closed channel with a variable, assign `nil` to the channel variable when the channel read comes back not ok, else the read will always happen
  * Functions
    - Since functions are first class in go, there is a `nil` value for functions
    - nil function values can be used to indicate default behavior
      * e.g., a function that takes a logger as a function can use a nil value to indicate lack of logging or logging to stdout

----------
### Handle Errors Gracefully
----------
  * http://dave.cheney.net/2016/04/27/dont-just-check-errors-handle-them-gracefully
  * github.com/pkg/errors
  * Errors are values
  * 3 core strategies for handling errors
    * Sentinel errors
      * if `err == ErrSomething {... }`
      * least graceful because a comparison needs to be done for each error being handled
      * you should never have to inspect `error.Error()` to find out what the error is
      * Sentinel errors become part of the API
      * Sentinel errors create dependency between two packages
    * Error types
      * An error that implements the error interface
      * Avoid error types in public APIs
      * Problem with error types
         - Error types must be exposed/exported
         - Creates tight coupling between caller and API, leading to a brittle API
    * Opaque Errors
      * Using `if err != nil { return err }`
      * As caller, treat errors as binary, it happened or it didn’t
      * Reduces coupling, allowing for flexibility of API and implementation
      * Most flexible approach
  * When binary handling is insufficient, assert for behavior, not type
    - Use interfaces to indicate behavior
    - Use of interfaces removes need to import package specific errors or types
  * Handle errors once, don’t log errors in libraries or methods that will pass the error up, instead annotate it (`errors.Wrap`) and pass it up

----------
### Go Vendoring Deconstructed
----------
  * Import resolution, as of vendoring support in 1.5 experimental, has the following resolution order
    - vendor/ (of project)
    - parent’s /vendor
    - parent’s parent’s /vendor
    - ...
    - until `GOPATH`
    - and `GOROOT`
  * Even standard library can be overridden via vendoring so avoid shadowing standard library names
  * Vendored Interfaces
    - Nested vendoring can cause issues when a vendored package A uses another vendored package B, and the main application needs to implement an interface in package B and vendors package B at the same level as A, the compiler sees a mismatch between the two interfaces.
    - For libraries, don’t expose any vendored types via interfaces, functions etc.  If another library needs to be vendored in the library, reimplement or wrap any vendored types/interfaces to avoid possibility of namespace mismatch when caller uses the same vendored library

----------
### Map Implementation in Go
----------
  * Maps are not comparable
  * Maps use buckets to store data, with each bucket having an overflow pointer that points to another bucket
  * Map reference points to metadata about the map
  * Map uses unsafe.Pointers to point to key and value types since there is no support for generics in Go at this time
  * Maps are dynamically resized like slices but copies are done incrementally
    - A few entries are copied over for each access until all entries have been copied over, at which point the old map is garbage collected
    - Reduces CPU impact of resizing large maps but does require memory overhead
  * Maps have about 100% overhead for each stored entry (for simple types) due to need to store metadata and key information
  * Uses about 100 cycles per lookup
  * Takeaway:
    - For small datasets lists/slices likely give the same performance as maps
    - For large datasets, maps provide good runtime performance but has a memory penalty

----------
### Go Without the Operating System
----------
  * Unikernels are a thing for go (yay!)
  * Operating systems are useful but duplicate much of what hypervisors do
  * Even when stripped down to the bare essentials, operating systems are large projects with millions of lines of code with a large attack surface due to the myriad of libraries and drivers
  * Operating Systems strongly impact
    - Boot time
    - Footprint (disk and memory)
    - Performance
    - Security
  * Hypervisor aware applications (unikernels) can interact directly with the hypervisor to perform only essential functions and minimize code
  * AtmanOS (https://github.com/atmanos/atmanos) is an experimental library for go that allows go programs to be compiled as unikernels that run under the Xen hypervisor

----------
### Design of the Go Assembler
----------
  * GCC and other common compilers take the following steps in compiling code
    - Compile -> assemble -> link
  * Plan 9 introduced changes to the process such that compilers generate representations that can be linked and a separate assembler takes assembly to generate representations that can also be linked
    - Compile -> link
    - Assemble -> link
  * Go leverages the same approach to generate object representations of code that can then be linked whether compiling go code or assembly instructions
    - Compile | obj -> link
    - Assemble | obj -> link
  * Go assembler leverages the fact that most assembly code is based on the same set of grammars, so most architectures in assembly are very similar
    - Utilizes lookup tables to account for differences between the architectures
    - Most of the assembler code is portable
  * https://golang.org/doc/asm
  * https://github.com/gophercon/2016-talks/tree/master/RobPike-TheDesignOfTheGoAssembler (slides not yet published as of 7/28/2016)

----------
### Practical Advice for Go Library Authors
----------
  * Generally imports are not renamed, so package names are used as part of the interface, thus stuttering struct names should be avoided
    - `client.Client` // bad, this provides no information about what the Client is
    - `context.Context` // stuttering but accepted
    - `http.Client` // optimal, this provides information about what the Client is a client of
  * Package function names should never stutter
  * Object construction
    - Avoid “New” function and use struct zero value if possible
    - Struct zero value is understandable from the declaration whereas a New function requires that the user inspect the implementation to understand
      * Positional arguments vs named arguments
  * Use config structs when configuring a library
  * Don’t log anything unless absolutely necessary, instead return errors in error cases
  * If logging is required, export a logger interface that the caller can pass an implementation of
  * Interfaces
    - Accept interfaces and return structs
    - Libraries that only expose interfaces tend to have large interfaces
    - Try to split large interfaces in to smaller structs that carry logic
  * Dealing with problems
    - You should usually never panic()
    - Panic is OK if
      * Function begins with MustXYZ()
      * Operations on nil
  * Checking errors
    - Check all errors on interface function calls, especially the ones you don’t expect to error
    - If you can’t return the error, always do something
      * Log it (possibly conflicts with advice about logging)
      * Increment something (a counter either in memory or a metric system)
    - Return errors when
      * A promise could not be kept
      * An answer could not be given
    - If an answer is desired and the answer is negative, it’s not an error
      * i.e. authorize() should not return error when unauthorized
  * Concurrency
    - Channels
      * Push what channels would be used for up the stack
      * Channels are rarely parameters to functions
      * E.g. stdlib has very few channels in the public API
    - Goroutines
      * Some libraries use New() to spawn goroutines, but this is not ideal, stdlib uses .Serve() functions that run inside goroutines
      * Push goroutines up the stack and provide .Serve() functions that allow callers to create goroutines
      * If goroutines/threads are absolutely necessary, .Close() should end all daemon threads/goroutines
    - context.Context
      * All blocking/long operations should be cancellable
      * Generally shouldn’t store context.Context
      * context.Value() should inform, not control, a program
        - i.e. don’t store synchronization primitives or other control structures in context
  * If something is hard to do, make someone else do it
    - Push problems up the stack
      * Logging
      * Goroutines
      * Locking
    - Keep things simple
  * Designing for efficiency
    - Correctness still trumps efficiency
    - Minimizing memory alloc is usually first priority
    - Avoid creating an API that forces memory allocations, e.g.:
      - forces an alloc since the function returns a byte array
      ```
      func Encrypt(... ) []byte
      ```
      - does not force an alloc and puts responsibility of allocating a `Writer` on the caller
      ```
      func WriteTo(Writer... )
      ```
  * Don’t vendor other libraries in a library
    - Try to depend on only `stdlib` and no third party libraries

----------
### Go for Data Science
----------
  * Python has problems with
    - Integrity of application – python is permissive in how input are taken and do dynamic type conversion that result in unexpected behavior when inputs are not formatted as expected, while go will error out when such things happen
    - Integrity of deployment – python requires complex execution environments with dependent libraries and associated versions while go removes these due to the compiled nature of go applications
  * Go has performance benefits over Python
  * Some tools to help data science with Go
    - Pachyderm – helps the process of ingesting data and provides ability to version data sets
    - github.com/gonum – provides packages to help with arithmetic and visualization
    - github.com/sajari/regression – provides packages for regression analysis

----------
### Go Guru
----------
  * golang.org/x/tools/cmd/guru
  * Provides simple command line utility to gain insight in to go code
  * Can perform pointer analysis to provide insight on what concrete or dynamic types a pointer may point to
