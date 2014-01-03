// Packing
package omg8583

import (
	"fmt"
	"strconv"
)

func unpack(msg string, id int) (*iso, error) {
	p := t[id]
	switch {
	case p.Type == "a" ||
		p.Type == "an" ||
		p.Type == "ans" ||
		p.Type == "n":
		return &iso{Data: msg[0:p.Length],
			RestData: msg[p.Length:]}, nil
	case p.Type == "lln":
		if length, err := strconv.ParseInt(msg[0:2], 10, 64); err == nil {
			if length > int64(p.Length) {
				length = int64(p.Length)
			}
			return &iso{Data: string(msg[2 : length+2]),
				Chunk:    string(msg[0 : length+2]),
				RestData: string(msg[length+2:])}, nil
		} else {
			return nil, fmt.Errorf("error %v\npacking data from bit %d\npackager: %v", err, id, t[id])
		}
	case p.Type == "lllan":
		if length, err := strconv.ParseInt(msg[0:3], 10, 64); err == nil {
			if length > int64(p.Length) {
				length = int64(p.Length)
			}
			return &iso{Data: string(msg[3 : length+3]),
				Chunk:    string(msg[0 : length+3]),
				RestData: string(msg[length+3:])}, nil
		} else {
			return nil, fmt.Errorf("error %v\npacking data from bit %d\npackager: %v", err, id, t[id])
		}
	case p.Type == "b":
		var chunk = msg
		var data, bitmap string
		var chunkBitmap string
		chunk = msg[0:p.Length]
		var ckb = 1
		for ckb > 0 {
			if n, err := strconv.ParseInt(chunk, 16, 64); err == nil {
				chunkBitmap = strconv.FormatInt(n, 2)
			} else {
				return nil, fmt.Errorf("error %v\npacking data from bit %d\npackager: %v", err, id, t[id])
			}
			for len(chunkBitmap) < (p.Length * 4) {
				chunkBitmap = "0" + chunkBitmap
			}
			ckb, _ = strconv.Atoi(string(chunkBitmap[0]))
			data += chunk
			bitmap += chunkBitmap
			msg = msg[p.Length:]
		}
		fieldIds := make([]int, 0)
		for i, _ := range bitmap {
			b, _ := strconv.Atoi(string(bitmap[i]))
			if i > 0 && b == 1 {
				fieldIds = append(fieldIds, i+1)
			}
		}
		return &iso{
			Data:     data,
			Bitmap:   bitmap,
			FieldIds: fieldIds,
			RestData: msg}, nil
	default:
		return nil, fmt.Errorf("error, id %d not implemented\n", id)
	}
}
