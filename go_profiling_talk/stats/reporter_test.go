package stats

import "testing"

func BenchmarkAddTagsToName(b *testing.B) {
	tags := map[string]string{
		"host":     "myhost",
		"endpoint": "hello",
		"os":       "OS X",
		"browser":  "Chrome",
	}
	for b.Loop() {
		addTagsToName("recv.calls", tags)
	}
}
