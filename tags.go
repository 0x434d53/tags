package tags

import (
	"fmt"
	"strings"
)

// SEP is the seperator to split the string into tags
var SEP string = ","

// Tags allows the managements of tags in a single string be comma seperation
type Tags string

// Add the tags in the string seperated by SEP
func (t *Tags) Add(tags string) {
	t.AddSlice(strings.Split(tags, SEP))
}

// AddSlice adds the strings in tags to the Tags
func (t *Tags) AddSlice(tags []string) error {
	collect := make(map[string]struct{}, 0)
	for _, tag := range tags {
		tag = strings.TrimSpace(tag)
		if strings.Contains(tag, SEP) {
			return fmt.Errorf("An tag cannot contain the sepereator")
		}
		if tag != "" {
			collect[tag] = struct{}{}
		}
	}

	for _, tag := range t.AsSlice() {
		collect[tag] = struct{}{}
	}

	*t = Tags(strings.Join(setToSlice(collect), SEP))
	return nil
}

// AsSlice returns the tags in Tags as a slice of tags
func (t *Tags) AsSlice() []string {
	if len(string(*t)) > 0 {
		return strings.Split(string(*t), SEP)
	}
	return []string{}
}

// Count returns the number of tags
func (t *Tags) Count() int {
	return len(t.AsSlice())
}

// String returns the Tags as a string
func (t *Tags) String() string {
	return string(*t)
}

// Contains checks if the given tag is in Tags
func (t *Tags) Contains(tag string) bool {
	for _, x := range t.AsSlice() {
		if x == tag {
			return true
		}
	}
	return false
}

// Remove removes the tags (seperated by SEP) from Tags
func (t *Tags) Remove(tags string) {
	t.RemoveSlice(strings.Split(tags, SEP))
}

// RemoveSlice removes the tags given in the slice from Tags
func (t *Tags) RemoveSlice(tags []string) {
	collect := make(map[string]struct{}, 0)

	for _, tag := range t.AsSlice() {
		collect[tag] = struct{}{}
	}

	for _, tag := range tags {
		tag = strings.TrimSpace(tag)
		delete(collect, tag)
	}

	*t = Tags(strings.Join(setToSlice(collect), SEP))
}

// Clear removes all tags from Tags
func (t *Tags) Clear() {
	*t = Tags("")
}

func setToSlice(set map[string]struct{}) []string {
	slice := make([]string, 0, len(set))

	for x := range set {
		slice = append(slice, x)
	}
	return slice
}
