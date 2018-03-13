package endpoints

import (
	"testing"
	"k8s.io/api/core/v1"
	"github.com/stretchr/testify/assert"
)

func TestReadyEndpoints(t *testing.T) {
	myEp := buildEndpoints()

	assert.Equal(t, []K8sEndpoint{{address: "1.2.3.4", port: 8012}, {address:"1.2.3.1", port: 8012}}, myEp.ReadyEndpoints(0))
}

func TestReadyEndpointsWithDefaultPort(t *testing.T) {
	myEp := buildEndpoints()

	assert.Equal(t, []K8sEndpoint{{address: "1.2.3.4", port: 1234}, {address:"1.2.3.1", port: 1234}}, myEp.ReadyEndpoints(1234))
}


func TestNotReadyEndpoints(t *testing.T) {
	myEp := buildEndpoints()

	assert.Equal(t, []K8sEndpoint{{address: "3.1.3.4", port: 8012}, {address:"3.1.3.1", port: 8012}}, myEp.NotReadyEndpoints(0))
}

func TestNotReadyEndpointsWithDefaultPort(t *testing.T) {
	myEp := buildEndpoints()

	assert.Equal(t, []K8sEndpoint{{address: "3.1.3.4", port: 9876}, {address:"3.1.3.1", port: 9876}}, myEp.NotReadyEndpoints(9876))
}

func buildEndpoints() *K8sEndpoints {
	return New(v1.Endpoints{
		Subsets: []v1.EndpointSubset{
			{
				Addresses: []v1.EndpointAddress{
					{IP: "1.2.3.4", TargetRef: &v1.ObjectReference{}},
					{IP: "1.2.3.1", TargetRef: &v1.ObjectReference{}},
				},
				NotReadyAddresses: []v1.EndpointAddress{
					{IP: "3.1.3.4", TargetRef: &v1.ObjectReference{}},
					{IP: "3.1.3.1", TargetRef: &v1.ObjectReference{}},
				},
				Ports: []v1.EndpointPort{
					{Port: 8012},
					{Port: 8082},
				},
			},
		},
	})
}

