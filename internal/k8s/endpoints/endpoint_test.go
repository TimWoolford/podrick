package endpoints

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/TimWoolford/podrick/internal/config"
)

func TestStatusUrlWithProvidedPath(t *testing.T) {
	endpoint := K8sEndpoint{
		Annotations: map[string]string{},
		address:     "1.2.3.4",
		port:        1122,
		config:      &config.Config{},
	}

	url := endpoint.StatusUrl("/myStatusPath")

	assert.Equal(t, "http://1.2.3.4:1122/myStatusPath", url)
}

func TestStatusUrlWithAnnotatedPath(t *testing.T) {
	endpoint := K8sEndpoint{
		Annotations: map[string]string{"podrick/status": "/aStatusPath"},
		address:     "1.2.3.4",
		port:        1122,
		config:      &config.Config{StatusPathAnnotation: "podrick/status"},
	}

	url := endpoint.StatusUrl("")
	assert.Equal(t, "http://1.2.3.4:1122/aStatusPath", url)
}

func TestStatusUrlWithDefaultPath(t *testing.T) {
	endpoint := K8sEndpoint{
		Annotations: map[string]string{"some": "annotation"},
		address:     "1.2.3.4",
		port:        1122,
		config:      &config.Config{StatusPathAnnotation: "podrick/status", DefaultStatusPath: "/defaultStatus"},
	}

	url := endpoint.StatusUrl("")
	assert.Equal(t, "http://1.2.3.4:1122/defaultStatus", url)
}
