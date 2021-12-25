package sf_formula_plant

import (
	"github.com/jjoykkm/ln-backend/common/config"
	"github.com/jjoykkm/ln-backend/common/models/model_db"
	"github.com/jjoykkm/ln-backend/common/models/model_other"
	"github.com/jjoykkm/ln-backend/modelsOld/model_databases"
)

type Servicer interface {
	// Service for API
	// status, language string
	GetPlantCategoryList(status string, reqModel *model_other.ReqModel) (*model_other.RespModel, error)
	// status, plantTypeId, language string, reqModel.Offset int
	GetPlantCategoryItem(status string, reqModel *model_other.ReqModel) (*model_other.RespOffsetModel, error)
	// status, reqModel.Uid, plantId string, reqModel.Offset int
	GetPlantOverviewByPlant(status string, reqModel *model_other.ReqModel) (*model_other.RespOffsetModel, error)
	// status, reqModel.Uid, language string, reqModel.Offset int
	GetPlantOverviewFavorite(status string, reqModel *model_other.ReqModel) (*model_other.RespOffsetModel, error)
	// status, reqModel.Uid, language string, reqModel.Offset int
	GetPlantOfMine(status string, reqModel *model_other.ReqModel) (*model_other.RespOffsetModel, error)
	// status, formulaPlasntId, language string
	GetFormulaPlantDetail(status string, reqModel *model_other.ReqModel) (*model_other.RespModel, error)

	// Function
	GetRateScoreAndPeople(formulaPlant model_databases.FormulaPlant) (float32, int)

	//Upsert
	AddChangeFormulaPlant(reqModel *ForPlantUS) error
	AddChangeFertilizer(req *model_db.FertilizerUS) error

	//Update
	AddFavoriteForPlant(reqModel *model_db.FavForPlantUS) error

	//Delete
	RemoveFavoriteForPlant(reqModel *model_db.FavForPlantUS) error
}

type Service struct {
	repo  Repositorier
}

func NewService(repo Repositorier) Servicer {
	return &Service{
		repo:  repo,
	}
}

func (s *Service) GetPlantCategoryList(status string, reqModel *model_other.ReqModel) (*model_other.RespModel, error) {
	plantTypeList, err := s.repo.FindAllPlantType(status)
	if err != nil{
		return nil, err
	}

	return &model_other.RespModel{
		Item: plantTypeList,
		Total: len(plantTypeList),
	}, nil
}

func (s *Service) GetPlantCategoryItem(status string, reqModel *model_other.ReqModel) (*model_other.RespOffsetModel, error) {
	joinList, err := s.repo.FindAllPlantWithPlantType(status, reqModel.PlantTypeId, reqModel.Offset)
	if err != nil{
		return nil, err
	}

	//for idx, join := range joinList {
	//	switch reqModel.Language {
	//	case config.GetLanguage().En:
	//		join.PlantType.PlantTypeTH = ""
	//		join.Plant.PlantNameTH = ""
	//		join.Plant.PlantDescTH = ""
	//	case config.GetLanguage().Th:
	//		join.PlantType.PlantTypeEN = ""
	//		join.Plant.PlantNameEN = ""
	//		join.Plant.PlantDescEN = ""
	//	}
	//	join.Plant.TotalItem = int(s.repo.GetCountFormulaPlant(config.GetStatus().Active, join.Plant.PlantId.UUID.String()))
	//	// Delete PlantTypeId in struct Plant
	//	join.Plant.PlantTypeId = nil
	//	// Modify list
	//	joinList[idx] = join
	//}
	//
	total := len(joinList)
	currentOffset := reqModel.Offset + total

	return &model_other.RespOffsetModel{
		Item: joinList,
		Offset: currentOffset,
		Total: total,
	}, nil
}

