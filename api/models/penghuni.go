package models

type Penghuni struct {
	ID           int    `json:"id" gorm:"type:int(11);column:ID;primaryKey;autoIncrement:true"`
	PenghuniName string `json:"penghuni_name" gorm:"type:varchar(255);column:PenghuniName"`
	Handphone    string `json:"handphone" gorm:"type:varchar(255);column:Handphone;uniqueIndex"`
	Kategori     int    `json:"kategori" gorm:"type:int(11);column:Level;comment:kategoriID"`
}
