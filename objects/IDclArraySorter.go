package objects

type IDclArraySorter struct {
	Data IDclArray
}

func (I *IDclArraySorter) Len() int {
	return len(I.Data)
}

func (I *IDclArraySorter) Less(i, j int) bool {
	return I.Data[i].GetOrderId() < I.Data[j].GetOrderId()
}

func (I *IDclArraySorter) Swap(i, j int) {
	t := I.Data[i]
	I.Data[i] = I.Data[j]
	I.Data[j] = t
}
