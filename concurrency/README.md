# Concurrency

Examples on Concurrency 

- Bank:  a simple Deposit - Withdrawal process.
Withdrawal fails if there are insufficient funds in the balance, and returns the result of withdrawal to the caller.

- Mutex Bank: same as _Bank_ but without using channels. Using only `sync.Mutex`