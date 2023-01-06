variable "printdetails" {
  type = list(string)
  default = ["Bisshwajit","Samanta"]
}

variable "onedetails" {
  type = list(string)
  default = ["code with Bisshwajit"]
}

variable "UserDetails" {
  type = list(object({
    region = string
    iamuser = string
    renewaldays= number
    userroles = list(string)
  }))
  default = [
    {
      region = "us-east-1"
      iamuser = "Bisshwajit"
      renewaldays = 30
      userroles =["s3fullaccess","admin"]
    },
    {
      region = "us-west-2"
      iamuser = "Samanta"
      renewaldays = 30
      userroles =["athenareadonlyaccess","admin"]
    }
  ]
}

variable "anydata" {
  type = list(any)
  default = [06,01,2023]
}