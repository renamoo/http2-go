package main
import (
"bytes"
"fmt"
"golang.org/x/net/http2/hpack"
)
func main() {
	s := "www.example.com"
	// 圧縮前のサイズ
	fmt.Println(len(s))
	// 圧縮後のサイズ
	fmt.Println(hpack.HuffmanEncodeLength(s))
	// 圧縮後の文字列
	b := hpack.AppendHuffmanString(nil, s)
	fmt.Printf("%x\n", b)
	// 解答されたの文字列
	var buf bytes.Buffer
	_, err := hpack.HuffmanDecode(&buf, b)
	if err != nil {
	fmt.Printf("ERROR: %s", err)
	}
	fmt.Printf("%s\n", buf.String())
}