package main

import (
	"flag"
	"fmt"
	"org.springcat/murmurhash/hash32"
)

func main() {
	slice := flag.Uint("s", 0, "hash slice")
	value := flag.String("v", "", "hash string")

	flag.Parse()

	b2 := []byte(*value)

	x86_32 := hash32.Murmurhash3_x86_32(b2, len(b2), 0)

	u := (uint32)(*slice)

	if u == 0 {
		fmt.Println(x86_32)
	}else {
		fmt.Println(x86_32 % u)
	}

}
