package main

import "fmt"

func main() {
	container := BuildContainer()
	err := container.Invoke(func(s *Server) {
		s.Start()
	})

	if err != nil {
		panic(err)
	}

	fmt.Printf(container.String())
}
