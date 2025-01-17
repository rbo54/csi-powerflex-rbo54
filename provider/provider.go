// Copyright © 2019-2022 Dell Inc. or its subsidiaries. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//      http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package provider

import (
	"github.com/dell/csi-md/md"
	nfs "github.com/dell/csi-nfs/nfs"
	"github.com/dell/csi-vxflexos/v2/service"
	"github.com/dell/gocsi"
	logrus "github.com/sirupsen/logrus"
)

// Log init
var Log = logrus.New()

// New returns a new Mock Storage Plug-in Provider.
func New() gocsi.StoragePluginProvider {
	svc := service.New()
	mdsvc := md.New(service.Name)
	service.PutMDService(mdsvc)
	md.PutVcsiService(svc)
	nfssvc := nfs.New(service.Name)
	service.PutNfsService(nfssvc)
	nfs.PutVcsiService(svc)
	return &gocsi.StoragePlugin{
		Controller:                svc,
		Identity:                  svc,
		Node:                      svc,
		BeforeServe:               svc.BeforeServe,
		RegisterAdditionalServers: svc.RegisterAdditionalServers,

		EnvVars: []string{
			// Enable request validation
			gocsi.EnvVarSpecReqValidation + "=true",

			// Enable serial volume access
			gocsi.EnvVarSerialVolAccess + "=true",
		},
	}
}
