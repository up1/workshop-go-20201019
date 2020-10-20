package demo

// SayHi :: SayHi
func SayHi() string {
	return "Hello sayHi"
}

func SayHi2() string {
	for i := 0; i < 10000; i++ {
		// fmt.Print(i)
	}
	return "Hello sayHi"
}
