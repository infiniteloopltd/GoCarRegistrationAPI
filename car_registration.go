package car_registration

import "net/http"
import "io/ioutil"
import "encoding/base64"	
import "encoding/json"

func basicAuth(username, password string) string {
   auth := username + ":" + password
   return base64.StdEncoding.EncodeToString([]byte(auth))
}

func Australia_lookup(registration_number, state, username, password string) interface{} {
	
        url := string("https://www.regcheck.org.uk/api/json.aspx/CheckAustralia/" + registration_number + "/" + state)
	return generic_lookup(url, username, password)
}

func USA_lookup(registration_number, state, username, password string) interface{} {
	
        url := string("https://www.regcheck.org.uk/api/json.aspx/CheckUSA/" + registration_number + "/" + state)
	return generic_lookup(url, username, password)
}

func European_lookup(endpoint, registration_number, username, password string) interface{} {
	
        url := string("https://www.regcheck.org.uk/api/json.aspx/" + endpoint + "/" + registration_number)
	return generic_lookup(url, username, password)
}

func generic_lookup(url, username, password string) interface{} {

	var client http.Client		

	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("Authorization","Basic "+basicAuth(username,password))

	if err != nil {}
	resp, err3 := client.Do(req)

	if err3 != nil {}

	defer resp.Body.Close()

	if resp.StatusCode == 200 { // OK
	    bodyBytes, err2 := ioutil.ReadAll(resp.Body)
	    bodyString := string(bodyBytes)
      	    var data interface{} 
	    json.Unmarshal([]byte(bodyString), &data)
	    return data
	    if err2 != nil {}
	} 	
	return nil
	
}
