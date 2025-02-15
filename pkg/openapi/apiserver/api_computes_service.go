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

package apiserver

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"go.uber.org/zap"

	"github.com/cisco-open/flame/pkg/openapi"
	"github.com/cisco-open/flame/pkg/restapi"
	"github.com/cisco-open/flame/pkg/util"
)

// ComputesApiService is a service that implements the logic for the ComputesApiServicer
// This service should implement the business logic for every endpoint for the ComputesApi API.
// Include any external packages or services that will be required by this service.
type ComputesApiService struct {
}

// NewComputesApiService creates a default api service
func NewComputesApiService() openapi.ComputesApiServicer {
	return &ComputesApiService{}
}

// DeleteCompute - Delete compute cluster specification
func (s *ComputesApiService) DeleteCompute(ctx context.Context, computeId string,
	xAPIKEY string) (openapi.ImplResponse, error) {
	// TODO - update DeleteCompute with the required logic for this service method.
	// Add api_computes_service.go to the .openapi-generator-ignore to avoid overwriting
	// this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, {}) or use other options such as http.Ok ...
	//return Response(200, nil),nil

	//TODO: Uncomment the next line to return response Response(404, {}) or use other options such as http.Ok ...
	//return Response(404, nil),nil

	//TODO: Uncomment the next line to return response Response(401, {}) or use other options such as http.Ok ...
	//return Response(401, nil),nil

	//TODO: Uncomment the next line to return response Response(0, Error{}) or use other options such as http.Ok ...
	//return Response(0, Error{}), nil

	return openapi.Response(http.StatusNotImplemented, nil), errors.New("DeleteCompute method not implemented")
}

// GetComputeConfig - Get configuration for a compute cluster
func (s *ComputesApiService) GetComputeConfig(ctx context.Context, computeId string,
	xAPIKEY string) (openapi.ImplResponse, error) {
	// TODO - update GetComputeConfig with the required logic for this service method.
	// Add api_computes_service.go to the .openapi-generator-ignore to avoid overwriting
	// this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, ComputeSpec{}) or use other options such as http.Ok ...
	//return Response(200, ComputeSpec{}), nil

	//TODO: Uncomment the next line to return response Response(0, Error{}) or use other options such as http.Ok ...
	//return Response(0, Error{}), nil

	return openapi.Response(http.StatusNotImplemented, nil), errors.New("GetComputeConfig method not implemented")
}

// GetComputeStatus - Get status of a given compute cluster
func (s *ComputesApiService) GetComputeStatus(ctx context.Context, computeId string,
	xAPIKEY string) (openapi.ImplResponse, error) {
	// TODO - update GetComputeStatus with the required logic for this service method.
	// Add api_computes_service.go to the .openapi-generator-ignore to avoid overwriting
	// this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, ComputeStatus{}) or use other options such as http.Ok ...
	//return Response(200, ComputeStatus{}), nil

	//TODO: Uncomment the next line to return response Response(0, Error{}) or use other options such as http.Ok ...
	//return Response(0, Error{}), nil

	return openapi.Response(http.StatusNotImplemented, nil), errors.New("GetComputeStatus method not implemented")
}

// GetDeploymentConfig - Get the deployment config for a job for a compute cluster
func (s *ComputesApiService) GetDeploymentConfig(ctx context.Context, computeId string,
	jobId string, xAPIKEY string) (openapi.ImplResponse, error) {
	// TODO - add logic later to validate the request coming from the deployer with ApiKey
	// Report error if apikey for the deployer does not match with the apiserver cache or from db

	// create controller request
	uriMap := map[string]string{
		"computeId": computeId,
		"jobId":     jobId,
	}
	url := restapi.CreateURL(HostEndpoint, restapi.GetDeploymentConfigEndpoint, uriMap)

	// send post request
	code, resp, err := restapi.HTTPGet(url)
	if err != nil {
		errMsg := fmt.Sprintf("failed to send get request to controller: %v", err)
		zap.S().Errorf(errMsg)
		return openapi.Response(http.StatusInternalServerError, nil), err
	}

	if err = restapi.CheckStatusCode(code); err != nil {
		return openapi.Response(code, nil), fmt.Errorf("%s", string(resp))
	}
	deploymentConfig := openapi.DeploymentConfig{}
	err = json.Unmarshal(resp, &deploymentConfig)
	if err != nil {
		errMsg := fmt.Sprintf("failed to send get deployment config from controller: %v", err)
		zap.S().Errorf(errMsg)
		return openapi.Response(http.StatusInternalServerError, nil), err
	}

	zap.S().Infof("Successfully got deployment config %v from controller for jobId %s and computeId %s", deploymentConfig, jobId, computeId)

	return openapi.Response(http.StatusOK, deploymentConfig), nil
}

