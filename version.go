// Copyright 2014 Oleg Dolya. All rights reserved.
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.
package privat24go

// The git commit that was compiled. This will be filled in by the compiler.
var GitCommit string

// The main version number that is being run at the moment.
const Version = "0.2.0"

// A pre-release marker for the version. If this is "" (empty string)
// then it means that it is a final release. Otherwise, this is a pre-release
// such as "dev" (in development), "beta", "rc1", etc.
const VersionPrerelease = ""
