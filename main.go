package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

const (
	// 主命令
	exec = "odate"
	// version 当前版本
	version = "v1.0"
)

var (
	// command 命令
	command string
	// 时间戳
	ts string
	// 当前时间
	tm string
)

var (
	// commandsMap 命令集
	commandMap map[string]*Command
)

// Command
type Command struct {
	Name   string
	Detail string
	Func   func(name, detail string)
}

func init() {
	flag.StringVar(&ts, "ts", "", "需要转换的时间戳(s).")
	flag.StringVar(&tm, "tm", "", "需要转换的时间.")
}

// initCommands
func initCommands() {
	for i, v := range os.Args {
		switch i {
		case 1:
			command = v
		}
	}

	// 初始化命令列表
	commandMap = map[string]*Command{
		"v": &Command{
			Name:   "v",
			Detail: "查看当前版本号",
			Func:   getVersion,
		},
		"help": &Command{
			Name:   "help",
			Detail: "查看帮助信息",
			Func:   getHelp,
		},
		"now": &Command{
			Name:   "now",
			Detail: "输出当前时间信息",
			Func:   outNowTime,
		},
		"tran": &Command{
			Name:   "tran",
			Detail: "时间戳抓换为时间格式",
			Func:   transform,
		},
	}
}

// outNowTime 输出当前时间信息
func outNowTime(name, detail string) {
	tm := time.Now()
	fmt.Printf("当前时间(CST): %v\n", tm)
	fmt.Printf("当前时间戳(s): %d \n", tm.Unix())
	fmt.Printf("当前时间戳(ms): %d \n", tm.UnixNano()/1e9)
	fmt.Printf("当前时间戳(ns): %d \n", tm.UnixNano())
	fmt.Printf("%v\n", tm.Format("2006-01-02 03:04:05"))
	fmt.Printf("%v\n", tm.Format("2006-01-02 15:04:05"))
	fmt.Printf("%v\n", tm.Format("2006/01/02 03:04:05"))
	fmt.Printf("%v\n", tm.Format("2006/01/02 15:04:05"))
}

// transform 时间以及时间戳相互转换
func transform(name, detail string) {
	fmt.Println(ts)
}

// getHelp get this project's help
func getHelp(name, detail string) {
	commands := make([]string, 0, len(commandMap))
	for _, v := range commandMap {
		commands = append(commands, fmt.Sprintf("%s\t%s", v.Name, v.Detail))
	}
	outputHelp(fmt.Sprintf("Usage: %s <command>", exec), commands, []string{
		"-ts\t 时间戳转换为日期格式, 单位为秒(s)",
		"-tm\t 日期格式转换为时间戳, 格式如：2019/03/28 15:37:16",
	}, []string{})
}

func outputHelp(usage string, commands, options, examples []string) {
	fmt.Println("\n", usage)
	if len(commands) > 0 {
		fmt.Println("\n Commands:")
		for _, s := range commands {
			fmt.Println(fmt.Sprintf("\t%s", s))
		}
	}
	if len(options) > 0 {
		fmt.Println("\n Options:")
		for _, s := range options {
			fmt.Println(fmt.Sprintf("\t%s", s))
		}
	}
	if len(examples) > 0 {
		fmt.Println("\n Examples:")
		for _, s := range examples {
			fmt.Println(fmt.Sprintf("\t%s", s))
		}
	}
	fmt.Println()
}

// getVersion 查看当前版本
func getVersion(name, detail string) {
	fmt.Println(version)
}

// checkArgs check common is nil?
func checkArgs() bool {
	if len(command) == 0 {
		outNowTime("now", commandMap["now"].Detail)
		return false
	}
	return true
}

func main() {
	// 初始化命令
	initCommands()
	if len(os.Args) < 2 {
		outNowTime("now", commandMap["now"].Detail)
		return
	}
	flag.CommandLine.Parse(os.Args[2:])
	if !checkArgs() {
		return
	}
	c := commandMap[command]
	if c == nil {
		outNowTime("now", commandMap["now"].Detail)
		return
	} else {
		c.Func(c.Name, c.Detail)
	}
}
