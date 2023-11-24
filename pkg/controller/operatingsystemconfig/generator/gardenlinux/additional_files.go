// Copyright (c) 2022 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file
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

package gardenlinux

import (
	"path/filepath"

	"github.com/gardener/gardener/extensions/pkg/controller/operatingsystemconfig/oscommon/generator"

	"github.com/gardener/gardener-extension-os-gardenlinux/pkg/gardenlinux"
)

var (
	fileList = []string{
		"g_functions.sh",
	}
)

// GetAdditionalScripts returns additional scripts that were provided as raw files in the embedded fs
func GetAdditionalScripts() ([]*generator.File, error) {
	files := []*generator.File{}

	for _, f := range fileList {
		scriptContent, err := gardenlinux.Templates.ReadFile(filepath.Join("scripts", f))
		if err != nil {
			return nil, err
		}

		additionalScript := &generator.File{
			Path:        filepath.Join(gardenlinux.ScriptLocation, f),
			Content:     scriptContent,
			Permissions: &gardenlinux.ScriptPermissions,
		}

		files = append(files, additionalScript)
	}

	return files, nil
}
