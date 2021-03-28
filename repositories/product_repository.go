package repositories

import (
	"database/sql"
	"go-product/common"
	"go-product/datamodels"
	"strconv"
)

//开发接口
//实现接口

type IProduct interface {
	//连接数据
	Conn() error
	Insert(product *datamodels.Product) (int64, error)
	Delete(int64) bool
	Update(product *datamodels.Product) error
	SelectByKey(int64) (*datamodels.Product, error)
	SelectAll() ([]*datamodels.Product, error)
}

type ProductManager struct {
	table     string
	mysqlConn *sql.DB
}

//func NewProductManager(table string, db *sql.DB) IProduct {
//	return &ProductManager{table: table, mysqlConn: db}
//}

// 数据连接
func (p *ProductManager) Conn() (err error) {
	if p.mysqlConn == nil {
		mysql, err := common.NewMysqlConn()
		if err != nil {
			return err
		}
		p.mysqlConn = mysql
	}
	if p.table == "" {
		p.table = "product"
	}
	return
}

//插入
func (p *ProductManager) Insert(product *datamodels.Product) (productId int64, err error) {
	if err = p.Conn(); err != nil {
		return
	}
	sql := "insert product set productName = ?, productNum = ?, productImage = ?, productUrl = ?"
	stmt, err := p.mysqlConn.Prepare(sql)
	if err != nil {
		return 0, err
	}
	result, err := stmt.Exec(product.ProductName, product.ProductNum, product.ProductImage, product.ProductUrl)
	if err != nil {
		return
	}
	return result.LastInsertId()
}

//删除
func (p *ProductManager) Delete(productId int64) bool {
	if err := p.Conn(); err != nil {
		return false
	}
	sql := "delete from product where id = ?"
	stmt, err := p.mysqlConn.Prepare(sql)
	if err == nil {
		return false
	}
	_, err = stmt.Exec(productId)
	if err != nil {
		return false
	}
	return true
}

//更新
func (p *ProductManager) Update(product *datamodels.Product) error {
	if err := p.Conn(); err != nil {
		return err
	}
	sql := "update product set productName = ?, productNum = ?, productImage = ?, productUrl = ?"
	stmt, err := p.mysqlConn.Prepare(sql)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(product.ProductName, product.ProductNum, product.ProductImage, product.ProductUrl)
	if err != nil {
		return err
	}
	return nil
}

// 根据上篇的id查询商品
func (p *ProductManager) SelectByKey(productId int64) (product *datamodels.Product, err error) {
	if err := p.Conn(); err != nil {
		return &datamodels.Product{}, err
	}
	sql := "select * from " + p.table + "where id = " + strconv.FormatInt(productId, 10)
	row, err := p.mysqlConn.Query(sql)
	if err != nil {
		return &datamodels.Product{}, err
	}
	result := common.GetResultRow(row)
	if len(result) == 0 {
		return &datamodels.Product{}, nil
	}
	common.DataToStructByTagSql(result, product)
	return
}
