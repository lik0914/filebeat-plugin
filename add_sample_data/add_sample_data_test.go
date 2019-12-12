// @author https://github.com/lik0914/filebeat-plugin
package main

import (
	"fmt"
	"github.com/elastic/beats/libbeat/beat"
	"github.com/elastic/beats/libbeat/common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestSampleDataRun(t *testing.T) {

	config := map[string]interface{}{
		"sample": 0.2,
	}

	testConfig, err := common.NewConfigFrom(config)
	assert.NoError(t, err)
	p, err := newSampleProcessor(testConfig)
	require.NoError(t, err)

	for i := 1; ; i++ {
		// 生成 event
		event := &beat.Event{
			Fields:    common.MapStr{},
			Timestamp: time.Now(),
		}
		// [05/Dec/2019:19:49:26 +0800] "POST /aaa/bbb/get_data_event HTTP/1.1" 200 200 591 "-" "okhttp/3.12.0" "117.136.22.52, 120.27.173.69" "0.005" "0.005" "172.16.33.32:9088"
		_, _ = event.Fields.Put("message",
			"[05/Dec/2019:19:49:26 +0800] "+
				"\"POST /aaa/bbb/get_data_event HTTP/1.1\""+
				" 200 200 591 "+
				"\"-\" "+
				"\"okhttp/3.12.0\" "+
				"\"117.136.22.52, 120.27.173.69\" "+
				"\"0.005\" \"0.005\" \"172.16.33.32:9088\"")

		newEvent, err := p.Run(event)

		fmt.Printf("i:%d counter:%#v\n", i, GetCounter())

		assert.NoError(t, err)
		if i%10 == 0 {
			assert.NotNil(t, newEvent, "")
		} else {
			assert.Nil(t, newEvent)
		}
	}

}
