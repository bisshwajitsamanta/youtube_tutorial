
output "FirstName" {
  value = "First Name is ${var.printdetails[0]}"
}

output "LastName" {
  value = "Last Name is ${var.printdetails[1]}"
}
output "Full_Name" {
  value = join(" ",var.printdetails)
}

output "Timestamp" {
  value = "Current Time is ${timestamp()}"
}

output "FirstUserRegion" {
  value = "${var.UserDetails[0].userroles}"
}
output "SecondUserDetails" {
  value = element(var.UserDetails,1).userroles
}
output "onevalue" {
  value = one(var.onedetails)
}
output "anydata" {
  value = join("-",var.anydata)
}