package main

func train(in_order, out_order []int) bool {
	if len(in_order) != len(out_order) {
		return false
	}
	st := make([]int, 0, len(in_order))
	j := 0
	for _, v := range in_order {
		if len(st) != 0 && st[len(st)-1] == out_order[j] {
			st = st[:len(st)-1]
			j++
		}
		st = append(st, v)
	}
	if j != len(out_order) {
		for i := j; i < len(out_order); i++ {
			if st[len(st)-1] == out_order[i] {
				st = st[:len(st)-1]
			} else {
				return false
			}
		}
	}
	return len(st) == 0
}
