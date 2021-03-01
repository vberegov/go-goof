package main
import (
	"fmt"
	"golang.org/x/text/encoding/unicode"
)
func main() {
		myEncoder := unicode.UTF16(unicode.BigEndian, unicode.UseBOM)
        myEncoder.NewDecoder()        
        gmt.Println("hello world")
}
