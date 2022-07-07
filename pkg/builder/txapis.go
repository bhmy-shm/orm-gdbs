package builder

import (
	"context"
	"fmt"
	"github.com/bhmy-shm/orm-gdbs/pbfiles"
	"github.com/mitchellh/mapstructure"
	"google.golang.org/grpc"
	"log"
)

//程序员在囧途(www.jtthink.com)出品 咨询群：98514334
//事务API对象
type Txapi struct {
	ctx    context.Context
	cancel context.CancelFunc
	client pbfiles.DBService_TxClient
}

func NewTxapi(ctx context.Context, client pbfiles.DBServiceClient, opts ...grpc.CallOption) *Txapi {
	apiCtx, cancel := context.WithCancel(ctx)
	txClient, err := client.Tx(apiCtx, opts...)
	if err != nil {
		panic(err)
	}
	return &Txapi{ctx: ctx, client: txClient, cancel: cancel}
}
func (this *Txapi) Exec(apiname string, paramBuilder *ParamBuilder, out interface{}) error {
	err := this.client.Send(&pbfiles.TxRequest{Api: apiname, Params: paramBuilder.Build(), Type: "exec"})
	//对于exec ,如果不出错， 会返回一个map，其中key=exec ,  值是一个interface切片 ，包含了两项
	//1、受影响的行  2 、selectkey(如果有的话)
	if err != nil {
		return err
	}
	rsp, err := this.client.Recv() //接收消息
	if err != nil {
		return err
	}
	if out != nil {
		if execRet, ok := rsp.Result.AsMap()["exec"]; ok { //execRet 是一个 切片 []interface{受影响的行，select}  .select 可能是nil
			if execRet.([]interface{})[1] != nil { //代表select 有值
				m := execRet.([]interface{})[1].(map[string]interface{})
				m["_RowsAffected"] = execRet.([]interface{})[0]
				return mapstructure.WeakDecode(m, out)
			} else { //没有select 情况 直接塞一个_RowsAffected 返回
				m := map[string]interface{}{"_RowsAffected": execRet.([]interface{})[0]}
				return mapstructure.WeakDecode(m, out)
			}

		}
	}
	return nil

}
func (this *Txapi) Query(apiname string, paramBuilder *ParamBuilder, out interface{}) error {
	err := this.client.Send(&pbfiles.TxRequest{Api: apiname, Params: paramBuilder.Build(), Type: "query"})
	// 对于查询，如果不出错，会返回一个map   其中key=query    值是查询结果
	if err != nil {
		return err
	}
	rsp, err := this.client.Recv()
	if err != nil {
		return err
	}
	if out != nil {
		if queryRet, ok := rsp.Result.AsMap()["query"]; ok {
			return mapstructure.WeakDecode(queryRet, out)
		} else {
			return fmt.Errorf("error query result ")
		}
	}
	return nil
}

func (this *Txapi) QueryForModel(apiname string, paramBuilder *ParamBuilder, out interface{}) error {
	err := this.client.Send(&pbfiles.TxRequest{Api: apiname, Params: paramBuilder.Build(), Type: "query"})
	// 对于查询，如果不出错，会返回一个map   其中key=query    值是查询结果
	if err != nil {
		return err
	}
	rsp, err := this.client.Recv()
	if err != nil {
		return err
	}
	if out != nil {
		if queryRet, ok := rsp.Result.AsMap()["query"]; ok {
			//queryRet 是 []interface类型。具体看 dbcore的一层转换--service.go #102
			if retForMap, ok := queryRet.([]interface{}); ok && len(retForMap) == 1 {
				return mapstructure.WeakDecode(retForMap[0], out)
			} else {
				return fmt.Errorf("error query model: no result ")
			}

		} else {
			return fmt.Errorf("error query result ")
		}
	}
	return nil
}

// 模仿 gorm
func (this *Txapi) Tx(fn func(tx *Txapi) error) error {
	err := fn(this)
	if err != nil {
		log.Println("tx error:", err)
		this.cancel() //取消
		return err
	}
	return this.client.CloseSend() //协程不安全
}
