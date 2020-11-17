## Maps

All the keys should be of the dame type and all the values should be of the same type. However the type of key and value shouldn't be the same

### Definition and assignment

map[KeyType]ValueType

Keynames are case sensitive

_Method 1_

```
color := map[string]string{
    "red": "#FF0000",
}
```

_Method 2_

```
var colors map[string]string
```

Map types are reference types, like pointers or slices, and so the value of m above is nil; it doesn't point to an initialized map. A nil map behaves like an empty map when reading, but attempts to write to a nil map will cause a runtime panic; don't do that. To initialize a map, use the built in make function

```
colors := make(map[string]string)
```

The above statemt can also be used without the declaration statement

_Assignment_

```
colors["white"] = "#FFFFFF"
```

The key-value pairs should have a trailing comma same as in the structs

_Deletion_

delete(mapVariable, keyName)

```
delete(colors, "white")
```

### Iteration