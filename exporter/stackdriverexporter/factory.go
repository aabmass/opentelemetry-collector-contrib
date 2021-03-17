// Copyright 2021, OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package stackdriverexporter

import (
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config/configmodels"

	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/googlecloudexporter"
)

const (
	// The value of "type" key in configuration.
	configType = configmodels.Type("stackdriver")
)

// NewFactory creates a factory for the stackdriver exporter
func NewFactory() component.ExporterFactory {
	return &factory{}
}

type factory struct{ component.ExporterFactory }

func (*factory) Type() configmodels.Type { return configType }

func (f *factory) CreateDefaultConfig() configmodels.Exporter {
	cfg := f.ExporterFactory.CreateDefaultConfig().(*googlecloudexporter.Config)
	cfg.TypeVal = configType
	return cfg
}

// Segfaults with:

// 2021-03-17T14:15:37.623Z	info	service/service.go:411	Starting OpenTelemetry Contrib Collector...	{"Version": "v0.22.0-34-g1dd45017", "GitHash": "1dd45017", "NumCPU": 16}
// 2021-03-17T14:15:37.623Z	info	service/service.go:255	Setting up own telemetry...
// 2021-03-17T14:15:37.625Z	info	service/telemetry.go:102	Serving Prometheus metrics	{"address": ":8888", "level": 0, "service.instance.id": "15341967-0260-48bd-abb5-cf1c007f743b"}
// panic: runtime error: invalid memory address or nil pointer dereference
// [signal SIGSEGV: segmentation violation code=0x1 addr=0x18 pc=0x2a04d2d]

// goroutine 1 [running]:
// github.com/open-telemetry/opentelemetry-collector-contrib/exporter/stackdriverexporter.(*factory).CreateDefaultConfig(0xc00062c9c0, 0x3f85b08, 0xc000902b40)
// 	/usr/local/google/home/aaronabbott/repo/opentelemetry-collector-contrib/exporter/stackdriverexporter/factory.go:39 +0x2d
// go.opentelemetry.io/collector/config/configcheck.ValidateConfigFromFactories(0xc00062e930, 0xc00062f710, 0xc00062ec30, 0xc00062e900, 0xc0005eea80, 0x0)
// 	/usr/local/google/home/aaronabbott/go/pkg/mod/go.opentelemetry.io/collector@v0.22.1-0.20210313012550-03904de3dd61/config/configcheck/configcheck.go:47 +0x3ef
// go.opentelemetry.io/collector/service.(*Application).setupConfigurationComponents(0xc0001e4900, 0x3f9f090, 0xc000050088, 0x3abdbf8, 0x0, 0x3)
// 	/usr/local/google/home/aaronabbott/go/pkg/mod/go.opentelemetry.io/collector@v0.22.1-0.20210313012550-03904de3dd61/service/service.go:288 +0x5a
// go.opentelemetry.io/collector/service.(*Application).execute(0xc0001e4900, 0x3f9f090, 0xc000050088, 0x3abdbf8, 0x0, 0x0)
// 	/usr/local/google/home/aaronabbott/go/pkg/mod/go.opentelemetry.io/collector@v0.22.1-0.20210313012550-03904de3dd61/service/service.go:429 +0x47e
// go.opentelemetry.io/collector/service.New.func1(0xc000640500, 0xc00059cd20, 0x0, 0x2, 0x0, 0x0)
// 	/usr/local/google/home/aaronabbott/go/pkg/mod/go.opentelemetry.io/collector@v0.22.1-0.20210313012550-03904de3dd61/service/service.go:155 +0xa7
// github.com/spf13/cobra.(*Command).execute(0xc000640500, 0xc00004c1c0, 0x2, 0x2, 0xc000640500, 0xc00004c1c0)
// 	/usr/local/google/home/aaronabbott/go/pkg/mod/github.com/spf13/cobra@v1.1.3/command.go:852 +0x472
// github.com/spf13/cobra.(*Command).ExecuteC(0xc000640500, 0xc00062f710, 0xc00062ec30, 0xc00062e900)
// 	/usr/local/google/home/aaronabbott/go/pkg/mod/github.com/spf13/cobra@v1.1.3/command.go:960 +0x375
// github.com/spf13/cobra.(*Command).Execute(...)
// 	/usr/local/google/home/aaronabbott/go/pkg/mod/github.com/spf13/cobra@v1.1.3/command.go:897
// go.opentelemetry.io/collector/service.(*Application).Run(...)
// 	/usr/local/google/home/aaronabbott/go/pkg/mod/go.opentelemetry.io/collector@v0.22.1-0.20210313012550-03904de3dd61/service/service.go:482
// main.runInteractive(0xc00062e930, 0xc00062f710, 0xc00062ec30, 0xc00062e900, 0x3989ecb, 0xe, 0x39bf3a3, 0x1f, 0x3ebd9a0, 0x14, ...)
// 	/usr/local/google/home/aaronabbott/repo/opentelemetry-collector-contrib/cmd/otelcontribcol/main.go:54 +0x11e
// main.run(...)
// 	/usr/local/google/home/aaronabbott/repo/opentelemetry-collector-contrib/cmd/otelcontribcol/main_others.go:22
// main.main()
// 	/usr/local/google/home/aaronabbott/repo/opentelemetry-collector-contrib/cmd/otelcontribcol/main.go:43 +0x238
