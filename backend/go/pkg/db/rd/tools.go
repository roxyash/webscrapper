package rd

import "time"

func setDefaultConfigValues(config *Config) {
	if config.MaxRetries == 0 {
		config.MaxRetries = 10 // default value
	}

	if config.RetryInterval == 0 {
		config.RetryInterval = time.Second // default value
	}
}
