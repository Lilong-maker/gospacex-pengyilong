package model

import "gorm.io/gorm"

type Goods struct {
	gorm.Model
	Name  string  `gorm:"type:varchar(30)"`
	Price float64 `gorm:"type:decimal(10,2)"`
	Num   int     `gorm:"type:int(11)"`
}

func (g Goods) GoodsAdd(db *gorm.DB) error {
	return db.Debug().Create(&g).Error
}

func (r Goods) FindGoods(db *gorm.DB, name string) error {
	return db.Debug().Where("name = ?", name).Find(&r).Error
}

func (g *Goods) FindGoodsById(db *gorm.DB, id any) interface{} {
	return db.Debug().Where("id = ?", id).First(&g).Error
}

type Order struct {
	gorm.Model
	OrderNo    string  `gorm:"type:varchar(30)"`
	UserID     int     `gorm:"type:int(11)"`
	TotalPrice float64 `gorm:"type:decimal(10,2)"`
	PayStatus  int     `gorm:"type:int;0 已下单 1 未下单"`
}

func (o *Order) OrderAdd(db *gorm.DB) interface{} {
	return db.Debug().Create(&o).Error
}

func (o *Order) OrderItemAdd(db *gorm.DB, items []*OrderItem) interface{} {
	return db.Debug().Create(items).Error
}

type OrderItem struct {
	gorm.Model
	OrderNo    string  `gorm:"type:varchar(30)"`
	GoodsID    int     `gorm:"type:int(11)"`
	GoodsName  string  `gorm:"type:varchar(30)"`
	GoodsPrice float64 `gorm:"type:decimal(10,2)"`
	Num        int     `gorm:"type:int(11)"`
}
