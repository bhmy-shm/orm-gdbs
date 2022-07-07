package builder

import (
	"github.com/bhmy-shm/dbsdk/pbfiles"
	"google.golang.org/protobuf/types/known/structpb"
)

//参数构建器
type ParamBuilder struct {
	param map[string]interface{}
}

func NewParamBuilder() *ParamBuilder {
	return &ParamBuilder{param: make(map[string]interface{})}
}

//传
func (this *ParamBuilder) Add(name string, value interface{}) *ParamBuilder {
	this.param[name] = value
	return this
}

func (this *ParamBuilder) Build() *pbfiles.SimpleParams {
	paramStruct, _ := structpb.NewStruct(this.param)
	return &pbfiles.SimpleParams{
		Params: paramStruct,
	}
}
