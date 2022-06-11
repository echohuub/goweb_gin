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

func (sd *ShopDao) QueryShopList(longitude, latitude float64, keyword string) []model.Shop {
	var list []model.Shop
	if keyword == "" {
		err := sd.Where(" longitude > ? and longitude < ? and latitude > ? and latitude < ? and status = 1",
			longitude-50, longitude+50, latitude-50, latitude+50).Find(&list)
		if err != nil {
			return nil
		}
	} else {
		err := sd.Where(" longitude > ? and longitude < ? and latitude > ? and latitude < ? and name like ? and status = 1",
			longitude-50, longitude+50, latitude-50, latitude+50, keyword).Find(&list)
		if err != nil {
			return nil
		}
	}
	return list
}

func (sd *ShopDao) QueryServiceByShopId(shopId int64) []model.Service {
	var services []model.Service
	err := sd.Table("service").Join("INNER", "shop_service", "service.id == shop_service.service_id and shop_service.shop_id = ?", shopId).Find(&services)
	if err != nil {
		return nil
	}
	return services
}
