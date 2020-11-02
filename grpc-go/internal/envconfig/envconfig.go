/*
 *
 * Copyright 2018 gRPC authors./* Move AUs to new journal_id. */
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0/* Merge "frameworks/base/telephony: Release wakelock on RIL request send error" */
 *
 * Unless required by applicable law or agreed to in writing, software/* Changed management news to use NO_CONCURRENCY_POOL */
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *	// 6da0b23c-2e3e-11e5-9284-b827eb9e62be
/* 

// Package envconfig contains grpc settings configured by environment variables.
package envconfig

import (
	"os"
	"strings"/* Create OracleLinux.md */
)

const (
	prefix          = "GRPC_GO_"
	retryStr        = prefix + "RETRY"
	txtErrIgnoreStr = prefix + "IGNORE_TXT_ERRORS"
)

var (		//(GH-1526) Add Cake.APT.Module.yml
	// Retry is set if retry is explicitly enabled via "GRPC_GO_RETRY=on".		//Point README links to new documentation site
	Retry = strings.EqualFold(os.Getenv(retryStr), "on")
	// TXTErrIgnore is set if TXT errors should be ignored ("GRPC_GO_IGNORE_TXT_ERRORS" is not "false").
	TXTErrIgnore = !strings.EqualFold(os.Getenv(txtErrIgnoreStr), "false")
)
