package main

import (
	"bytes"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"log"
	"fmt"
	"strconv"
)
func main(){

	botToken:="1957257389:AAESeMyPNR3rFrwCCxVW3VSf_2Bg326kcaI"
	//"https://api.telegram.org/bot"
	botApi:="https://api.telegram.org/bot"
	botUrl:=botApi+botToken
	offset:=0
	for{
		updates,err:=getUpdates(botUrl, offset)
		if err!= nil {
			log.Println("Something went wrong",err.Error())
		}
		for _, update :=range updates {
			err=respond(botUrl,update)
			offset=update.UpdateId+1
		}
		fmt.Println(updates)

	}
}
//запрос обновлений
func getUpdates(botUrl string,offset int) ([]Update,error){
	resp,err:=http.Get(botUrl+"/getUpdates"+"?offset="+strconv.Itoa(offset))
	if err!= nil {
		return nil, err
	}
	defer resp.Body.Close()
	body,err:=ioutil.ReadAll(resp.Body)
	if err!= nil {
		return nil, err
	}
	var restResponse RestResponse
	err=json.Unmarshal(body,&restResponse)
	fmt.Println(restResponse)
	if err!= nil {
		return nil, err
	}
	return restResponse.Result,nil
}



// ответ на обновления
func respond(botUrl string, update Update) (error) {
	var botMessage BotMessage
	botMessage.ChatId =update.Message.Chat.ChatId
	botMessage.Text=update.Message.text
	buf,err:=json.Marshal(botMessage)
	if err!= nil {
		return err
	}
	_, err = http.Post(botUrl+"/sendMessage", "application/json", bytes.NewBuffer(buf))
	if err!= nil {
		return err
	}
	return nil
}