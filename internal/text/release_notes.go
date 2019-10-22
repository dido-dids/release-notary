package text

import (
	"strings"
)

// ReleaseNotes holds the required settings for generating ReleaseNotes
type ReleaseNotes struct {
	Simple bool
}

// Generate generates the output mentioned in the expected-output.md
func (r *ReleaseNotes) Generate(sections map[string][]Commit) string {
	builder := strings.Builder{}
	// Extra lines at the start to make sure formatting starts correctly
	builder.WriteString("\n\n")

	if len(sections["features"]) > 0 {
		builder.WriteString(r.buildSection("features", sections["features"]))
	}

	if len(sections["bugs"]) > 0 {
		builder.WriteString(r.buildSection("bugs", sections["bugs"]))
	}

	if len(sections["chores"]) > 0 {
		builder.WriteString(r.buildSection("chores", sections["chores"]))
	}

	if len(sections["others"]) > 0 {
		builder.WriteString(r.buildSection("others", sections["others"]))
	}

	return builder.String()
}