package categories

var Known = []string{
	"architecture",
	"concurrency",
	"language",
	"testing",
	"io",
	"net",
	"encoding",
	"configuration",
	"error-handling",
	"os",
	"performance",
	"security",
	"design",
	"database",
	"logging",
	"templating",
	"crypto",
	"observability",
	"datastore",
	"web",
}

func IsKnown(category string) bool {
	for _, k := range Known {
		if k == category {
			return true
		}
	}
	return false
}
