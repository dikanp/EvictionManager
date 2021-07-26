package models

type EvictionManager interface {
	push(key string)
	pop() string
	clear()
}