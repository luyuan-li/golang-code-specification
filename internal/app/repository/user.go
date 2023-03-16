package repository

const (
	CollectionNameUser = "user"
)

type IUserRepo interface {
}

func NewUserRepo() IUserRepo {
	return &userRepo{ /*coll: cli.Database(database).Collection(CollectionNameUser)*/ }
}

type userRepo struct {
	//coll *qmgo.Collection
}
