//  service/publish_topic.go
package service

import (
	"errors"
	"time"
	"unicode/utf8"

	"github.com/Moonlight-Zhao/go-project-example/repository"
)

func PublishTopic(title string, content string) (int64, error) {
	return NewPublishTopicFlow(title, content).Do()
}

func NewPublishTopicFlow(title string, content string) *PublishTopFlow {
	return &PublishTopFlow{

		title:   title,
		content: content,
	}
}

type PublishTopFlow struct {
	topicId int64
	title   string
	content string
}

func (f *PublishTopFlow) Do() (int64, error) {
	if err := f.checkParam(); err != nil {
		return 0, err
	}
	//publish之后，才会产生真正的topicId
	if err := f.publish(); err != nil {
		return 0, err
	}
	return f.topicId, nil
}

func (f *PublishTopFlow) checkParam() error {
	//title不大于50，content不大于500
	if utf8.RuneCountInString(f.title) >= 50 {
		return errors.New("title length must be less than 50")
	}
	if utf8.RuneCountInString(f.content) >= 500 {
		return errors.New("content length must be less than 500")
	}
	return nil
}

func (f *PublishTopFlow) publish() error {
	topic := &repository.Topic{
		//Id:         f.topicId,传不传都行，现在还不是确定的
		Title:      f.title,
		Content:    f.content,
		CreateTime: time.Now().Unix(),
	}
	//在repository层分配topicid
	if err := repository.NewTopicDaoInstance().CreateTopic(topic); err != nil { ////////////////////////
		return err
	}
	f.topicId = topic.Id
	return nil
}
