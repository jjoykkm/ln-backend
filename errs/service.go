package errs

type Servicer interface {
	//GetFarmAreaAndName(farmAreaId string) (*model_db.FarmArea, string)
}

type Service struct {
	repo Repositorier
}

func NewService(repo Repositorier) Servicer {
	return &Service{
		repo:  repo,
	}
}

//func (s *Service) GetFarmAreaAndName(farmAreaId string) (*model_db.FarmArea, string) {
//	farmAreaModel, err := s.repo.FindOneFarmArea(config.GetStatus().Active, farmAreaId)
//	if err != nil {
//		return nil, ""
//	}
//	return farmAreaModel, farmAreaModel.FarmAreaName
//}
