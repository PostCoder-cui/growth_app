package service

import (
	"context"
	"user_growth/dao"
	"user_growth/models"
)

/*
数据服务
*/

type CoinDetailService struct {
	ctx           context.Context
	daoCoinDetail *dao.CoinDetailDao
}

func NewCoinDetailService(ctx context.Context) *CoinDetailService {
	return &CoinDetailService{ctx: ctx, daoCoinDetail: dao.NewCoinDetailDao(ctx)}
}

func (s *CoinDetailService) Get(id int) (*models.TbCoinDetail, error) {
	return s.daoCoinDetail.Get(id)
}
func (s *CoinDetailService) FindByUid(uid, page, size int) ([]models.TbCoinDetail, int64, error) {
	return s.daoCoinDetail.FindByUid(uid, page, size)
}
func (s *CoinDetailService) Save(data *models.TbCoinDetail, musDetails ...string) error {
	return s.daoCoinDetail.Save(data, musDetails...)
}
