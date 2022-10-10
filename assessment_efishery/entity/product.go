package entity

type Products struct {
	ID        int    `json:"id" gorm:"primary_key"`
	Nama      string `json:"nama" gorm:"typevarchar(100);not_null;default:null"`
	Foto      string `json:"foto" gorm:"typevarchar(100);not_null;default:null"`
	Harga     int    `json:"harga" gorm:"int64;not_null;not_null"`
	Stok      int    `json:"stok" gorm:"int64"`
	Kategori  string `json:"kategori" gorm:"char(30);not_null;default:null"`
	Deskripsi string `json:"deskripsi" gorm:"varchar(400)"`
}

type CreateProductRequest struct {
	Nama      string `json:"nama"`
	Foto      string `json:"foto"`
	Harga     int    `json:"harga"`
	Stok      int    `json:"stok"`
	Kategori  string `json:"kategori"`
	Deskripsi string `json:"deskripsi"`
}

type UpdateProductRequest struct {
	Nama      string `json:"nama"`
	Foto      string `json:"foto"`
	Harga     int    `json:"harga"`
	Stok      int    `json:"stok"`
	Kategori  string `json:"kategori"`
	Deskripsi string `json:"deskripsi"`
}

type ProductResponse struct {
	ID       int    `json:"id"`
	Nama     string `json:"nama"`
	Foto     string `json:"foto"`
	Harga    int    `json:"harga"`
	Stok     int    `json:"stok"`
	Kategori string `json:"kategori"`
}

type DetailedProductResponse struct {
	Nama      string `json:"nama"`
	Foto      string `json:"foto"`
	Harga     int    `json:"harga"`
	Stok      int    `json:"stok"`
	Kategori  string `json:"kategori"`
	Deskripsi string `json:"deskripsi"`
}

type ProducttoCart struct {
	Nama  string `json:"nama"`
	Harga int    `json:"harga"`
}
