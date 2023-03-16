package repository

var (
	UserRepo IUserRepo
)

func InitRepo() {
	UserRepo = NewUserRepo()
}
