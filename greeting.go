package main

// Greet returns a greeting message for the provided name.
// If name is empty, it defaults to "World".
func Greet(name string) string {
	if name == "" {
		name = "World"
	}
	return "Hello, " + name + "!"
}
