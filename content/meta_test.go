package content

import "testing"

func TestMetaData_GenerateHeader(t *testing.T) {
	var meta = MetaData{}

	// no metadata
	if meta.GenerateHeader() != "" {
		t.Fatal("Faild no metadata test")
	}

	// only tag
	meta.Tag = "sample tag"
	var tagHeader = `---
tag: sample tag
---`
	if meta.GenerateHeader() != tagHeader {
		t.Fatal("Faild tag only metadata test.")
	}

	// only description
	meta = MetaData{
		Description: "description",
	}
	var descriptionHeader = `---
description: description
---`
	if meta.GenerateHeader() != descriptionHeader {
		t.Fatal("Faild description only metadata test.")
	}

	// full data
	meta.Tag = "sample tag"
	var header = `---
tag: sample tag
description: description
---`
	if meta.GenerateHeader() != header {
		t.Fatal("Faild full metadata test.")
	}
}
