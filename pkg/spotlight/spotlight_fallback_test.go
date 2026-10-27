//go:build !darwin

package spotlight

import "testing"

func TestUnsupportedPlatform(t *testing.T) {
	for _, query := range []string{"", "readme"} {
		t.Run(query, func(t *testing.T) {
			opts := SearchOptions{Query: query, Scope: []string{"/"}, MaxResults: 1}
			const want = "spotlight search is unsupported on this platform (requires macOS)"

			results, err := Search(opts)
			if results != nil || err == nil || err.Error() != want {
				t.Errorf("Search(%+v) = %v, %v; want nil, %q", opts, results, err, want)
			}

			files, err := SearchWithMetadata(opts)
			if files != nil || err == nil || err.Error() != want {
				t.Errorf("SearchWithMetadata(%+v) = %v, %v; want nil, %q", opts, files, err, want)
			}
		})
	}
}
