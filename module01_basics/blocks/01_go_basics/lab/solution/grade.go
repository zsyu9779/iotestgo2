package grade

import "errors"

var ErrScoreOutOfRange = errors.New("score must be between 0 and 100")

func Grade(score int) (string, error) {
	if score < 0 || score > 100 {
		return "", ErrScoreOutOfRange
	}

	switch {
	case score >= 90:
		return "A", nil
	case score >= 80:
		return "B", nil
	case score >= 70:
		return "C", nil
	case score >= 60:
		return "D", nil
	default:
		return "F", nil
	}
}
