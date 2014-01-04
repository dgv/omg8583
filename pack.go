package omg8583

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

func pack(row map[int]string, id int) (string, error) {
	p := t[id]
	switch {
	case p.Type == "a" ||
		p.Type == "an" ||
		p.Type == "ans" ||
		p.Type == "n":
		return row[id], nil
	case p.Type == "lln":
		length := len(row[id])
		if length > p.Length {
			length = p.Length
		}
		msg := strconv.Itoa(length)
		if length < 10 {
			msg = "0" + msg
		}
		return msg + row[id][0:length], nil
	case p.Type == "lllan":
		length := len(row[id])
		if length > p.Length {
			length = p.Length
		}
		msg := strconv.Itoa(length)
		if length < 10 {
			msg = "00" + msg
		} else if length < 100 {
			msg = "0" + msg
		}
		return msg + row[id], nil
	case p.Type == "b":
		var msg, bmp string
		inx := 0
		var a []int
		for i, _ := range row {
			a = append(a, i)
		}
		sort.Ints(a)
		for _, i := range a {
			if i > 1 {
				offset := i - inx - 1
				for j := 0; j < offset; j++ {
					bmp += "0"
				}
				bmp += "1"
				inx = i
			}
		}
		length := math.Ceil(float64(len(bmp))/(float64(p.Length*4))) * float64(p.Length*4)
		blength := float64(len(bmp))
		for i := 0; float64(i) < length-blength; i++ {
			bmp += "0"
		}
		if n, err := strconv.ParseInt(bmp, 2, 64); err == nil {
			msg = strconv.FormatInt(n, 16)
			return strings.ToUpper(msg), nil
		} else {
			return "", fmt.Errorf("error %v\npacking data from bit %d\npackager: %v", err, id, t[id])
		}
	default:
		fmt.Errorf("error, type %d not implemented")
	}
	return "", nil
}
