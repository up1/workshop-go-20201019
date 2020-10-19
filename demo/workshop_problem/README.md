# Workshop :: find memory leak or cpu

```
$go test -run none -bench . -benchtime 3s -benchmem -cpuprofile cpu.out

$go tool pprof workshop_problem.test cpu.out
> top
> web
> top --cum
> list handleAdd

```

## Fix

```
var htmlTemplate = template.Must(template.ParseFiles("template.html"))
```

And run again

```
$go test -run none -bench . -benchtime 3s -benchmem -cpuprofile cpu.out
```
