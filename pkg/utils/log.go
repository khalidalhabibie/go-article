package utils

import (
	log "github.com/sirupsen/logrus"
)

func LogFormat(layer, service string, message interface{}) log.Fields {

	return log.Fields{
		"layer":   layer,
		"service": service,
		"message": message,
	}

}
