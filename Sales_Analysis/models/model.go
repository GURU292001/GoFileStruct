package models

// Customer Details
type CustomerStruc struct {
	Id               int    `gorm:"primaryKey"`
	Customer_Id      string `gorm:"column:customer_id"`
	Customer_name    string `gorm:"column:customer_name"`
	Customer_email   string `gorm:"column:customer_email"`
	Customer_address string `gorm:"column:customer_address"`
	Category         string `gorm:"column:category"`
	CreatedDate      string `gorm:"column:createdDate"`
	UpdatedDate      string `gorm:"column:updatedDate"`
}

// Product Details
type ProductStruc struct {
	Id          int    `gorm:"primaryKey"`
	ProductId   string `gorm:"column:product_id"`
	ProductName string `gorm:"column:product_Name"`
	Category    string `gorm:"column:category"`
	CreatedDate string `gorm:"column:createdDate"`
	UpdatedDate string `gorm:"column:updatedDate"`
}

// Order Details
type OrderStruc struct {
	Id             int    `gorm:"primaryKey"`
	Order_Id       string `gorm:"column:order_id"`
	Product_Id     string `gorm:"column:product_id"`
	Customer_Id    string `gorm:"column:customer_id"`
	Date_of_sale   string `gorm:"column:date_of_sale"`
	Quantity_sold  int    `gorm:"column:quantity_sold"`
	Unit_price     int    `gorm:"column:unit_price"`
	Discount       int    `gorm:"column:discount"`
	Shipping_cost  int    `gorm:"column:shhipping_cost"`
	Payment_method string `json:"column:payment_method"`
	Region         string `json:"column:region"`
	Category       string `gorm:"column:category"`
	CreatedDate    string `gorm:"column:createdDate"`
	UpdatedDate    string `gorm:"column:updatedDate"`
}

type ResponseStruc struct {
	Status string `json:"status"`
	ErrMsg string `json:"errMsg"`
}
