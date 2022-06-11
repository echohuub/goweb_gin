package dao

import (
	"goweb_gin/model"
	"goweb_gin/tool"
)

type GoodsDao struct {
	*tool.Orm
}

func NewGoodsDao() GoodsDao {
	return GoodsDao{tool.DBEngine}
}

func (gd *GoodsDao) GetGoods(shopId int64) []model.Goods {
	var goods []model.Goods
	err := gd.Where("shop_id = ?", shopId).Find(&goods)
	if err != nil {
		return nil
	}
	return goods
}
