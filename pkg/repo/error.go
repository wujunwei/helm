/*
Copyright The Helm Authors.

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

package repo

import (
	"fmt"
)

type ChartNotFoundError struct {
	RepoURL string
	Chart   string
}

func (e ChartNotFoundError) Error() string {
	return fmt.Sprintf("%s not found in %s repository", e.Chart, e.RepoURL)
}

func (e ChartNotFoundError) Is(err error) bool {
	_, ok := err.(ChartNotFoundError)
	return ok
}
