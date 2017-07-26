package models

type Product struct {
	ID          int     `json:"pro_id"`
	Name        string  `json:"pro_name"`
	Description string  `json:"pro_description"`
	Price       float32 `json:"pro_price"`
	Image       string  `json:"pro_image"`
	Status      string  `json:"pro_status"`
}

type Products []Product

func (this *Product) Create(id int, name, description, price float32, image, status string) {
	sql := "INSERT INTO pro_id, pro_name, pro_description, pro_price, pro_image, pro_status from tb_product where pro_id=?"
	OpenConnection()
	Exec(sql, &id, &name, &description, &price, &image, &status)
	CloseConnection()
}

func (this *Product) Find(id int) {
	sql := "select pro_id, pro_name, pro_description, pro_price, pro_image, pro_status from tb_product where pro_id=?"
	OpenConnection()
	rows, _ := Query(sql, id)
	CloseConnection()
	for rows.Next() {
		rows.Scan(&this.ID, &this.Name, &this.Description, &this.Price, &this.Image, &this.Status)
	}
}

func (this *Products) FindAll() {
	sql := "select pro_id, pro_name, pro_description, pro_price, pro_image, pro_status from tb_product"
	OpenConnection()
	rows, _ := Query(sql)
	CloseConnection()

	for rows.Next() {
		var product Product
		rows.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.Image, &product.Status)
		*this = append(*this, product)
	}
}
