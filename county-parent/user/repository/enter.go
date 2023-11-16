package repository

type Repository struct {
	UserRepository UserRepo
}

var RepositoryInstance = new(Repository)
