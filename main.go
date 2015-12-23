package main

import (
	"bytes"
	"fmt"
	"net/http"
	"net/url"
	"os"

	"github.com/drone/drone-go/drone"
	"github.com/drone/drone-go/plugin"
)

var (
	build                       string
	buildDate                   string
	newRelicDeploymentsEndpoint string = "https://api.newrelic.com/deployments.xml"
)

func main() {
	fmt.Printf("Drone New Relic Plugin built at %s\n", buildDate)

	build := drone.Build{}
	vargs := Params{}

	plugin.Param("build", &build)
	plugin.Param("vargs", &vargs)
	plugin.MustParse()

	if len(vargs.LicenseKey) == 0 {
		fmt.Println("Please provide your New Relic licencse key.")
		os.Exit(1)
		return
	}

	if len(vargs.ApplicationId) == 0 && len(vargs.AppName) == 0 {
		fmt.Println("Please provide either application_id or app_name option.")
		os.Exit(2)
		return
	}

	if len(vargs.User) == 0 {
		vargs.User = "drone-ci"
	}

	data := url.Values{}
	data.Set("application_id", vargs.ApplicationId)
	data.Set("app_name", vargs.AppName)
	data.Set("description", vargs.Description)
	data.Set("user", vargs.User)
	data.Set("revision", build.Commit)

	payload := bytes.NewBufferString(data.Encode())
	request, _ := http.NewRequest("POST", newRelicDeploymentsEndpoint, payload)
	request.Header.Add("X-Api-Key", vargs.LicenseKey)
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	// request.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	fmt.Println("=> Posting deployment notification to New Relic")
	fmt.Printf("%s\n", data.Encode())

	client := &http.Client{}
	resp, _ := client.Do(request)
	fmt.Printf("New Relic response code: %s\n", resp.Status)
}
