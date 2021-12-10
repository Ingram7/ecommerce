package repository

type User interface {
	Select()
	//SelectMany(query Query, limit int) (results []datamodels.User)
	//
	//InsertOrUpdate(user datamodels.User) (updatedUser datamodels.User, err error)
	//Delete(query Query, limit int) (deleted bool)
}

func NewUser() User {
	return &user{}
}

type user struct {
}

func (repo *user) Select() {

}
