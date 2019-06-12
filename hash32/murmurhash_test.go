package hash32

import (
	"log"
	"os"
	"strconv"
	"testing"
)

func TestMurmurhash(t *testing.T) {
	b := []byte("31000151032")
	x86_32 := Murmurhash3_x86_32(b, len(b), 0)
	log.Printf(">%d",x86_32%32)
}

func TestMurmurhash2(t *testing.T) {
	b2 := []byte("teste2")
	x86_32 := Murmurhash3_x86_32(b2, len(b2), 0)
	log.Printf(">%d", x86_32%32)
}

func TestMurmurhash3(t *testing.T) {
	start := 41000133040
	n := 0

	fd,_:=os.OpenFile("a.txt",os.O_RDWR|os.O_CREATE|os.O_APPEND,0644)
	for i:= start; i < 99999999999; i++ {
		itoa := strconv.Itoa(i)
		b2 := []byte(itoa)
		x86_32 := Murmurhash3_x86_32(b2, len(b2), 0)
		if x86_32%32 == 0 {
			n++

			b2 = []byte(itoa+"\r\n")
			fd.Write(b2)
		}

		if n == 10000000 {
			fd.Close()
			break
		}
	}
}

