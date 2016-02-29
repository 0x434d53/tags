package tags

import "testing"

func TestAddRemoveTag(t *testing.T) {
	tags := new(Tags)

	firstTag := "tag1"
	tags.Add(firstTag)
	if string(*tags) != firstTag {
		t.Fatalf("Tags should be |%v|, but is |%v|", firstTag, tags)
	}

	tags.Remove(firstTag)
	if string(*tags) != "" {
		t.Fatalf("Tags should be |%v|, but is |%v|", "", tags)
	}
}

func TestAddRemoveMany(t *testing.T) {
	tags := new(Tags)
	tags.Add("tag1,tag1,tag2,tag3")

	if !tags.Contains("tag1") || !tags.Contains("tag2") || !tags.Contains("tag3") {
		t.Fatalf("Tags does not contain all added elements, added |%v|, contains |%v|", "tag1,tag2,tag3", tags)
	}

	tags.Remove("tag1,tag2")

	if tags.Contains("tag1") || tags.Contains("tag2") {
		t.Fatalf("Tags does still contain the remove elements: Expected |%v|, Actucal: |%v|", "tag3", tags)
	}

	if !tags.Contains("tag3") {
		t.Fatalf("Tags should contain tag3, doesnt. Tags is |%v|", tags)
	}

	if tags.Count() != 1 {
		t.Fatalf("Tags should contain 1 tag, contains %v", tags.Count())
	}
}

func TestClear(t *testing.T) {
	tags := new(Tags)
	tags.Add("tag1")
	tags.Clear()
	if tags.Count() != 0 {
		t.Fatalf("Tags should be empty, has %v elements", tags.Count())
	}
}
