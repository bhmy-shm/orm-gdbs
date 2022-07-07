package examples

import (
	"context"
	"fmt"
	"github.com/bhmy-shm/dbsdk/pkg/builder"
	"google.golang.org/grpc"
	"log"
)
//- name: adduser
//sql: "insert into users(user_name,user_pass) values(@userName,@userPass)"
//select:
//sql: "SELECT LAST_INSERT_ID() as user_id"
//- name: adduserscore
//sql: "insert into userscores(user_id,user_score) values(@userID,@userScore)"
type User struct {
    UserID int  `mapstructure:"id"`
    UserName string
    UserPass string
}
type ExecResult struct {
	RowsAffected int `mapstructure:"_RowsAffected"`
}

func Txtest()  {
	    client,cerr:=builder.NewClientBuilder("localhost:8082").WithOption(grpc.WithInsecure()).Build()
	    if cerr!=nil{
	    	log.Fatal(cerr)
		}
	    //创建 事务API
		txApi:=builder.NewTxapi(context.Background(),client)
		// 模仿了 Gorm
		err:=txApi.Tx(func(tx *builder.Txapi) error {
			//构建 用户实体
			user:=&User{UserName:"zhangsan",UserPass:"123"}
			addUser_Param:=builder.NewParamBuilder().
				Add("userName",user.UserName).
				Add("userPass",user.UserPass)

			//执行新增 用户，user.UserID 会自动赋值
			err:=tx.Exec("adduser",addUser_Param,user)
			if err!=nil{
				return  err
			}

			log.Println("新增用户成功，用户ID是:",user.UserID)

			//给用户赠送积分
			addUserScore_Param:=builder.NewParamBuilder().
				Add("userID",user.UserID).Add("userScore",10)

			execRet:=&ExecResult{}
			err=tx.Exec("adduserscore",addUserScore_Param,execRet)
			if err!=nil || execRet.RowsAffected!=1{
				return  fmt.Errorf("adduserscore error")
			}
			return nil
		})
		log.Println(err)
	select {

		}
}
