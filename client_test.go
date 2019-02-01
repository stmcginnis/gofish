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

package gofish

import "testing"

// TestClientGetRoot tests getting the root data from the server.
func TestClientGetRoot(t *testing.T) {
	client := &Client{
		Host: "localhost",
		Port: 5000}

	client.GetRoot()
}
