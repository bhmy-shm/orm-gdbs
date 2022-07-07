package examples

import (
	"context"
	"fmt"
	"github.com/bhmy-shm/orm-gdbs/pkg/builder"
	"google.golang.org/grpc"
	"log"
	"time"
)

type ProdStock struct {
	StockId int `mapstructure:"stock_id"`
	Stock   int `mapstructure:"stock"`
	Version int `mapstructure:"version"`
}

func TxStocktest(pid, num int, delay int) {
	client, cerr := builder.NewClientBuilder("localhost:8082").WithOption(grpc.WithInsecure()).Build()
	if cerr != nil {
		log.Fatal(cerr)
	}
	//创建 事务API
	txApi := builder.NewTxapi(context.Background(), client)
	// 模仿了 Gorm
	//程序员在囧途(www.jtthink.com)出品 咨询群：98514334
	err := txApi.Tx(func(tx *builder.Txapi) error {
		ps := &ProdStock{StockId: pid}
		psParam := builder.NewParamBuilder().
			Add("stock_id", pid)

		err := tx.QueryForModel("getstock", psParam, ps)
		if err != nil {
			return err
		}
		fmt.Println("当前库存为:", ps.Stock)
		if ps.Stock < num {
			return fmt.Errorf("库存不够了")
		}
		// 很可能还有其他业务

		// 故意延迟， 模拟 业务很繁忙
		if delay > 0 {
			time.Sleep(time.Second * time.Duration(delay))
		}
		setStockParam := builder.NewParamBuilder().
			Add("stock_id", pid).
			Add("stock", ps.Stock-num).
			Add("version", ps.Version)

		execRet := &ExecResult{} // 增删改执行结果
		err = tx.Exec("setstock", setStockParam, execRet)
		if err != nil {
			return fmt.Errorf("扣减库存失败:%s", err.Error())
		}
		if execRet.RowsAffected != 1 {
			return fmt.Errorf("扣减库存没有成功")
		}

		return nil
	})
	log.Println(err)

}

//程序员在囧途(www.jtthink.com)出品 咨询群：98514334
