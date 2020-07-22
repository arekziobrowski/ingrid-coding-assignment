package route

import (
	"sort"
	"testing"
)

var dummyData = []Route{{
	Dest:     "1.1",
	Duration: 40,
	Distance: 20,
}, {
	Dest:     "2.2",
	Duration: 20,
	Distance: 10,
}, {
	Dest:     "3.3",
	Duration: 30,
	Distance: 10,
}}

var dummyDataWithEqualDuration = []Route{{
	Dest:     "1.1",
	Duration: 40,
	Distance: 20,
}, {
	Dest:     "2.2",
	Duration: 30,
	Distance: 10,
}, {
	Dest:     "3.3",
	Duration: 40,
	Distance: 30,
}}

func TestSort(t *testing.T) {
	tables := []struct {
		given []Route
		when  string
		then  []Route
	}{
		{dummyData, "asc", []Route{dummyData[1], dummyData[2], dummyData[0]}},
		{dummyData, "desc", []Route{dummyData[0], dummyData[2], dummyData[1]}},
		{[]Route{}, "asc", []Route{}},
		{[]Route{dummyData[0]}, "desc", []Route{dummyData[0]}},
		{dummyDataWithEqualDuration, "asc", []Route{dummyDataWithEqualDuration[1], dummyDataWithEqualDuration[0], dummyDataWithEqualDuration[2]}},
		{dummyDataWithEqualDuration, "desc", []Route{dummyDataWithEqualDuration[2], dummyDataWithEqualDuration[0], dummyDataWithEqualDuration[1]}},
	}

	for _, table := range tables {
		actual := table.given
		if table.when == "asc" {
			sort.Sort(ByDurationAndDistance(table.given))
		} else {
			sort.Sort(sort.Reverse(ByDurationAndDistance(table.given)))
		}
		for i := range actual {
			if actual[i] != table.then[i] {
				t.Errorf("Sorting in %s manner was incorrect, got: %v, want: %v.", table.when, actual, table.then)
			}
		}

	}
}
