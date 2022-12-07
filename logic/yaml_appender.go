package logic

import "github.com/shiro8613/ConfigGenerator/paser"

func YamlAppender(commmon, any paser.HogeYaml) paser.HogeYaml {
	anyCopy := make(paser.HogeYaml)

	for k, v := range any {
		anyCopy[k] = v
	}

	for ck, cv := range commmon {
		for ak := range any {
			if ck == ak {
				anyCopy[ck] = cv
			}
		}
	}

	return anyCopy
}
