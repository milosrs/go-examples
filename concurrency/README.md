# Concurrency

Examples on Concurrency 

- Bank:  a simple Deposit - Withdrawal process.
Withdrawal fails if there are insufficient funds in the balance, and returns the result of withdrawal to the caller.

- Mutex Bank: same as _Bank_ but without using channels. Using only `sync.Mutex`

- Mutex Bank 2: same as _Mutex Bank_ but with using `RLock` and `RUnlock` while reading balances. This demonstrates the _Multiple readers, one writer_ principle.

- Racy: provides a concurrency problem which the _Race Detector_ will catch in runtime.

- TimeRace: Same as _Racy_

- Non Blocking Cache: URL Body fetcher, which memoizes the responses

## Notes

### _Race detector_
- `go build -race` builds the project so the race detector is included in runtime
- `go run -race` & `go test -race` detects race conditions in runtime
- Race detector is costly, it can use 10x of CPU power. Use it wisely.