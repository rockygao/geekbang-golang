package ginex

import (
	"context"
	"net/http"
	"strconv"
	"strings"

	"go-project/pkg/ginex/internal"
	"go-project/pkg/ut"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/jinzhu/copier"
	log "github.com/sirupsen/logrus"
)

type BaseHdl struct {
}

func GetContext(c *gin.Context) context.Context {
	return nil
}

func (p *BaseHdl) GetContext(c *gin.Context) context.Context {
	return GetContext(c)
}

func (p *BaseHdl) Copy(toValue interface{}, fromValue interface{}) error {
	err := copier.Copy(toValue, fromValue)
	if err != nil {
		return ut.WrapInternalCusError(err, "内部错误")
	}

	return nil
}

func (p *BaseHdl) HandleError(c *gin.Context, err error) bool {
	return ut.HandleError(c, err)
}

func (p *BaseHdl) ParseIntParam(c *gin.Context, key string) (int, bool) {
	s := c.Param(key)
	i, err := strconv.Atoi(s)
	if err != nil {
		log.WithField("key", key).WithField("value", s).Errorln("bad int param")
		c.String(http.StatusBadRequest, "参数错误")
		c.Abort()
		return 0, false
	}

	return i, true
}

func (p *BaseHdl) Binding(c *gin.Context, obj interface{}, b ...binding.Binding) bool {
	var err error
	if len(b) == 0 {
		err = c.Bind(obj)
	} else {
		err = c.MustBindWith(obj, b[0])
	}

	if err == nil {
		return true
	}

	log.WithError(err).Errorln("bind error")

	var msg string
	errs, ok := err.(validator.ValidationErrors)
	if ok {
		ret := internal.Translate(errs)
		var arr []string
		for _, v := range ret {
			arr = append(arr, ut.LowerFirst(v))
		}

		msg = strings.Join(arr, "\n")
	} else {
		msg = err.Error()
	}

	c.String(http.StatusBadRequest, msg)
	c.Abort()
	return false

}

func (p *BaseHdl) Valid(c *gin.Context, v Validator) bool {
	if err := v.Valid(); err != nil {
		log.WithError(err).Errorln("valid form error")

		e, ok := ut.IsCusError(err)
		if ok {
			c.String(http.StatusBadRequest, e.Msg())
		} else {
			c.String(http.StatusBadRequest, err.Error())
		}

		c.Abort()
		return false
	}

	return true
}

type Validator interface {
	Valid() error
}
