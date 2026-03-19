package request

type GoodsAdd struct {
	Name  string  `form:"name"  binding:"required"`
	Price float64 `form:"price" binding:"required"`
	Num   int     `form:"num" binding:"required"`
}
