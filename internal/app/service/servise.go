package service

import "time"

type Service struct {
}

func New() *Service {
	return &Service{}
}

func (s *Service) WhatTime() string {
	return time.Now().Local().String()
}
