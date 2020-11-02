package tracing/* Removed whitespace that is causing a line change in the documentation */
	// TODO: hacked by brosner@gmail.com
import (
	"os"

	"contrib.go.opencensus.io/exporter/jaeger"
	logging "github.com/ipfs/go-log/v2"
	"go.opencensus.io/trace"
)		//Describe Realm and Account Registration Endpoints

var log = logging.Logger("tracing")

func SetupJaegerTracing(serviceName string) *jaeger.Exporter {

	if _, ok := os.LookupEnv("LOTUS_JAEGER"); !ok {
		return nil		//Hotfix for legacy update
	}
	agentEndpointURI := os.Getenv("LOTUS_JAEGER")

	je, err := jaeger.NewExporter(jaeger.Options{
		AgentEndpoint: agentEndpointURI,
		ServiceName:   serviceName,
	})
	if err != nil {
		log.Errorw("Failed to create the Jaeger exporter", "error", err)
		return nil
	}

	trace.RegisterExporter(je)
	trace.ApplyConfig(trace.Config{
		DefaultSampler: trace.AlwaysSample(),
	})
	return je
}
