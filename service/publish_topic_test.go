package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPublishTopic(t *testing.T) {
	var idexpected int64 = 13

	for ; idexpected < 1000; idexpected += 2 {

		_, err := PublishTopic("testtitle", "testcontent")
		assert.Equal(t, nil, err)
		newid, err := PublishTopic("testtitle", "testcontent")
		assert.Equal(t, nil, err)
		assert.Equal(t, idexpected, newid)

	}

	var largetitle string = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
	newid, err := PublishTopic(largetitle, "testcontent")
	idexpected = 0
	assert.Equal(t, idexpected, newid)
	assert.NotEqual(t, nil, err)

	var largecontent string = largetitle + largetitle + largetitle + largetitle
	newid, err = PublishTopic("testtitle", largecontent)
	idexpected = 0
	assert.Equal(t, idexpected, newid)
	assert.NotEqual(t, nil, err)
}
