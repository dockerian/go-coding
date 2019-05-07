# Golang Gotchas


## Contents

  * [Go](#lang)
  * [Common](#common)
  * [Defer](#defer)
  * [Pointer](#pointer)
  * [Nil](#nil)

## References

  * [Avoid gotchas](https://divan.github.io/posts/avoid_gotchas/)
  * [50 Shades of Go](http://devs.cloudimmunity.com/gotchas-and-common-mistakes-in-go-golang/)
  * [Golang gotcha](https://yourbasic.org/golang/gotcha/)
  * [Go format cheatsheet](https://yourbasic.org/golang/fmt-printf-reference-cheat-sheet/)
  * [Go bitwise](https://yourbasic.org/golang/bitwise-operator-cheat-sheet/)


<a name="lang"><br/></a>
## Language

  * Dependency management still in development.
  * Generic typed function (like sorting), or Template (e.g. in C# and Java).
  * Concat: `"H" + "i"` vs `'H' + 'i'`
  * Caution on `for x := range []int{3, 1, 4}` where `x` is the index.
  * Circumflex `x^y` denotes bitwise `XOR` (e.g. `1001 ^ 0011 == 1010`) in Go. Not the `math.Pow(x, y)`.
  * Diff: `strings.TrimSuffix()` vs `strings.TrimRight()`
  * Ensure a struct type, e.g. `SomeType`, implements an interface, e.g. `SomeIf`, at compile time:

    ```go
    var _ SomeIf = SomeType{} // verify SomeType implements SomeIf
    var _ SomeIf = (*SomeType)(nil) // verify *SomeType implements SomeIf
    ```

  * Go increment and decrement operations cannot be used as expressions,
    only as in statements, and only the postfix notation is allowed: `i++`.
  * Go strings are immutable and behave like read-only byte slices. Use `[]rune` instead.
  * Go **reuses** the XOR operator (`^`) as the unary "NOT" operator (aka
    bitwise complement), while many other languages use `~`.
  * Go uses a special "AND NOT" bitwise operator (`&^`), as a special feature/hack
    to support "x AND (NOT y)"" without requiring parentheses.
  * In a multi-line slice, array or map literal, every line **must end with a comma**.
  * It may seem like Go supports multi-dimensional arrays and slices, but it doesn't.

  * JSON: marshal/encode/serialize vs unmarshal/decode/deserialize
  * JSON: `json.Marshal` does not produce private members in a struct.
    **note**: Marshal (serializing) vs Unmarshal (deserializing).
  * JSON encoder adds a newline (`\n`) character.

    ```go
    import (
    	"bytes"
    	"fmt"
    	"encoding/json"
    )
    data := map[string]string{"key": "value"}
    var buf bytes.Buffer
    json.NewEncoder(&buf).Encode(data)
    raw, _ := json.Marshal(data)
    fmt.Printf("%s != %s\n", raw, buf.String())
    ```
  * JSON package escapes special HTML characters in keys and string values.
    - use `SetEscapeHTML(false)` but it is not available for `json.Marshel`.
    - such HTML encoding is NOT sufficient to protect against XSS vulnerabilities in all web applications.
    - assume the primary use case for JSON is web page, which breaks the
      configuration libraries and the REST/HTTP APIs by default.

      ```go
      data := "x < y"
      raw, _ := json.Marshal(data) // "<" becomes "\u003c"
      var buf bytes.Buffer
      encoder := json.NewEncoder(&buf)
      encoder.SetEscapeHTML(false) // important to keep "&", "<", and ">"
      encoder.Encode(data)
      ```

  * The function `regexp.MatchString` (as well as most functions in the
    regexp package) does substring matching.
  * Use `vet -shadow` to detect shadowed variables.
  * Overflow integer.


<a name="common"><br/></a>
## Common Gotchas

  * appending slice may or may not create a new instance copy; check `cap` first.

    ```go
    a := []byte("abc")   // `cap(a)` == 32
    a1 := append(a, 'd') // `a1` becomes "abcd"
    a2 := append(a, 'e') // `a1` and `a2` end up as same slice `[]bytes("abce")`
    ```
  * appending slice vs item: `append(a, slice...)` vs `append(a, item)`.
  * range expression is evaluated once before beginning the loop and a copy
    of the array is used to generate the iteration values.
    range a slice (e.g. `range a[:]`) instead.

    ```go
    var a [2]int
    for _, x := range a {
        fmt.Println("x =", x) // always printing `0`
        a[1] = 8
    }
    fmt.Println("a =", a)
    ```
  * map: using `make` to initialize a map before adding any elements.
  * slice: using `make` to allocate before calling `copy(dst, src)`.

  * no mixing of numeric types.

    ```go
    t := 200 // `t` is `int`; use `const` instead.
    time.Sleep(n * time.Millisecond) // won't compile
    ```

  * closing http connection
    * some HTTP servers keep network connections open for a while (based on the
      HTTP 1.1 spec and the server "keep-alive" configurations).
    * by default, the standard http library will close the network connections
      only when the target HTTP server asks for it -- app may run out of
      sockets/file descriptors under certain conditions.
    * ok to keep the network connection open if sending a lot of requests to the
      same HTTP server.
    * close the network connections right after your app receives the responses,
      or increase the open file limit if app sends 1+ requests to many different
      HTTP servers in a short period of time.
      - ask the http library to close the connection after the request is done
        by setting the `Close` field in the request variable to `true`.

        ```go
        req, err := http.NewRequest("GET", "http://golang.org", nil)
        if rep != nil {
            req.Close = true
        }
        ```
      - or add a `Connection` request header and set it to close. The target
        HTTP server should respond with a `Connection: close` header too. When the
        http library sees this response header it will also close the connection.

        ```go
        req.Header.Add("Connection", "close")
        ```
      - or disable http connection reuse globally.

        ```go
        tr := &http.Transport{DisableKeepAlives: true}
        client := &http.Client{Transport: tr}
        ```


<a name="defer"><br/></a>
## defer

  * defer functions are stacked (LIFO)
  * ensure non `nil` defer function (runtime panic)
  * ensure non `nil` resource before defer close function

    ```go
    func process() error {
      res, err := http.Get("http://notexists")
      if res != nil {
        defer res.Body.Close() // check res before defer
      }
      // continue with non-nil resource

      f, err := os.Open("book.txt")
      if err != nil {
        return err // check err before defer (assuming f is not nil)
      }
      defer func() {
        if ferr := f.Close(); ferr != nil {
          // log closing error
        }
      }()
      // continue with non-nil file
    }
    ```
  * ensure to close HTTP Response Body even if not to use/read.

    ```go
    resp, err := http.Get("https://reqres.in/api/users")
    if resp != nil {
        defer resp.Body.Close()
    }
    // read and discard the remaining response data in order to
    // reuse http connection for another request if the keepalive is enabled.
    // default HTTP client's Transport may not reuse HTTP/1.x "keep-alive" TCP
    // connections if the Body is not read to completion and closed.
    // see https://golang.org/pkg/net/http/#Response.
    _, err = io.Copy(ioutil.Discard, resp.Body)
    ```
  * call recover() always DIRECTLY inside a deferred func

    ```go
    func process() (err error) {
      defer func() {
        if r := recover(); r != nil { // called directly
          err = r.(error)
        }
      }()
      // some process will panic
    }
    ```
  * avoid unintended in-loop defer (since it is registered for the func not block)

    ```go
    func() {
      for {
        func() {
          // defer a func
          // some work
        }
      }
    }
    ```
  * defer does not need return but can change named result values

    ```go
    import errors
    func getNewError() (err error) {
      defer func() {
        err = errors.New("new error")
      }()
      return nil // it will return err instead of `nil`
    }
    ```
  * defer as a wrapper

    ```go
    import fmt
    type database struct{}
    func (db *database) connect() (disconnect func()) {
      fmt.Println("connect")
      return func() {
        fmt.Println("disconnect")
      }
    }
    func() {
      db := &database()
      closeFunc = db.connect()
      defer closeFunc()
      // ...
    }
    ```
  * defer method with pointers (https://play.golang.org/p/7ufFz3o8dAG)

    ```go
    import fmt
    type Foobar struct {
      name string
    }
    func (f Foobar) PrintInfo() {
      fmt.Println(f.name)
    }
    func (f *Foobar) PrintInfoByPointer() {
      fmt.Println(f.name)
    }
    func main() {
      o := Foobar{name: "origin"}
      defer fmt.Println(o.name)    // "origin" since `o` was evaluated ad hoc
      defer o.PrintInfoByPointer() // "change" since `o` was copied by pointer
      defer o.PrintInfo()          // "origin" since `o` was copied by value
      o.name = "change"
    }
    ```
  * defer in closure

    ```go
    for i := 0; i < 3; i++ {
      defer fmt.Println(i) // evaluated ad hoc
      a := i
      defer func() {
        fmt.Println(a) // copied inside closure
      }()
    }
    ```


<a name="function"><br/></a>
## function and go routine

  * function with variadic parameter

    ```go
    func fprint(format string, values ...interface{}) {
      fmt.Printf(format, values)    // wrong: only print a slice
      fmt.Printf(format, values...) // correct
    }
    ```

  * closure

    ```go
    var wg sync.WaitGroup
    for i := 0; i < 3; i++ {
      wg.Add(1)
      go func() {
        fmt.Println(i, "in go routine") // data race
        wg.Done()
      }()
    }
    wg.Wait()
    ```


<a name="interface"><br/></a>
## interface

  * interfaces are not pointers even though they may look like pointers.
    interface variable will be `nil` only when both type and value are `nil`.

    ```go
    var x interface{}
    fmt.Printf("x = [%5T, %+v], is nil : %v\n", x, x, x == nil)
    var p *int
    fmt.Printf("p = [%5T, %+v], is nil : %v\n", p, p, p == nil)
  	x = p
    fmt.Printf("x = [%5T, %+v], is nil : %v\n", x, x, x == nil)
    ```

<a name="pointer"><br/></a>
## pointer

  * for range pointers (https://play.golang.org/p/WGGJgIAQMT)

    ```go
    data := []string{"a", "b", "c"}
    copy := []*string{}
    for _, num := range data {
      copy = append(copy, &num) // adding the same pointer to the last item
    }
    ```

  * it's ok to return references to local variables (the Go compiler decides
    where the variable will be allocated even if the new() or make() functions
    are used), which is not ok in other languages like C or C++.

  * map elements are NOT addressable; (slice elements are addressable)

    ```go
    package main

    import (
    	"fmt"
    )
    type IFoobar interface {
      print()
    }
    type Data struct {
      v string
    }
    func (d *Data) print() {
      fmt.Println(d.v)
    }

    func main() {
      a1 := []Data{{"test"}}
    	a1[0].print()

      m1 := map[string]Data{"x1": {"x1-data"}}
      m1["x1"].print() // error: cannot take the address of m1["x1"]
      mv := m1["x1"] // assigning to another variable
    	mv.print() // ok

      m2 := map[string]*Data{"x2": &Data{"x2-data"}}
      m2["x2"].print() // ok
    }
    ```

  * passing by pointer is required

    ```go
    package main
    import (
      "fmt"
      "sync"
      "time"
    )
    func main() {
      wg = sync.WaitGroup{} // comparing `var wg sync.WaitGroup`
      defer func() {
        wg.Wait()
        fmt.Println("all done.")
      }()
      for i := 0; i < 3; i++ {
        wg.Add(1)
        doWork(&wg, fmt.Sprintf("do work #%d\n", i))
      }
    }
    func doWork(wg *sync.WaitGroup, params interface{}) {
      time.Sleep(250 * time.Millisecond)
      fmt.Printf("%+v\n", params)
      wg.Done()
    }
    ```


<a name="nil"><br/></a>
## nil

  * untyped `nil` is a compiler error: `var foobar = nil`
  * send and receive operations on a `nil` channel block forever.
  * string cannot be `nil`, but pointer to string (`*string`) can.
  * nil interface (`nil` implies an unset dynamic type `<type, nil>`)

    ```go
    func testNil(i interface{}) {
    	fmt.Printf("interface %+v == nil? %v\n", i, i == nil)
    }
    func equalToNil(i interface{}) {
      return i == nil
    }
    type foobar struct {
      f1 *int
      f2 *string
    }
    func test() {
      var i interface{}
      var f *foobar
      i = f
      fmt.Println(i == f)   // true
      fmt.Println(i == nil) // false
    }
    ```
