// Copyright 2023 Harald Albrecht.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package once

import (
	"sync"
)

// Oncer executes the registered function only once. It is designed along the
// idea of DRY, in contrast to [sync.Once].
type Oncer struct {
	oh sync.Once
	fn func()
}

// Once returns a oncer that will execute the passed function at most once when
// calling its Do method.
func Once(fn func()) *Oncer {
	return &Oncer{fn: fn}
}

// Do the function that was registered when creating this oncer exactly once, no
// more, no less. One times shall be the number thou shalt count, and the number
// of counting shall be one. Two shalt thou not count, neither count thou zero,
// except that you proceed to one. Three is right out.
func (o *Oncer) Do() {
	o.oh.Do(o.fn)
}
