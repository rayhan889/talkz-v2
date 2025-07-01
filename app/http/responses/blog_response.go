package responses

type (
	ComposeBlogRespone struct {
		ID      string `json:"id"`
		Title   string `json:"title"`
		Content string `json:"content"`
	}
	BlogsResponse struct {
		Blogs []BlogResponse `json:"blogs"`
	}
	BlogResponse struct {
		Title     string `json:"title"`
		Content   string `json:"content"`
		Author    string `json:"author"`
		CreatedAt string `json:"created_at"`
	}
)
