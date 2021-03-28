package services

import (
	"encoding/base64"
	"encoding/binary"
	"time"

	"github.com/macduyhai/SmartHomeVer2/config"
	"github.com/macduyhai/SmartHomeVer2/daos"
	"github.com/macduyhai/SmartHomeVer2/dtos"
	"github.com/macduyhai/SmartHomeVer2/middlewares"
	"github.com/macduyhai/SmartHomeVer2/models"
)

type UserService interface {
	Login(request dtos.LoginRequest) (*dtos.LoginResponse, error)
	Create(user models.User) (*models.User, error)
	CreateLog(log models.Log) (*models.Log, error)
	//GetAverageMonth(request dtos.AvrMoneyPerMonthRequest) (*dtos.AvrMoneyPerMonthResponse, error)
	AnalysisByTag(userID int64, begin *time.Time, end *time.Time) (map[string]int64, error)
	AnalysisByDay(userID int64, begin *time.Time, end *time.Time) (map[string]int64, error)
	GetLogsByTime(userID int64, begin *time.Time, end *time.Time) ([]models.Log, error)
}

type userServiceImpl struct {
	config  *config.Config
	userDao daos.UserDao
	jwt     middlewares.JWT
}

func NewUserService(conf *config.Config, userDao daos.UserDao, jwt middlewares.JWT) UserService {
	return &userServiceImpl{config: conf,
		userDao: userDao,
		jwt:     jwt,
	}
}

func (service *userServiceImpl) Login(request dtos.LoginRequest) (*dtos.LoginResponse, error) {
	user, err := service.userDao.Login(request.Username, request.Password)
	if err != nil {
		return nil, err
	}
	array_byte := make([]byte, 8)
	binary.LittleEndian.PutUint64(array_byte, uint64(user.ID))
	// encode, err := RsaEncrypt([]byte(user.Username), config.PublicKey)
	encode, err := RsaEncrypt(array_byte, config.PublicKey)
	if err != nil {
		panic(err)
	}
	token_str := base64.StdEncoding.EncodeToString(encode) // convert int64  to []byte

	// tocken, err := service.jwt.CreateTocken(user.ID)
	tocken, err := service.jwt.CreateTockenPrivate(token_str)
	if err != nil {
		return nil, err
	}

	response := dtos.LoginResponse{
		UserID:   user.ID,
		Username: user.Username,
		Name:     user.Name,
		Tocken:   tocken,
		Money:    user.Money,
		Phone:    user.Phone,
	}
	return &response, nil
}

func (service *userServiceImpl) Create(user models.User) (*models.User, error) {
	return service.userDao.Create(user)
}

func (service *userServiceImpl) CreateLog(log models.Log) (*models.Log, error) {
	return service.userDao.CreateLog(log)
}

//func (service *userServiceImpl) GetAverageMonth(request dtos.AvrMoneyPerMonthRequest) (*dtos.AvrMoneyPerMonthResponse, error) {
//	return nil, nil
//}

func (service *userServiceImpl) AnalysisByTag(userID int64, begin *time.Time, end *time.Time) (map[string]int64, error) {
	logs, err := service.userDao.GetLog(userID, begin, end)
	if err != nil {
		return nil, err
	}

	m := make(map[string]int64)
	for _, item := range logs {
		if val, ok := m[item.Tag]; ok {
			m[item.Tag] = val + item.Money
		} else {
			m[item.Tag] = item.Money
		}
	}
	return m, nil
}

func (service *userServiceImpl) AnalysisByDay(userID int64, begin *time.Time, end *time.Time) (map[string]int64, error) {
	logs, err := service.userDao.GetLog(userID, begin, end)
	if err != nil {
		return nil, err
	}

	m := make(map[string]int64)
	for _, item := range logs {
		timeStr := item.UpdateAt.Format("2006-01-02")
		if val, ok := m[timeStr]; ok {
			m[timeStr] = val + item.Money
		} else {
			m[timeStr] = item.Money
		}
	}

	return m, nil
}

func (service *userServiceImpl) GetLogsByTime(userID int64, begin *time.Time, end *time.Time) ([]models.Log, error) {
	return service.userDao.GetLog(userID, begin, end)
}
