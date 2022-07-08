package builder

import (
	"context"
	"github.com/bhmy-shm/orm-gdbs/pbfiles"
	"github.com/bhmy-shm/orm-gdbs/pkg/helpers"
	"github.com/mitchellh/mapstructure"
	"log"
)

const (
	APITYPE_QUERY = iota
	APITYPE_EXEC
)

type ApiBuilder struct {
	name    string //api 名称
	apitype int
}

func NewApiBuilder(name string, apitype int) *ApiBuilder {
	return &ApiBuilder{name: name, apitype: apitype}
}

//普通执行， 不是事务
func (this *ApiBuilder) Invoke(ctx context.Context, paramBuilder *ParamBuilder,
	client pbfiles.DBServiceClient, out interface{}) error {
	if this.apitype == APITYPE_QUERY { //查询
		req := &pbfiles.QueryRequest{Name: this.name, Params: paramBuilder.Build()}
		rsp, err := client.Query(ctx, req)
		if err != nil {
			log.Println(err)
			return err
		}
		log.Println("rsp=", rsp)
		mapList := helpers.PbstructsToMapList(rsp.GetResult())
		log.Println(mapList)
		if out != nil {
			return mapstructure.WeakDecode(mapList, out)
		}
		return nil
	} else {
		//增删改查
		req := &pbfiles.ExecRequest{Name: this.name, Params: paramBuilder.Build()}
		rsp, err := client.Exec(ctx, req)
		if err != nil {
			return err
		}
		if out != nil { //如果 out 没有传值 不做转换
			var m map[string]interface{}
			if rsp.Select != nil {
				m = rsp.Select.AsMap()
				m["_RowsAffected"] = rsp.RowsAffected
			} else {
				m = map[string]interface{}{"_RowsAffected": rsp.RowsAffected}
			}
			return mapstructure.WeakDecode(m, out)
		}
		return nil
	}
}
