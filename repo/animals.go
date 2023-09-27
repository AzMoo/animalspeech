package repo

import "github.com/AzMoo/animalspeech/domain"

type CatRepo struct{}

func (CatRepo) Fetch(name string) domain.Cat {
	// the repo layer is responsible for constructing the domain object
	return domain.Cat{Name: name}
}

type DogRepo struct{}

func (DogRepo) Fetch(name string) domain.Dog {
	// the repo layer is responsible for constructing the domain object
	return domain.Dog{Name: name}
}
