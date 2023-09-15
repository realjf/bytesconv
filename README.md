# bytesconv

Byte unit conversion library

## v2

### Features

- Support for conversions between B/KB/MB/GB/TB/PB/EB/ZB/YB/RB/QB
- Support for conversions between B/KiB/MiB/GiB/Tib/PiB/EiB/ZiB/YiB/RiB/QiB

### Examples

```go
func main() {
 fmt.Println(bytesconv.ByteCountSI(big.NewInt(1), 2))
 fmt.Println(bytesconv.ByteCountSI(big.NewInt(10), 2))
 fmt.Println(bytesconv.ByteCountSI(big.NewInt(1000), 2))
 fmt.Println(bytesconv.ByteCountSI(big.NewInt(10000), 2))
 fmt.Println(bytesconv.ByteCountSI(big.NewInt(1000000), 2))
 fmt.Println(bytesconv.ByteCountSI(big.NewInt(1000000000), 2))
 fmt.Println(bytesconv.ByteCountIEC(big.NewInt(10), 2))
 fmt.Println(bytesconv.ByteCountIEC(big.NewInt(1024), 2))
 fmt.Println(bytesconv.ByteCountIEC(big.NewInt(9104342), 3))
 fmt.Println(bytesconv.ByteCountIEC(big.NewInt(143421343), 3))
 fmt.Println(bytesconv.ToKiB("1B", 6))
 fmt.Println(bytesconv.ToKiB("10B", 4))
 fmt.Println(bytesconv.ToKiB("10K", 1))
 fmt.Println(bytesconv.ToKiB("10M", 1))
 fmt.Println(bytesconv.ToKiB("10G", 1))
 fmt.Println(bytesconv.ToKiB("10T", 1))
 fmt.Println(bytesconv.ToKiB("10P", 1))
 fmt.Println(bytesconv.ToKB("1B", 6))
 fmt.Println(bytesconv.ToKB("10B", 4))
 fmt.Println(bytesconv.ToKB("10K", 1))
 fmt.Println(bytesconv.ToKB("10M", 1))
 fmt.Println(bytesconv.ToKB("10G", 1))
 fmt.Println(bytesconv.ToKB("10T", 1))
 fmt.Println(bytesconv.ToKB("10P", 1))
}
```

output

```bash
1.00 B
10.00 B
1.00 kB
10.00 kB
1.00 MB
1.00 GB
10.00 B
1.00 KiB
8.683 MiB
136.777 MiB
0.000977 KiB
0.0098 KiB
10.0 KiB
10240.0 KiB
10485760.0 KiB
10737418240.0 KiB
10995116277760.0 KiB
0.001000 kB
0.0100 kB
10.2 kB
10485.8 kB
10737418.2 kB
10995116277.8 kB
11258999068426.2 kB
```

## v1

### Features

- Support for conversions between B/KB/MB/GB/TB/PB/EB
- Support for conversions between B/KiB/MiB/GiB/Tib/PiB/EiB

### Examples

```go
func main() {
 fmt.Println(bytesconv.ByteCountSI(1, 2))
 fmt.Println(bytesconv.ByteCountSI(10, 2))
 fmt.Println(bytesconv.ByteCountSI(1000, 2))
 fmt.Println(bytesconv.ByteCountSI(10000, 2))
 fmt.Println(bytesconv.ByteCountSI(100000, 2))
 fmt.Println(bytesconv.ByteCountSI(1000000, 2))
 fmt.Println(bytesconv.ByteCountIEC(98, 2))
 fmt.Println(bytesconv.ByteCountIEC(1024, 2))
 fmt.Println(bytesconv.ByteCountIEC(98765, 3))
 fmt.Println(bytesconv.ByteCountIEC(987654321, 3))
 fmt.Println(bytesconv.ToKiB("1B", 6))
 fmt.Println(bytesconv.ToKiB("10B", 4))
 fmt.Println(bytesconv.ToKiB("10K", 1))
 fmt.Println(bytesconv.ToKiB("10M", 1))
 fmt.Println(bytesconv.ToKiB("10G", 1))
 fmt.Println(bytesconv.ToKiB("10T", 1))
 fmt.Println(bytesconv.ToKiB("10P", 1))
 fmt.Println(bytesconv.ToKB("1B", 6))
 fmt.Println(bytesconv.ToKB("10B", 4))
 fmt.Println(bytesconv.ToKB("10K", 1))
 fmt.Println(bytesconv.ToKB("10M", 1))
 fmt.Println(bytesconv.ToKB("10G", 1))
 fmt.Println(bytesconv.ToKB("10T", 1))
 fmt.Println(bytesconv.ToKB("10P", 1))
}
```

output

```bash
1.00 B
10.00 B
1.00 kB
10.00 kB
100.00 kB
1.00 MB
98.00 B
1.00 KiB
96.450 KiB
941.901 MiB
0.000977 KiB
0.0098 KiB
10.0 KiB
10240.0 KiB
10485760.0 KiB
10737418240.0 KiB
10995116277760.0 KiB
0.001000 kB
0.0100 kB
10.2 kB
10485.8 kB
10737418.2 kB
10995116277.8 kB
11258999068426.2 kB
```
