package tcp

import "strconv"

type EsbTcp struct {
}

func (et *EsbTcp) Hit(data [][]byte) (name []byte, ok bool) {
	header := data[0]
	hb := header[:8]
	if hl, err := strconv.Atoi(string(hb)); err == nil {
		num := 0
		for _, d := range data {
			num += len(d)
		}
		bl := num - 8
		if hl == bl {
			return []byte("esb"), true
		}
	}
	return nil, false
}
