package main; import "fmt"; var source = "package main; import \"fmt\"; var source = %#v; func main() { fmt.Printf(source, source) }\n"; func main() { fmt.Printf(source, source) }
