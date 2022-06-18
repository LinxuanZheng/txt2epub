package src

import (
	"gopkg.in/ini.v1"
	"io/ioutil"
	"os"
)

const ConfigFile = "project.ini"
const DefaultConfig = `[book]
cover-img=cover.png
title=Title
author=佚名
regex=^第{0,1}[零〇一二三四五六七八九十百千万a-zA-Z0-9]{1,7}[章节卷集部篇回]`

type Config struct {
	Cover  string
	Title  string
	Author string
	Desc   string
	Regex  string
	UUID   string
}

func InitConfig() {
	currDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	cfgPath := currDir + "/" + ConfigFile

	if !Exists(cfgPath) {
		err = ioutil.WriteFile(cfgPath, []byte(DefaultConfig), 0444)
		if err != nil {
			panic(err)
		}
	}
}

func LoadConfig() Config {
	currDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	cfg, err := ini.Load(currDir + "/" + ConfigFile)
	if err != nil {
		panic(err)
	}

	// 获取mysql分区的key
	cvr := cfg.Section("book").Key("cover-img").String()
	title := cfg.Section("book").Key("title").String()
	author := cfg.Section("book").Key("author").String()
	desc := cfg.Section("book").Key("desc").String()
	regex := cfg.Section("book").Key("regex").String()

	return Config{
		Cover:  cvr,
		Title:  title,
		Author: author,
		Desc:   desc,
		Regex:  regex,
	}
}
