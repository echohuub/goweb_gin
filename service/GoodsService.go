package service

import (
	"goweb_gin/dao"
	"goweb_gin/model"
)

type GoodsService struct {
}

func (gs *GoodsService) GetGoods(shopId int64) []model.Goods {
	goodsDao := dao.NewGoodsDao()
	return goodsDao.GetGoods(shopId)
}
