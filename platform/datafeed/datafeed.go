package datafeed

//
/*
type Item struct {
	Title string `json:"title"`
	Post  string `json:"post"`
}*/
type Item struct {
	prodid int `json:"prodid"`
	prodname string `json:"prodname"`
	prodpic string `json:"prodpic"`
	price  int `json:"price"`
}
//
type Repo struct {
	Items []Item
}

//
func New() *Repo {
	return &Repo{}
}

//
func (r *Repo) Add(item Item) {
	r.Items = append(r.Items, item)
}

//
func (r *Repo) GetAll() []Item {
	return r.Items
}
