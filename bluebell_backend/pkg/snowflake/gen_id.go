package snowflake

import (
	"fmt"
	"time"

	"github.com/sony/sonyflake"
)

var (
	sonyFlake     *sonyflake.Sonyflake	// 实例
	sonyMachineID uint16				// 机器ID
)

func getMachineID() (uint16, error) {  // 返回全局定义的机器ID
	return sonyMachineID, nil
}

// 需传入当前的机器ID
func Init(machineId uint16) (err error) {
	sonyMachineID = machineId
	t, _ := time.Parse("2006-01-02", "2022-02-09")	// 初始化一个开始的时间
	settings := sonyflake.Settings{		// 生成全局配置
		StartTime: t,
		MachineID: getMachineID,								// 指定机器ID
	}
	sonyFlake = sonyflake.NewSonyflake(settings)	// 用配置生成sonyflake节点
	return
}

// GetID 返回生成的id值
func GetID() (id uint64, err error) {		// 拿到sonyflake节点生成id值
	if sonyFlake == nil {
		err = fmt.Errorf("snoy flake not inited")
		return
	}

	id, err = sonyFlake.NextID()
	return
}

//func main()  {
//	if err := Init(1);err!=nil{
//		fmt.Printf("Init failed,err:%v\n",err)
//		return
//	}
//	id,_ := GetID()
//	fmt.Println(id)
//}
