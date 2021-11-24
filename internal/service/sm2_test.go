package service

import (
	"fmt"
	"testing"
)

const (
	privateKey = "D5F2AFA24E6BA9071B54A8C9AD735F9A1DE9C4657FA386C09B592694BC118B38"
	publicKey  = "MFkwEwYHKoZIzj0CAQYIKoEcz1UBgi0DQgAE1xPq3B3Cw2U+t+R7Fb0JCJvy87/LDbUDFilGjkQU89VLl57pbUPLKUwP2jnAyOEKmJS9USsz+VwXNd4/bjdIFA=="
)

func TestSign(t *testing.T) {
	//privateKey := "D5F2AFA24E6BA9071B54A8C9AD735F9A1DE9C4657FA386C09B592694BC118B38"
	//src := "biz_content={\"merId\":\"xxxxxxxx\",\"userId\":\"xxxxxxxx\",\"orderId\":\"202110160150337498\",\"notifyUrl\":\"https:\\/\\/baidu.com\",\"txnAmt\":10,\"tradeScene\":\"OFFLINE\"}&encoding=UTF-8&signMethod=02&version=0.0.1"
	src := "{\"biz_content\":\"{\\\"orderId\\\": \\\"STtest0519080001TESTABAB1\\\", \\\"body\\\": \\\"\\\\u5355\\\\u7b14\\\\u652f\\\\u4ed8body\\\", \\\"mchReserved\\\": \\\"\\\\u5355\\\\u7b14\\\\u652f\\\\u4ed8\\\\u5546\\\\u6237\\\\u4fdd\\\\u7559\\\\u57df\\\", \\\"notifyUrl\\\": \\\"http://99.12.95.160:30040/mch/api/notifyYTH\\\", \\\"currencyCode\\\": \\\"156\\\", \\\"userId\\\": \\\"N003525987\\\", \\\"payValidTime\\\": null, \\\"tradeScene\\\": \\\"OFFLINE\\\", \\\"serviceFee\\\": \\\"111\\\", \\\"termId\\\": \\\"00000001\\\", \\\"txnAmt\\\": \\\"1111\\\", \\\"plantformNo\\\": \\\"308999170120FWX\\\", \\\"merId\\\": \\\"308999170120FWY\\\"}\",\"encoding\":\"UTF-8\",\"signMethod\":\"02\",\"version\":\"0.0.1\",\"sign\":\"MEUCIQDbGeUpR00VA7vUdHDcukYd9pyDsWjQ+tOpRJcOoMYd2AIgPdhcIT+1bDaoAetsbhHEIpPt0Dy3D+PFk1FCPt03HuI=\"}"

	sign, err := Sign(privateKey, src)
	if err != nil {
		t.Errorf("sign error %s", err)
	}
	fmt.Println("sign:", sign)
}

func TestVerify(t *testing.T) {
	//publicKey := "MFkwEwYHKoZIzj0CAQYIKoEcz1UBgi0DQgAE6Q+fktsnY9OFP+LpSR5Udbxf5zHCFO0PmOKlFNTxDIGl8jsPbbB/9ET23NV+acSz4FEkzD74sW2iiNVHRLiKHg=="

	sign := "MEQCIFZV8MUlpMpPtm6ojT14jcW1rV/diXiH7es5oe3ka2k0AiBQao6QtYghQkPdVhRHQkk6/6cRUHP3QzyN7QyFMD8Kgw=="

	src := "{\"biz_content\":\"{\\\"orderId\\\": \\\"STtest0519080001TESTABAB1\\\", \\\"body\\\": \\\"\\\\u5355\\\\u7b14\\\\u652f\\\\u4ed8body\\\", \\\"mchReserved\\\": \\\"\\\\u5355\\\\u7b14\\\\u652f\\\\u4ed8\\\\u5546\\\\u6237\\\\u4fdd\\\\u7559\\\\u57df\\\", \\\"notifyUrl\\\": \\\"http://99.12.95.160:30040/mch/api/notifyYTH\\\", \\\"currencyCode\\\": \\\"156\\\", \\\"userId\\\": \\\"N003525987\\\", \\\"payValidTime\\\": null, \\\"tradeScene\\\": \\\"OFFLINE\\\", \\\"serviceFee\\\": \\\"111\\\", \\\"termId\\\": \\\"00000001\\\", \\\"txnAmt\\\": \\\"1111\\\", \\\"plantformNo\\\": \\\"308999170120FWX\\\", \\\"merId\\\": \\\"308999170120FWY\\\"}\",\"encoding\":\"UTF-8\",\"signMethod\":\"02\",\"version\":\"0.0.1\",\"sign\":\"MEUCIQDbGeUpR00VA7vUdHDcukYd9pyDsWjQ+tOpRJcOoMYd2AIgPdhcIT+1bDaoAetsbhHEIpPt0Dy3D+PFk1FCPt03HuI=\"}"

	status, err := Verify(src, sign, publicKey)
	if err != nil {
		t.Errorf("verify error %s", err)
	}

	if status != true {
		t.Error("verify fail")
	}

	fmt.Println("status", status)

}

func TestGenerateKey(t *testing.T) {
	pri, pub, err := GenerateKey()
	if err != nil {
		t.Error(err)
	}

	src := "encoding=UTF-8&signMethod=02&version=0.0.1"

	// sign
	sign, err := Sign(pri, src)
	if err != nil {
		t.Error(err)
	}

	// verify
	status, err := Verify(src, sign, pub)
	if err != nil {
		t.Error(err)
	}

	if status != true {
		t.Error("verify fail")
	}
}
