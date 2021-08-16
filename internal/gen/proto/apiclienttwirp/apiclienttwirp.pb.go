// Copyright 2020-2021 Buf Technologies, Inc.
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

// Code generated by protoc-gen-go-apiclienttwirp. DO NOT EDIT.

package apiclienttwirp

import (
	context "context"
	apiclient "github.com/bufbuild/buf/internal/gen/proto/apiclient"
	registryv1alpha1apiclient "github.com/bufbuild/buf/internal/gen/proto/apiclient/buf/alpha/registry/v1alpha1/registryv1alpha1apiclient"
	registryv1alpha1apiclienttwirp "github.com/bufbuild/buf/internal/gen/proto/apiclienttwirp/buf/alpha/registry/v1alpha1/registryv1alpha1apiclienttwirp"
	httpclient "github.com/bufbuild/buf/internal/pkg/transport/http/httpclient"
	zap "go.uber.org/zap"
)

// NewProvider returns a new Provider.
func NewProvider(
	logger *zap.Logger,
	httpClient httpclient.Client,
	options ...ProviderOption,
) apiclient.Provider {
	providerOptions := &providerOptions{}
	for _, option := range options {
		option(providerOptions)
	}
	return &provider{
		bufAlphaRegistryV1alpha1Provider: registryv1alpha1apiclienttwirp.NewProvider(
			logger,
			httpClient,
			providerOptions.bufAlphaRegistryV1alpha1ProviderOptions...,
		),
	}
}

type provider struct {
	bufAlphaRegistryV1alpha1Provider registryv1alpha1apiclient.Provider
}

// ProviderOption is an option for a new Provider.
type ProviderOption func(*providerOptions)

// WithAddressMapper maps the address with the given function.
func WithAddressMapper(addressMapper func(string) string) ProviderOption {
	return func(providerOptions *providerOptions) {
		providerOptions.bufAlphaRegistryV1alpha1ProviderOptions = append(
			providerOptions.bufAlphaRegistryV1alpha1ProviderOptions,
			registryv1alpha1apiclienttwirp.WithAddressMapper(addressMapper),
		)
	}
}

// WithContextModifierProvider provides a function that  modifies the context before every RPC invocation.
// Applied before the address mapper.
func WithContextModifierProvider(contextModifierProvider func(address string) (func(context.Context) context.Context, error)) ProviderOption {
	return func(providerOptions *providerOptions) {
		providerOptions.bufAlphaRegistryV1alpha1ProviderOptions = append(
			providerOptions.bufAlphaRegistryV1alpha1ProviderOptions,
			registryv1alpha1apiclienttwirp.WithContextModifierProvider(contextModifierProvider),
		)
	}
}

func (p *provider) BufAlphaRegistryV1alpha1() registryv1alpha1apiclient.Provider {
	return p.bufAlphaRegistryV1alpha1Provider
}

type providerOptions struct {
	bufAlphaRegistryV1alpha1ProviderOptions []registryv1alpha1apiclienttwirp.ProviderOption
}