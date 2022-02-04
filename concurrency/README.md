# Concurrency

Examples on Concurrency 

- Bank:  a simple Deposit - Withdrawal process.
Withdrawal fails if there are insufficient funds in the balance, and returns the result of withdrawal to the caller.

- Mutex Bank: same as _Bank_ but without using channels. Using only `sync.Mutex`

- Mutex Bank 2: same as _Mutex Bank_ but with using `RLock` and `RUnlock` while reading balances. This demonstrates the _Multiple readers, one writer_ principle.

- Racy: provides a concurrency problem which the _Race Detector_ will catch in runtime.

- TimeRace: Same as _Racy_

- Non Blocking Cache: URL Body fetcher, which memoizes the responses

- Ping-Pong: Measures the throughput of gouritnes. Data Race free.

## Notes

### _Race detector_
- `go build -race` builds the project so the race detector is included in runtime
- `go run -race` & `go test -race` detects race conditions in runtime
- Race detector is costly, it can use 10x of CPU power. Use it wisely.

### GOMAXPROCS=N
- `Go scheduler` parameter which sets a number (N) of maximum active OS threads for a number (M) of `go` routines. Remember, go has it's own scheduler that uses `m:n` scheduling. It multiplexes `M x go` to `N x OS Threads`

### Goroutines have no identity
- OS Threads have an identity, such as a pointer to a thread or an integer. This allows developers to build a `thread-local storage`, which is a global map keyed by thread identity.
- `go` routines have _NO IDENTITY_ by design. This is because `thread-local storage` is mostly abused by developers, in a way that function execution depends no on function arguments alone, but on the thread identity that executes it. This often leads to bugs and undefined behaviours.