package sf_my_farm

import (
	"errors"
	"github.com/jjoykkm/ln-backend/common/config"
	"github.com/jjoykkm/ln-backend/common/models/model_db"
	"github.com/jjoykkm/ln-backend/common/models/model_other"
	"github.com/jjoykkm/ln-backend/errs"
	"gorm.io/gorm"
)

type Servicer interface {
	// FarmId
	GetOverviewFarm(status string, reqModel *model_other.ReqModel) (*model_other.RespModel, error)
	// FarmId
	GetManageRole(status string, reqModel *model_other.ReqModel) (*model_other.RespModel, error)
	// FarmId
	GetManageFarmArea(status string, reqModel *model_other.ReqModel) (*model_other.RespModel, error)
	// FarmId
	GetManageMainbox(status string, reqModel *model_other.ReqModel) (*model_other.RespModel, error)

	CheckMainboxIsInactivated(serialNo string) (bool, error)
	ActivateMainbox(reqModel *model_db.MainboxSerialUS) error
	ConfigFarm(reqModel *ReqConfFarm) error
	ConfigMainbox(reqModel *ReqConfMainbox) error
	ConfigAddSensor(reqModel *ReqConfMainbox) error
	ConfigDeleteSocket(reqModel *ReqDeleteConfig) error
	ConfigDeleteMainbox(reqModel *ReqDeleteConfig) error
	ConfigDeleteFarm(reqModel *ReqDeleteConfig) error
	ConfigDeleteFarmArea(reqModel *ReqDeleteConfig) error
	ConfigFarmArea(reqModel *ReqConfFarmArea) error
	RemoveSocketLinkedFarm(reqModel *ReqRemoveLink) error
}

type Service struct {
	repo  Repositorier
}

func NewService(repo Repositorier) Servicer {
	return &Service{
		repo:  repo,
	}
}

func (s *Service) GetOverviewFarm(status string, reqModel *model_other.ReqModel) (*model_other.RespModel, error) {
	farm, err := s.repo.FindOneFarm(status, reqModel.FarmId)
	if err != nil{
		return nil, err
	}
	// Get Mainbox count
	farm.MainboxCount, err = s.repo.GetCountMainbox(status, reqModel.FarmId)
	if err != nil{
		return nil, err
	}
	// Get Farm area count
	farm.FarmAreaCount, err = s.repo.GetCountFarmArea(status, reqModel.FarmId)
	if err != nil{
		return nil, err
	}
	return &model_other.RespModel{
		Item: farm,
		Total: 1,
	}, nil
}

func (s *Service) GetManageRole(status string, reqModel *model_other.ReqModel) (*model_other.RespModel, error) {
	// Check auth for edit
	//isAuth, err := Servicer.GetAuthorizeCheckForManageFarm(s, reqModel.User.Uid, reqModel.FarmId)
	//if err != nil{
	//	return nil, err
	//}
	//// No Auth
	//if isAuth != true {
	//	return nil, &errs.ErrContext{
	//		Code: ERROR_4002005,
	//		Err:  err,
	//		Msg:  MSG_NO_AUTH,
	//	}
	//}

	manageRole, err := s.repo.FindAllManageRole(status, reqModel.FarmId)
	if err != nil{
		return nil, err
	}
	return &model_other.RespModel{
		Item: manageRole,
		Total: len(manageRole),
	}, nil
}

func (s *Service) GetManageFarmArea(status string, reqModel *model_other.ReqModel) (*model_other.RespModel, error) {
	manageFarmArea, err := s.repo.FindAllManageFarmArea(status, reqModel.FarmId)
	if err != nil{
		return nil, err
	}
	return &model_other.RespModel{
		Item: manageFarmArea,
		Total: len(manageFarmArea),
	}, nil
}

func (s *Service) GetManageMainbox(status string, reqModel *model_other.ReqModel) (*model_other.RespModel, error) {
	manageFarmArea, err := s.repo.FindAllManageMainbox(status, reqModel.FarmId)
	if err != nil{
		return nil, err
	}
	return &model_other.RespModel{
		Item: manageFarmArea,
		Total: len(manageFarmArea),
	}, nil
}

func (s *Service) CheckMainboxIsInactivated(serialNo string) (bool, error) {
	mainbox, err := s.repo.FindOneMainboxBySerialNo(serialNo)
	if err != nil{
		// Check serial no has been found in DB
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, &errs.ErrContext{
					Code: ERROR_4001002,
					Err:  err,
					Msg:  MSG_WRONG_MB,
				}
		}else {
			return false, err
		}
	}
	// Check serial no is inactive
	if mainbox.StatusId == config.GetStatus().Inactive {
		return true, nil
	}else {
		return false, &errs.ErrContext{
			Code: ERROR_4001001,
			Err:  err,
			Msg:  MSG_DUP_MB,
		}
	}
}

