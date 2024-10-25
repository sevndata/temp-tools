package sm3

import (
	"fmt"
	"testing"
)

func byteToString(b []byte) string {
	ret := ""
	for i := 0; i < len(b); i++ {
		ret += fmt.Sprintf("%02x", b[i])
	}
	fmt.Println("ret = ", ret)
	return ret
}
func TestSm3(t *testing.T) {
	msg := []byte("test")
	hw := New()
	hw.Write(msg)
	hash := hw.Sum(nil)
	fmt.Println(hash)
	fmt.Printf("hash = %d\n", len(hash))
	fmt.Printf("%s\n", byteToString(hash))
	hash1 := Sm3Sum(msg)
	fmt.Println(hash1)
	fmt.Printf("%s\n", byteToString(hash1))

}

func BenchmarkSm3(t *testing.B) {
	t.ReportAllocs()
	msg := []byte("test")
	hw := New()
	for i := 0; i < t.N; i++ {
		hw.Sum(nil)
		Sm3Sum(msg)
	}
}
