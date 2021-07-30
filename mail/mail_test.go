package mail

import "testing"

func TestMail(t *testing.T) {
	config := `{"username":"gitlab@dianjiu.cc","password":"Xj576NSKATxxCJiY","host":"smtp.exmail.qq.com","port":465}`
	mail := NewEMail(config)
	if mail.Username != "gitlab@dianjiu.cc" {
		t.Fatal("email parse get username error")
	}
	if mail.Password != "Xj576NSKATxxCJiY" {
		t.Fatal("email parse get password error")
	}
	if mail.Host != "smtp.exmail.qq.com" {
		t.Fatal("email parse get host error")
	}
	if mail.Port != 465 {
		t.Fatal("email parse get port error")
	}
	mail.To = []string{"dianjiu@dianjiu.cc"}
	mail.From = "dianjiu@dianjiu.cc"
	mail.Subject = "hi, just from gotool!"
	mail.Text = "Text Body is, of course, supported!"
	mail.HTML = "<h1>Fancy Html is supported, too!</h1>"
	//mail.AttachFile("/c/Users/dianjiu/Pictures/gzh.jpg")
	mail.Send()
}
