package model

type Article struct {
	Id      int       `json:"id" form:"id"`
	Cid     int       `json:"cid" form:"cid"`
	Title   string    `json:"title" form:"title"`
	Author  string    `json:"author" form:"author"`
	Content string    `json:"content" form:"content"`
	Hits    int       `json:"hits" form:"hits"`
	Ctime   time.Time `json:"ctime" form:"ctime"` // CreateTime
}

func ArticlPage(pi int, ps int) ([]Article, error){
	mod := make([]Article, 0 ,ps)
	err := DB.Select(&mod, `select * from article limit ?,?`, (pi-1)*ps,ps)
	return mod, err
}

func ArticleCount() (int, error) {
	var count int
	err := DB.Get(&count, "SELECT COUNT(id) as count FROM article")
	return count, err
}
