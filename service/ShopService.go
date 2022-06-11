package service

import (
	"goweb_gin/dao"
	"goweb_gin/model"
	"strconv"
)

type ShopService struct {
}

func (ss *ShopService) GetShopList(lon, lat string) []model.Shop {
	longitude, err := strconv.ParseFloat(lon, 10)
	if err != nil {
		return nil
	}
	latitude, err := strconv.ParseFloat(lat, 10)
	if err != nil {
		return nil
	}
	shopDao := dao.NewShopDao()
	return shopDao.QueryShopList(longitude, latitude)
}
