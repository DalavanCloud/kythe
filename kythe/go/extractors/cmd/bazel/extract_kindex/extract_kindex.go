/*
 * Copyright 2017 Google Inc. All rights reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

// Program extract_kindex implements a Bazel extra action that captures a Kythe
// compilation record for a "spawn" action.
package main

import (
	"context"
	"flag"
	"log"
	"time"

	"kythe.io/kythe/go/extractors/bazel"
)

var (
	outputPath = flag.String("output", "", "Path of output index file (required)")

	settings bazel.Settings
)

func init() {
	flag.Usage = settings.SetFlags(nil, "")
}

func main() {
	flag.Parse()

	// Verify that required flags are set.
	if *outputPath == "" {
		log.Fatal("You must provide a non-empty --output file path")
	}

	config, info, err := bazel.NewFromSettings(settings)
	if err != nil {
		log.Fatalf("Invalid config settings: %v", err)
	}

	start := time.Now()
	ai, err := bazel.SpawnAction(info)
	if err != nil {
		log.Fatalf("Invalid extra action: %v", err)
	}
	cu, err := config.Extract(context.Background(), ai)
	if err != nil {
		log.Fatalf("Extraction failed: %v", err)
	}
	log.Printf("Finished extracting [%v elapsed]", time.Since(start))
	if err := bazel.Write(cu, *outputPath); err != nil {
		log.Fatalf("Writing index: %v", err)
	}
}
