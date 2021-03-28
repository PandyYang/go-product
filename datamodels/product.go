package datamodels

type Product struct {
	ID           int64  `json:"id" sql:"id" go:"id"`
	ProductName  string `json:"ProductName" sql:"productName" go:"ProductName"`
	ProductNum   int64  `json:"ProductNum" sql:"productNum" go:"ProductNum"`
	ProductImage string `json:"ProductImage" sql:"productImage" go:"ProductImage"`
	ProductUrl   string `json:"ProductUrl" sql:"productUrl" go:"ProductUrl"`
}