func (s *Service) GetPlantOverviewByPlant(status string, reqModel *model_other.ReqModel) (*model_other.RespOffsetModel, error) {
	forPlant, err := s.repo.FindAllFormulaPlantByPlant(status, reqModel.PlantId, reqModel.Offset)
	if err != nil{
		return nil, err
	}
	// Get favorite formula plant
	_, favMap, _ := s.repo.FindAllFavForPlantId(status, config.GetResType().Map, reqModel.Uid)
	// Get planted formula plant
	_, plantedMap, _ := s.repo.FindAllPlantedForPlantId(status, config.GetResType().Map, reqModel.Uid)
	for idx, wa := range forPlant {
		// Check is favorite
		wa.IsFavorite = favMap[wa.FormulaPlant.FormulaPlant.FormulaPlantId.UUID.String()]
		// Check planted
		wa.IsPlanted = plantedMap[wa.FormulaPlant.FormulaPlant.FormulaPlantId.UUID.String()]
		forPlant[idx] = wa
	}
	total := len(forPlant)
	currentOffset := reqModel.Offset + total
	return &model_other.RespOffsetModel{
		Item: forPlant,
		Offset: currentOffset,
		Total: total,
	},nil
}

func (s *Service) GetPlantOverviewFavorite(status string, reqModel *model_other.ReqModel) (*model_other.RespOffsetModel, error) {
	forPlant, err := s.repo.FindAllFormulaPlantFavorite(status, reqModel.Uid, reqModel.Offset)
	if err != nil{
		return nil, err
	}
	// Get planted formula plant
	_, plantedMap, _ := s.repo.FindAllPlantedForPlantId(status, config.GetResType().Map, reqModel.Uid)
	for idx, wa := range forPlant {
		wa.IsFavorite = true
		// Check planted
		wa.IsPlanted = plantedMap[wa.FormulaPlant.FormulaPlant.FormulaPlantId.UUID.String()]
		forPlant[idx] = wa
	}
	total := len(forPlant)
	currentOffset := reqModel.Offset + total
	return &model_other.RespOffsetModel{
		Item: forPlant,
		Offset: currentOffset,
		Total: total,
	}, nil
}

func (s *Service) GetPlantOfMine(status string, reqModel *model_other.ReqModel) (*model_other.RespOffsetModel, error) {
	forPlant, err := s.repo.FindAllMyFormulaPlant(status, reqModel.Uid, reqModel.Offset)
	if err != nil{
		return nil, err
	}
	// Get favorite formula plant
	_, favMap, _ := s.repo.FindAllFavForPlantId(status, config.GetResType().Map, reqModel.Uid)
	// Get planted formula plant
	_, plantedMap, _ := s.repo.FindAllPlantedForPlantId(status, config.GetResType().Map, reqModel.Uid)
	for idx, wa := range forPlant {
		// Check is favorite
		wa.IsFavorite = favMap[wa.FormulaPlant.FormulaPlant.FormulaPlantId.UUID.String()]
		// Check planted
		wa.IsPlanted = plantedMap[wa.FormulaPlant.FormulaPlant.FormulaPlantId.UUID.String()]
		forPlant[idx] = wa
	}
	total := len(forPlant)
	currentOffset := reqModel.Offset + total
	return &model_other.RespOffsetModel{
		Item: forPlant,
		Offset: currentOffset,
		Total: total,
	}, nil
}

func (s *Service) GetFormulaPlantDetail(status string, reqModel *model_other.ReqModel) (*model_other.RespModel, error) {
	forPlantDetail, err := s.repo.FindAllFormulaPlantDetail(status, reqModel.FormulaPlantId)
	if err != nil{
		return nil, err
	}
	// Get favorite formula plant
	_, favMap, _ := s.repo.FindAllFavForPlantId(status, config.GetResType().Map, reqModel.Uid)
	// Get planted formula plant
	_, plantedMap, _ := s.repo.FindAllPlantedForPlantId(status, config.GetResType().Map, reqModel.Uid)
	for idx, wa := range forPlantDetail {
		// Check is favorite
		wa.IsFavorite = favMap[wa.FormulaPlant.FormulaPlant.FormulaPlantId.UUID.String()]
		// Check planted
		wa.IsPlanted = plantedMap[wa.FormulaPlant.FormulaPlant.FormulaPlantId.UUID.String()]
		forPlantDetail[idx] = wa
	}
	return &model_other.RespModel{
		Item: forPlantDetail,
		Total: len(forPlantDetail),
	}, nil
}

