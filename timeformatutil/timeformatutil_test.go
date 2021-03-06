package timeformatutil

import (
	"testing"
	"time"
)

func TestTimeEncoder(t *testing.T) {
	t1, t2 := TimeEncoder(time.Now())
	t.Log(t1, t2)
}

func TestTimeDecoder(t *testing.T) {
	t1, t2 := TimeEncoder(time.Now())
	t.Log(TimeDecoder(t1))
	t.Log(TimeDecoder(t2))
}

func TestWeekDateTimeInteger(t *testing.T) {
	t.Log(WeekDateTimeInteger(time.Now(), 'w'))
}
