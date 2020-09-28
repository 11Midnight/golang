package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"github.com/ghodss/yaml"
)
func scanner() string {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		return scanner.Text()
	}
	return scanner.Text()
}
func main() {
	fileNameYaml:=scanner()
	yamlData, err := ioutil.ReadFile(fileNameYaml)
	if err != nil {
		fmt.Printf("Error reading YAML file: %s\n", err)
		return
	}
	/*j := []byte(`{"gateways":{"sepp-ingressgateway":{"autoscaleEnabled":false,"replicaCount":1}},"global":{"config":{"localsepp":{"domain":"remote.sepp2.com","scp":{"ip":"10.178.254.185","port":30001},"securityCapability":"TLS","tls":{"certificate":"-----BEGIN CERTIFICATE-----\nMIIDyDCCArCgAwIBAgIJAJ/NlN2hRmiiMA0GCSqGSIb3DQEBCwUAMHkxCzAJBgNV\nBAYTAlVTMQswCQYDVQQIDAJDQTELMAkGA1UEBwwCU1YxDTALBgNVBAoMBENHQlUx\nDzANBgNVBAsMBk9yYWNsZTESMBAGA1UEAwwJbG9jYWxob3N0MRwwGgYJKoZIhvcN\nAQkBFg14eXpAZ21haWwuY29tMB4XDTE5MDQwMTExMTQxNloXDTIwMDMzMTExMTQx\nNloweTELMAkGA1UEBhMCVVMxCzAJBgNVBAgMAkNBMQswCQYDVQQHDAJTVjENMAsG\nA1UECgwEQ0dCVTEPMA0GA1UECwwGT3JhY2xlMRIwEAYDVQQDDAlsb2NhbGhvc3Qx\nHDAaBgkqhkiG9w0BCQEWDXh5ekBnbWFpbC5jb20wggEiMA0GCSqGSIb3DQEBAQUA\nA4IBDwAwggEKAoIBAQC/zV8n5kOOKWLuqlZD1cf7KAtgSAH1MB8biR9mKAJ/sH1M\n1/OR11FXrEs4Sq3yB+6HcNgwg/le3XAncGXbKRSI8MeuLftewwfj2q9q8QnIUftX\nDlyq92Zd675wpPzp4JjRewPgJFsSBMa8svZBZJzIDWLaIke90Bx9bJNzhs+aLijW\n10rZVeO5bRzKaU8ze/7Py1IGc/HWhR+S+DRVInMwMui5ozqyZZaprVHjDititSoS\nHcWphI0ef8e6w/mfe2AcUKNugI4GGfK+egJ3SXSe5BzJSx+EJgKQaSP3/DTGCSju\nC33uSKch43IhX+zWL2FOQD0I/Y8YCIYHTA0GJevPAgMBAAGjUzBRMB0GA1UdDgQW\nBBS5wMSwF048S+pa+onXF7t5LBw3EDAfBgNVHSMEGDAWgBS5wMSwF048S+pa+onX\nF7t5LBw3EDAPBgNVHRMBAf8EBTADAQH/MA0GCSqGSIb3DQEBCwUAA4IBAQBQEG9o\nIAIOwGkpbsqbFfFti9gSilaaP8rzSN8Thj7PvHxem6oc/Gj9ioyWcCAKk3t6gZgd\ndtXo5SJp+DXFpNnzmmLkHFLbloZwc8QwPoIE4aqiYknBlKdvmXGZSaRYI8CEk8dK\nuH5Dot5mJWGE391zAuRPqvJSl4rborTZGfk1Nl8KQPi4P2GnoFxZHzM8Pt6LYYJq\nlOn7BgF8rGMoSwVB1+Znt2R5KaeMANE4WOW9qrqE2fd0Oq23cOBXckizCzc5HwDF\n6Ny3JHxs/sf0ZX8teyXlHgv153O/KMaQijYcU1WCJ5ETAC/UbzkRgWUbcoDFFUSr\nsRpvmOpp7/KKGh3c\n-----END CERTIFICATE-----","cipherSuites":["[ECDHE-ECDSA-AES128-GCM-SHA256|ECDHE-ECDSA-CHACHA20-POLY1305]","[ECDHE-RSA-AES128-GCM-SHA256|ECDHE-RSA-CHACHA20-POLY1305]","ECDHE-ECDSA-AES128-SHA256","ECDHE-RSA-AES128-SHA256","ECDHE-ECDSA-AES128-SHA","ECDHE-RSA-AES128-SHA","AES128-GCM-SHA256","AES128-SHA256","AES128-SHA","ECDHE-ECDSA-AES256-GCM-SHA384","ECDHE-RSA-AES256-GCM-SHA384","ECDHE-ECDSA-AES256-SHA384","ECDHE-RSA-AES256-SHA384","ECDHE-ECDSA-AES256-SHA","ECDHE-RSA-AES256-SHA","AES256-GCM-SHA384","AES256-SHA256","AES256-SHA"],"maxProtocolVersion":"TLSV1_3","minProtocolVersion":"TLSV1_3","privateKey":"-----BEGIN RSA PRIVATE KEY-----\nMIIEogIBAAKCAQEAv81fJ+ZDjili7qpWQ9XH+ygLYEgB9TAfG4kfZigCf7B9TNfz\nkddRV6xLOEqt8gfuh3DYMIP5Xt1wJ3Bl2ykUiPDHri37XsMH49qvavEJyFH7Vw5c\nqvdmXeu+cKT86eCY0XsD4CRbEgTGvLL2QWScyA1i2iJHvdAcfWyTc4bPmi4o1tdK\n2VXjuW0cymlPM3v+z8tSBnPx1oUfkvg0VSJzMDLouaM6smWWqa1R4w4rYrUqEh3F\nqYSNHn/HusP5n3tgHFCjboCOBhnyvnoCd0l0nuQcyUsfhCYCkGkj9/w0xgko7gt9\n7kinIeNyIV/s1i9hTkA9CP2PGAiGB0wNBiXrzwIDAQABAoIBABQHonRMO6nQ8J/z\n8eZLoQK9f9KiMnIkgFsjyr+kVf53s94efzUvwzLgr0pFojywT1ydkDYg0h3yChBU\nUI6+j4JtgqzR6HpQ67YD1GXanibpSNIikdLY2zvbrpuc3zydU9gsfI8fvNL45Lg4\nCDd0Bltt6nDVO6C65v7CI/PglDWJhia2oz3HfR8ruDPw3EP8W+L7qtSVcKprVj82\nv8LQC/SVhzwCAyA9RlpU/z6NHD5wqn3a4+f3fNnAzFWfVp8S8ZxYemiw7q0mowob\n2ogf9eHuIMy2P20bqHcU4LeJlaGz1PkPiUtXuRPwRJNkA1FiR+2dm/35BofO/XpZ\nm9UYoSECgYEA6q4PoHYHB/8yQv+2N2xFZmgW83fFAsfCojPteSxb7lUBf/9EEkbB\nPuY30PCeyOceEBy+Iz7PbkxBzcwkdPos1qs82knGMOM7TCDFlVm1kJRWg/M+LBP9\nc14mj3+PbazWQJTIo2fPzbJUSY29qBzyUWXtOaQIZ8pSxTVVc8lhTWsCgYEA0Tob\nOufE8x5LVRdqOF0b97za50z9cAkS9d6hb9eJ+ApEkTe07DBgnhmDzup98WMEvPM6\nKd2jb0U9ACAo3Dudz1dDtfLHjdWd27QKInS2hw5xXceStMV55NOAFNQ0oF+dfJa8\naBua8m9b+BzQQAh6NSqytcHVak1aliDY5GUl8C0CgYBePOlfiSXHq022o9K7LMot\nlkzpFcZSGH9q/Sk2SH7eFNqWWJ4wLKN7K51sVMOeb1ieJbygkJ9hblPaXUiihLeu\nuHKmQYvJo+Xy6xIMqEqFtVuMfwgPUTOsuTkI1LN22jnrExQCsjQ7KIo6QyXOtVkF\nIMYDKICLlS5prMIUzeA54wKBgHfBAWry05iv40Bd+Y8vQ93Fe6neOEVS/EY8WjyQ\nqsiM3/gaYXS6r+JuCjJ5pwJtwX2A3e6ujGgYwjR7M8fyW34cnYXb4vo8pXDmGNLl\n6L9eteaOX1sWmJEvuWSynTiZ4aM5B7ey7ToMISDfJRcxgvlBai58NnH0un+pZ1s6\nxb81AoGAB7sIVYgdJXXCS1KMtYlhpxhOOOkvuJBe0YSfzYXb9kjktoKOK8xSqKGE\nmLQsZzMWP5CubRY3PiJ7GhcV1RvEJsj5sR493Mk8ukEP/LTWlZMh7BR5YXr4tg2U\n9gCKExnTlMkeDJNRuw43MslPxJiaXMds7BDgkCvcf3Av9OEvEvk=\n-----END RSA PRIVATE KEY-----"}},"remotesepps":[{"caCertificate":"-----BEGIN CERTIFICATE-----\nMIIDyjCCArKgAwIBAgIJANO6mMns+CWEMA0GCSqGSIb3DQEBCwUAMHoxCzAJBgNV\nBAYTAklOMQswCQYDVQQIDAJLQTEMMAoGA1UEBwwDQkxSMQ8wDQYDVQQKDAZPcmFj\nbGUxDTALBgNVBAsMBENHQlUxEjAQBgNVBAMMCWxvY2FsaG9zdDEcMBoGCSqGSIb3\nDQEJARYNYWJjQGdtYWlsLmNvbTAeFw0xOTA0MDExMTExMTJaFw0yMDAzMzExMTEx\nMTJaMHoxCzAJBgNVBAYTAklOMQswCQYDVQQIDAJLQTEMMAoGA1UEBwwDQkxSMQ8w\nDQYDVQQKDAZPcmFjbGUxDTALBgNVBAsMBENHQlUxEjAQBgNVBAMMCWxvY2FsaG9z\ndDEcMBoGCSqGSIb3DQEJARYNYWJjQGdtYWlsLmNvbTCCASIwDQYJKoZIhvcNAQEB\nBQADggEPADCCAQoCggEBALau4TZylSvrxNtuk6ec65zv83gnd++Lioei3QDutU/J\nvVyl0t51rkf1tyqmqjXhYetFwAAH3sNJ5etGKuq0N4opHZRsUaKhxVgPuZ+5bOzT\nOD/iCW+CEsGPGyGN9agUJNp7TH0R3+lruGz+Z7u0M90rlEl/dJkWgMwzZ69O17pB\nzML3fBQ+p5do7p5jzvMMLT7JzzVW+k+2+kEzuIMh0Yp1IpA/e9Qk9vI7uKYI7Uxy\nUuaPTECZrDffe+eTVTkrXRt7v0JaUM6LptI7RgL6bnSpTGz0kwN0gO7fZCxGfPSv\n7KpY7gOenZ4QtEhxg9dTNGEGEU4P7/ju2gAyWMfOjMkCAwEAAaNTMFEwHQYDVR0O\nBBYEFNQn+TW2QBzOFi2XG/IF4EBe7kePMB8GA1UdIwQYMBaAFNQn+TW2QBzOFi2X\nG/IF4EBe7kePMA8GA1UdEwEB/wQFMAMBAf8wDQYJKoZIhvcNAQELBQADggEBACgW\nkXWYuwv9OA9nScdbInXubPFpOBDK+V7LA7AnnFElxQldcowzPQvzhSmBB3UBSxfK\nJAsG4eIztVldYavECKDzk9pP1/pYtGA1SRhw2wocqKtORlZoerx7ZcVo8ogtlJIj\n+4K0AIerxREalobAkZb/vambw+Hmj6YTKFG69GTtTAJVK3A33f9xnbCP0mplYvsq\nqitUXmon713P/Zcadwz/kGfwUGzzAFO8VZcKHGS8mxaRKSHlwZMETpE2Zlj9LZDc\n+uWKg29PssujJ17T0IGNxn0C6sH2R2IEPAq4xWYwf0vkHO4SfoacpDJgfCXZMEE3\nwlMx41TEYSgcSqCcp8s=\n-----END CERTIFICATE-----","domain":"remote.sepp1.com","isProducer":true,"n32c":{"ip":"10.178.254.170","port":31390},"n32f":{"ip":"10.178.254.161","port":8433}}]},"mysql":{"image":"mysql","pullPolicy":"IfNotPresent","tag":8},"n32Client":{"image":"ocsepp/ocsepp-client","pullPolicy":"IfNotPresent","tag":"1.1.0"},"n32Server":{"image":"ocsepp/ocsepp-server","pullPolicy":"IfNotPresent","tag":"1.1.0"},"nsRegistration":{"image":"ocsepp-nsregistration","pullPolicy":"IfNotPresent","tag":"1.0.0"},"pilot":{"image":"istio/pilot","pullPolicy":"IfNotPresent","tag":"1.1.3"},"proxy":{"image":"istio/proxyv2","pullPolicy":"IfNotPresent","tag":"1.1.3"},"sds":{"image":"istio/node-agent-k8s","pullPolicy":"IfNotPresent","tag":"1.1.3"}},"nsregistration":{"configFiles":{"nf.profile":"{\n  \"plmn\": {\"mcc\": \"310\", \"mnc\": \"14\"},\n  \"fqdn\": \"ocsepp-nsgateway.ocsepp.svc.us.lab.oracle.com\",\n  \"ipv4Addresses\": [ \"127.0.0.1\", \"10.0.0.1\" ],\n  \"priority\": 1\n}"},"nrf":{"host":"http://10.75.157.63","port":31605}},"ocsepp-client":{"enabled":false},"ocsepp-server":{"enabled":true}}`)
	y, err := yaml.JSONToYAML(j)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}*/
	//fmt.Println(string(yamlData))
	jsonData, err := yaml.YAMLToJSON(yamlData)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	//fmt.Println(string(jsonData))
	fileNameJson:=strings.Replace(fileNameYaml,".yaml",".json", -1)
	jsonFile, err := os.Create("./"+fileNameJson)
	if err != nil {
		panic(err)
	}
	defer jsonFile.Close()
	jsonFile.Write(jsonData)
	jsonFile.Close()
	fmt.Println("JSON data written to ", jsonFile.Name())
}