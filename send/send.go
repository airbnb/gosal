package send
//
// import (
//   "net/http"
//   "log"
//   "io/ioutil"
//   "bytes"
// )

// func SalPostReq(data []byte) []byte {
//     client := &http.Client{}
//
//     // load config from config.json
//     url := LoadConfig("./config.json").URL + "checkin/"
//     key := LoadConfig("./config.json").Key
//
//     //pass data to the request's body
//     req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
//     req.SetBasicAuth("sal", key)
//     resp, err := client.Do(req)
//     if err != nil {
//         log.Fatal(err)
//     }
//     bodyText, err := ioutil.ReadAll(resp.Body)
//
//     return bodyText
// }
