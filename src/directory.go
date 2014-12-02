package main


type directory struct {
	size  int64
	path  string
}

type directoryArray []directory

func (d directoryArray) Len() int {
	return len(d)
}

func (d directoryArray) Less(i, j int) bool {
	return d[i].size < d[j].size
}

func (d directoryArray) Swap(i, j int) {
	d[i],d[j] = d[j], d[i]
}


