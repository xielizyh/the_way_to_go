package main

// #include <stdlib.h>
// #include <stdio.h>
/*
int random(void) {
	return random();
}
*/
/*
void srandom(unsigned int i) {
	srandom(i);
}
*/
/*
void print(char *str) {
    printf("%s\n", str);
}
*/
import "C"

// Random example
func Random() int {
	return int(C.random())
}

// Seed example
func Seed(i int) {
	C.srandom(C.uint(i))
}

func main() {
	// Seed(1)
	// fmt.Println("The random is: ", Random())
	print("Hello World")
}
