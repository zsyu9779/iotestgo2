package solution

import "fmt"

type Relation struct{ AuthorID, PostID, TagID uint }

func ValidateRelation(r Relation) error {
	if r.AuthorID == 0 || r.PostID == 0 || r.TagID == 0 {
		return fmt.Errorf("relation keys must be non-zero")
	}
	return nil
}
