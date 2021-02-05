package upload

import (
	//"os"
	//"path"
	//"log"
	//"fmt"
	//"strings"
	//"mime/multipart"
	//"wechatNotify/pkg/logging"
	"wechatNotify/pkg/setting"
)

func GetImageFullUrl(name string) string {
	return setting.ImagePrefixUrl + "/" + GetImagePath() + name
}
//
//func GetImageName(name string) string {
//	ext := path.Ext(name)
//	fileName := strings.TrimSuffix(name, ext)
//	fileName = util.EncodeMD5(fileName)
//
//	return fileName + ext
//}
//
func GetImagePath() string {
	return setting.ImageSavePath
}
//
func GetImageFullPath() string {
	return setting.RuntimeRootPath + GetImagePath()
}
//
//func CheckImageExt(fileName string) bool {
//	ext := file.GetExt(fileName)
//	for _, allowExt := range setting.AppSetting.ImageAllowExts {
//		if strings.ToUpper(allowExt) == strings.ToUpper(ext) {
//			return true
//		}
//	}
//
//	return false
//}
//
//func CheckImageSize(f multipart.File) bool {
//	size, err := file.GetSize(f)
//	if err != nil {
//		log.Println(err)
//		logging.Warn(err)
//		return false
//	}
//
//	return size <= setting.AppSetting.ImageMaxSize
//}
//
//func CheckImage(src string) error {
//	dir, err := os.Getwd()
//	if err != nil {
//		return fmt.Errorf("os.Getwd err: %v", err)
//	}
//
//	err = file.IsNotExistMkDir(dir + "/" + src)
//	if err != nil {
//		return fmt.Errorf("file.IsNotExistMkDir err: %v", err)
//	}
//
//	perm := file.CheckPermission(src)
//	if perm == true {
//		return fmt.Errorf("file.CheckPermission Permission denied src: %s", src)
//	}
//
//	return nil
//}