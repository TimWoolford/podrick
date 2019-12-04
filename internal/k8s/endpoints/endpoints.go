package endpoints

import (
	"github.com/TimWoolford/podrick/internal/config"
	"k8s.io/api/core/v1"
)

type K8sEndpoints struct {
	endpoints *v1.Endpoints
	config    *config.Config
}

func New(eps *v1.Endpoints, config *config.Config) *K8sEndpoints {
	return &K8sEndpoints{endpoints: eps, config: config}
}

func Empty() *K8sEndpoints {
	return &K8sEndpoints{}
}

func (eps K8sEndpoints) ReadyEndpoints(expectedPort int32) []K8sEndpoint {
	return eps.endpointsFor(expectedPort, readyEndpoints)
}

func (eps K8sEndpoints) NotReadyEndpoints(expectedPort int32) []K8sEndpoint {
	return eps.endpointsFor(expectedPort, notReadyEndpoints)
}

func readyEndpoints(subset v1.EndpointSubset) []v1.EndpointAddress {
	return subset.Addresses
}

func notReadyEndpoints(subset v1.EndpointSubset) []v1.EndpointAddress {
	return subset.NotReadyAddresses
}

func (eps K8sEndpoints) endpointsFor(expectedPort int32, srcFunc func(subset v1.EndpointSubset) []v1.EndpointAddress) []K8sEndpoint {
	port := firstPort(expectedPort, eps.endpoints.Subsets)

	var addresses []K8sEndpoint
	for _, ss := range eps.endpoints.Subsets {
		for _, address := range srcFunc(ss) {
			addresses = append(addresses, K8sEndpoint{
				Name:        address.TargetRef.Name,
				Namespace:   eps.endpoints.Namespace,
				Annotations: eps.endpoints.Annotations,
				address:     address.IP,
				port:        port,
				config:      eps.config,
			})
		}
	}
	return addresses
}

func firstPort(suggestedPort int32, subsets []v1.EndpointSubset) int32 {
	if suggestedPort > 0 {
		return suggestedPort
	}

	for _, ss := range subsets {
		for _, port := range ss.Ports {
			return port.Port
		}
	}
	return 0
}
