package explore

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/magiconair/properties"
)

// Exploration to be explored
type Exploration struct {
	Path      string
	Collect   bool
	Stats     *Stats
	Collected []*properties.Properties
	Filters   []string
}

func (exploration *Exploration) visit(path string, f os.FileInfo, err error) error {
	// fmt.Printf("Visited: %s\n", path)
	if strings.HasSuffix(path, "properties") && !strings.Contains(path, "metrics") {
		props := properties.MustLoadFile(path, properties.UTF8)

		if !exploration.include(props) {
			return nil
		}

		exploration.Stats.TotalBlobs++
		exploration.Stats.TotalSize += props.GetInt64("size", 0)
		if props.GetBool("deleted", false) {
			exploration.Stats.SoftDeleted++
			exploration.Stats.TotalSizeDeleted += props.GetInt64("size", 0)
		}
		if exploration.Collect {
			exploration.Collected = append(exploration.Collected, props)
		}
	}
	return nil
}

func (exploration *Exploration) include(props *properties.Properties) bool {
	if len(exploration.Filters) == 0 {
		return true
	}
	for _, val := range exploration.Filters {
		property := strings.Split(val, "=")
		matches, err := regexp.MatchString(property[1], props.GetString(property[0], ""))
		if err != nil {
			fmt.Println(err)
		}
		if !matches {
			return matches
		}
	}
	return true

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
	fmt.Printf("Exploring %s", exploration.Path)
	filepath.Walk(exploration.Path, exploration.visit)
	fmt.Println("Stats:")
	fmt.Println(exploration.Stats)
	if exploration.Collect {
		for _, prop := range exploration.Collected {
			fmt.Println(prop)
		}
	}
}

// Explore defines the Blob Store directory to explore
func Explore(path string, collect bool, filters []string) Exploration {
	return Exploration{
		Path:      path,
		Collect:   collect,
		Stats:     &Stats{0, 0, 0, 0},
		Collected: make([]*properties.Properties, 0),
		Filters:   filters}
}
