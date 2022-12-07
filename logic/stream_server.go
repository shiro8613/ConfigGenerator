package logic

import (
	"github.com/gin-gonic/gin"
	"github.com/shiro8613/ConfigGenerator/config"
)

func StreamServer(stream_conf_key []string) error {
	config := config.GetConfig()
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	for _, k := range stream_conf_key {
		for _, ck := range config.Stream.Stream_Conf {
			if ck == k {
				router.StaticFile("/export/", config.Generate[k].ExportDir)
			}
		}
	}

	router.Run(config.Stream.Bind)

	return nil
}
