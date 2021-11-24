package singleton

import "sync"

var (
	once sync.Once
	instance *singleton
)

type singleton struct {
	Title string
}

func GetInstance(title string) *singleton {
	once.Do(func() {
		instance = &singleton{Title:title}
	})
	return instance
}