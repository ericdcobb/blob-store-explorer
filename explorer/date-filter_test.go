package explore

import (
	"fmt"
	"testing"
)

func TestBefore(t *testing.T) {
	var isBefore = IsBefore("2017-03-09T15:12:27-07:00", "1486675424191")
	assertTrue(t, isBefore)
}

func TestAfter(t *testing.T) {
	var isAfter = IsAfter("2017-01-09T15:12:27-07:00", "1486675424191")
	assertTrue(t, isAfter)
}

func assertTrue(t *testing.T, a bool) {
	if a {
		return
	}
	var message = fmt.Sprintf("%v is not true", a)
	t.Fatal(message)
}
