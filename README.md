# golog
本日志库初期为本人的golang提供日志支持

基础代码基于golang的官方日志库修改

# 主要功能

* 增加了log级别设置

* 去掉默认全局日志接口，使用日志需要实例化给出模块名方便跟踪

* 可根据模块名设定所有模块日志的层级方便切换想看到的日志

* 根据内容标记日志颜色

# fork-修改说明(odwodw)

* 基于github.com/davyxu/golog修改

* 增加Pid,GoId日志输出

* 增加字符串格式设置parts
** %L LogPart_Level
** %T LogPart_TimeMS
** %t LogPart_Time
** %F LogPart_LongFileName
** %f LogPart_ShortFileName
** %N LogPart_Name
** %P LogPart_Pid
** %G LogPart_Gid


# 安装方法

	go get github.com/odwodw/golog

# 使用方法

* 基本使用

	var log *golog.Logger = golog.New("test")

	log.Debugln("hello world")

* 层级设置

	golog.SetLevelByString( "test", "info")


# 备注

感觉不错请star, 谢谢!

开源讨论群: 527430600

知乎: [http://www.zhihu.com/people/sunicdavy](http://www.zhihu.com/people/sunicdavy)
