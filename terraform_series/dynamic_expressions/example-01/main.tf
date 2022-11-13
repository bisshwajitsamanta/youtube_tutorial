
resource "random_id" "some_value" {
  byte_length = 8
}

output "hello" {
  value = (var.count_value == true ? "Hello Youtubers!!": "${var.Joined}..${random_id.some_value.hex}" )
}