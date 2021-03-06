package elastic

import (
	"encoding/json"
	"testing"
)

func TestFuzzyLikeThisQuery(t *testing.T) {
	q := NewFuzzyLikeThisQuery().Fields("name.first", "name.last").LikeText("text like this one").MaxQueryTerms(12)
	data, err := json.Marshal(q.Source())
	if err != nil {
		t.Fatalf("marshaling to JSON failed: %v", err)
	}
	got := string(data)
	expected := `{"fuzzy_like_this":{"fields":["name.first","name.last"],"like_text":"text like this one","max_query_terms":12}}`
	if got != expected {
		t.Errorf("expected\n%s\n,got:\n%s", expected, got)
	}
}
