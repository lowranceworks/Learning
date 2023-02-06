locals {
  // map(map(string))
  fruit_map = {
    "apple" = {
        color = "red"
        count = 1
    }

    "orange" = {
        color = "orange"
        count = 1
    }

    "pear" = {
        color = "green"
        count = 1
    }
  }

  list_of_maps = [for map in local.fruit_map: map]

  list_of_map_key_attributes = [for k, v in local.fruit_map: k]

  list_of_map_value_attributes = [for k, v in local.fruit_map: v]

  list_of_each_attribute_inside_each_map_value_attribute = [for k, v in local.fruit_map: v.color]

}

output list_of_maps {
  value       = local.list_of_maps
  sensitive   = false
  description = "Returns a list of maps without the key/value."
  depends_on  = []
}

output list_of_map_key_attributes {
  value       = local.list_of_map_key_attributes
  sensitive   = false
  description = "Returns a list of strings taken from the key attribute of each map."
  depends_on  = []
}

output list_of_map_value_attributes {
  value       = local.list_of_map_value_attributes
  sensitive   = false
  description = "Returns the same output as local.list_of_maps."
  depends_on  = []
}

output list_of_each_attribute_inside_each_map_value_attribute {
  value       = local.list_of_each_attribute_inside_each_map_value_attribute
  sensitive   = false
  description = "Returns a list of strings taken from the value attributed of each map nested inside the fruit_map."
  depends_on  = []
}