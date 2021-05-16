package analytics

import "time"

type ResultObjectExample struct {
	Time    time.Time
	Id      string
	Count   int
	Err     string
	Message string
}

func (b *ResultObjectExample) Set(currTime time.Time, id string, cnt int, err string, message string) {
	b.Reset()
	b.Time = currTime
	b.Id = id
	b.Count = cnt
	b.Err = err
	b.Message = message
}

func (b *ResultObjectExample) Reset() {
	b = &ResultObjectExample{}
}
