package main

import (
	"fmt"
	"goapi/pkg/notice/bark"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

// 监控
func main() {
	for {
		check()
		time.Sleep(time.Second * 60)
	}
}

func check() {
	// 查询当前服务器ip
	res, err := exec.Command("curl", "ip.sb").Output()
	if err != nil {
		fmt.Println("执行命令出错：", err)
		return
	}
	ip := strings.Trim(string(res), "\n\n")
	// 使用 exec.Command 函数创建一个命令对象
	cmd := exec.Command("df", "-h")

	// 执行命令并捕获输出
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("命令执行失败:", err)
		return
	}
	// 将输出拆分为行
	lines := strings.Split(string(output), "\n")
	// 遍历每一行，跳过标题行
	for i, line := range lines {
		if i == 0 {
			// 跳过标题行
			continue
		}
		// 使用空格分割每一行
		fields := strings.Fields(line)
		if len(fields) < 6 {
			continue
		}
		// 提取挂载点和使用率
		mountPoint := fields[5]
		usageStr := fields[4]
		// 去除百分号并转换为整数
		usageStr = strings.TrimSuffix(usageStr, "%")
		usage, err := strconv.Atoi(usageStr)
		if err != nil {
			fmt.Printf("无法解析使用率：%v\n", err)
			continue
		}
		fmt.Printf("挂载点: %s, 使用率: %d%%\n", mountPoint, usage)
		// 如果使用率超过90%，触发报警
		if usage > 70 {
			// 在这里触发报警操作
			bark.Notice(fmt.Sprintf("服务器【%s】磁盘告警", ip), fmt.Sprintf("警告：磁盘使用率超过70%%！挂载点：%s\n", mountPoint))
		} else {
			fmt.Println("一切正常")
		}
	}
}
