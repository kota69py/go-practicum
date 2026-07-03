package main

func main() {
	// Should report: os.Exit in main function
	os.Exit(0)
}

func other() {
	// Should NOT report: not in main function
	os.Exit(1)
}
