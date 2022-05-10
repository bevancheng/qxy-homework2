package repository

import (
	"bufio"
	"encoding/json"
	"os"
)

var (
	topicIndexMap map[int64]*Topic
	postIndexMap  map[int64][]*Post
	nextTopicId   int64
)

func Init(filePath string) error {
	if err := initTopicIndexMap(filePath); err != nil {
		return err
	}
	if err := initPostIndexMap(filePath); err != nil {
		return err
	}
	return nil
}

func initTopicIndexMap(filePath string) error {
	open, err := os.Open(filePath + "topic")
	if err != nil {
		return err
	}
	scanner := bufio.NewScanner(open)
	topicTmpMap := make(map[int64]*Topic)
	var i int64

	for scanner.Scan() {
		text := scanner.Text()
		var topic Topic
		if err := json.Unmarshal([]byte(text), &topic); err != nil {
			return err
		}
		topicTmpMap[topic.Id] = &topic
		i++

	}
	i++
	nextTopicId = i
	topicIndexMap = topicTmpMap
	return nil
}
func FlushTopic(filePath string, topic *Topic) error {
	open, err := os.OpenFile(filePath+"topic", os.O_WRONLY|os.O_APPEND, 0777) ///第三个参数？
	if err != nil {
		return err
	}
	//encode := json.NewDecoder()
	bytes, err := json.Marshal(*topic)
	if err != nil {
		return err
	}
	_, err = open.Write(bytes)
	open.WriteString("\n")
	if err != nil {
		return err
	}

	return nil
}

func FlushTopicAll(filePath string) error {
	open, err := os.OpenFile(filePath+"topic", os.O_RDWR, 0777) ///第三个参数？

	if err != nil {
		return err
	}
	//encode := json.NewDecoder()
	for _, topicItem := range topicIndexMap {
		bytes, err := json.Marshal(*topicItem)
		if err != nil {
			return err
		}
		_, err = open.Write(bytes)
		open.WriteString("\n")
		if err != nil {
			return err
		}
	}
	err = open.Close()
	if err != nil {
		return err
	}
	return nil

}

func initPostIndexMap(filePath string) error {
	open, err := os.Open(filePath + "post")
	if err != nil {
		return err
	}
	scanner := bufio.NewScanner(open)
	postTmpMap := make(map[int64][]*Post)
	for scanner.Scan() {
		text := scanner.Text()
		var post Post
		if err := json.Unmarshal([]byte(text), &post); err != nil {
			return err
		}
		posts, ok := postTmpMap[post.ParentId]
		if !ok {
			postTmpMap[post.ParentId] = []*Post{&post}
			continue
		}
		posts = append(posts, &post)
		postTmpMap[post.ParentId] = posts
	}
	postIndexMap = postTmpMap
	return nil
}
