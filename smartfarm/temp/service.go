package sf_dashboard

import (
	"github.com/jjoykkm/ln-backend/common/models/model_other"
)

type Servicer interface {
	// status, uid string
	GetFarmList(status string, reqModel *model_other.ReqModel) (*model_other.RespModel, error)
	//status, farmId string
	GetFarmAreaDashboardList(status string, reqModel *model_other.ReqModel) (*model_other.RespModel, error)
}

type Service struct {
	repo  Repositorier
}

func NewService(repo Repositorier) Servicer {
	return &Service{
		repo:  repo,
	}
}

func (s *Service) GetFarmList(status string, reqModel *model_other.ReqModel) (*model_other.RespModel, error) {
	myFarm, err := s.repo.FindAllMyFarm(status, reqModel.User.Uid)
	if err != nil{
		return nil, err
	}
	return &model_other.RespModel{
		Item: myFarm,
		Total: len(myFarm),
	}, nil
}

func (s *Service) GetFarmAreaDashboardList(status string, reqModel *model_other.ReqModel) (*model_other.RespModel, error) {
	dashboardList, err := s.repo.FindAllFarmAreaDashboard(status, reqModel.FarmId)
	if err != nil{
		return nil, err
	}
	return &model_other.RespModel{
		Item: dashboardList,
		Total: len(dashboardList),
	}, nil
}
