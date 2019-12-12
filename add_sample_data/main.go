// @author https://github.com/lik0914/filebeat-plugin
// 抽样采集日志
// 基于URL维度计数，进行采样控制
package main

import (
	"github.com/elastic/beats/libbeat/plugin"
	"github.com/elastic/beats/libbeat/processors"
)

var Bundle = plugin.Bundle(
	processors.Plugin(processorName, newSampleProcessor),
)

/*func main() {
	str := "[05/Dec/2019:19:49:26 +0800] \"GET /activity/fission/task/switchs HTTP/1.1\" 200 200 779 \"http://api.yizhibo.com\" \"-\" \"47.95.124.115, 120.55.177.27\" \"0.003\" \"0.003\" \"172.16.34.174:80\""

	reg := regexp.MustCompile(`(?:GET|POST) ([\w\/]+) HTTP`)

	match := reg.FindStringSubmatch(str)

	if len(match) < 2 {
		fmt.Println("error~")
	} else {
		fmt.Println(match[1])
	}

	fmt.Println(0.2*float64(5) > 1)

	counter := make(map[string]int)

	fmt.Println(counter["a"])
	counter["a"] += 1

	fmt.Println(counter["a"])

	fmt.Printf("counter:%#v", counter)
}*/

/*func main() {
	data := []byte("/activity/fission/task/switchs")
	fmt.Printf("%x \n", sha1.Sum(data))

	str := md5.Sum(data)
	s := string(str[:])
	fmt.Printf("%x \n", str)
	fmt.Println(s)
}*/
