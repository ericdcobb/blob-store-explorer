package explore

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/magiconair/properties"
)

// Exploration to be explored
type Exploration struct {
	Path  string
	Stats *Stats
}

func (exploration *Exploration) visit(path string, f os.FileInfo, err error) error {
	// fmt.Printf("Visited: %s\n", path)
	if strings.HasSuffix(path, "properties") && !strings.Contains(path, "metrics") {
		props := properties.MustLoadFile(path, properties.UTF8)
		exploration.Stats.TotalBlobs++
		exploration.Stats.TotalSize += props.GetInt64("size", 0)
		if props.GetBool("deleted", false) {
			exploration.Stats.SoftDeleted++
			exploration.Stats.TotalSizeDeleted += props.GetInt64("size", 0)
		}
	}
	return nil
}

// Stats about the exploration
type Stats struct {
	TotalBlobs       int64
	TotalSize        int64
	SoftDeleted      int64
	TotalSizeDeleted int64
}

func (stats *Stats) String() string {
	return fmt.Sprintf("Total blobs: %d, Total size: %d, Soft Deleted: %d, Total Size Deleted %d",
		stats.TotalBlobs, stats.TotalSize, stats.SoftDeleted, stats.TotalSizeDeleted)
}

// Run executes the exploration
func (exploration *Exploration) Run() {
	fmt.Printf("exploring %s \n", exploration.Path)
	filepath.Walk(exploration.Path, exploration.visit)
	fmt.Println(exploration.Stats)
}

// Explore defines the Blob Store directory to explore
func Explore(path string) Exploration {
	return Exploration{path, &Stats{0, 0, 0, 0}}
}
