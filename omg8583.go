package omg8583

import "fmt"
import "sort"

type iso struct {
	Data     string
	Bitmap   string
	FieldIds []int
	Chunk    string
	RestData string
}

func Pack(isomsg map[int]string) (string, error) {
	var data string
	// http://blog.golang.org/go-maps-in-action#iteration_order
	var keys []int
	for k := range isomsg {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	for _, k := range keys {
		if msg, err := pack(isomsg, k); err == nil {
			data += msg
		} else {
			fmt.Errorf("%v", err)
		}
	}
	return data, nil
}

func Unpack(msg string) (isomsg map[int]string, err error) {
	isomsg = make(map[int]string)
	r, err := unpack(msg, 0)
	if err != nil {
		return nil, fmt.Errorf("Unpack: %v", err)
	}
	isomsg[0] = r.Data
	r, err = unpack(r.RestData, 1)
	if err != nil {
		return nil, fmt.Errorf("Unpack: %v", err)
	}
	isomsg[1] = r.Data
	ids := r.FieldIds
	chunk := r.RestData
	for _, i := range ids {
		if m, err := unpack(chunk, i); err == nil {
			chunk = m.RestData
			isomsg[i] = m.Data
		} else {
			return nil, fmt.Errorf("Unpack: %v", err)
		}
	}
	return
}
