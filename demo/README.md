### demo
```go
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
```
### output/results
```xml
<isomsg>
  <field id="0" value="0210"/>
  <field id="1" value="723A00010A808400"/>
  <field id="2" value="593600141001099999"/>
  <field id="3" value="011000"/>
  <field id="4" value="000010000000"/>
  <field id="7" value="1007021533"/>
  <field id="11" value="000001"/>
  <field id="12" value="191533"/>
  <field id="13" value="1006"/>
  <field id="15" value="1007"/>
  <field id="32" value="565656"/>
  <field id="37" value="100609010224"/>
  <field id="39" value="00"/>
  <field id="41" value="00000901"/>
  <field id="49" value="360"/>
  <field id="54" value="100236C0102240000000"/>
</isomsg>
<dump length="72"/>
  02 10 72 3A 00 01 0A 80 84 00 18 59 36 00 14 10 
  01 09 99 99 01 10 00 00 00 10 00 00 00 10 07 02 
  15 33 00 00 01 19 15 33 10 06 10 07 06 56 56 56 
  10 06 09 01 02 24 00 00 00 09 01 36 00 20 10 02 
  36 C0 10 22 40 00 00 00 
<dump/>
<isomsg>
  <field id="0" value="0210"/>
  <field id="1" value="723A00010A808400"/>
  <field id="2" value="593600141001099999"/>
  <field id="3" value="011000"/>
  <field id="4" value="000010000000"/>
  <field id="7" value="1007021533"/>
  <field id="11" value="000001"/>
  <field id="12" value="191533"/>
  <field id="13" value="1006"/>
  <field id="15" value="1007"/>
  <field id="32" value="565656"/>
  <field id="37" value="100609010224"/>
  <field id="39" value="00"/>
  <field id="41" value="00000901"/>
  <field id="49" value="360"/>
  <field id="54" value="100236C0102240000000"/>
</isomsg>

real    0m0.006s
user    0m0.002s
sys     0m0.002s
```
*HW-CPU:[T5600](http://ark.intel.com/products/27254/Intel-Core2-Duo-Processor-T5600-2M-Cache-1_83-GHz-667-MHz-FSB)/RAM:3GB-DDR2
