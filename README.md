# chucknorris Provider

The chucknorris provider is a utility provider for providing random chuck norris jokes.

This provider requires no configuration.

#### chucknorris Data Source

The chucknorris data source exports random chucknorris jokes.

#### Example Usage

```
data "chucknorris" "example" {

}

output "joke" {
    value = data.chucknorris.example.body
}
```

#### Build

```
go build -o terraform-provider-chucknorris
```
