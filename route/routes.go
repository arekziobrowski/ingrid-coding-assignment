package route

type Route struct {
	Dest     string
	Duration float64
	Distance float64
}

type ByDurationAndDistance []Route

func (a ByDurationAndDistance) Len() int {
	return len(a)
}

func (a ByDurationAndDistance) Less(i, j int) bool {
	if a[i].Duration == a[j].Duration {
		return a[i].Distance < a[j].Distance
	} else {
		return a[i].Duration < a[j].Duration
	}
}

func (a ByDurationAndDistance) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
