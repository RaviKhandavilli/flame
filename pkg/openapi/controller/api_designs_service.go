// Copyright 2022 Cisco Systems, Inc. and its affiliates
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// SPDX-License-Identifier: Apache-2.0

/*
 * Flame REST API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package controller

import (
	"context"
	"fmt"
	"net/http"

	"go.uber.org/zap"

	"github.com/cisco-open/flame/cmd/controller/app/database"
	"github.com/cisco-open/flame/pkg/openapi"
)

// DesignsApiService is a service that implents the logic for the DesignsApiServicer
// This service should implement the business logic for every endpoint for the DesignsApi API.
// Include any external packages or services that will be required by this service.
type DesignsApiService struct {
	dbService database.DBService
}

// NewDesignsApiService creates a default api service
func NewDesignsApiService(dbService database.DBService) openapi.DesignsApiServicer {
	return &DesignsApiService{
		dbService: dbService,
	}
}

// CreateDesign - Create a new design template.
func (s *DesignsApiService) CreateDesign(ctx context.Context, user string, designInfo openapi.DesignInfo) (openapi.ImplResponse, error) {
	var d = openapi.Design{
		Name:        designInfo.Name,
		Description: designInfo.Description,
		Id:          designInfo.Id,
		UserId:      user,
		Schemas:     []openapi.DesignSchema{},
	}

	err := s.dbService.CreateDesign(user, d)
	if err != nil {
		return openapi.Response(http.StatusInternalServerError, nil), fmt.Errorf("failed to create new design: %v", err)
	}

	return openapi.Response(http.StatusCreated, nil), nil
}

// DeleteDesign - Delete design template
func (s *DesignsApiService) DeleteDesign(ctx context.Context, user string, designId string) (openapi.ImplResponse, error) {
	zap.S().Debugf("Received DeleteDesign Delete request: %s | %s", user, designId)

	err := s.dbService.DeleteDesign(user, designId)
	if err != nil {
		return openapi.Response(http.StatusInternalServerError, nil),
			fmt.Errorf("failed to delete design: %v", err)
	}
	return openapi.Response(http.StatusOK, nil), nil
}

// GetDesign - Get design template information
func (s *DesignsApiService) GetDesign(ctx context.Context, user string, designId string) (openapi.ImplResponse, error) {
	info, err := s.dbService.GetDesign(user, designId)
	if err != nil {
		return openapi.Response(http.StatusInternalServerError, nil), fmt.Errorf("failed to get design: %v", err)
	}

	return openapi.Response(http.StatusOK, info), nil
}

// GetDesigns - Get list of all the designs created by the user.
func (s *DesignsApiService) GetDesigns(ctx context.Context, user string, limit int32) (openapi.ImplResponse, error) {
	designList, err := s.dbService.GetDesigns(user, limit)
	if err != nil {
		return openapi.Response(http.StatusInternalServerError, nil), fmt.Errorf("failed to get list of designs: %v", err)
	}

	return openapi.Response(http.StatusOK, designList), nil
}
