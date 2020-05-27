data "chucknorris" "example" {

}

output "joke" {
    value = data.chucknorris.example.body
}