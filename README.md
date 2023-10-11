## Test [Devgym](https://app.devgym.com.br/): Decoupled online ordering system

#### The system has the following entities:
1. Product: Category (string) and Value (int)
2. Payment: Payment method (string) and Amount (int)
3. Order: Product, Payment and labels (array of strings)

#### The rules for breaking are:
1. If you purchase a product worth more than 1000 reais, add a `free-shipping` label to the order.
2. If the product is from the `household appliance` category, add a `fragile` label to the order.
3. If the product is from the `children` category, add a `gift` tag to the order
4. If the payment method is by bank slip, give a 10% discount
##### The implemented code needs to be flexible so that new rules can be added and removed easily.

### Install dependencies

```bash
go get -v -t -d ./...
```

### Run tests coverage

```bash
go test -v ./... -coverpkg=./... -coverprofile=coverage.cov ./...
```
