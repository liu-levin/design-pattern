package Singleton

import "sync"

type Singleton struct {
}

var eagerSingleton *Singleton

func init() {
	eagerSingleton = &Singleton{}
}

func GetEagerInstance() *Singleton {
	return eagerSingleton
}

var (
	lazySingleton *Singleton
	once          = &sync.Once{}
)

func GetLazyInstance() *Singleton {
	if lazySingleton == nil {
		once.Do(func() {
			lazySingleton = &Singleton{}
		})
	}
	return lazySingleton
}
