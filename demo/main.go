// Demo in jPOS.org-ISO-8583-C-lightweight-library style
// http://www.jpos.org/products/iso_c_lib
package main

import (
	"fmt"
	"sort"

	"github.com/danielvargas/omg8583"
)

func main() {
	isomsg := make(map[int]string)

	isomsg[0] = "0210"
	isomsg[1] = "723A00010A808400"
	isomsg[2] = "593600141001099999"
	isomsg[3] = "011000"
	isomsg[4] = "000010000000"
	isomsg[7] = "1007021533"
	isomsg[11] = "000001"
	isomsg[12] = "191533"
	isomsg[13] = "1006"
	isomsg[15] = "1007"
	isomsg[32] = "565656"
	isomsg[37] = "100609010224"
	isomsg[39] = "00"
	isomsg[41] = "00000901"
	isomsg[49] = "360"
	isomsg[54] = "100236C0102240000000"

	isomsg_dump(isomsg)
	if data, err := omg8583.Pack(isomsg); err == nil {
		dumpbuf(data)
		if isomsg2, err := omg8583.Unpack(data); err == nil {
			isomsg_dump(isomsg2)
		} else {
			panic(err)
		}
	} else {
		panic(err)
	}
}

func isomsg_dump(msg map[int]string) {
	fmt.Println("<isomsg>")
	var keys []int
	for k := range msg {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	for _, k := range keys {
		fmt.Printf("  <field id=\"%d\" value=\"%s\"/>\n", k, msg[k])
	}
	fmt.Println("</isomsg>")
}

func dumpbuf(data string) {
	fmt.Printf("<dump length=\"%d\"/>\n  ", len(data)/2)
	i:=0
	for i< len(data) {        
		fmt.Printf("%s ",string(data[i:i+2]))
		i+=2
		if i%32==0 {
		fmt.Printf("\n  ")
		} 
	}
	fmt.Printf("\n<dump/>\n")
}
