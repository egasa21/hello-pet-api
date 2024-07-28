package product_request

type ProductRequest struct {
	Name  string  `json:"name"`
	Price float32 `json:"price"`
	Stock int16   `json:"stock"`
}
