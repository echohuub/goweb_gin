package dao

import (
	"goweb_gin/model"
	"goweb_gin/tool"
)

type ShopDao struct {
	*tool.Orm
}

func NewShopDao() ShopDao {
	return ShopDao{tool.DBEngine}
}

func (sd *ShopDao) QueryShopList(longitude, latitude float64) []model.Shop {
	var list []model.Shop
	sd.Where(" longitude > ? and longitude < ? and latitude > ? and latitude < ?",
		longitude-50, longitude+50, latitude-50, latitude+50).Find(&list)
	return list
}
