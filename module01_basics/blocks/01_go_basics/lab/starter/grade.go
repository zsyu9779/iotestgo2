package grade

import "errors"

var ErrScoreOutOfRange = errors.New("score must be between 0 and 100")

func Grade(score int) (string, error) {
	return "", nil
}
