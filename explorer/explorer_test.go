package explore

import (
	"fmt"
	"testing"
)

func TestFunctional_noCollect_noFilters(t *testing.T) {
	explorer := Explore("testdata", false, nil)
	explorer.Run()

	assertEqual(t, explorer.Stats.TotalBlobs, int64(31), "")
	assertEqual(t, explorer.Stats.TotalSize, int64(23984), "")
	assertEqual(t, explorer.Stats.SoftDeleted, int64(10), "")
	assertEqual(t, explorer.Stats.TotalSizeDeleted, int64(9737), "")

	assertEqual(t, len(explorer.Collected), 0, "")
}

func TestFunctional_noCollect_withFilters(t *testing.T) {
	explorer := Explore("testdata", false, []string{"@BlobStore.blob-name=pom"})
	explorer.Run()

	assertEqual(t, explorer.Stats.TotalBlobs, int64(11), "")
	assertEqual(t, explorer.Stats.TotalSize, int64(7987), "")
	assertEqual(t, explorer.Stats.SoftDeleted, int64(2), "")
	assertEqual(t, explorer.Stats.TotalSizeDeleted, int64(3108), "")

	assertEqual(t, len(explorer.Collected), 0, "")
}

func TestFunctional_withCollect_noFilters(t *testing.T) {
	explorer := Explore("testdata", true, nil)
	explorer.Run()
	assertEqual(t, len(explorer.Collected), 31, "")
}

func assertEqual(t *testing.T, a interface{}, b interface{}, message string) {
	if a == b {
		return
	}
	if len(message) == 0 {
		message = fmt.Sprintf("%v != %v", a, b)
	}
	t.Fatal(message)
}
