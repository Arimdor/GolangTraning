package models

type Entities interface {
	FindAll()
	Find()
}

type Model Entities
