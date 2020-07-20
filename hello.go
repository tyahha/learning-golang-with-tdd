package main

const englishHelloPrefix = "Hello, "

func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	return englishHelloPrefix + name
}

//func main() {
//	fmt.Println(Hello("world"))
//}