//-------------------------------------------------------------------------------//
//							Update data
//-------------------------------------------------------------------------------//
func (s *Service) ConfigFarm(reqModel *ReqConfFarm) error {
	// Prepare model before upsert data
	// Assign status active
	reqModel.Farm.StatusId = config.GetStatus().Active
	// Create Sensor
	err := s.repo.UpsertFarm(reqModel.Farm)
	if err != nil{
		return err
	}
	return nil
}

func (s *Service) ActivateMainbox(reqModel *model_db.MainboxSerialUS) error {
	isInactive, err := s.CheckMainboxIsInactivated(reqModel.MainboxSerialNo)
	if !isInactive || err != nil {
		return err
	}
	err = s.repo.UpdateOneMainboxBySerialNo(reqModel)
	if err != nil{
		return err
	}
	return nil
}

func (s *Service) ConfigMainbox(reqModel *ReqConfMainbox) error {
	// Update Mainbox detail
	tx := s.repo.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if !((model_db.MainboxUS{}) == *reqModel.Mainbox) {	// Check model has value
		err := tx.UpdateOneMainbox(reqModel.Mainbox)
		if err != nil{
			tx.Rollback()
			return err
		}
	}
	// Prepare model before upsert data
	for idx, _ := range reqModel.Socket {
		// Assign status active
		reqModel.Socket[idx].StatusId = config.GetStatus().Active
	}
	// Upsert Socket
	err := tx.UpsertSocket(reqModel.Socket)
	if err != nil{
		tx.Rollback()
		return err
	}
	return tx.db.Commit().Error
}

func (s *Service) ConfigAddSensor(reqModel *ReqConfMainbox) error {
	// Prepare model before upsert data
	// Assign status pending
	reqModel.Sensor.StatusId = config.GetStatus().Pending
	// Create Sensor
	err := s.repo.CreateOneSensor(reqModel.Sensor)
	if err != nil{
		return err
	}
	return nil
}

func (s *Service) ConfigFarmArea(reqModel *ReqConfFarmArea) error {
	// Prepare model before upsert data
	reqModel.FarmArea.StatusId = config.GetStatus().Active //Assign status active

	tx := s.repo.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	// Upsert FarmArea detail
	err, farmAreaId := tx.UpsertFarmArea(reqModel.FarmArea)
	if err != nil{
		return err
		tx.Rollback()
	}

	// Check array is not empty
	if len(reqModel.LinkedSocFarmArea.SocketId) > 0 {
		// Assign FarmAreaId before update data
		reqModel.LinkedSocFarmArea.FarmAreaId = *farmAreaId
		// Update Socket
		err = tx.UpdateAllSocket(reqModel.LinkedSocFarmArea)
		if err != nil{
			return err
			tx.Rollback()
		}
	}
	return tx.db.Commit().Error
}

//-------------------------------------------------------------------------------//
//							Delete data
//-------------------------------------------------------------------------------//
func (s *Service) ConfigDeleteSocket(reqModel *ReqDeleteConfig) error {
	// Delete Socket
	err := s.repo.DeleteOneSocket(reqModel.SocketId)
	if err != nil{
		return err
	}
	return nil
}

func (s *Service) ConfigDeleteMainbox(reqModel *ReqDeleteConfig) error {
	// Deactivate Mainbox
	err := s.repo.DeactivateOneMainbox(reqModel.MainboxId)
	if err != nil{
		return err
	}
	return nil
}

func (s *Service) ConfigDeleteFarm(reqModel *ReqDeleteConfig) error {
	// Delete Farm
	err := s.repo.DeleteOneFarm(reqModel.FarmId)
	if err != nil{
		return err
	}
	return nil
}

func (s *Service) ConfigDeleteFarmArea(reqModel *ReqDeleteConfig) error {
	// Delete FarmArea
	err := s.repo.DeleteOneFarmArea(reqModel.FarmAreaId)
	if err != nil{
		return err
	}
	return nil
}

func (s *Service) RemoveSocketLinkedFarm(reqModel *ReqRemoveLink) error {
	err := s.repo.UpdateAllSocketNullFarmArea(reqModel.LinkedSocFarmArea.SocketId)
	if err != nil{
		return err
	}
	return nil
}