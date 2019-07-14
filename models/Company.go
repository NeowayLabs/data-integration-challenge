package models

import "time"

// Company modelo para armazenar empresa
type Company struct {
	ID        int       `form:"id" json:"id" xorm:"int pk autoincr notnull unique 'compan_id'"`
	Name      string    `form:"name" json:"name" xorm:"varchar(100) notnull unique 'compan_name'"`
	Zipcode   string    `form:"zip" json:"zip" xorm:"varchar(5) notnull 'compan_zipcode'"`
	Website   string    `form:"website" json:"website" xorm:"varchar(100) 'compan_website'"`
	CreatedAt time.Time `json:"-" xorm:"notnull created 'compan_created'"`
	UpdatedAt time.Time `json:"-" xorm:"notnull updated 'compan_updated'"`
	Version   int       `json:"-" xorm:"notnull version 'compan_version'"`
}

// TableName retorna nome fisico que quero para minha tabela
func (u *Company) TableName() string {
	return "companies"
}
