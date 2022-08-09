// Copyright (c) 2022 The Parca Authors
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package metadata

import (
	"github.com/prometheus/common/model"
)

type Provider struct {
	name      string
	labelFunc func(pid int) (model.LabelSet, error)
}

func (p *Provider) Labels(pid int) (model.LabelSet, error) {
	return p.labelFunc(pid)
}

func (p *Provider) Name() string {
	return p.name
}