// GetDeploymentStatus - Get the deployment status for a job on a compute cluster
func (s *ComputesApiService) GetDeploymentStatus(ctx context.Context, computeId string,
	jobId string, xAPIKEY string) (openapi.ImplResponse, error) {
	// TODO - update GetDeploymentStatus with the required logic for this service method.
	// Add api_computes_service.go to the .openapi-generator-ignore to avoid overwriting
	// this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, DeploymentStatus{}) or use other options such as http.Ok ...
	//return Response(200, DeploymentStatus{}), nil

	//TODO: Uncomment the next line to return response Response(0, Error{}) or use other options such as http.Ok ...
	//return Response(0, Error{}), nil

	return openapi.Response(http.StatusNotImplemented, nil), errors.New("GetDeploymentStatus method not implemented")
}

// GetDeployments - Get all deployments within a compute cluster
func (s *ComputesApiService) GetDeployments(ctx context.Context, computeId string,
	xAPIKEY string) (openapi.ImplResponse, error) {
	// TODO - update GetDeployments with the required logic for this service method.
	// Add api_computes_service.go to the .openapi-generator-ignore to avoid overwriting
	// this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, []DeploymentStatus{}) or use other options such as http.Ok ...
	//return Response(200, []DeploymentStatus{}), nil

	//TODO: Uncomment the next line to return response Response(0, Error{}) or use other options such as http.Ok ...
	//return Response(0, Error{}), nil

	return openapi.Response(http.StatusNotImplemented, nil), errors.New("GetDeployments method not implemented")
}

// PutDeploymentStatus - Add or update the deployment status for a job on a compute cluster
func (s *ComputesApiService) PutDeploymentStatus(ctx context.Context, computeId string,
	jobId string, xAPIKEY string, requestBody map[string]openapi.AgentState) (openapi.ImplResponse, error) {
	// create controller request
	uriMap := map[string]string{
		"computeId": computeId,
		"jobId":     jobId,
	}
	url := restapi.CreateURL(HostEndpoint, restapi.PutDeploymentStatusEndpoint, uriMap)
	code, resp, err := restapi.HTTPPut(url, requestBody, "application/json; charset=utf-8")
	if err != nil {
		zap.S().Errorf("failed to post deployment status to controller: %v", err)
		return openapi.Response(http.StatusInternalServerError, nil), err
	}

	if err = restapi.CheckStatusCode(code); err != nil {
		return openapi.Response(code, nil), fmt.Errorf("%s", string(resp))
	}
	zap.S().Infof("Successfully sent deployment status %v to controller for jobId %s and computeId %s",
		requestBody, jobId, computeId)

	return openapi.Response(http.StatusOK, string(resp)), nil
}

// RegisterCompute - Register a new compute cluster
func (s *ComputesApiService) RegisterCompute(ctx context.Context, computeSpec openapi.ComputeSpec) (openapi.ImplResponse, error) {
	zap.S().Debugf("Registering new compute cluster with computeSpec: %v", computeSpec)

	// create controller request
	uriMap := map[string]string{}
	url := restapi.CreateURL(HostEndpoint, restapi.RegisterComputeEndpoint, uriMap)

	// send post request
	code, resp, err := restapi.HTTPPost(url, computeSpec, "application/json")
	if err != nil {
		errMsg := fmt.Sprintf("failed to send post to controller: %v", err)
		zap.S().Errorf(errMsg)
		return openapi.Response(http.StatusInternalServerError, nil), err
	}

	if err = restapi.CheckStatusCode(code); err != nil {
		return openapi.Response(code, nil), fmt.Errorf("%s", string(resp))
	}

	// Once everything goes well, the ComputeStatus is updated
	computeStatus := openapi.ComputeStatus{}
	err = util.ByteToStruct(resp, &computeStatus)
	if err != nil {
		errMsg := fmt.Sprintf("failed to construct response message: %v", err)
		zap.S().Errorf(errMsg)
		return openapi.Response(http.StatusInternalServerError, nil), fmt.Errorf(errMsg)
	}

	return openapi.Response(http.StatusCreated, computeStatus), nil
}

// UpdateCompute - Update a compute cluster&#39;s specification
func (s *ComputesApiService) UpdateCompute(ctx context.Context, computeId string,
	xAPIKEY string, computeSpec openapi.ComputeSpec) (openapi.ImplResponse, error) {
	// TODO - update UpdateCompute with the required logic for this service method.
	// Add api_computes_service.go to the .openapi-generator-ignore to avoid overwriting
	// this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, {}) or use other options such as http.Ok ...
	//return Response(200, nil),nil

	//TODO: Uncomment the next line to return response Response(401, {}) or use other options such as http.Ok ...
	//return Response(401, nil),nil

	//TODO: Uncomment the next line to return response Response(0, Error{}) or use other options such as http.Ok ...
	//return Response(0, Error{}), nil

	return openapi.Response(http.StatusNotImplemented, nil), errors.New("UpdateCompute method not implemented")
}
