package main

import (
    "flag"
    "github.com/pkg/profile"
    "math/rand"
    gi"github.com/iotaledger/giota"
    "log"
    "fmt"
)

var (
    iSeed            gi.Trytes
)

var  iTrs = []gi.Transfer{
    {
        Address: "BXHANKTHPJUPUVZOLJPZPQLDZPWVSBPGLMLSOYFZM9RSHVZRRBZJZJDZYTNRHXBVMQKFT9DVKVNDPCGC9ZXXTZCTMB",
        Value:   1500000,
        Tag:     "RPROOF",
    },
}

func main() {
    ts := "DOZREH9PKR9UCVWUBMFFHLHYVMJXRFQAUJSEBEWV9DAQZC9FCAYKQRVBKPMONBRKDRGKROBGQMGVRXKFA"
    s, err := gi.ToTrytes(ts)
    if err != nil {
        log.Fatal(err)
    } else {
        iSeed = s
    }

    // This code was token from the godocs at https://godoc.org/github.com/pkg/profile
    // use the flags package to selectively enable profiling.
    mode := flag.String("profile.mode", "", "enable profiling mode, one of [cpu, mem, mutex, block]")
    flag.Parse()
    switch *mode {
    case "cpu":
        defer profile.Start(profile.CPUProfile).Stop()
    case "mem":
        defer profile.Start(profile.MemProfile).Stop()
    case "mutex":
        defer profile.Start(profile.MutexProfile).Stop()
    case "block":
        defer profile.Start(profile.BlockProfile).Stop()
    default:
        // do nothing
    }


    var exBundle gi.Bundle

    for i := 0; i < 100; i++ {
        api := gi.NewAPI("http://localhost:14265", nil)
        randVal := rand.Int63n(1279530283277761)
        iTrs[0].Value = randVal
        bdl, err := gi.PrepareTransfers(api, iSeed, iTrs, nil, "", 2)
        if err != nil {
            log.Fatal(err)
        }

        if bdl.IsValid() != nil {
            log.Fatal(err)
        }
        if i == 0 {
            exBundle = bdl
        }
    }

    tryteLen := 0
    for _, tran := range exBundle {
        tryteLen += len(tran.Trytes())
    }

    fmt.Printf("The length of a bundle in Trytes is %v\n", tryteLen)
}