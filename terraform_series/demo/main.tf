// Dynamic Expression
// conditional value -- if the value is true then do something or else block

resource "random_id" "thisRandom" {
  byte_length = 8
}
output "printOut" {
  value = (var.output_Value ==true ? "Hello Youtubers !!": "${var.RandomOutput}..${random_id.thisRandom.hex}")
}