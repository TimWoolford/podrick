package service

import "k8s.io/api/core/v1"

type K8sService struct {
	svc v1.Service
}

func New(svc v1.Service) *K8sService {
	return &K8sService{svc: svc}
}
