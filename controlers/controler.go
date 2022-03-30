package controlers

import (
	"errors"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/macduyhai/SmartHomeVer2/common"

	"github.com/gin-gonic/gin"

	"github.com/macduyhai/SmartHomeVer2/dtos"
	"github.com/macduyhai/SmartHomeVer2/models"
	"github.com/macduyhai/SmartHomeVer2/services"
	"github.com/macduyhai/SmartHomeVer2/utilitys"
)

type Controller struct {
	userService   services.UserService
	deviceService services.DeviceService
	mediaService  services.MediaService
}

func NewController(provider services.Provider) Controller {
	return Controller{
		userService:   provider.GetUserService(),
		deviceService: provider.GetDeviceService(),
		mediaService:  provider.GetMediaService(),
	}
}

//-----------------------------------------------------------

//ListMedia
func (ctl *Controller) ListMedia(context *gin.Context) {
	var request dtos.ListMediaRequest
	err := context.ShouldBindJSON(&request)
	if err != nil {
		utilitys.ResponseError400(context, err.Error())
		return
	}
	// fmt.Println(request)
	data, err := ctl.mediaService.ListMedia(request)

	if err != nil {
		utilitys.ResponseError400(context, err.Error())
	} else {
		utilitys.ResponseSuccess200(context, data, "success")
	}
}

// Delete media
func (ctl *Controller) DeleteMedia(context *gin.Context) {
	var request dtos.DeleteMediaRequest

	err := context.ShouldBindJSON(&request)
	if err != nil {
		utilitys.ResponseError400(context, err.Error())
		return
	}
	data, err := ctl.mediaService.DeleteMedia(request)
	if err != nil {
		utilitys.ResponseError400(context, err.Error())
	} else {
		utilitys.ResponseSuccess200(context, data, "success")
	}
}

//--------------------------DEVICE----------------------------------
func (ctl *Controller) AddDevice(context *gin.Context) {
	var request dtos.AddRequest
	err := context.ShouldBindJSON(&request)
	if err != nil {
		utilitys.ResponseError400(context, err.Error())
		return
	}
	// fmt.Println(request)
	data, err := ctl.deviceService.Add(request)

	if err != nil {
		utilitys.ResponseError400(context, err.Error())
	} else {
		utilitys.ResponseSuccess200(context, data, "success")
	}
}

//ListDevice
func (ctl *Controller) ListDevice(context *gin.Context) {
	var request dtos.ListRequest
	err := context.ShouldBindJSON(&request)
	if err != nil {
		utilitys.ResponseError400(context, err.Error())
		return
	}
	// fmt.Println(request)
	data, err := ctl.deviceService.List(request)

	if err != nil {
		utilitys.ResponseError400(context, err.Error())
	} else {
		utilitys.ResponseSuccess200(context, data, "success")
	}
}

// Delete device
func (ctl *Controller) DeleteDevice(context *gin.Context) {
	var request dtos.DeleteRequest

	err := context.ShouldBindJSON(&request)
	if err != nil {
		utilitys.ResponseError400(context, err.Error())
		return
	}
	data, err := ctl.deviceService.Delete(request)
	if err != nil {
		utilitys.ResponseError400(context, err.Error())
	} else {
		utilitys.ResponseSuccess200(context, data, "success")
	}
}

// Edit device
func (ctl *Controller) EditDevice(context *gin.Context) {
	var request dtos.EditRequest

	err := context.ShouldBindJSON(&request)
	if err != nil {
		utilitys.ResponseError400(context, err.Error())
		return
	}
	data, err := ctl.deviceService.Edit(request)

	if err != nil {
		utilitys.ResponseError400(context, err.Error())
	} else {
		utilitys.ResponseSuccess200(context, data, "success")
	}
}

//
func (ctl *Controller) GetstatusDevice(context *gin.Context) {
	var request dtos.GetstatusRequest
	err := context.ShouldBindJSON(&request)
	if err != nil {
		utilitys.ResponseError400(context, err.Error())
		return
	}
	data, err := ctl.deviceService.Getstatus(request)
	if err != nil {
		utilitys.ResponseError400(context, err.Error())
	} else {
		utilitys.ResponseSuccess200(context, data, "success")
	}
}

