package image_service

import (
	"AfterEnd/global"
	"AfterEnd/model"
	"AfterEnd/plugin/qiniu"
	"AfterEnd/utils"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"io/ioutil"
	"mime/multipart"
	"path/filepath"
	"strings"
)

// ImageWhiteList 图片上传白名单
var ImageWhiteList = []string{
	".jpg",
	".png",
	".apng",
	".jpeg",
	".tiff",
	".gif",
	".ico",
	".svg",
	".webp",
}

// ImageUploadService 处理图片文件上传的方法
func (ImageService) ImageUploadService(FileHeader *multipart.FileHeader, c *gin.Context) (req model.FileUploadResponse) {
	var size int64 = global.Config.QiNiu.Size

	fileName := FileHeader.Filename
	// 提取文件的扩展名字
	ext := strings.ToLower(filepath.Ext(fileName))

	// 如果用户上传的文件不在白名单中，则直接判断下一个文件
	if !utils.IsInStringList(ext, ImageWhiteList) {
		req = GenerateFileUploadReq(0, fileName, false, fmt.Sprintf("上传文件类型错误，当前文件后缀为%s", ext))
		return req
	}

	// 判断文件大小是否大于指定最大文件大小，大于则直接判断下一个文件
	fileSize := float64(FileHeader.Size) / float64(1024*1024)
	if fileSize > float64(size) {
		req = GenerateFileUploadReq(0, fileName, false, fmt.Sprintf("上传图片大小大于设定大小，设定大小为 %.3f MB，当前图片大小为 %.3f MB", float64(size), fileSize))
		return req
	}

	// 先调用Open函数打开`*multipart.FileHeader`对应的文件，用fileObj接收文件对应的`multipart.File`
	// 将文件中的内容读出，存入`[]byte`，便于调用函数将文件存入七牛云服务器
	fileObj, err := FileHeader.Open()
	fileObjContent, err := ioutil.ReadAll(fileObj)
	if err != nil {
		global.Log.Error(err.Error())
	}
	// 调用MD5对读取出的[]byte内容进行hash
	imageHash := utils.MD5(fileObjContent)

	// 去数据库中查询该数据是否存在，若存在则直接下一个文件，不存在则上传和入库
	//fmt.Println(imageHash)
	var photo model.ProfilePhoto
	result := global.Db.Where("hash = ?", imageHash).First(&photo)
	if result.Error != nil {
		if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
			// 查询数据库出错，返回错误信息
			return GenerateFileUploadReq(0, fileName, false, fmt.Sprintf("查询数据库出错: %v", result.Error))
		}
		// 数据库中不存在该图片，继续上传
	} else {
		// 数据库中存在该图片，直接返回图片已存在的信息
		return GenerateFileUploadReq(photo.ID, photo.Path, false, "图片已存在")
	}

	// 上传图片文件至七牛云存储空间
	fileURL, err := qiniu.UploadFileToQiNiu(fileObjContent, "cinema", fileName)

	// 图片上传是否失败
	if err != nil { // 图片上传本地或云服务器失败
		global.Log.Error(err.Error())
		var msg string
		if global.Config.QiNiu.Enable {
			msg = fmt.Sprintf("上传图片保存到七牛云服务器失败，错误信息:%s", err.Error())
		} else {
			msg = fmt.Sprintf("上传图片保存到本地失败，错误信息:%s", err.Error())
		}
		req = GenerateFileUploadReq(0, fileName, false, msg)
		return req
	} else { // 图片已经上传成功
		// 将上传的图片记录存入数据库
		// 将上传的图片记录存入数据库
		var image = model.ProfilePhoto{
			Path: fileURL,
			Hash: imageHash,
			Name: fileName,
		}
		err = global.Db.Create(&image).Error
		if err != nil {
			global.Log.Error(fmt.Sprintf("图片文件写入数据库出错，报错信息:%s", err.Error()))
		}

		var pictureID uint
		var msg string
		pictureID = image.ID
		msg = fmt.Sprintf("图片上传七牛云服务器成功，当前图片大小为 %.3f MB", float64(FileHeader.Size)/(2<<20))
		// 为每个图片文件构造返回信息
		req = GenerateFileUploadReq(pictureID, fileURL, true, msg)
		err = fileObj.Close()
		if err != nil {
			global.Log.Warn(fmt.Sprintf("文件关闭失败，报错信息:%s", err.Error()))
		}
		return req
	}
}

func GenerateFileUploadReq(pictureID uint, filePath string, isSuccess bool, msg string) (req model.FileUploadResponse) {
	return model.FileUploadResponse{
		PictureID: pictureID,
		FilePath:  filePath,
		IsSuccess: isSuccess,
		Msg:       msg,
	}
}
