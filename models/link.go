package models

import (
	"github.com/astaxie/beego/orm"
)

//友情链接
type Link struct {
	Id       int64  `orm:"pk"  json:"id"`
	SiteName string `json:"siteName"`
	Url      string `json:"url"`
	Rank     int8   `json:"rank"`
}

func (m *Link) TableName() string {
	return TableName("link")
}

func (m *Link) Insert() error {
	if _, err := orm.NewOrm().Insert(m); err != nil {
		return err
	}
	Cache.Delete("links")
	return nil
}

func (m *Link) Read(fields ...string) error {
	if err := orm.NewOrm().Read(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *Link) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(m, fields...); err != nil {
		return err
	}
	Cache.Delete("links")
	return nil
}

//删除
func (m *Link) Delete() error {
	if _, err := orm.NewOrm().Delete(m); err != nil {
		return err
	}
	Cache.Delete("links")
	return nil
}

//表查询
func (m *Link) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(m)
}
