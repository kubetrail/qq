# qq
`qq` is a Diffie-Hellman based peer to peer secure messaging app.
The intended use is to establish a shared key for sending and receiving
short messages, hence the name `qq`, short for quick question!

## installation
Download the code and cd to the folder, then run
```bash
go install
```

## usage
### initialize shared key
Let's assume we have Alice and Bob needing to communicate with each 
other. They both first need to run `qq init` and walk through steps
of sending and receiving each other's public keys:

When Alice and Bob run `qq init` command in their own terminals, they 
are requested to send to the other a public key and enter what they
received from the other.

Alice:
```bash
Alice $ qq init
Send to peer: AjgyeiQBvbed1L3ySG82fraMqTQGe1hkdhFTo6fqTJNKygzEUkZhnn7p9NUT1yWUjtgQMpRPixMPWkwFqDHzmZe8zPdoFw8t4MpcRzmAWQ92YVYWd1v2Pww92MHpCQBpzonCQgN8a8iEibsmFz6XzyAEXSugbQCZRGKw6quLzXxicXo31zob66cBcLrub7QfSAh63fxQPePfYRgsTDE6rpK5Fkby566Z5Nd6yjHUqaaV75xHydwJHfLm7S1KnohQ6kidg8sm97MKRME6zE872TNw7YcGDAh5mFiVHB4fmTsZt5rX1NgUsTKz6GrmHoq4hzjRP1fDx4GyH2a6CqB5fwqCrNyxdw
Enter peer key: 
```
Bob:
```bash
Bob $ qq init
Send to peer: Bz5VPsJyDuEpp9gppAuSE1d5AKc1Q6KrHL93sPBcc9E323uDgrXvFsLG8ePmNQTNb2N7YHvb4mADgj9tJFmvJSjNJ29KLpiFewgcTfsH6Jdj6aqkRUSPFBuZ22JkxMeCYjLTsGrCH9HuueXjoQN8rQY4n8s4QTJo5MyYEf3GpHJ5k7V4Ywkm4fSGMpjzjeeKGqncCUBWERyRcBDcXX8ioeSFzxB6GR1yaMoY9hm88dr7UUEqdgJa4h8de1gTu4QSUNMYVJCEiThNdNoxpPHqAhMpcFCmYND4Umwa41svnGw7dhL2mkH4XSqJ3oCP26Xy1np77HwRqtJ1NXjWwrYQ87fw5s5nXU
Enter peer key: 
```

When they both enter what they receive from each other, they both receive
a shared key that can then be used for communicating

Alice:
```bash
Alice $ qq init
Send to peer: AjgyeiQBvbed1L3ySG82fraMqTQGe1hkdhFTo6fqTJNKygzEUkZhnn7p9NUT1yWUjtgQMpRPixMPWkwFqDHzmZe8zPdoFw8t4MpcRzmAWQ92YVYWd1v2Pww92MHpCQBpzonCQgN8a8iEibsmFz6XzyAEXSugbQCZRGKw6quLzXxicXo31zob66cBcLrub7QfSAh63fxQPePfYRgsTDE6rpK5Fkby566Z5Nd6yjHUqaaV75xHydwJHfLm7S1KnohQ6kidg8sm97MKRME6zE872TNw7YcGDAh5mFiVHB4fmTsZt5rX1NgUsTKz6GrmHoq4hzjRP1fDx4GyH2a6CqB5fwqCrNyxdw
Enter peer key: Bz5VPsJyDuEpp9gppAuSE1d5AKc1Q6KrHL93sPBcc9E323uDgrXvFsLG8ePmNQTNb2N7YHvb4mADgj9tJFmvJSjNJ29KLpiFewgcTfsH6Jdj6aqkRUSPFBuZ22JkxMeCYjLTsGrCH9HuueXjoQN8rQY4n8s4QTJo5MyYEf3GpHJ5k7V4Ywkm4fSGMpjzjeeKGqncCUBWERyRcBDcXX8ioeSFzxB6GR1yaMoY9hm88dr7UUEqdgJa4h8de1gTu4QSUNMYVJCEiThNdNoxpPHqAhMpcFCmYND4Umwa41svnGw7dhL2mkH4XSqJ3oCP26Xy1np77HwRqtJ1NXjWwrYQ87fw5s5nXU
export AES_KEY=GT15cXvXqEz4RnMjLMKh3CxBMEmdXeYXsfYeVCuRaoX3
```
Bob:
```bash
Bob $ qq init
Send to peer: Bz5VPsJyDuEpp9gppAuSE1d5AKc1Q6KrHL93sPBcc9E323uDgrXvFsLG8ePmNQTNb2N7YHvb4mADgj9tJFmvJSjNJ29KLpiFewgcTfsH6Jdj6aqkRUSPFBuZ22JkxMeCYjLTsGrCH9HuueXjoQN8rQY4n8s4QTJo5MyYEf3GpHJ5k7V4Ywkm4fSGMpjzjeeKGqncCUBWERyRcBDcXX8ioeSFzxB6GR1yaMoY9hm88dr7UUEqdgJa4h8de1gTu4QSUNMYVJCEiThNdNoxpPHqAhMpcFCmYND4Umwa41svnGw7dhL2mkH4XSqJ3oCP26Xy1np77HwRqtJ1NXjWwrYQ87fw5s5nXU
Enter peer key: AjgyeiQBvbed1L3ySG82fraMqTQGe1hkdhFTo6fqTJNKygzEUkZhnn7p9NUT1yWUjtgQMpRPixMPWkwFqDHzmZe8zPdoFw8t4MpcRzmAWQ92YVYWd1v2Pww92MHpCQBpzonCQgN8a8iEibsmFz6XzyAEXSugbQCZRGKw6quLzXxicXo31zob66cBcLrub7QfSAh63fxQPePfYRgsTDE6rpK5Fkby566Z5Nd6yjHUqaaV75xHydwJHfLm7S1KnohQ6kidg8sm97MKRME6zE872TNw7YcGDAh5mFiVHB4fmTsZt5rX1NgUsTKz6GrmHoq4hzjRP1fDx4GyH2a6CqB5fwqCrNyxdw
export AES_KEY=GT15cXvXqEz4RnMjLMKh3CxBMEmdXeYXsfYeVCuRaoX3
```

As you can see both of them now have `AES_KEY=GT15cXvXqEz4RnMjLMKh3CxBMEmdXeYXsfYeVCuRaoX3` as a 
shared secret, which can then be used for communicating.

> The shared secret needs to be an env. var for commands shown below

### communicating messages
Alice can now send message to Bob
```bash
Alice $ export AES_KEY=GT15cXvXqEz4RnMjLMKh3CxBMEmdXeYXsfYeVCuRaoX3
Alice $ qq e hi Bob this is Alice
71TX811EMMVztMgtwkhSUtZ5DuRwMu3JNwbcXNLyjBDhvKjVuNBTrAeN1J9r8uzoEV
```
Bob can then decipher the message from Alice
```bash
Bob $ export AES_KEY=GT15cXvXqEz4RnMjLMKh3CxBMEmdXeYXsfYeVCuRaoX3
Bob $ qq d 71TX811EMMVztMgtwkhSUtZ5DuRwMu3JNwbcXNLyjBDhvKjVuNBTrAeN1J9r8uzoEV
hi Bob this is Alice
```
