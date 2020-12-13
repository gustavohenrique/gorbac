## GoRBAC

> Role Base Access Control for Go

[![Coverage Status](https://coveralls.io/repos/github/gustavohenrique/gorbac/badge.svg?branch=main)](https://coveralls.io/github/gustavohenrique/gorbac?branch=main)

## Usage

Create a file containing roles.

```sh
# A role to allow read all orders and update only orders paid and with amount less than 1000
cat > /tmp/roles.yaml <<EOF
- name: Attendant
  permissions:
  - action: read
    target: order
  - action: update
    target: order
    rules:
    - allowed: true
      key: paid
      operator: eq
      value: true
    - allowed: true
      key: amount_paid
      operator: lte
      value: 1000
EOF
```

Verify access.

```go
import (
    rbac "github.com/gustavohenrique/gorbac"
)

type Order struct {
		ProductName string `json:"product_name"`
		Amount      int    `json:"amount_paid"`
		Paid        bool   `json:"paid"`
}

func main() {
    roles := rbac.FromFile("/tmp/roles.yaml")
	  canRead := rbac.With(roles).HasPermission(rbac.READ, "order")
    if !canRead {
        log.Fatalf("You do not have permission to read orders")
    }
	  canUpdate := rbac.With(roles).HasPermission(rbac.UPDATE, "order")
    if !canUpdate {
        log.Fatalf("You do not have permission to update orders")
    }
    order := Order{ProductName: "Notebook", Amount: 1000, Paid: true}
    isAllowedToChange := rbac.With(roles).HasPermission(rbac.UPDATE, "order").IsAllowed(order)
    if !isAllowedToChange {
        log.Fatalf("You are not allowed to change orders with amount greater than 1000 or not paid")
    }
}
```

## License

Copyright 2020 Gustavo Henrique

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
