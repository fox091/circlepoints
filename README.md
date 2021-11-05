# circlepoints
Provides some functions to generate evenly distributed points pseudo-randomly in a circle using a few different techniques.

Heavily inspired by https://www.youtube.com/watch?v=4y_nmpv-9lI.

I wanted to determine the performance of some of these methods using Go and came to the same conclusion that the lame rejection method is the fastest.

Here are the results from testing on my gaming PC:

```
cpu: AMD Ryzen 7 5800X 8-Core Processor
BenchmarkGeneratePointSqrt-16           45206934                24.96 ns/op
BenchmarkGeneratePointRejection-16      63149917                17.08 ns/op
BenchmarkGeneratePointTriangle-16       34750174                34.29 ns/op
BenchmarkGeneratePointMax-16            34232604                33.93 ns/op
```

[![Go Reference](https://pkg.go.dev/badge/github.com/fox091/circlepoints.svg)](https://pkg.go.dev/github.com/fox091/circlepoints)
