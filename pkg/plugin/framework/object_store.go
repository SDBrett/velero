/*
Copyright 2019 the Velero contributors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package framework

import (
	"github.com/hashicorp/go-plugin"
	"google.golang.org/grpc"

	proto "github.com/heptio/velero/pkg/plugin/generated"
)

// ObjectStorePlugin is an implementation of go-plugin's Plugin
// interface with support for gRPC for the cloudprovider/ObjectStore
// interface.
type ObjectStorePlugin struct {
	plugin.NetRPCUnsupportedPlugin
	*pluginBase
}

// GRPCClient returns an ObjectStore gRPC client.
func (p *ObjectStorePlugin) GRPCClient(c *grpc.ClientConn) (interface{}, error) {
	return newClientDispenser(p.clientLogger, c, newObjectStoreGRPCClient), nil

}

// GRPCServer registers an ObjectStore gRPC server.
func (p *ObjectStorePlugin) GRPCServer(s *grpc.Server) error {
	proto.RegisterObjectStoreServer(s, &ObjectStoreGRPCServer{mux: p.serverMux})
	return nil
}
