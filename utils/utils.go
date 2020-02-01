package utils

import (
	"bytes"
	"fmt"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/go-playground/validator.v9"
	zhTranslations "gopkg.in/go-playground/validator.v9/translations/zh"
	"reflect"
	"time"
)

var validate *validator.Validate

//根据Json格式设置obj对象
func SetObjByJson(obj interface{}, data map[string]interface{}) error {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	for key, value := range data {
		if err := setField(obj, key, value); err != nil {
			log.Error().Msg("SetObjByJson set field fail.")
			return err
		}
	}
	return nil
}

//设置结构体中的变量
func setField(obj interface{}, name string, value interface{}) error {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	structData := reflect.TypeOf(obj).Elem()
	fieldValue, result := structData.FieldByName(name)
	if !result {
		log.Error().Str("No such field ", name)
		return fmt.Errorf("No such field %s", name)
	}

	//结构体中变量的类型
	fieldType := fieldValue.Type
	//参数的值
	val := reflect.ValueOf(value)
	//参数的类型
	valTypeStr := val.Type().String()
	//结构体中变量的类型
	fieldTypeStr := fieldType.String()
	//float64 to int
	if valTypeStr == "float64" && fieldTypeStr == "int" {
		val = val.Convert(fieldType)
	}

	//类型必须匹配
	if fieldType != val.Type() {
		return fmt.Errorf("value type %s didn't match obj field type %s ", valTypeStr, fieldTypeStr)
	}

	//fieldValue.Set(val)

	return nil
}



func ValidateData(data interface{}) string {
	validate = validator.New()
	zhCh := zh.New()
	uni := ut.New(zhCh)
	trans, _ := uni.GetTranslator("zh")
	_ = zhTranslations.RegisterDefaultTranslations(validate, trans)
	err := validate.Struct(data)
	if err != nil {
		var errString bytes.Buffer
		for _, err := range err.(validator.ValidationErrors) {
			errString.WriteString(err.Translate(trans))
			errString.WriteString("|")
		}
		return errString.String()[0 : (len(errString.String()) - 1)]
	}
	return "success"
}

func Api(code int64, message string, data interface{}) mvc.Result {
	return mvc.Response{
		Object: map[string]interface{}{
			"status": code,
			"message": message,
			"data":  data,
		},
	}
}

func HashAndSalt (pwd string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.MinCost)
	if err != nil {
		panic(err.Error())
	}
	return string(hash)
}

func ComparePasswords(hashedPwd string, Pwd string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPwd), []byte(Pwd))
	if err != nil {
		return false
	}
	return true
}

func LogInfo(app *iris.Application, v ...interface{}) {
	app.Logger().Info(v)
}

func LogError(app *iris.Application, v ...interface{}) {
	app.Logger().Error(v)
}

func LogDebug(app *iris.Application, v ...interface{}) {
	app.Logger().Debug(v)
}

/**
 * 格式化数据
 */
func FormatDatetime(time time.Time) string {
	return time.Format("2006-01-02 03:04:05")
}

type ReturnData struct {
	Message string `json:"message"`
	Data    interface{}
	Code    int8
}
