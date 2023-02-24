package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"

	// "time"

	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
)

func main() {
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
    if err := driver.Get("chrome-extension://jcacnejopjdphbnjgfaaobbfafkihpep/popup.html"); err != nil {
        fmt.Printf("Failed to load page: %s\n", err)
        os.Exit(1)
    }


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