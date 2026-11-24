# Stock Span Problem

> Striver SDE Sheet 089 · Day 14 Stack & Queue II — the span ends at the previous strictly greater price; a monotonic stack keeps exactly those candidates. Go, stdlib only, tested.

## Problem

For each day's stock price, its *span* is the number of consecutive days ending today whose price was at most today's. `[100,80,60,70,60,75,85]` spans to `[1,1,1,2,1,4,6]`.

## Approach

The span is just the distance to the **previous strictly greater element** — so this is the previous-greater twin of Day 13's next-greater. Keep a stack of indexes whose prices strictly decrease upward; pop everything ≤ today's price (those days are inside today's span and, being smaller, can never terminate a future span either — safe to discard forever). The surviving top is the span's boundary; empty stack means the span reaches day one. Pushed once, popped once: amortized O(1) per day, which matters because spans are naturally computed online, day by day.

**Complexity:** O(n) time, O(n) stack.

## Run

    go test ./...   # verify
    go run .        # demo

---

One of 191 repos solving Striver's SDE Sheet in Go — stdlib only, one problem per repo, every solution tested. Browse the set: [`dsa-*` repos](https://github.com/Sreenivas-Sadhu-Prabhakara?tab=repositories&q=dsa-&type=source).
