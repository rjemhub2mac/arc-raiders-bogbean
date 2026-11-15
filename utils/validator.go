package main

import (
    "fmt"
    "sync"
    "time"
    "crypto/sha256"
)

var ( appVersion = "5.9" )

func uYegRDGJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gTXGg9dnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SkGRMLosWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 18
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pinMDUExWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3IgV5LklWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IlcIDn7IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func e8iLBLbNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 36
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qTssxz4yWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uwdj1L6IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1P95h9zjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PldrDGe0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func B8flbvCAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LZSKB2cnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func orcC52u2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3dlo4mK1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PHZUUDd8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LOVenCLYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PBLrnuuwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func P4sdOrPxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vDbLWQiyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2GcJ2waNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func R2azGDYvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xdA3QJfBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pQdTytS5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NxMLIv5gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 80
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NUiKgImjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pQZDUOa8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 16
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Lh17lasyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func m5kNdTTmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 76
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9O5UYZ1GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jmggaoaVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 23
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CzcNDxHYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XCPIDflLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zD4NFpijWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func leXgDZeIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BWk2sW8JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dxAuJ2syWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Sv5mUcAZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RT4D0uHbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 53
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func g0OSDc8DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 155
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8Q1v5PntWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4FqclUPoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ND8W9BUoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1d72bvFMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dMfz0MhrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 255
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YatuMh5CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HKfWjT73Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QC0AHdsKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func k41B0REaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VyqpBsO6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hC8StjvqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func I11vPRK0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func j3c1McmVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Sf7XQqzgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kARob8isWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 152
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DH8w3Ww0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 197
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func r2kUJmzAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uKGpxKZAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8MjiuIFgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Aqxl43LuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 32NYmaEcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 103
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4fkt4U5CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9UOafo70Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RUgAr0MjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jbK5eTTFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zyDqD1u8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gRKoKwLMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NqqfrWeuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 218
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func w0i3diqOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 104
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func y3uDbyRDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zBgdSveSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 179
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4DZAeeDSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bLZ9FSVaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PZ4GS0BQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kY1ExF0eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 283
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PVvtHhUGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hVXpER6xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oQCPrbaUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GdoxNJnxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7cKj6tFzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sWbTCvkeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 53jvWsloWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func A5syzxZiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IIFYFDU0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YZY5xHPgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6ToCl57rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SahqKy5oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IPq062A5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 88
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FdAHiIiTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 199
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9RlOnQ6ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LuobvClLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Na1FtSWIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6rMfvEErWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func E25LFD04Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qbe2WO9hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 199
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ja26pumXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func v20Oix2OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cX9o97lEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HBOJcHyuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2oDHfCSDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func h3R580gFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 221
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func frkVePiCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 197
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LK25RVgzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DLjhwBWyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 88
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wLKb0qoDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cEAFlN7HWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KdUOvPm9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Iz8ToVsQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8PGQHKDfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 217
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9cEwhJpfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MSO4tY2kWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 110
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mWrB85dLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yxmzURfgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 39
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oYlQbEVWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZopQnK1lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 244
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func H4FFLuP8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1cJ8AxMNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 86
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yyHJiWoyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 107
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func W78LEYHfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 155
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ksPv4DR8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func A2iNDyUhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 107
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func L7AMzJWkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ybUKJIoXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 104
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gnEORKgSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nLoJpcFkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uGHvPLJuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 62
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3JnvQ3zSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AxuzZgxEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 229
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iNeulFlDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func j11sk0ieWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WnhzxigoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 152
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1EgCpMmFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Hm2cfffzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 244
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kEmXZXprWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FbQM45WdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pge9G8L6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BpgRaoNqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aTDXOXxOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YhVjMtzxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pIMbgQEzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nBficzIYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 41
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kkcJaUAaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7ngW3LP7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func T2HP11LdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ltgkSL3hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func N3SYGISXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dmqtp18uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cuuh3HAnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 89
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sx95huh4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qbQugxr9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YzlONaPKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KSaEPlaUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lJkEJePQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ABMovCgZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tSewfDUnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DgrvG1slWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 14
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cKcCXQCtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ihI2o4REWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QmnPIDJhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Rvh7FyimWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gRhxpDcDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OthjDBUEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Zc8J8XAMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pZFB6Z0iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 175
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OTMz23u6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func leBh4iDKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 175
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5md28rSFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GF9Psyf8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3TOluFvwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ref9W7LGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VzZjIv1LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5xfYNpItWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Glf3y9YrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 200
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func S30vig9QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 53
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RLVrm59BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 195
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qNak7ou0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 91
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6dOhlQwHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func key7tzVBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dExWMdUjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func q8KuHhCUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 197
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3bSv4z1qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6rdolqoXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oq95JBhgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 196
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func deENOg6JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func O4Z4YZ99Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JAmsff2SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YPAfJJGzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZqI9M83MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func k53VYTibWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DQ1kbxzoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aO3knnpZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func llRILKZRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eFgdWksKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EH2lkeUNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oOw25ejpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 179
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Fdd28jnwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MVLFk4H9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZeqVh5rqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 297
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Um6sf7c8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ClQyioVZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AcHbZqvZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RzsVqSlNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BU9S3SJBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AELGvFBiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cM9nhYV8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ppyRojccWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tR5IATlIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IhSa8XvEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SQGn0kYyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 69
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func A4Aoe7YwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Fxy1qL1tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JIUcOK1YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 35ovE7qTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 297
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Hm7IYVzKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nBV5aFohWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8Iy7u8tKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func l7GAV0IMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Zxm37fWGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func y94E1g7jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zDprs6uUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Nkt68bkqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IjMNrizgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BtOHWMGnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func teW5NT9ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Gxneykj5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RsPkZcPmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QCTJqwpAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vFF7nwC7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 36
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OC72UB9IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 161
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WSiUuBbwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 163
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mSGZ0GiJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 199
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pvxCOsuxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Igu8Z1KZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func i1zLuL6iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pxHr0Ml7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sCguLnc5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func l7VyOTRyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DACGIWjeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 291
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xEu5KwyZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 283
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func d41278Y8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vXSd8XzYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DWyNY7Y1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 141
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UoRUkyikWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HFKL4DTbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sJNqqHzFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 115
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 78iu9K3VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bWveqegdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func u9NzqETEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func to9doOtAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func y5rRA01PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func H8EkTSDlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 62
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func m9xL6uO0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 39
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Mh21jMpoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vx3tTWRbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iSHQE70IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5EeRYURZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BFXeWu2NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lXN15J5VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mFPnov7dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 88
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xgGcXdXdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ChiDRai0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GCKAYApdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 81NVR7arWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kJceX53OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 62
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EDzVjXmAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KL90PTynWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5hAZvq67Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 15
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WrQEXLVHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6pS91nruWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 209
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func luE6HTE2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 62
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MQhiDA6GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func d158gPEvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LL2vVS4IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3INz30BiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wLp8mQcuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ji6g6x3aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 236
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HJ1aDNgnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cpRGkI8SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 246
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1mwKENSwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 140
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func te2i6JW2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 107
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1G8l2HTCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 287
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wTP9u6YMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 13CIx4YjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iSQXSuFxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 103
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gTDJHRKnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iy5NT3z8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 208
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OkZhejkLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oKIQx6oWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IkS5Dx7rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SSglGKCgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FZ3DlkdiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jD0ZFumkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0gaAkBm1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 39
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qsWVDIPeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 260
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Qh0Yq0ydWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qQzIsxUMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pMkMgKkGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 112
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NCuf2tLaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Wpme2InZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 110
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LrP7FjJLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Dovxelg1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HxP8JHEFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 221
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func j4IxfKXPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SYIqslYWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tKFWMA3WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cOkHvYi8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1sKonrLnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6KsQNMFcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BxWSVvgAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 16
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func id7d2XmKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 226
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func h7g4muUqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 36d8JohfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XLUxvhwYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 208
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func m95bll7iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LDLeICCSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0nfumA8YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 214
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dn7upeBCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6h1AS3QtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fPjVEAjvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Rxpc6IWaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func W2A4lH3DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ynDY8leiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HRnoNOtPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Va07X6nYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func o9eX6gGmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 179
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UdpNqTpAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func i5hPxoyxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func II5vHYWPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cYH1quacWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VClTJ3ljWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FfIpRprtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zOX2TTgnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4MPswMKKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rnuLhI5bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WOeczGAzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7nh6USfSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ivKU9dmFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func L8VOgfiCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kYkILtOyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Lf1aEhNlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 163
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Mai9xqbdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Mjl1YFhLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func n4NYS0rBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5lUR50zyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8PIu71CFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bEf6aCjIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZjdmVu0rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 37
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Jo5zT7p3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2reabwZnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ylbZ7SfYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 54
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mIC1omPxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 136
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func G0wRbhI9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ja25Hh4JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZVMcRazLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qJYxZujtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3OfsN8nTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 59
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zMtXBzEzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GEJBTThCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MOiHtmy1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NYHB35fuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 41
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hfJkV5JZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func V8vGTAFtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func i4QQjIl8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 104
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6IHMeN5sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 274
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nxDsdbAmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zpRuyvJSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZoRFKvk7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sG31Rf2sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LIp4BEQCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WlMAXVBDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7XhHY62JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gc2KCA6FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lGMNdXDvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vsSn5PCdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IzmwFdccWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GGLCkHJpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ck92ndPvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OEPzcFAvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 175
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ackLmVfaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 253
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MT0ixIJhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 38
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bK2SWLPhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NZV4pAmQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 86
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fSbkIUXqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 62
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nr1hYajCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 59
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func N6WpYspuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QigbxxioWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bEL4RccBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dK6BOO7NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 80
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aVaTVUgiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 162
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func h1H4DFgZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func flUOw2BNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oCeDEIKIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 39
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UZDCM7QqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ERJFcoflWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jGHtVV8OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0ceVqd1jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tyVSSeG0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FfxefQO9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 203
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WCPVGW4UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YYXtLPwzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7CgdE7CJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func om35jRzOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 141
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zGrlyk9UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 217
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zh9W4EbXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dIubM3yIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VbBNO0jJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8nusLtN4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NusDK5GEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func M3Ath3wEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3kjLUpwJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func onNSVI4XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cxur4QiJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3JWxs5smWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Eaz3nScSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 195
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SkhyHZKBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WMs0DLPtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func P0HcwIIYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dL4zjdnQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PScriyRDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 20
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Dh2KBQMuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 141
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func utsiEM0xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5dAkYDSUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 90
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tFmhB5RlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func e4DIwySQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XdunFZIhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Cn6RdQtPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 53
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8pBqjCZtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WTPku7GBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xcygnX0HWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WwUUFLv1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3tx0AXn8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 226
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nzkbzJhAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5UA2p2LhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 152
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pgB1zbe4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QYWkMSwzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1VSNdSZeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2XhurgpFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SHlZk83iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func G84PMKQdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JdTLfgfbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func j8y1536TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hecibmuoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 38
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eVfCJ9cQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9hYMOhHMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dMPjf4NDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YUTXeSsFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4V5EvjZ9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func r2GwnFXuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 155
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func q7oCxi3bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Kt1SpBJ3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 79
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bVLQcwU3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func e8CQs3gNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VXxkaZ1eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ouVFCDmRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func f6LZrsL4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 255
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OfooEJ1AWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FSSJS9i3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6xiIYfSFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CPZSe1reWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XxfdTsv0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VmCVvKbKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2SraIhngWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func s8h86D4rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 25
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kod1nRFTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0po89oqoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 200
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yuFIsaUKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func P1tflpJXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BhgM1ZoRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qC22AOfKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 685buMF7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 104
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func i0TWtoNuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TplfwKyLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eqkyHSVTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZSVQChSkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RXK5SDfRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 287
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MvyDQOJBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QerPnxJmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 86
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hKBIHnwpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rBi5mEB5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 80
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DGORNxt8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rCkpPvDlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 253
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func l3JIzRCbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 200
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cu8TyRKKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eI5QLpmnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1tD6wvFPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Y9DGLSvoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func scRwmk19Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VoCUeWQ4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func H6EKzwFoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bwBKFPZPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Fv3892TZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RKCCp0eFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 260
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kgQ7KG7GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fcZNQThUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 38
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kyuEwy3EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Vlau4dLGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func i2DxGH1NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 59
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rYRmD4JeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0BrrmjH3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aPvKQi3yWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 16
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XCvGuP6gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 208
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3bNhvFfrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SVamMTT7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3MJ4bxMkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 112
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QyJLGVgxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PvKEkZ09Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LH4DBwVeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 214
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9kTQzVtTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mhvHNxI9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0eVw7mCcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UaxsfKYrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func T4M9aAzkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 38Du48LiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XKAn0hcYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 287
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TSx3NjdcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XPGtp6oRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func W1gA61bkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 217
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bmwBM4YkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 16
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func m2pu2RP6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func l74L52B9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bkdloH7rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func y4nAbSIbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SUcTMhfWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 152
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xV6eNHcsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ebeNIwhDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ydlflYdbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 218
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iRaBiPtrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8uVQc8HoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 283
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KJSmSxsIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LIiVFoRfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iGlJOrTCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0em1vfMIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func m7oaHF6YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vG6qDG7xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KX6DEqldWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rQYzOErMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VQJmb4UOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FMOkjxAVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ykxmAsrlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pJRmi8KwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ztJlxzPRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 141
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yKrtBJcrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ybg0c97xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 246
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iY1SVfjKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 80
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vOxQRjpOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BTQAM56wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func j7BgsewCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5X6Qt4Q9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 62
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XDtHT0qbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func j8ItCvknWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func H91nw044Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nIp8mA57Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DAoOJPnZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zpKDQNr7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 79
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JshxHjiVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1NBqtr3MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KnknsMvDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FNTzAwOKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rzVjJGKjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7zKxBdc3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aPLk6uWGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lmwmDNVfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JCbTDoHMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RIhpb3c4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 110
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func H6fWAUdhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9Tw5CLZ0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CmP2Vx9wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 283
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1yGv3i5sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 291
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Lqc2binPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func risw6QR1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 260
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wMdEW6uYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 263
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zI1A57U9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wPfZXYm9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 161
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1SlpHXfWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jgTNxauFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ovHS6j3aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CcuAjIw6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2ejTAflBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 226
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func L1et2LNgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FuMH00HOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 140
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FPj9EOuEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EIROJEDpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func p6r4eTVkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zaL3kbdpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func r09vHDtJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 299
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YrGC2MPVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HtDDMukgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func B0N67l5jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eqzEMKvIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WhUMtMNnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 14
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ss1cJtc4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 152
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func APXQiAhfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RKS1p2ieWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func s6zt5ULrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fbiTPrPxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dTqXSKOdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 273
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ucvfoCmxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func acCOQFn3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1VbX0UIeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 260
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UA2TYTukWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2KLWaGz0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 236
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EgloulgBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fvcZTaxbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HghO2rZbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 299
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CZ51eRjwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 22
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fLQwXql5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OyTFi5tNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func G8i4mOnQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func y1BVtNLNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kOKY9cgdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qWq1ZdUSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 208
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bJL4tEz2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ik4xFLSOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Bgq7RJ2SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3ZTlMNVPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zuUWv5kmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 69
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oAvlq6NyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func o9IEYtB6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YJNQQkWKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hiwb6zRdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9bqOItssWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 76
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func g6QGgB2SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func M8LecL7aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 175
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ypDJya5eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QCx8QT5BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sJkY14FeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vSnKME2TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func f2hzekNYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ch9KqUGWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 73
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XvOFuW7KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pSat2byHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func O3B8bIfBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 14
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dzTkIXhDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 253
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

