// @author https://github.com/lik0914/filebeat-plugin
package main

import (
	"crypto/md5"
	"fmt"
	"github.com/elastic/beats/libbeat/beat"
	"github.com/elastic/beats/libbeat/common"
	"github.com/elastic/beats/libbeat/logp"
	"github.com/elastic/beats/libbeat/processors"
	"github.com/pkg/errors"
	"regexp"
)

type SampleData struct {
	config Config
}

var (
	counter = make(map[string]int)
	uriReg  = regexp.MustCompile(`(?:GET|POST) ([\w\/]+) HTTP`)
)

const (
	processorName = "add_sample_data"
)

/*func init() {
	processors.RegisterPlugin("add_sample_data", newSampleProcessor)
}*/

// 获取URI统计信息
func GetCounter() map[string]int {
	return counter
}

func newSampleProcessor(cfg *common.Config) (processors.Processor, error) {
	config := defaultConfig()

	if err := cfg.Unpack(&config); err != nil {
		return nil, errors.Wrapf(err, "fail to unpack the %v configuration", processorName)
	}

	sampleData := &SampleData{
		config: config,
	}
	return sampleData, nil
}

func (sampleData SampleData) Run(event *beat.Event) (*beat.Event, error) {

	message, ok := event.Fields.GetValue("message")
	if ok != nil {
		return event, nil
	}

	// 文本内容需要固定格式
	/*uri := strings.Fields(message.(string))[3]
	if uri == "" {
		return nil, nil
	}*/

	// 使用正则匹配相对实用性相对好点
	uris := uriReg.FindStringSubmatch(message.(string))
	if len(uris) < 2 {
		return nil, nil
	}

	uri := uris[1]

	// logp.Debug(processorName, "sample start msg:%s", uri)

	//根据采样配置进行采样
	if sampleData.sample(uri) {
		return event, nil
	}

	// 抛弃
	// libbeat/publisher/pipeline/client.go publish 调用 c.onFilteredOut(e)
	return nil, nil
}

func (sampleData *SampleData) sample(uri string) bool {
	h := sampleData.md5Hash(uri)
	counter[h] += 1
	logp.Debug(processorName, "sample counter:%#v", counter)
	if sampleData.config.Sample*float64(counter[h]) >= 1 {
		logp.Debug(processorName, "sample successful uri:%s", uri)
		delete(counter, uri)
		return true
	}
	logp.Debug(processorName, "sample abnormal uri:%s", uri)
	return false

}

// 实时统计数据按照URL维度
func (sampleData *SampleData) GetCounter() map[string]int {
	return GetCounter()
}

// MD5Hash工具
func (sampleData *SampleData) md5Hash(s string) string {
	h := md5.Sum([]byte(s))
	return fmt.Sprintf("%x", h)
}

func (sampleData SampleData) String() string {
	return "sampleData"
}
