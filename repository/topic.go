package repository

import (
	"sync"
)

type Topic struct {
	Id         int64  `json:"id"`
	Title      string `json:"title"`
	Content    string `json:"content"`
	CreateTime int64  `json:"create_time"`
}
type TopicDao struct {
}

var (
	topicDao  *TopicDao
	topicOnce sync.Once
	topicLock sync.Mutex
)

func NewTopicDaoInstance() *TopicDao {
	topicOnce.Do(
		func() {
			topicDao = &TopicDao{}
		})
	return topicDao
}
func (*TopicDao) QueryTopicById(id int64) *Topic {
	return topicIndexMap[id]
}

//  repository/topic.go新增函数
func (*TopicDao) CreateTopic(topic *Topic) error {
	//lock
	//在内存中的map创建新的topic，再写回topic文件
	//topicLock.Lock()
	topicLock.Lock()
	////从topicIndexMap中找到没用过的id
	topic.Id = nextTopicId

	topicIndexMap[nextTopicId] = topic
	nextTopicId++
	FlushTopic("./data/", topic)
	topicLock.Unlock()
	//unlock
	return nil
}
