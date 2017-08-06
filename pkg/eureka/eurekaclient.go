package eureka

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"time"

	"github.com/carllhw/go-eureka-client-sample/pkg/util"
)

const registerTextTpl = `
{
  "instance": {
    "hostName":"{{.ipAddress}}",
    "app":"vendor",
    "ipAddr":"{{.ipAddress}}",
    "vipAddress":"vendor",
    "status":"UP",
    "port":{
      "$": "{{.port}}",
      "@enabled": "true"
    },
    "securePort" : {
      "$": "8443",
      "@enabled": "false"
    },
    "homePageUrl" : "http://{{.ipAddress}}:{{.port}}/",
    "statusPageUrl": "http://{{.ipAddress}}:{{.port}}/info",
    "healthCheckUrl": "http://{{.ipAddress}}:{{.port}}/health",
    "dataCenterInfo" : {
      "name": "MyOwn",
      "@class": "com.netflix.appinfo.InstanceInfo$DefaultDataCenterInfo"
    },
    "metadata": {
      "instanceId" : "vendor:{{.instanceId}}"
    }
  }
}
`

var (
	instanceId string
	serviceUrl string = "http://localhost:8000/eureka/apps/vendor/"
)

func Register() {
	instanceId = util.GetUUID()

	registerTpl := template.Must(template.New("RegisterInfo").Parse(registerTextTpl))

	var buf bytes.Buffer
	vals := map[string]string{
		"ipAddress":  util.GetLocalIP(),
		"port":       "8080",
		"instanceId": instanceId,
	}
	if err := registerTpl.Execute(&buf, vals); err != nil {
		log.Fatal(err)
	}
	registerInfo := buf.String()

	// Register.
	registerAction := HttpAction{
		Url:    "http://localhost:8000/eureka/apps/vendor",
		Method: "POST",
		Body:   registerInfo,
		Headers: map[string]string{
			"Content-Type":                    "application/json",
			"x-netflix-discovery-replication": "false",
		},
	}
	var result bool
	for {
		result = DoHttpRequest(registerAction)
		if result {
			break
		} else {
			time.Sleep(time.Second * 5)
		}
	}
}

func StartHeartbeat() {
	for {
		time.Sleep(time.Second * 30)
		heartbeat()
	}
}

func heartbeat() {
	heartbeatAction := HttpAction{
		Url:    serviceUrl + util.GetLocalIP() + ":vendor:" + instanceId,
		Method: "PUT",
	}
	DoHttpRequest(heartbeatAction)
}

func Deregister() {
	fmt.Println("Trying to deregister application...")
	// Deregister
	deregisterAction := HttpAction{
		Url:    serviceUrl + util.GetLocalIP() + ":vendor:" + instanceId,
		Method: "DELETE",
	}
	DoHttpRequest(deregisterAction)
	fmt.Println("Deregistered application, exiting. Check Eureka...")
}
