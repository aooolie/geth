package main

import (
    "crypto/sha256"
    "module"
    "util"
)

func main() {
    util.Start()
    inputs := module.ReadTable()
    module.RedisWrite(inputs)
    for country := range inputs {
        key:=sha256.New()
        key.Write([]byte(inputs[country]))
        r:=key.Sum(nil)
        hash := module.SendGeth("0x" + util.Convert(r))
        module.RedisSingleWrite("hash" + country, hash)
    }

    /**key <-> value
    *  key <-> hash (the key here is "hash" + key)
    */
    println(module.RedisRead("024764f06056bea60eac2b9697db9988d25d8cb39bd92ffa190599803ecb6ba5"))
    println(module.RedisRead("hash024764f06056bea60eac2b9697db9988d25d8cb39bd92ffa190599803ecb6ba5"))

    // check1: assert module.RedisRead(key) equals sha256(value) will be ojbk
    // check2 depends on how geth auth works
}
