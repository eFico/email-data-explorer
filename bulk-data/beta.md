go mod init github.com/eFico/mandelbrot-trace
go mod tidy
go mod vendor


defer profile.Start(profile.TraceProfile, profile.ProfilePath(".")).Stop()
1-> build
2-> run
3-> go tool trace trace.out




defer profile.Start(profile.CPUProfile, profile.ProfilePath(".")).Stop()
1-> 
2-> build and run
3-> go tool pprof -http=:8080 cpu.pprof



defer profile.Start(profile.MemProfile, profile.ProfilePath("."), profile.MemProfileRate(1)).Stop()
1-> 
2-> build and run
3-> go tool pprof -http=:8080 mem.pprof