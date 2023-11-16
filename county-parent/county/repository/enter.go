package repository

type Repository struct {
	CountyRepository CountyRepo
}

var RepositoryInstance = new(Repository)
