func count64(n : UInt64) -> UInt64 {
    var i = n
    i = i - ((i >> 1) & 0x5555555555555555)
    i = (i & 0x3333333333333333) + ((i >> 2) & 0x3333333333333333)
    i = ((i &+ (i >> 4)) & 0x0F0F0F0F0F0F0F0F)
    return (i&*(0x0101010101010101))>>56
}

func count32(n : UInt32) -> UInt32 {
    var x = n
    x = x - ((x >> 1) & 0x55555555);
    x = (x & 0x33333333) + ((x >> 2) & 0x33333333);
    x = (x + (x >> 4)) & 0x0F0F0F0F;
    x = x + (x >> 8);
    x = x + (x >> 16);
    return x & 0x0000003F;
}

var arr = readLine()!.characters.split(" ").map{Int(String($0))!}
var n = arr[0]
var m = arr[1]
var bits = Array(count:n, repeatedValue:Array(count:8,repeatedValue:UInt64(0)))
var counts = Array(count:n, repeatedValue:UInt64(0))
for i in 0..<n {
    let b = (readLine()!).characters.map{String($0)}
    for j in 0..<m {
        if b[j] == "1" {
            bits[i][j/64] = bits[i][j/64] | (1 << (UInt64(j) % 64))
        }
    }
}

var maxC = UInt64(0)
var t = 0
var c = UInt64(0)
for i in 0..<n-1 {
    for j in i+1..<n {
        c = 0
        for k in 0..<8 {
            c += count64(bits[i][k] | bits[j][k])
        }
        if c > maxC {
            maxC = c
            t = 1
        } else if c == maxC {
            t += 1
        }
    }
}
print(maxC)
print(t)
