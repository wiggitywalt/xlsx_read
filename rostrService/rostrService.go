package rostr
import (
	"net/http"
	"fmt"
	"io/ioutil"
	"net/url"
    // "encoding/json"
)

func RetrieveInfo(name string) []byte {
  //map[string]interface {}
	name = url.QueryEscape(name)
	api_token := "81f97f0925378a849b601b5546c77439";
	url := "https://rostr.disney.com/api/v2/people/search?query=" + name + "&token=" + api_token
	resp, err := http.Get(url)

	htmlData, err := ioutil.ReadAll(resp.Body) //<--- here!

	if err != nil {
		fmt.Println(string(htmlData))
	}

	// print out
	//fmt.Println(string(htmlData)) //<-- here !

  // empMap := make(map[string]interface{})
  // err = json.Unmarshal([]byte(htmlData), &empMap)
  //
  // if err != nil {
	// 	// handle error
	// }
	defer resp.Body.Close()

	return htmlData
}
