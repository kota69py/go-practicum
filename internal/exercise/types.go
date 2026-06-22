package exercise

type Exercise struct {
	Name       string   `json:"name"`
	Title      string   `json:"title"`
	Category   string   `json:"category"`
	Difficulty int      `json:"difficulty"`
	Topics     []string `json:"topics"`
	Hints      []string `json:"hints"`
	Files      []string `json:"files"`
}

type Progress struct {
	Completed  []string `json:"completed"`
	InProgress string   `json:"in_progress,omitempty"`
}
