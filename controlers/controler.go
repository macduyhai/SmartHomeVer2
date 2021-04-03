package controlers

import (
	"log"
	"os"
	"strconv"
	"strings"
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
	avrService    services.AverageService
	deviceService services.DeviceService
	mediaService  services.MediaService
}

func NewController(provider services.Provider) Controller {
	return Controller{userService: provider.GetUserService(),
		avrService:    provider.GetAverageService(),
		deviceService: provider.GetDeviceService(),
		mediaService:  provider.GetMediaService(),
	}
}

//------------------------------------------------------------
func (ctl *Controller) AddMedia(context *gin.Context) {
	var request dtos.AddMediaRequest
	err := context.ShouldBindJSON(&request)
	if err != nil {
		log.Println("Lôi encode Json request")
		log.Println(err)
		utilitys.ResponseError400(context, err.Error())
		return
	}
	log.Println(request)
	data, err := ctl.mediaService.AddMedia(request)

	if err != nil {
		utilitys.ResponseError400(context, err.Error())
	} else {
		utilitys.ResponseSuccess200(context, data, "success")
	}
}

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
func (ctl *Controller) Download(context *gin.Context) {
	url := context.Request.URL.Path
	log.Println(url)
	p := strings.Split(url, "/")
	log.Println(p)
	path := "./storage/" + p[5] + "/" + p[6]
	log.Println(path)
	//log.Println("Opening a file ")
	// var file, err = os.OpenFile(path, os.O_RDWR, 0644)
	// if err != nil {
	// 	log.Println(err)
	// } else {
	// 	log.Println("MOpen file done")
	// }
	// defer file.Close()
	// context.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", p[6])) //fmt.Sprintf("attachment; filename=%s", filename) Downloaded file renamed
	// context.Writer.Header().Add("Content-Type", "application/octet-stream")
	// context.FileAttachment(path, p[6])

	// http.ServeFile(context.Writer, context.Request, path)
	context.FileAttachment(path, p[6])
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
	files := form.File["file"]
	userID := form.Value["user_id"]
	key := form.Value["key"]

	log.Printf("%T", files)

	// log.Println(userID[0])

	var request dtos.UploadRequest
	request.User_ID, _ = strconv.ParseInt(userID[0], 10, 64)
	request.Key = key[0]

	for _, file := range files {
		var f dtos.FileUpload
		f.Video_name = file.Filename
		f.Video_size = file.Size
		f.Video_time = 0
		request.Files = append(request.Files, f)
	}
	log.Println(request.Files)
	// err := context.ShouldBindJSON(&request)
	// if err != nil {
	// 	utilitys.ResponseError400(context, err.Error())
	// 	return
	// }
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
		Max_size: 629145600,
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

func (ctl *Controller) AnalysisByTag(context *gin.Context) {
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

	data, err := ctl.userService.AnalysisByTag(userID, beginTime, endTime)
	if err != nil {
		utilitys.ResponseError400(context, err.Error())
	} else {
		utilitys.ResponseSuccess200(context, data, "success")
	}
}

func (ctl *Controller) AnalysisByDay(context *gin.Context) {
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

	data, err := ctl.userService.AnalysisByDay(userID, beginTime, endTime)
	if err != nil {
		utilitys.ResponseError400(context, err.Error())
	} else {
		utilitys.ResponseSuccess200(context, data, "input format YYYY-MM-DD")
	}
}

func (ctl *Controller) GetAverageByDay(context *gin.Context) {
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

	money, err := ctl.avrService.CalculateByDay(userID, beginTime, endTime)
	if err != nil {
		utilitys.ResponseError400(context, err.Error())
	} else {
		result := dtos.AverageMoneyPerDayResponse{Money: money}
		utilitys.ResponseSuccess200(context, result, "success")
	}
}

func (ctl *Controller) Ping(context *gin.Context) {
	context.JSON(200, gin.H{
		"message": "Pong Pong",
	})
}
