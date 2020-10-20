# Workshop tracing with Go program

Step 1 :: Go code in `demo_01.go`

```
package main

import (
	"log"
	"os"
	"runtime/trace"
)

func main() {
	_ = trace.Start(os.Stdout)
	defer trace.Stop()
}
```

Run trace tool

```
$go run demo_01.go > trace.out
$go tool trace trace.out

or 
$go test <package name> -trace trace.out
$go tool trace trace.out
```
