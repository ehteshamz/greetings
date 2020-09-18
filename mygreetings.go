package greetings

var greetingsString = "Hello World"

func PrintGreetings(name string) string {
	return greetingsString + "-" + name
}
