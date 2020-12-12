// Copyright © 2019 Ryan Ciehanski <ryan@ciehanski.com>
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

package libgen

import "net/url"

// SearchMirrors contains all valid and tested mirrors used for
// querying against Library Genesis.
var SearchMirrors = []url.URL{
	{
		Scheme: "http",
		Host:   "gen.lib.rus.ec",
	},
	{
		Scheme: "https",
		Host:   "libgen.rs",
	},
	{
		Scheme: "https",
		Host:   "libgen.unblockit.red",
	},
	{
		Scheme: "http",
		Host:   "libgen.unblockall.org",
	},
	{
		Scheme: "https",
		Host:   "93.174.95.27",
	},
	
}
