package main

import (
	"flag"
	"log"
)

/*
Reflection in programming is the ability of a program to inspect and modify its own structure and
behavior at runtime. In Go, reflection is primarily handled by the reflect package, which allows
programs to dynamically inspect the types and values of variables. This capability is powerful
for developing generic functions that can operate on a wide variety of types, implementing
serialization and deserialization routines (e.g., for JSON or XML parsing), and building
frameworks that require type introspection without knowing the types at compile time.

Key uses of reflection include:

1)Identifying the type of an object at runtime.
2)Inspecting the structure of types, including fields, methods, and tags.
3)Dynamically calling methods on objects.
4)Working with interfaces to determine the underlying concrete types.

Reflection should be used judiciously, as it can make code harder to understand and maintain,
and it often comes with a performance cost compared to direct, type-safe operations.
*/
func main() {
	mode := flag.String("mode", "", "start in client or server mode")
	flag.Parse()

	switch *mode {
	case "server":
		StartServer()
	case "client":
		RunClient()
	default:
		log.Fatalf("Unknown mode %s", *mode)
	}
}

// runnign the pogram
// go run main.go -mode=server
// go run main.go -mode=client
// run them in seprate terminals
