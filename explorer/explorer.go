package explore

import (
	"fmt"
	"os"
	"path/filepath"
)

// Exploration to be explored
type Exploration struct {
	Path string
}

func (exploration *Exploration) visit(path string, f os.FileInfo, err error) error {
	fmt.Printf("Visited: %s\n", path)
	return nil
}

// Run executes the exploration
func (exploration *Exploration) Run() {
	fmt.Printf("exploring %s \n", exploration.Path)
	err := filepath.Walk(exploration.Path, exploration.visit)
	fmt.Printf("filepath.Walk() returned %v\n", err)
}

// Explore defines the Blob Store directory to explore
func Explore(path string) Exploration {
	return Exploration{path}
}
