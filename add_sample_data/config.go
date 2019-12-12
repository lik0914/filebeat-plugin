// @author https://github.com/lik0914/filebeat-plugin
package main

// Config for sample processor.
type Config struct {
	Sample float64 `config:"sample"` //采样率
}

func defaultConfig() Config {
	return Config{
		Sample: 0.2,
	}
}
