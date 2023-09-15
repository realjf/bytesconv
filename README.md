# bytesconv

Byte unit conversion library

## v2

### Features

- Support for conversions between B/KB/MB/GB/TB/PB/EB/ZB/YB/BB/NB/DB
- Support for conversions between B/KiB/MiB/GiB/Tib/PiB/EiB/ZiB/YiB/BiB/BiB/NiB/DiB

### Examples

```go
func main() {
 fmt.Println(bytesconv.ByteCountIEC(big.NewInt(1000), 2))
 fmt.Println(bytesconv.ByteCountIEC(big.NewInt(987654321), 3))
 fmt.Println(bytesconv.ToKiB("10M", 1))
 fmt.Println(bytesconv.ToKiB("1b", 6))
}
```

output

```bash
1000 B
941.901 MiB
10.0 KiB
0.000001 KiB
```

## v1

### Features

- Support for conversions between B/KB/MB/GB/TB/PB/EB
- Support for conversions between B/KiB/MiB/GiB/Tib/PiB/EiB

### Examples

```go
func main() {
 fmt.Println(bytesconv.ByteCountIEC(1000, 2))
 fmt.Println(bytesconv.ByteCountIEC(987654321, 3))
 fmt.Println(bytesconv.ToKiB("10M", 1))
 fmt.Println(bytesconv.ToKiB("1b", 6))
}
```

output

```bash
1000 B
941.901 MiB
10.0 KiB
0.000001 KiB
```
