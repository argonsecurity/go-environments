# Environments

This util package provides a generic way to integrate with CI environments.
It collects data from environment variables, git commands, and local files to create a data object that contains relevant data of the current CI run and the repository.

The data object collects data both on the CI platform and the SCM repository.

---

## Supported Environments

| SCM               | CI Platform         |
| ----------------- | ------------------- |
| GitHub            | GitHub Workflows    |
| GitLab            | GitLab CI           |
| Azure Devops      | Azure Pipelines     |
| Bitbucket         | Bitbucket Pipelines |
| GitHub            | Jenkins             |
| GitLab            | Jenkins             |
| Bitbucket         | Jenkins             |
| GitHub Enterprise | GitHub Workflows    |
| GitLab Server     | GitLab CI           |
| GitLab Server     | Jenkins             |
| Bitbucket Server  | Jenkins             |

---

## Testing

To test the package, run `make test` in the root directory of the package.

---

## Usage

```go
package main

import (
	"fmt"

	"github.com/argonsecurity/go-utils/environments"
)

func main() {
	env := environments.DetectEnvironment()
	configuration, err := env.GetConfiguration()
	if err != nil {
		panic(err)
	}

	fmt.Println(configuration.Repository.Id)
	fmt.Println(configuration.SCMApiUrl)
	...
}
```
