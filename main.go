package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
)

func main() {
    fmt.Println(time.Now())
    // Initialize options
    // Read the extension file contents
    extensionData, err := ioutil.ReadFile("data/hivekeychain.crx")
    if err != nil {
        println((1))
        // Handle error
    }

    // Encode the extension data as base64
    extensionBase64 := base64.StdEncoding.EncodeToString(extensionData)
    chromeOptions := chrome.Capabilities{
        Path: "",
        Args: []string{
            "--no-sandbox",
            "--disable-dev-shm-usage",
            "--disable-setuid-sandbox",
            "--disable-backgrounding-occluded-windows",
            "--disable-background-timer-throttling",
            "--disable-translate",
            "--disable-popup-blocking",
            "--disable-infobars",
            "--disable-gpu",
            "--disable-blink-features=AutomationControlled",
            "--mute-audio",
            "--ignore-certificate-errors",
            "--allow-running-insecure-content",
            "--window-size=300,600",
            "--headless=new",
        },
        Extensions: []string{extensionBase64},
        Prefs: map[string]interface{}{
            "profile.managed_default_content_settings.images":          1,
            "profile.managed_default_content_settings.cookies":         1,
            "profile.managed_default_content_settings.javascript":      1,
            "profile.managed_default_content_settings.plugins":         1,
            "profile.default_content_setting_values.notifications":     2,
            "profile.managed_default_content_settings.stylesheets":     2,
            "profile.managed_default_content_settings.popups":          2,
            "profile.managed_default_content_settings.geolocation":     2,
            "profile.managed_default_content_settings.media_stream":    2,
        },
        ExcludeSwitches: []string{
            "enable-automation",
            "enable-logging",
        },
    }

    caps := selenium.Capabilities{}
    caps.AddChrome(chromeOptions)


    // Start a new ChromeDriver instance
    wd, err := selenium.NewChromeDriverService("webdrivers/chromedriver.exe", 9515)
    if err != nil {
        fmt.Printf("Failed to create ChromeDriver service: %s\n", err)
        os.Exit(1)
    }
    defer wd.Stop()

    // Create a new WebDriver instance
    driver, err := selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d/wd/hub", 9515))
    if err != nil {
        fmt.Printf("Failed to create WebDriver: %s\n", err)
        os.Exit(1)
    }
    defer driver.Quit()

    // Navigate to a web page
    login("","",driver,err)



    screenshot, err := driver.Screenshot()
    if err != nil {
        fmt.Printf("Failed to take screenshot: %s\n", err)
        os.Exit(1)
    }

    // write the screenshot to a file
    if err := ioutil.WriteFile("screenshot.png", screenshot, 0644); err != nil {
        fmt.Printf("Failed to write screenshot to file: %s\n", err)
        os.Exit(1)
    }
    
}
func elementWaitAndClick(wd selenium.WebDriver, xpath string){
    byXpath := selenium.ByXPATH
    for {
        element, err := wd.FindElement(byXpath, xpath)
        if err != nil {
            panic(err)
        }
        isEnabled, err := element.IsEnabled()
        if err != nil {
            panic(err)
        }
        if isEnabled {
            err = element.Click()
            if err != nil {
                panic(err)
            }
            break
        }
        time.Sleep(1 * time.Second)
    }
}
func login(userName string, postingKey string, wd selenium.WebDriver,err error) {
	
    err = wd.SetImplicitWaitTimeout(5 * time.Second)
    if err != nil {
        panic(err)
    }
    err = wd.Get("chrome-extension://jcacnejopjdphbnjgfaaobbfafkihpep/popup.html")
    if err != nil {
        panic(err)
    }
    
    elementWaitAndClick(wd,"/html/body/div/div/div[4]/div[2]/div[5]/button")

	el, _ := wd.FindElement(selenium.ByXPATH, "/html/body/div/div/div[1]/div/div[1]/div/input")
	el.SendKeys("Aa123Aa123!!")
	el, _ = wd.FindElement(selenium.ByXPATH, "/html/body/div/div/div[1]/div/div[2]/div/input")
	el.SendKeys("Aa123Aa123!!")
	el, _ = wd.FindElement(selenium.ByXPATH, "/html/body/div/div/div[1]/button/div")
	el.Click()
	el, _ = wd.FindElement(selenium.ByXPATH, "/html/body/div/div/div[1]/div[2]/div/div[2]/button[1]/div")
	el.Click()
	el, _ = wd.FindElement(selenium.ByXPATH, "/html/body/div/div/div[1]/div[2]/div/div[2]/div[1]/div/input")
	el.SendKeys(userName)
	el, _ = wd.FindElement(selenium.ByXPATH, "/html/body/div/div/div[1]/div[2]/div/div[2]/div[2]/div/input")
	el.SendKeys(postingKey)
	el, _ = wd.FindElement(selenium.ByXPATH, "/html/body/div/div/div[1]/div[2]/div/div[2]/div[2]/div/input")
	time.Sleep(1*time.Second)
    el.SendKeys("\ue007")
    err = wd.ResizeWindow("bigger",1565,1080)
    if err != nil{
        println("can not change size")
    }

	
	// wd.SetWindowSize(1565, 1080)
	wd.Get("https://splinterforge.io/#/")
    el, _ = wd.FindElement(selenium.ByXPATH, "/html/body/app/div[1]/div[1]/app-header/success-modal/section/div[1]/div[4]/div/button")
    el.Click()
    el, _ = wd.FindElement(selenium.ByXPATH, "/html/body/app/div[1]/div[1]/app-header/section/div[4]/div[2]/div/div/a/div[1]")
    el.Click()
    el, _ = wd.FindElement(selenium.ByXPATH, "/html/body/app/div[1]/login-modal/div/div/div/div[2]/div[2]/input")
    el.SendKeys(userName)
    el, _ = wd.FindElement(selenium.ByXPATH, "/html/body/app/div[1]/login-modal/div/div/div/div[2]/div[3]/button")
    el.Click()
    for {
        handles, _ := wd.WindowHandles()
        if len(handles) == 2 {
            break
        }
    }
    handles, _ := wd.WindowHandles()
    wd.SwitchWindow(handles[1])
    el, _ = wd.FindElement(selenium.ByXPATH, "/html/body/div/div/div/div[3]/div[1]/div/div")
    el.Click()
    el, _ = wd.FindElement(selenium.ByXPATH, "/html/body/div/div/div/div[3]/div[2]/button[2]/div")
    el.Click()
    wd.SwitchWindow(handles[0])
    println("success log in")
    fmt.Println(time.Now())
    time.Sleep(5*time.Second)

}