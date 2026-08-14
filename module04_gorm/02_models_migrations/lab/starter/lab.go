package starter

import "errors"

var ErrNotImplemented = errors.New("not implemented")

type Relation struct{ AuthorID, PostID, TagID uint }

func ValidateRelation(Relation) error { return ErrNotImplemented }
