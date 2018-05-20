package main

import (
    "testing"
    "math/rand"
    "log"
    gi "github.com/iotaledger/giota"
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

func init()  {
    ts := "DOZREH9PKR9UCVWUBMFFHLHYVMJXRFQAUJSEBEWV9DAQZC9FCAYKQRVBKPMONBRKDRGKROBGQMGVRXKFA"
    s, err := gi.ToTrytes(ts)
    if err != nil {
        log.Fatal(err)
    } else {
        iSeed = s
    }
}

func BenchmarkBundle(b *testing.B) {

    var err error

    if err != nil {
        b.Errorf("There was an error initialising the test: %s", err)
    }

    var exBundle gi.Bundle

    for i := 0; i < 10; i++ {
        api := gi.NewAPI("http://localhost:14265", nil)
        randVal := rand.Int63n(1279530283277761)
        iTrs[0].Value = randVal
        bdl, err := gi.PrepareTransfers(api, iSeed, iTrs, nil, "", 2)

        if err != nil {
            b.Error(err)
        }

        if i == 0 {
            exBundle = bdl
        }
    }

    tryteLen := 0
    for _, tran := range exBundle {
        tryteLen += len(tran.Trytes())
    }

    b.Logf("The length of a bundle in Trytes is %v\n", tryteLen)
}

func BenchmarkBundleValid(b *testing.B) {

    var err error

    if err != nil {
        b.Errorf("There was an error initialising the test: %s", err)
    }

    for i := 0; i < 10; i++ {
        api := gi.NewAPI("http://localhost:14265", nil)
        randVal := rand.Int63n(1279530283277761)
        iTrs[0].Value = randVal
        bdl, err := gi.PrepareTransfers(api, iSeed, iTrs, nil, "", 2)

        if err != nil {
            b.Error(err)
        }

        if err = bdl.IsValid(); err != nil {
            b.Error(err)
        }
    }
}