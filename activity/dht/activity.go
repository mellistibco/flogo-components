package dht

import (
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/d2r2/go-dht"
)

const (
	ivType = "type"
	ivPin = "pin"
  ivRetries = "retries"
	ovTemp = "temp"
	ovHumidity = "humidity"
)

// MyActivity is a stub for your Activity implementation
type MyActivity struct {
	metadata *activity.Metadata
}

// NewActivity creates a new activity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &MyActivity{metadata: metadata}
}

// Metadata implements activity.Activity.Metadata
func (a *MyActivity) Metadata() *activity.Metadata {
	return a.metadata
}

// Eval implements activity.Activity.Eval
func (a *MyActivity) Eval(context activity.Context) (done bool, err error) {

	deviceType := context.GetInput(ivType).(string)
	gpioPin := context.GetInput(ivPin).(int)
	retries := context.GetInput(ivRetries).(int)

	sensorType := dht.DHT22

	if deviceType=="DHT11" {
		sensorType = dht.DHT11
	}

	temperature, humidity, _, err := dht.ReadDHTxxWithRetry(sensorType, gpioPin, false, retries)

  if err != nil {
  	return false, err
  }

	context.SetOutput(ovTemp,temperature)
	context.SetOutput(ovHumidity,humidity)
	return true,nil

}
