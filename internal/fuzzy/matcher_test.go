package fuzzy

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewMatcher(t *testing.T) {
	items := []string{"foo", "bar", "baz"}
	m := NewMatcher(items)

	assert.NotNil(t, m)
	assert.Equal(t, items, m.items)
}

func TestMatcher_Match(t *testing.T) {
	items := []string{
		"feature/auth",
		"feature/payment",
		"bugfix/login",
		"main",
		"develop",
	}
	m := NewMatcher(items)

	tests := []struct {
		name        string
		pattern     string
		wantCount   int
		wantFirst   string
		wantContain []string
	}{
		{
			name:      "empty pattern returns all items",
			pattern:   "",
			wantCount: 5,
		},
		{
			name:        "exact match",
			pattern:     "main",
			wantCount:   1,
			wantFirst:   "main",
			wantContain: []string{"main"},
		},
		{
			name:        "partial match",
			pattern:     "feat",
			wantCount:   2,
			wantContain: []string{"feature/auth", "feature/payment"},
		},
		{
			name:        "fuzzy match",
			pattern:     "fauth",
			wantCount:   1,
			wantFirst:   "feature/auth",
			wantContain: []string{"feature/auth"},
		},
		{
			name:      "no match",
			pattern:   "xyz123",
			wantCount: 0,
		},
		{
			name:        "case insensitive",
			pattern:     "MAIN",
			wantCount:   1,
			wantContain: []string{"main"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			matches := m.Match(tt.pattern)

			assert.Len(t, matches, tt.wantCount)

			if tt.wantFirst != "" && len(matches) > 0 {
				assert.Equal(t, tt.wantFirst, matches[0].Str)
			}

			if len(tt.wantContain) > 0 {
				strs := make([]string, len(matches))
				for i, m := range matches {
					strs[i] = m.Str
				}
				for _, want := range tt.wantContain {
					assert.Contains(t, strs, want)
				}
			}
		})
	}
}

func TestMatcher_BestMatch(t *testing.T) {
	items := []string{
		"feature/auth",
		"feature/payment",
		"bugfix/login",
	}
	m := NewMatcher(items)

	tests := []struct {
		name    string
		pattern string
		want    *string
	}{
		{
			name:    "returns best match",
			pattern: "auth",
			want:    strPtr("feature/auth"),
		},
		{
			name:    "returns nil for no match",
			pattern: "xyz123",
			want:    nil,
		},
		{
			name:    "empty pattern returns first item",
			pattern: "",
			want:    strPtr("feature/auth"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := m.BestMatch(tt.pattern)

			if tt.want == nil {
				assert.Nil(t, result)
			} else {
				require.NotNil(t, result)
				assert.Equal(t, *tt.want, result.Str)
			}
		})
	}
}

func TestMatch_Fields(t *testing.T) {
	items := []string{"hello", "world"}
	m := NewMatcher(items)

	matches := m.Match("hello")
	require.Len(t, matches, 1)

	match := matches[0]
	assert.Equal(t, 0, match.Index)
	assert.Equal(t, "hello", match.Str)
	assert.Greater(t, match.Score, 0)
	assert.NotEmpty(t, match.Indexes, "should have matched character indexes")
}

func strPtr(s string) *string {
	return &s
}
