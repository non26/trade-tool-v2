package service

import (
	entity "tradetoolv2/internal/app/domain/entity/okx"
)

func (o *okxFutureService) PlaceAPosition(req *entity.PlaceSingleOrderEntity) (*entity.PlaceOrderEntity, error) {
	data, err := o.okxExtService.PlaceASinglePosition(req)
	if err != nil {
		return nil, err
	}

	return data, nil
}
