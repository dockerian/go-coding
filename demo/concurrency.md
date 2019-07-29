# Golang Concurrency Programming

> Don't (let computations) communicate by sharing memory, (let them) share memory by communicating (through channels). -- Rob Pike


## Contents

  * [Goroutines](#goroutine)
  * [Concurrency programming](#concurreny)
  * [Channels](#channels)



<br/><a name="goroutine"></a>
## Goroutines

### Goroutine States

  * A live goroutine may stay in (and switch between) two states, "running" and "blocking".
  * Goroutines can **only** exit from running state, and never from blocking state.
    For any reason, if a goroutine stays in blocking state forever (which should
    be avoid, except for some rare case, in concurrent programming), then it will never exit.
  * A blocking goroutine can only be unblocked by an operation made in another goroutine.
  * At any given time, the maximum number of goroutines being executed will not
    exceed the number of the logical CPUs ([`runtime.NumCPU`](https://golang.org/pkg/runtime/#NumCPU)) available for the current program.
  * Go runtime adopts the M-P-G model.
    - **M** represents OS threads.
      * Each OS thread can only be attached to at most one goroutine at any given time.
    - **P** represents logical/virtual processors (not logical CPUs)
      * get and set the number by [`runtime.GOMAXPROCS`](https://golang.org/pkg/runtime/#GOMAXPROCS))
      * default initial value (equal to the number of logical CPUs, or `1` before Go 1.5) is the best choice for most programs.
      * a `GOMAXPROCS` value larger than `runtime.NumCPU()` may be helpful for some file IO heavy programs.
    - **G** represents goroutines.
      * A goroutine can only get executed when it is attached to an OS thread.
      * At any time, the number of goroutines in the executing sub-state is no
        more than the smaller one of `runtime.NumCPU()` and `runtime.GOMAXPROCS()`.

  ```
              ╭────────────────────────────────╮
              │           running              │
              │ ╭┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈╮ │
              │ ┊ sleeping, system call, ... ┊ │
              │ ╰┈┈┈┈┬┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈↑┈┈┈┈┈╯ │
              │      ┊                 ┊       │
              │ ╭┈┈┈┈↓┈┈┈┈╮      ╭┈┈┈┈┈┴┈┈┈┈┈╮ │
  (create) ┄┄┄┼→┊ queuing ┊ <==> ┊ executing ┊ ├┄┄┄┄→(exit)
              │ ╰┈┈┈┈↑┈┈┈┈╯      ╰┈┈┈┈┈┬┈┈┈┈┈╯ │
              ╰──────┼─────────────────┼───────╯
                     ┊                 ┊
              ┌──────┴─────────────────↓───────┐
              │           blocking             │
              └────────────────────────────────┘
  ```
  Goroutine states by [ASCII box drawing](https://en.wikipedia.org/wiki/List_of_Unicode_characters#Box_Drawing).



<br/><a name="concurreny"></a>
## Concurrency programming

### Data races

  In Go, generally, each computation is a goroutine.

  * At the time one computation is writing data to a memory segment, while
    another computation is reading data from the same memory segment,
    the integrity of the data read by the other computation might be not preserved.
  * At the time one computation is writing data to a memory segment, while
    another computation is also writing data to the same memory segment,
    the integrity of the data stored at the memory segment might be not preserved.

### Concurrent programming duties

  Most operations (including value assignments, argument passing and container
  element manipulations, etc) in Go are not synchronized (meaning "not concurrency-safe").
  Concurrent programming duties are to:

  * ensure concurrency/data synchronization: control resource sharing among
    concurrent computations, so that data races will not happen.
  * determine how many computations are needed.
  * determine when to start, block, unblock and end a computation.
  * determine how to distribute workload among concurrent computations.



<br/><a name="channels"></a>
## Channels

### Channel Basic

  * channels make goroutines share memory by communicating.
  * channel is first-class citizen in Go, comparing other traditional concurrency synchronization techniques, e.g. mutex.
  * a channel can be viewed as an internal FIFO (first in, first out) queue within a program. Some goroutines send values to the queue (the channel) and some other goroutines receive values from the queue.
  * a (logic) ownership view between goroutines:
    - goroutine releases the ownership of data by sending it to a channel.
    - goroutine acquires the ownership of data by receiving it from a channel.
  * critical techniques
    - use "Write Barrieres" to send data to destination (mem move by assembly code, very fast, not interrupted by garbage collection - creates a bitmap to force writes between the source and destination to happen in sequence).
    - use "CAS" (CompareAndSwap) operation (to solve "lazy barber" race condition).

### Channel Types and Values

  * use `chan<- T` to denote a send-only (producer) channel type.
  * use `<-chan T` to denote a receive-only (consumer) channel type.
  * use `chan T` to denote a bidirectional channel type, which can be implicitly converted to both send-only type `chan<- T` and receive-only type `<-chan T`, but not vice versa.
  * caution: there is only `<-`
  * each channel can have a capacity.
    - bufferred channel has non-zero capacity, e.g. `make(chan int, 10)`.
    - unbuffered channel has zero capacity, e.g. `make(chan int)`.
  * the size of channel element types must be smaller than 65536.
  * channel element values are transferred by copy -
    avoid cost on large value copy; use pointer type instead.
  * three categories of channels:
    - nil channels
    - non-nil but closed channels
    - non-closed non-nil channels
  * channel's (internal) data structure:
    - receiving goroutine queue (recvq) - a double linked list, w/o size limit,
      of blocked goroutines trying to read data from the channel.
    - sending gorouting queue (sendq) - a double linked list, w/o size limit,
      of blocked goroutines tyring to send data to the channel.
    - the value buffer queue - a circular queue, size as the channel capacity,
      or 0 for unbuffered channel.
    - a mutex lock to avoid data races in all kinds of operations.

### Channel Operations

  All the following operations are already synchronized:

  * Close the channel: `close(ch)`.
  * Send a value: `ch <- v` (for `ch` must be `chan T` or `chan<- T`).
  * Receive a value: `<- ch` (for `chan T` or `<-chan T`, and `v` is optional).
  * Query buffer capacity: `cap(ch)` (return 0 for `nil` channel).
  * Query the current number of values in the buffer: `len(ch)` (return 0 for `nil` channel).

  |operation   |a nil channel |a closed channel|a non-closed non-nil channel |
  |:-----------|:------------:|:--------------:|:---------------------------:|
  |close       | panic        | panic          | succeed                     |
  |send ch<-v  | block forever| panic          | block until `v` is received |
  |receive <-ch| block forever| never block    | block or succeed to receive |
  |cap(ch)     | 0            | capacity       | capacity                    |
  |len(ch)     | 0            | 0              | number of values in buffer  |

  Like most other operations in Go, channel value assignments are not synchronized. Similarly, assigning the received value to another value (`v, ok := <- ch`) is also not synchronized, though any channel receive operation is synchronized.

  * for-range:

    ```go
    for v := range ch { // for closed channel, otherwise blocking
    	// use v
    }
    ```
    equivalent to

    ```go
    for {
    	if v, ok := <-ch; ok {
    		// use v
    		continue
    	}
      break
    }
    ```

  * select-case

    ```go
    var chnil chan struct{}
    timeout := time.After(15 * time.Second) // <-chan Time
    Loop:
    for {
      select {
      // a case will only be considered if it is not nil.
      case v := <-ch:
        // use v
      case _, ok := <-chnil:
        if !ok {
          chnil = nil // to be ignored in next loop.
        }
      case <-timeout // <-chan Time
        break Loop
      default:
        // here if blocking on ch
      }
    }
    ```

    multiple channels

    ```go
    for i := range channels {
      select {
      case channels[i] <- value:
      default:
      }
    }
    ```

## References

  * https://go101.org/article/channel.html
  * https://blog.gauravagarwalr.com/posts/2019-01-23-summarizing-go-channels/
  * https://itnext.io/diving-into-golang-channels-e9e610d586e8
  * https://stackedco.de/understanding-golang-channels
