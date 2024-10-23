package main

import (
	"fmt"
	"github.com/Leviathangk/go-unselenium/unselenium"
	"github.com/golang/glog"
	"github.com/stitch-june/selenium"
	"log"
	"os"
	"time"
)

func main() {
	// 启动 Driver
	driver, err := unselenium.NewDriver(unselenium.NewConfig(
		unselenium.SetDriverPath("./chromedriver.exe"),
		//unselenium.SetHeadless(),
		unselenium.SetLogLevel(1),
		unselenium.SetUserDataDir("F:\\go_workspace\\unselenium\\datas"),
		unselenium.SetShowLog(),
		unselenium.SetArgs("--user-agent=Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/130.0.0.0 Safari/537.36"),
	))
	if err != nil {
		glog.Fatalln(err)
	}

	// 关闭浏览器及其服务（该方法被重写了）
	defer driver.Quit()

	// 测试 cdp 命令
	driver.ExecuteCDPScript("window.GK = 123;")

	// 检测点通过性查看
	//driver.Get("https://www.google.com/")
	//driver.Get("https://bot.sannysoft.com/")

	// 访问测试
	//driver.Get("https://nowsecure.nl/")

	search(driver)

	// 延迟关闭
	time.Sleep(10 * time.Second)

	// 保存图片测试，验证无头模式正常运行
	file, _ := os.Create("slap.png")
	content, _ := driver.Screenshot()
	file.Write(content)
	time.Sleep(30 * time.Second)
}
func search(wd *unselenium.Driver) {
	err := wd.Get("https://www.google.com/")
	if err != nil {
		panic(err)
	}
	wd.SetImplicitWaitTimeout(10 * time.Second)
	searchBox, err := wd.FindElement(selenium.ByCSSSelector, "textarea[name='q']")
	if err != nil {
		// log.Fatalf("here - 1")
		log.Fatalf("Failed to find search box: %v", err)
	}

	if err != nil {
		// log.Fatalf("here - 4")
		log.Fatalf("Failed to wait: %v", err)
	}

	if err := searchBox.SendKeys("apple"); err != nil {
		log.Fatalf("Failed to enter search query: %v", err)
	}

	if err := searchBox.SendKeys(selenium.EnterKey); err != nil {
		log.Fatalf("Failed to submit: %v", err)
	}

	time.Sleep(2 * time.Second)

	title, err := wd.Title()
	if err != nil {
		log.Fatalf("Failed to get the page tile: %v", err)
	}

	fmt.Printf("Page title is: %s\n", title)

	es, e := wd.FindElements(selenium.ByCSSSelector, "a.gb_z")
	if e != nil {
		log.Fatalf("Failed to find search box: %v", e)
	}
	for _, v := range es {
		fmt.Printf("v: %v\n", v)
		v.Click()
	}
	time.Sleep(30 * time.Second)
}
