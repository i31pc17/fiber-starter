package post

type Repository struct {
	rdb   interface{}
	redis interface{}
}

func NewPostRepository(RDB interface{}, Redis interface{}) *Repository {
	return &Repository{RDB, Redis}
}
