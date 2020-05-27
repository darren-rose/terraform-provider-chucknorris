# chucknorris Data Source

#### Build

```
go build -o terraform-provider-chucknorris
```

#### Example Usage

```
data "chucknorris" "example" {

}

output "joke" {
    value = data.chucknorris.example.body
}
```