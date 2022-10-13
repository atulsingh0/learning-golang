package main

func main() {

	for i := 0; i < 10; i++ {
		go Square(i)
	}

	for i := 0; i < 10; i++ {
		go Cube(i)
	}
	select {}

}

func Square(n int) int {
	return n * n
}

func Cube(n int) int {
	return n * n * n
}