func (s *Service) GetRateScoreAndPeople(formulaPlant model_databases.FormulaPlant) (float32, int) {
	var rateScore float32
	var ratePeople int

	ratePeople = formulaPlant.Recommend1 + formulaPlant.Recommend2 + formulaPlant.Recommend3 + formulaPlant.Recommend4 + formulaPlant.Recommend5

	rateScore += float32(formulaPlant.Recommend1)
	rateScore += (float32(formulaPlant.Recommend2) * 2)
	rateScore += (float32(formulaPlant.Recommend3) * 3)
	rateScore += (float32(formulaPlant.Recommend4) * 4)
	rateScore += (float32(formulaPlant.Recommend5) * 5)
	rateScore = rateScore / float32(ratePeople)

	return rateScore, ratePeople
}

//-------------------------------------------------------------------------------//
//				 	   				Upsert
//-------------------------------------------------------------------------------//
func (s *Service) AddChangeFormulaPlant(reqModel *ForPlantUS) error {
	var (
		forPlantId *string
		err	error
	)

	//Prepare data before process
	reqModel.FormulaPlant.StatusId = config.GetStatus().Active
	tx := s.repo.Begin()
	//Upsert Formula Plant
	if (model_db.FormulaPlantUS{}) != reqModel.FormulaPlant {
		err, forPlantId = tx.UpsertFormulaPlant(&reqModel.FormulaPlant)
		if err != nil{
			tx.Rollback()
			return err
		}
	}
	//Upsert Sensor Value
	if len(reqModel.SensorValue) > 0 {
		if *forPlantId != "" {
			//Assign FormulaPlantId
			for idx, wa := range reqModel.SensorValue {
				wa.FormulaPlantId = reqModel.FormulaPlant.FormulaPlantId
				wa.StatusId = config.GetStatus().Active
				reqModel.SensorValue[idx] = wa
			}
			err := tx.UpsertForPlantSensor(reqModel.SensorValue)
			if err != nil{
				tx.Rollback()
				return err
			}
		}
	}
	//Upsert Fertilizer Ratio
	if len(reqModel.FertRatio) > 0 {
		if *forPlantId != "" {
			//Assign FormulaPlantId
			for idx, wa := range reqModel.FertRatio {
				wa.FormulaPlantId = reqModel.FormulaPlant.FormulaPlantId
				wa.StatusId = config.GetStatus().Active
				reqModel.FertRatio[idx] = wa
			}
			err := tx.UpsertForPlantFert(reqModel.FertRatio)
			if err != nil{
				tx.Rollback()
				return err
			}
		}
	}
	tx.Commit()

	return nil
}

func (s *Service) AddChangeFertilizer(req *model_db.FertilizerUS) error {
	req.StatusId = config.GetStatus().Active
	err := s.repo.UpsertFertilizer(req)
	if err != nil{
		return err
	}
	return nil
}
//-------------------------------------------------------------------------------//
//									Create
//-------------------------------------------------------------------------------//
func (s *Service) AddFavoriteForPlant(reqModel *model_db.FavForPlantUS) error {
	if (model_db.FavForPlantUS{}) == *reqModel  {
		//return err.ErrContext
	}
	err := s.repo.CreateFavFormulaPlant(reqModel)
	if err != nil{
		return err
	}
	return nil
}

//-------------------------------------------------------------------------------//
//									Delete
//-------------------------------------------------------------------------------//
func (s *Service) RemoveFavoriteForPlant(reqModel *model_db.FavForPlantUS) error {
	if (model_db.FavForPlantUS{}) == *reqModel  {
		//return err.ErrContext
	}
	err := s.repo.DeleteFavFormulaPlant(reqModel)
	if err != nil{
		return err
	}
	return nil
}