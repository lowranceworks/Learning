locals {
  fruits = {
    "apple" = {
        color = "red"
        count = 1
    }

    "tangerine" = {
        color = "tangerine"
        count = 2
    }

    "pear" = {
        color = "green"
        count = 3
    }
  }
}

output list_of_maps {
  value       = [for e in local.fruits : e]
  sensitive   = false
  description = "description"
  depends_on  = []
}

output list_of_fruit_colors {
  value       = [for k, v in local.fruits : v.color]
  sensitive   = false
  description = "description"
  depends_on  = []
}
