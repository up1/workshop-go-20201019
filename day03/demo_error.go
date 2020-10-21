package main

type UsernameInvalidError struct {
	code    int
	message string
}

func (u UsernameInvalidError) Error() string {
	return "UsernameInvalidError"
}

func main() {
	e := UsernameInvalidError{}

}
