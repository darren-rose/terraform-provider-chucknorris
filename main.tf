data "chucknorris" "example" {

}

output "example_joke" {
    value = data.chucknorris.example.joke
}