package content

import (
	"fmt"
	"strings"
)

// MetaData is note metadata struct.
type MetaData struct {
	Tag         string
	Description string
}

func (m *MetaData) GenerateHeader() string {
	var metadata = []string{
		"---",
	}

	if m.Tag != "" {
		metadata = append(metadata, fmt.Sprintf("tag: %s", m.Tag))
	}
	if m.Description != "" {
		metadata = append(metadata, fmt.Sprintf("description: %s", m.Description))
	}

	if len(metadata) == 1 {
		return ""
	} else {
		metadata = append(metadata, "---")
		return strings.Join(metadata, "\n")
	}
}
