/*
Copyright SecureKey Technologies Inc. All Rights Reserved.
SPDX-License-Identifier: Apache-2.0
*/

package discovery

import (
	"github.com/trustbloc/trustbloc-did-method/pkg/vdri/trustbloc/models"
)

// MockDiscoveryService implements a mock discovery service
type MockDiscoveryService struct {
	GetEndpointsFunc func(domain string) ([]*models.Endpoint, error)
}

// GetEndpoints discover endpoints from a consortium
func (m *MockDiscoveryService) GetEndpoints(domain string) ([]*models.Endpoint, error) {
	if m.GetEndpointsFunc != nil {
		return m.GetEndpointsFunc(domain)
	}

	return nil, nil
}
