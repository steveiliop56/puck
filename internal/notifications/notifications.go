package notifications

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/briandowns/spinner"
	"github.com/fatih/color"
)

func NotifyNtfy(skipped []string, upgradable []string, uptodate []string, url string) {
	var body = ""
	var spinnerAnimation = []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"}
	var spinner = spinner.New(spinnerAnimation, 100*time.Millisecond, spinner.WithColor("blue"))
	
	spinner.Suffix = " Sending ntfy notification...\n"
	spinner.Start()

	for _, element := range skipped {
		body = body + fmt.Sprintf("⚠️ Server %s skipped, unsupported distro!\n", element)
	}

	for _, element := range upgradable {
		body = body + fmt.Sprintf("🔁 Server %s has an update!\n", element)
	}

	for _, element := range uptodate {
		body = body + fmt.Sprintf("✅ Server %s is up to date!\n", element)
	}

	var req, _ = http.NewRequest("POST", url, strings.NewReader(body))

	req.Header.Set("Title", "Update check finished!")
	req.Header.Set("Tags", "tada")

	var resp, err = http.DefaultClient.Do(req)

	fmt.Println(resp)

	spinner.Stop()

	if err != nil {
		color.Set(color.FgRed)
		fmt.Print("✗ ")
		color.Unset()
		fmt.Printf("Error in making request! Error: %s", err)
		os.Exit(1)
	}

	color.Set(color.FgGreen)
	fmt.Print("✔ ")
	color.Unset()
	fmt.Println("Ntfy notification sent!")
}