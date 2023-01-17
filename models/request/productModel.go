package request

type Product struct {
	Name    string
	Gambar  string
	Details []ProductDetail
}

type ProductDetail struct {
	Id     int
	Name   string
	Stock  int
	Gambar string
}
