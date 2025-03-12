package localSpike

type LocalSpike struct {
	LocalInStock     int64 //本地库存
	LocalSalesVolume int64 //已售出数量
}

// 本地扣库存,返回bool值
func (spike *LocalSpike) LocalDeductionStock() bool {
	spike.LocalSalesVolume = spike.LocalSalesVolume + 1
	return spike.LocalSalesVolume <= spike.LocalInStock
}
