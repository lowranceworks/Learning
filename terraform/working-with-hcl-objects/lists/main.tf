locals {
  fruits = [
    {
      fruit = "apple"
      color = "red"
      count = 1
    },
    {
      fruit = "tangerine"
      color = "orange"
      count = 2
    },
    {
      fruit = "pear"
      color = "green"
      count = 3
    }
  ]
}

output "list_of_fruit_colors" {
  value       = [for f in local.fruits : f.color]
  sensitive   = false
  description = "description"
  depends_on  = []
}
