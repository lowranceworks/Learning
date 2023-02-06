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

  //
  map_of_maps_to_list_of_maps = [for map in local.fruit_map: map]

  //
  map_keys_to_list_of_keys = [for k, v in local.fruit_map: k]

  //
  map_values_to_list_of_values = [for k, v in local.fruit_map: v]

  //
  map_value_attributes_to_list_of_attributes = [for k, v in local.fruit_map: v.color]

}

output ist_of_maps {
  value       = local.map_of_maps_to_list_of_maps
  sensitive   = false
  description = "a list of maps."
  depends_on  = []
}

output list_of_map_keys {
  value       = local.map_keys_to_list_of_keys
  sensitive   = false
  description = "a list of keys."
  depends_on  = []
}

output list_of_map_values {
  value       = local.map_values_to_list_of_values
  sensitive   = false
  description = "a list of values."
  depends_on  = []
}

output list_of_map_values_attributes {
  value       = local.map_value_attributes_to_list_of_attributes
  sensitive   = false
  description = "a list of values."
  depends_on  = []
}