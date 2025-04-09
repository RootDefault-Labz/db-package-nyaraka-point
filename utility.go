package nyarakadb

import (
    "reflect"
    "strings"
)

// StructToMap converts a struct to a map[string]interface{} using reflection.
// It uses json tags if available, otherwise uses the field name.
// Zero values are omitted from the resulting map.
func StructToMap(obj interface{}) map[string]interface{} {
    result := make(map[string]interface{})
    
    // Get reflected value and type of the struct
    value := reflect.ValueOf(obj)
    if value.Kind() == reflect.Ptr {
        value = value.Elem()
    }
    
    if value.Kind() != reflect.Struct {
        return result
    }
    
    typ := value.Type()
    for i := 0; i < value.NumField(); i++ {
        field := value.Field(i)
        
        // Skip zero values
        if reflect.DeepEqual(field.Interface(), reflect.Zero(field.Type()).Interface()) {
            continue
        }
        
        // Get the field tag or name
        fieldType := typ.Field(i)
        tag := fieldType.Tag.Get("json")
        name := fieldType.Name
        
        if tag != "" {
            // Split the tag to handle cases like `json:"name,omitempty"`
            name = strings.Split(tag, ",")[0]
            if name == "-" {
                continue // Skip fields with json:"-"
            }
        }
        
        result[name] = field.Interface()
    }
    
    return result
}