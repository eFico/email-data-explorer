# email-data-explorer

## GO project init


```sh
go mod init
go mod tidy
go mod vendor
go build name
go run name
```


## Profiler method

```sh
Profiler method
defer profile.Start(profile.TraceProfile, profile.ProfilePath(".")).Stop()
1-> build
2-> ejecutar
3-> go tool trace trace.out


defer profile.Start(profile.MemProfile, profile.MemProfileRate(1)).Stop()
1-> build and run
2-> go tool pprof -http=:8080 cpu.pprof



defer profile.Start(profile.MemProfile, profile.MemProfileRate(1)).Stop()
1-> build and run
2-> go tool pprof -http=:8080 mem.pprof

```

## Vue 3


```sh
1-> npm i
2-> npm run dev
```
