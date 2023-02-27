package tools

import (
	"luckydraw-backend/internal/types"
	"luckydraw-backend/model"
)

func AwardsToModel(req *types.Awards) *model.Awards {
	if req == nil {
		return nil
	}
	return &model.Awards{
		Id:         req.Id,
		ActivityId: req.ActivityId,
		Uuid:       req.Uuid,
		Grade:      req.Grade,
		Name:       req.Name,
		Image:      req.Image,
		Price:      req.Price,
		Prob:       req.Prob,
		Quantity:   req.Quantity,
		Count:      req.Count,
		IsWin:      req.IsWin,
	}
}

func AwardsToResp(data *model.Awards) *types.Awards {
	if data == nil {
		return nil
	}
	return &types.Awards{
		Id:         data.Id,
		ActivityId: data.ActivityId,
		Uuid:       data.Uuid,
		Grade:      data.Grade,
		Name:       data.Name,
		Image:      data.Image,
		Price:      data.Price,
		Prob:       data.Prob,
		Quantity:   data.Quantity,
		Count:      data.Count,
		IsWin:      data.IsWin,
	}
}

func ActivityToModel(req *types.Activity) *model.Activity {
	if req == nil {
		return nil
	}
	return &model.Activity{
		Id:     req.Id,
		Uuid:   req.Uuid,
		Name:   req.Name,
		Status: req.Status,
	}
}

func ActivityToResp(data *model.Activity) *types.Activity {
	if data == nil {
		return nil
	}
	return &types.Activity{
		Id:     data.Id,
		Uuid:   data.Uuid,
		Name:   data.Name,
		Status: data.Status,
	}
}

func AdminerToModel(req *types.Adminer) *model.Adminer {
	if req == nil {
		return nil
	}
	return &model.Adminer{
		Id:       req.Id,
		Username: req.Username,
		Password: req.Password,
		Access:   req.Access,
		IsSuper:  req.IsSuper,
	}
}

func AdminerToResp(data *model.Adminer) *types.Adminer {
	if data == nil {
		return nil
	}
	return &types.Adminer{
		Id:       data.Id,
		Username: data.Username,
		Password: data.Password,
		Access:   data.Access,
		IsSuper:  data.IsSuper,
	}
}

func UserToModel(req *types.Users) *model.Users {
	if req == nil {
		return nil
	}
	return &model.Users{
		Id:              req.Id,
		ActivityId:      req.ActivityId,
		AvailableAwards: req.AvailableAwards,
		Name:            req.Name,
		Passport:        req.Passport,
		Count:           req.Count,
		Total:           req.Total,
	}
}

func UserToResp(data *model.Users) *types.Users {
	return &types.Users{
		Id:              data.Id,
		ActivityId:      data.ActivityId,
		AvailableAwards: data.AvailableAwards,
		Name:            data.Name,
		Passport:        data.Passport,
		Count:           data.Count,
		Total:           data.Total,
	}
}

func WinningRecordsToModel(req *types.WinningRecords) *model.WinningRecords {
	return &model.WinningRecords{
		Id:         req.Id,
		UserId:     req.UserId,
		AwardId:    req.AwardId,
		ActivityId: req.ActivityId,
		Ip:         req.Ip,
		Platform:   req.Platform,
	}
}

func WinningRecordsToResp(data *model.WinningRecords) *types.WinningRecords {
	if data == nil {
		return nil
	}
	return &types.WinningRecords{
		Id:         data.Id,
		UserId:     data.UserId,
		AwardId:    data.AwardId,
		ActivityId: data.ActivityId,
		Ip:         data.Ip,
		Platform:   data.Platform,
	}
}
