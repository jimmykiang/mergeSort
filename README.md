# mergeSort

```bash
go test ./multithreadNaive -run=^$$ -bench=. -trace=trace.out

go tool trace trace.out

go test ./multithreadNaive -run=^$$ -bench=. -benchmem -cpuprofile cpu.prof -count 5

go tool pprof cpu.prof
    >top
```
