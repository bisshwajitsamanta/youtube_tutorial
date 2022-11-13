terraform {
  required_version = "~>1.2.3" // Allows 1.2.4 or 1.2.5 Deny - 1.3.3 or 1.3.5
}
output "first_program" {
  value = "Hello World !"
}