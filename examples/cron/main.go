package main

import (
	"fmt"
	"time"

	cron "github.com/robfig/cron/v3"
)

func main() {
	c := cron.New(cron.WithSeconds())
	// 	常见cron表达式：
	// 表达式如果算上年份，共7位：
	// 秒 分 时 日 月 周 年
	// 实际应用中很少用到年份，所以一般表达式都是6位：
	// 秒 分 时 日 月 周
	// */1 * * * * ? 每秒
	// 00 * * * * ? 每分
	// 00 30 22 * * ? 每天晚上22:30
	// 00 30 22 * * 2 每周二晚上22:30
	// 00 30 22 * * 1,3 每周一和周三晚上22:30
	// */5 35 10 * * 1-3 每周一到周三上午10点35分00秒开始第一次，然后每5秒一次
	// 00 10,40 14 ? 3 4 每年三月的星期四的下午14:10和14:40

	id, err := c.AddFunc("*/1 * * * * ?", func() {
		fmt.Println("123")
	})
	if err != nil {
		panic(err)
	}
	c.Start()
	go func() {
		time.Sleep(3 * time.Second)
		c.Remove(id)
		time.Sleep(1 * time.Second)
		// c.Stop()
	}()

	select {}
}
