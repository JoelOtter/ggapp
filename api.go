package ggapp

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/JoelOtter/ggapp/internal/api"
	"github.com/JoelOtter/ggapp/internal/graphql"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
)

const (
	url = "https://api.ggapp.io"
)

type ListGamesForStatusesVariables struct {
	StatusIds []PlayStatusCode `json:"statusIds"`
	UserId    int              `json:"userId"`
	Limit     int              `json:"limit,omitempty"`
	Offset    int              `json:"offset,omitempty"`
}

func ListGamesForStatuses(logger *logrus.Entry, variables ListGamesForStatusesVariables) ([]GameWithStatus, error) {
	requestBody := api.Request{
		OperationName: "listGamesForStatuses",
		Query:         graphql.QueryListGamesForStatuses,
		Variables:     variables,
	}
	requestJson, err := json.Marshal(requestBody)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request JSON: %w", err)
	}
	request, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(requestJson))
	if err != nil {
		return nil, fmt.Errorf("failed to create HTTP request: %w", err)
	}
	request.Header.Add("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, fmt.Errorf("failed to make HTTP request: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("failed to read error body: %w", err)
		}
		logger.Error(string(body))
		return nil, fmt.Errorf("failed to make HTTP request, got response code %d", resp.StatusCode)
	}

	var result struct {
		Data struct {
			Games []GameWithStatus `json:"listGamesForStatuses"`
		} `json:"data"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}
	return result.Data.Games, nil
}
