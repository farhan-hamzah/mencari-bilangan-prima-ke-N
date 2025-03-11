package main

import "fmt"

func bilanganPrimaKeN(masukan int, i, j *int) {
    *i = 2
    *j = 3
    if masukan == 1 {
        fmt.Println(2)
        return
    }

    hitung := 1
    for hitung < masukan {
        isPrima := true

        for *i = 2; *i * *i <= *j; (*i)++ {
            if *j % *i == 0 {
                isPrima = false
                break
            }
        }

        if isPrima {
            hitung++
            if hitung == masukan {
                fmt.Println(*j)
                return
            }
        }
        (*j) += 2
    }
}

func main() {
    var n, a, b int
    fmt.Scan(&n)
    bilanganPrimaKeN(n, &a, &b)
}