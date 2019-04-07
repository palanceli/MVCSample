package model

import (
	"fmt"
	"time"
)

// WorkerData 保存到数据库的结构体
type WorkerData struct {
	ID       int       `xorm:"'id' pk autoincr"`
	DataType string    `xorm:"data_type" json:"data_type"`
	Content  string    `xorm:"data_content" json:"data_content"`
	Time     time.Time `xorm:"uptdate_time"`
}

// GetDataType ...
func GetDataType(dataType int32) string {
	return fmt.Sprintf("data_type_%04d", dataType)
}
