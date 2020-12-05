package impl

import "fmt"

// 渠道补单
type ChannelSupplementOrder struct {
}

func (r *ChannelSupplementOrder) SupplementOrder(params map[string]interface{}) map[string]interface{} {
	fmt.Println("创建渠道补单")
	response := make(map[string]interface{})
	return response
}