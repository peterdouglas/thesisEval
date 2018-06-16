package main

import (
    "testing"
    "math/rand"
    "log"
    gi "github.com/iotaledger/giota"
)

const ENDPOINT  = "http://localhost:14266"

func init()  {
    ts := "DOZREH9PKR9UCVWUBMFFHLHYVMJXRFQAUJSEBEWV9DAQZC9FCAYKQRVBKPMONBRKDRGKROBGQMGVRXKFA"
    s, err := gi.ToTrytes(ts)
    if err != nil {
        log.Fatal(err)
    } else {
        iSeed = s
    }
}

func benchmarkBundle(i int, api *gi.API) gi.Bundle {
    randVal := rand.Int63n(1279530283277761)
    iTrs[0].Value = randVal
    bdl, err := gi.PrepareTransfers(api, iSeed, iTrs, nil, "", 2)

    if err != nil {
        log.Fatal(err)
    }

    if i == 0 {
        return bdl
    } else {
        return nil
    }
}

func BenchmarkBundle(b *testing.B) {
    b.ReportAllocs()
    var err error

    if err != nil {
        b.Errorf("There was an error initialising the test: %s", err)
    }

    var exBundle gi.Bundle
    api := gi.NewAPI(ENDPOINT, nil)

    for i := 0; i < b.N; i++ {
        tempBdl := benchmarkBundle(i, api)
        if tempBdl != nil {
            exBundle = tempBdl
        }
    }

    tryteLen := 0
    for _, tran := range exBundle {
        tryteLen += len(tran.Trytes())
    }

    b.Logf("The length of a bundle in Trytes is %v\n", tryteLen)
}

func BenchmarkAddressing(b *testing.B) {
    seed := gi.NewSeed()
    b.ReportAllocs()

    for i := 0; i < b.N; i++  {
        _, err := gi.NewAddress(seed, i, 2)
        if err != nil {
            b.Error(err)
        }
    }
}

func BenchmarkBundleValid(b *testing.B) {
    b.ReportAllocs()
    var err error

    if err != nil {
        b.Errorf("There was an error initialising the test: %s", err)
    }

    for i := 0; i < b.N; i++ {
        api := gi.NewAPI(ENDPOINT, nil)
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

func BenchmarkGetBalances(b *testing.B) {
    b.ReportAllocs()
    api := gi.NewAPI(ENDPOINT, nil)
    for i := 0; i < b.N; i++  {
        // If inputs with enough balance
        _, err := gi.GetInputs(api, iSeed, 0, 100, 100, 2)
        if err != nil {
            b.Error(err)
        }
    }
}