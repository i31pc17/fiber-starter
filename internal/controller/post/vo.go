package post

type createPostRequest struct {
	Title string `query:"title"`
	Body  string `query:"body"`
}

type getAllPostsRequest struct {
	Offset int    `query:"offset"`
	Limit  int    `query:"limit"`
	Title  string `query:"title"`
}

type getPostDetailRequest struct {
	Id int `params:"id"`
}

type updatePostDetailRequest struct {
	Id    int    `params:"id"`
	Title string `query:"title"`
	Body  string `query:"body"`
}

type deletePostDetailRequest struct {
	Id int `params:"id"`
}
