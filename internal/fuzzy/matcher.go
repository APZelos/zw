package fuzzy

import (
	"github.com/sahilm/fuzzy"
)

// Match represents a fuzzy match result.
type Match struct {
	Index   int
	Score   int
	Str     string
	Indexes []int
}

// Matcher provides fuzzy matching functionality.
type Matcher struct {
	items []string
}

// NewMatcher creates a new fuzzy matcher with the given items.
func NewMatcher(items []string) *Matcher {
	return &Matcher{items: items}
}

// Match returns items that match the given pattern, sorted by score.
func (m *Matcher) Match(pattern string) []Match {
	if pattern == "" {
		// Return all items with zero score
		matches := make([]Match, len(m.items))
		for i, item := range m.items {
			matches[i] = Match{
				Index: i,
				Score: 0,
				Str:   item,
			}
		}
		return matches
	}

	results := fuzzy.Find(pattern, m.items)
	matches := make([]Match, len(results))
	for i, r := range results {
		matches[i] = Match{
			Index:   r.Index,
			Score:   r.Score,
			Str:     r.Str,
			Indexes: r.MatchedIndexes,
		}
	}
	return matches
}

// BestMatch returns the best match for the given pattern, or nil if no match.
func (m *Matcher) BestMatch(pattern string) *Match {
	matches := m.Match(pattern)
	if len(matches) == 0 {
		return nil
	}
	return &matches[0]
}
