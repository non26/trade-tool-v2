package service

import entity "tradetoolv2/internal/app/domain/entity/okx"

func (o *okxFutureService) SetNewLeverage(data *entity.SetLeverageFuture) (*entity.SetLeverageFuture, error) {
	e, err := o.okxExtService.SetLeverage(
		data,
	)
	if err != nil {
		return nil, err
	}

	return e, nil
}
