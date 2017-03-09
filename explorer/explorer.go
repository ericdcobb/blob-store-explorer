package explore

import (
	"encoding/json"
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
	Format    string
	Before    string
	After     string
}

func (exploration *Exploration) visit(path string, f os.FileInfo, err error) error {
	// fmt.Printf("Visited: %s\n", path)
	if strings.HasSuffix(path, "properties") && !strings.Contains(path, "metrics") && !strings.Contains(path, "metadata") {
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
			props.Set("filePath", path)
			exploration.Collected = append(exploration.Collected, props)
		}
	}
	return nil
}

func (exploration *Exploration) include(props *properties.Properties) bool {

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

	if len(exploration.Before) > 0 && !IsBefore(exploration.Before, props.GetString("creationTime", "0000000000000")) {
		return false
	}

	if len(exploration.After) > 0 && !IsAfter(exploration.After, props.GetString("creationTime", "0000000000000")) {
		return false
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
	fmt.Printf("Exploring %s\n", exploration.Path)
	filepath.Walk(exploration.Path, exploration.visit)
	fmt.Println("Stats:")
	fmt.Println(exploration.Stats)
	if exploration.Collect {
		if exploration.Format == "json" {
			allProps := make([]map[string]string, 0)
			for _, prop := range exploration.Collected {
				allProps = append(allProps, getMap(prop))
			}
			data, _ := json.MarshalIndent(allProps, "", "\t")
			fmt.Println(string(data))
		} else {
			for _, prop := range exploration.Collected {
				fmt.Println(prop)
			}
		}
	}
}

func getMap(prop *properties.Properties) map[string]string {
	m := make(map[string]string, 0)
	for _, key := range prop.Keys() {
		value, ok := prop.Get(key)
		if ok {
			m[key] = value
		}
	}
	return m
}

// Explore defines the Blob Store directory to explore
func Explore(path string, collect bool, filters []string, format string,
	beforeDate string, afterDate string) Exploration {
	return Exploration{
		Path:      path,
		Collect:   collect,
		Stats:     &Stats{0, 0, 0, 0},
		Collected: make([]*properties.Properties, 0),
		Filters:   filters,
		Format:    format,
		Before:    beforeDate,
		After:     afterDate}
}
