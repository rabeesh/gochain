# Learn basic of Block chain

This code heavily inspired from "Code your own blockchain in less than [200 lines of Go](https://medium.com/@mycoralhealth/code-your-own-blockchain-in-less-than-200-lines-of-go-e296282bcffc)"

### Technologies used

* negroni - For application split ups
* spew - Print formatted structs in  console

### Deployment steps:
- `git clone https://github.com/rabeesh/gochain.git`
- navigate to this directory and rename the example file `mv config/.env-sample config/.env`
- `go run main.go`
- open a web browser and visit `http://localhost:8082/api/blocks`
- to write new blocks, send a `POST` request to `http://localhost:8080/api/blocks` with a JSON payload with `BPM` as the key and an integer as the value. For example:
```
{"BPM":50}
```
- Send as many requests as you like and refresh your browser to see your blocks grow! Use your actual heart rate (Beats Per Minute) to track it over time.
