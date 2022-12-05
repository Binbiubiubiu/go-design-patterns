package main

import (
	"fmt"
	"log"
)

func main() {
	// otp := Otp{}

	// smsOTP := &Sms{
	// 	Otp: otp,
	// }

	// smsOTP.genAndSendOTP(smsOTP, 4)

	// emailOTP := &Email{
	// 	Otp: otp,
	// }
	// emailOTP.genAndSendOTP(emailOTP, 4)
	// fmt.Scanln()

	smsOTP := &Sms{}
	o := Otp{
		iOtp: smsOTP,
	}
	err := o.genAndSendOTP(4)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("")
	emailOTP := &Email{}
	o = Otp{
		iOtp: emailOTP,
	}
	err = o.genAndSendOTP(4)
	if err != nil {
		log.Fatal(err)
	}
}
