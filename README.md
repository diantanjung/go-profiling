# go-profiling

# Build and run the go project 
- $ go build -o app_name
- $ ./app_name

# Command for stand alone
- Run the pprof 
  $ go tool pprof app_name cpu.prof
- List Top CPU Usage
  $ top N, top -cum, top -sum, top -flat
- Export graphical .svg
  $ web
- Seek To detail cpu usage Source Code
  $ list functionName
  $ top N, top -cum, top -sum, top -flat

# Command for http API. 
- Make huge request using Apache Benchmark tool
  $ ab -k -c 8 -n 100000 http://localhost:9090/hi
- start profiling 
  cpu             : $ go tool pprof app http://localhost:9090/debug/pprof/profile
  heap            : $ go tool pprof app http://localhost:9090/debug/pprof/heap
  goroutine       : $ go tool pprof app http://localhost:9090/debug/pprof/goroutine
  block           : $ go tool pprof app http://localhost:9090/debug/pprof/block
- List Top CPU Usage
  $ top N, top -cum, top -sum, top -flat
- Export graphical .svg
  $ web
- Seek To detail cpu usage Source Code
  $ list functionName
