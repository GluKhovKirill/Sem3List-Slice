package model

type Storage interface {
	Add(data any) (index int64)
	Get(index int64) (data any)
	Delete(index int64) (ok bool)
	String() string
	SortIncrease(more func(i, j any) bool)
	SortDecrease(more func(i, j any) bool)
	Print()
}
