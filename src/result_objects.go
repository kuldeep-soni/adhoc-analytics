package src

type IResultObject interface {
	ToJson() string
	ToDollarSV() string
	Reset()
}

type Result1 struct {
	Name string
	Id   string
}

func (r Result1) ToJson() string {
	panic("implement me")
}

func (r Result1) ToDollarSV() string {
	panic("implement me")
}

func (r *Result1) Reset() {
	panic("implement me")
}
