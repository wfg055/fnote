// Copyright 2023 chenmingyong0423

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//     http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"os"

	"github.com/chenmingyong0423/fnote/backend/ioc"
	"github.com/pkg/errors"
)

func main() {
	if len(os.Args) < 3 {
		panic(errors.New("missing parameters"))
	}
	username := os.Args[1]
	password := os.Args[2]
	app, err := initializeApp(ioc.Username(username), ioc.Password(password))
	if err != nil {
		panic(err)
	}
	err = app.Run()
	if err != nil {
		panic(err)
	}
}
