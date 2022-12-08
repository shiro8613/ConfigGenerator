package logic

import (
	"fmt"
	"path/filepath"

	"github.com/labstack/echo/v4"
	"github.com/shiro8613/ConfigGenerator/config"
)

func StreamServer(stream_conf_key []string) error {
	config := config.GetConfig()
	e := echo.New()
	e.HideBanner = true

	fmt.Println("ExportFileStreamServer - Started!!")

	for _, k := range stream_conf_key {
		for _, ck := range config.Stream.Stream_Conf {
			if ck == k {
				fmt.Printf("http://%s/export/%s/somefiles", config.Stream.Bind, k)
				path, _ := filepath.Abs(config.Generate[k].ExportDir)
				e.Static(fmt.Sprintf("/export/%s/", k), path)
			}
		}
	}

	err := e.Start(config.Stream.Bind)
	if err != nil {
		return err
	}

	return nil

}
