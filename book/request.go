package book

type BookRequest struct {
	Title    string `binding:"required"`
	Price    int    `json:"price" binding:"required,number"`
	Description string `binding:"required"`
	Rating int `json:"rating" binding:"required,number"`
}

type UpdateBookRequest struct {
	Title string
	Price int
	Description string
	Rating int
}
