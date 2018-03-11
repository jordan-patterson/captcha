# captcha
Easily implement captchas in your web applications to prevent spam.

## Installation

```bash
 $ go get github.com/jordan-patterson/captcha
```
## Usage

###server
```go
package main

import(
    "fmt"
    "net/http"
    "github.com/jordan-patterson/captcha"
)
func captchaHandler(w http.ResponseWriter,r *http.Request){
    capt:=captcha.New()
    //get session variable here
    //session.Values["captcha"]=capt.Text
    //save session here then write captcha image as response
    err:=capt.WriteImage(w) 
    if err!=nil{
        fmt.Println(err.Error())
    }  
}
func main(){
    http.HandleFunc("/captcha",captchaHandler)
    http.ListenAndServe(":5050",nil)
}
```

###client
```html
    <form method="POST">
        Name:<input type="text" name="name" requierd>
        Email:<input type="email" name="email" requierd>
        Message:<input type="text" name="message" required>
        Are you human? 
        <img src="localhost:5050/captcha">
        <input type="text" name="captcha">
        <input type="submit" value="Send">
    </form>
```