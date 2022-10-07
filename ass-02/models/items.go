package models

type Item struct {
	Item_id     uint `gorm:"primaryKey"`
	Item_code   uint
	Description string `gorm:"nit null;type:varchar(191)"`
	Quantity    uint
	Order_id    uint
}
