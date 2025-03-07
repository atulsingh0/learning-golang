package main

// compiler does not know the size of the struct
// hence, it will use heap memory instead of stack memory

// OS stack memory size: 1 mb
// Go stack memory size: 2 kb
// GO will create a new continuous memory stack and copy the old stack to the new stack

func main() {

}
