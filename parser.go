package main

// Parser is the basic interface that needs to be implemented for generic handling
type Parser interface {
	GetName() string
	IsAvailable() (bool, error)
	GetURL() string
	GetShortURL() string
}
