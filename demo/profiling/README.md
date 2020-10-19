# Workshop :: Performacne testing and profiling

## Step 1 :: Preparing for performance testing

Start web server

```
$go run web_server.go
```

Open url `http://localhost:8080/` in web browser

Install performance testing tool, using [go-wrk](https://github.com/tsliwowicz/go-wrk)

```
$go get github.com/tsliwowicz/go-wrk
$go-wrk

Usage: go-wrk <options> <url>
Options:
	-H 	 header line, joined with ';' (Default )
	-M 	 HTTP method (Default GET)
	-T 	 Socket/request timeout in ms (Default 1000)
	-body 	 request body string or @filename (Default )
	-c 	 Number of goroutines to use (concurrent connections) (Default 10)
	-ca 	 CA file to verify peer against (SSL/TLS) (Default )
	-cert 	 CA certificate file to verify peer against (SSL/TLS) (Default )
	-d 	 Duration of test in seconds (Default 10)
	-f 	 Playback file name (Default <empty>)
	-help 	 Print help (Default false)
	-host 	 Host Header (Default )
	-http 	 Use HTTP/2 (Default true)
	-key 	 Private key file name (SSL/TLS (Default )
	-no-c 	 Disable Compression - Prevents sending the "Accept-Encoding: gzip" header (Default false)
	-no-ka 	 Disable KeepAlive - prevents re-use of TCP connections between different HTTP requests (Default false)
	-redir 	 Allow Redirects (Default false)
	-v 	 Print version details (Default false)
```

## Step 2 :: Run performance testing

```
$go-wrk -d 5 http://localhost:8080

Running 5s test @ http://localhost:8080
  10 goroutine(s) running concurrently
236232 requests in 4.890416068s, 25.68MB read
Requests/sec:		48305.09
Transfer/sec:		5.25MB
Avg Req Time:		207.017µs
Fastest Request:	55.407µs
Slowest Request:	22.412925ms
Number of Errors:	0
```

## Step 3 :: Performance with [pprof](https://golang.org/pkg/runtime/pprof/)

Install pprof

```
$go get github.com/google/pprof
```

try_test.go

```
package main

import "testing"

var name string
var ok bool

func BenchmarkIsGopher(b *testing.B) {
	for i := 0; i < b.N; i++ {
		name, ok = isGopher("demo@golang.org")
	}
}

```

Run benchmark with CPU and Memory

```
$go test -bench=. -cpuprofile pprof.cpu -memprofile prof.mem
```

Visualize the analysis as a graph

```
$go tool pprof pprof.cpu
> top
> web

or

$pprof profiling.test pprof.cpu
$pprof profiling.test pprof.mem
```

Show in server mode

```
$pprof -http=:6060
```

## Step 4 :: Generating profiles from running HTTP servers

Add pprof package in `web_server.go`

```
import _ "net/http/pprof"
```

Start web server

```
$go run web_server.go
```

Performance testing again

```
$go-wrk -d 30 http://localhost:8080
```

Analyze ther performance of Web server

```
$pprof -seconds 5 http://localhost:8080/debug/pprof/profile

> top
> web

or

$pprof -http=:6060 http://localhost:8080/debug/pprof/profile
```
