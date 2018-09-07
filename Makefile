# Copyright © 2018 Krishna Iyer Easwaran.  All Rights Reserved.

# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
# 	http:#www.apache.org/licenses/LICENSE-2.0

# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

BUILD_DIR = build
PKGS := $(shell go list ./... | grep -v /vendor)
#GOBUILD = CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o "$@"
GOBUILD = go build -o "$@"
SRC_FILES = *.go
SRC := gateway

init: 
	@mkdir -p $(BUILD_DIR)
	dep init -v 

build: clean $(SRC:%=$(BUILD_DIR)/%)

$(BUILD_DIR)/%:
	$(GOBUILD) pkg/$(@:$(BUILD_DIR)/%=%)/$(SRC_FILES)

deps:
	dep ensure -v --update

test:
	go test -v $(PKGS)

clean:
	@rm -r $(BUILD_DIR)
	@mkdir -p $(BUILD_DIR)