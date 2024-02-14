# `functional` - Go basic functional programming
This is a module I made to make it easy to do functional stuff in Go. It is just for my personal use, and I make no guarantees about stability.
- `Map`, `Do`, `Make`, `Filter`
- `PMap`, `PDo`, `PMake`, `PFilter` (parallel versions)
- `Pair[T,U]` and a bunch of methods around that
- `Must`, `Must2`, ... for panicking if the last return of the function is a non-nil error