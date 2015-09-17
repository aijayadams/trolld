package main

import (
    "bufio"
    "io"
    "io/ioutil"
    "log"
    "os"
    "strings"
)

func LoadTrolls(a *Assets, dir string) () {
    files, _ := ioutil.ReadDir(dir)
    for _, f := range files {
        if strings.Contains(f.Name(), "static.troll") {
            troll_asset := LoadFile(dir + "/" + f.Name())
            troll_name := strings.SplitN(f.Name(), ".", 2)[0]
            a.AddAsset(Asset{asset: troll_asset, name: troll_name})
        }
    }
    return
}

func LoadFile(filename string) (bb []byte) {
    f, err := os.Open(filename)
    if err != nil {
        log.Fatal(err)
        return
    }
    filereader := bufio.NewReader(f)
    bb, err = filereader.ReadBytes('\x00')

    if err != io.EOF {
        log.Fatal(err)
        return
    }
    f.Close()
    return
}