// PushDevice
func (ctl *Controller) PushDevice(context *gin.Context) {
	var request dtos.PushRequest
	err := context.ShouldBindJSON(&request)
	if err != nil {
		utilitys.ResponseError400(context, err.Error())
		return
	}
	// fmt.Println(request)
	data, err := ctl.deviceService.Push(request)

	if err != nil {
		utilitys.ResponseError400(context, err.Error())
	} else {
		utilitys.ResponseSuccess200(context, data, "success")
	}

}
func (ctl *Controller) Upload(context *gin.Context) {
	start := time.Now()
	log.Println(start)
	// Multipart form
	form, _ := context.MultipartForm()
	log.Println(form)

	files := form.File["files"]
	userID := form.Value["user_id"]
	key := form.Value["key"]
	log.Println(userID)
	log.Println(files)
	if userID == nil || key == nil || files == nil {
		err_up := errors.New("Params upload not null")
		utilitys.ResponseError400(context, err_up.Error())
		return
	}
	// log.Println(userID[0])

	var request dtos.UploadRequest
	request.User_ID, _ = strconv.ParseInt(userID[0], 10, 64)
	request.Key = key[0]

	for _, file := range files {
		var f dtos.FileUpload
		f.Video_name = file.Filename
		fType := filepath.Ext(f.Video_name)
		if fType != ".mp4" && fType != ".jpg" && fType != ".png" {
			err_up := errors.New("Only supported mp4, jpeg, png File  ")
			utilitys.ResponseError400(context, err_up.Error())
			return
		}
		f.Video_size = file.Size / 1024
		f.Video_time = 0
		request.Files = append(request.Files, f)
	}
	log.Println(request.Files)

	// process data
	t1 := time.Since(start)
	log.Println("T1: ")
	log.Println(t1)
	data, err := ctl.deviceService.Upload(request)
	if err != nil {
		utilitys.ResponseError400(context, err.Error())
	} else {
		t2 := time.Since(start)
		log.Println("T2: ")
		log.Println(t2)
		//Tao thu muc
		var path = "./storage/" + string(userID[0])
		if _, err := os.Stat(path); os.IsNotExist(err) {
			err = os.MkdirAll(path, 0755)
			if err != nil {
				log.Println(err)
			}
		}

		for _, file := range files {
			log.Println(file.Filename)
			// err := context.SaveUploadedFile(file, "./storage/"+path+"/"+file.Filename)
			err := context.SaveUploadedFile(file, path+"/"+file.Filename)
			if err != nil {
				log.Println(err)
			} else {
				log.Println("save thanh cong")
			}

		}
		t3 := time.Since(start)
		log.Println("T3: ")
		log.Println(t3)
		utilitys.ResponseSuccess200(context, data, "success")
	}
}

//-------------------------------------------------------------

func (ctl *Controller) Login(context *gin.Context) {
	log.Println(context.Request.Header)
	var request dtos.LoginRequest
	err := context.ShouldBindJSON(&request)
	if err != nil {
		log.Println(err)
		utilitys.ResponseError400(context, "login error")
		return
	}

	data, err := ctl.userService.Login(request)
	if err != nil {
		log.Println(err)
		utilitys.ResponseError400(context, err.Error())
	} else {
		utilitys.ResponseSuccess200(context, data, "login success")
	}
}

func (ctl *Controller) CreateUser(context *gin.Context) {
	var request dtos.CreateUserRequest
	err := context.ShouldBindJSON(&request)
	if err != nil {
		utilitys.ResponseError400(context, err.Error())
		return
	}

	acc := models.User{
		Name:     request.Name,
		Username: request.Username,
		Password: request.Password,
		Phone:    request.Phone,
		Max_size: 600000, //KB 1MB=1000KB
	}
	data, err := ctl.userService.Create(acc)
	if err != nil {
		utilitys.ResponseError400(context, err.Error())
	} else {
		utilitys.ResponseSuccess200(context, data, "success")
	}
}

func (ctl *Controller) CreateLog(context *gin.Context) {
	userID, err := utilitys.GetUserID(context)
	if err != nil {
		utilitys.ResponseError400(context, "get userid error")
		return
	}

	var logRequest dtos.CreateLogRequest
	err = context.ShouldBindJSON(&logRequest)
	if err != nil || logRequest.Money <= 0 {
		utilitys.ResponseError400(context, "json decode error")
		return
	}

	log := models.Log{
		Detail: logRequest.Detail,
		Money:  logRequest.Money,
		Tag:    logRequest.Tag,
		UserID: userID,
	}

	data, err := ctl.userService.CreateLog(log)
	if err != nil {
		utilitys.ResponseError400(context, err.Error())
	} else {
		utilitys.ResponseSuccess200(context, data, "create log success")
	}

}

func (ctl *Controller) GetLogs(context *gin.Context) {
	userID, err := utilitys.GetUserID(context)
	if err != nil {
		utilitys.ResponseError400(context, "get userID error")
		return
	}

	beginTime, err := stringYYYYMMDD2Time(context.Query(common.BeginTimeStampKey))
	if err != nil {
		utilitys.ResponseError400(context, "input format YYYY-MM-DD")
		return
	}

	endTime, err := stringYYYYMMDD2Time(context.Query(common.EndTimeStampKey))
	if err != nil {
		utilitys.ResponseError400(context, "input format YYYY-MM-DD")
		return
	}

	data, err := ctl.userService.GetLogsByTime(userID, beginTime, endTime)
	if err != nil {
		utilitys.ResponseError400(context, err.Error())
	} else {

		utilitys.ResponseSuccess200(context, data, "success")
	}
}

func stringYYYYMMDD2Time(input string) (*time.Time, error) {
	layout := "2006-01-02"
	result, err := time.Parse(layout, input)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (ctl *Controller) Ping(context *gin.Context) {
	context.JSON(200, gin.H{
		"message": "Pong Pong",
	})
}
