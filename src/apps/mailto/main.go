// Copyright FreeWheel Inc. All Rights Reserved.
// Author: jjli@freewheel.tv(Li JingJie)
// Created Time: Wed Jan 11 10:45:57 2017

package main

import (
	"log"
	"net/smtp"
)

func main() {
	send("hello there")
}
func send(body string) {
	from := "adstools_monitor@freewheel.tv"
	pass := "..."
	to := "jjli@freewheel.tv"
	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: Hello there\n\n" +
		body
	err := smtp.SendMail("smtp.fwmrm.net:25",
		smtp.PlainAuth("", from, pass, "smtp.fwmrm.net"),
		from, []string{to}, []byte(msg))
	if err != nil {
		log.Printf("smtp error: %s", err)
		return
	}
	log.Print("sent, visit http://foobarbazz.mailinator.com")
}
