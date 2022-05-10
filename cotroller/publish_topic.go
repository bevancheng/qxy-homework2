//  cotroller/publish_topic.go
package cotroller

import "github.com/Moonlight-Zhao/go-project-example/service"

func PublishTopic(title, content string) *PageData {
	//获取service层的结果
	topicId, err := service.PublishTopic(title, content)
	if err != nil {
		return &PageData{
			Code: -1,
			Msg:  err.Error(),
		}
	}
	return &PageData{
		Code: 0,
		Msg:  "success",
		Data: map[string]int64{
			"topic_id": topicId,
		},
	}
}
