package explore

import "fmt"

// Exploration to be explored
type Exploration struct {
	Path string
}

// Run executes the exploration
func (exploration *Exploration) Run() {
	fmt.Printf("exploring %s", exploration.Path)
}

// Explore defines the Blob Store directory to explore
func Explore(path string) Exploration {
	return Exploration{path}
}
