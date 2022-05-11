package util

import (
	"io"
	"io/ioutil"
	"wechatNotify/pkg/logging"

	"github.com/emersion/go-imap"
	"github.com/emersion/go-imap/client"
	"github.com/emersion/go-message/charset"
	"github.com/emersion/go-message/mail"
)

type MailConfig struct {
	User       string `json:"user"`
	Pass       string `json:"pass"`
	MailServer string `json:"server"`
}

type ServerGmail struct {
	MailServer string
	User       string
	Pass       string
	client     *client.Client
}

type MailInfo struct {
	Uid      uint32 `json:"uid"`
	Date     string `json:"date"`
	From     string `json:"from"`
	To       string `json:"to"`
	Subject  string `json:"subject"`
	Body     string `json:"body"`
	Filename string `json:"filename"`
}

func NewServerGmail(MailServer string, User string, Pass string) *ServerGmail {
	server := &ServerGmail{
		MailServer: MailServer,
		User:       User,
		Pass:       Pass,
	}
	if err := server.connect(); err != nil {
		logging.Error("连接失败", err)
		return nil
	}
	if ok, err := server.login(); !ok {
		logging.Error("登录失败", err)
		return nil
	}

	return server
}

func (t *ServerGmail) connect() error {
	logging.Info("连接服务器中...")
	c, err := client.DialTLS(t.MailServer, nil)
	if err != nil {
		logging.Error(err)
		return err
	}
	logging.Info("连接成功")
	t.client = c
	return nil
}

func (t *ServerGmail) login() (bool, error) {
	logging.Info("登录中...")
	if err := t.client.Login(t.User, t.Pass); err != nil {
		logging.Error("登录失败", err)
		return false, err
	}
	return true, nil
}

func (t *ServerGmail) GetTotalMail() (int, error) {
	logging.Info("获取邮件总数...")
	mbox, err := t.client.Select("INBOX", false)
	if err != nil {
		logging.Error("获取邮件总数失败", err)
		return 0, err
	}
	return int(mbox.Messages), nil
}

func (t *ServerGmail) GetMail(from uint32, to uint32) ([]MailInfo, error) {
	logging.Info("获取邮件...")
	var messages []MailInfo
	seqset := new(imap.SeqSet)
	seqset.AddRange(from, to)

	messageCh := make(chan *imap.Message, 10)
	section := imap.BodySectionName{}
	items := []imap.FetchItem{section.FetchItem()}
	done := make(chan error, 1)
	go func() {
		done <- t.client.Fetch(seqset, items, messageCh)
	}()

	imap.CharsetReader = charset.Reader
	for msg := range messageCh {
		var message MailInfo
		var err error
		r := msg.GetBody(&section)
		if r == nil {
			logging.Warn("获取邮件失败", err)
			return nil, err
		}
		mr, err := mail.CreateReader(r)
		if err != nil {
			logging.Error("获取邮件失败", err)
			return nil, err
		}

		message.Uid = msg.SeqNum
		header := mr.Header
		var subject string
		if date, err := header.Date(); err == nil {
			message.Date = date.String()
		}
		if from, err := header.AddressList("From"); err == nil {
			message.From = from[0].String()
		}
		if to, err := header.AddressList("To"); err == nil {
			message.To = to[0].String()
		}
		if subject, err = header.Subject(); err == nil {
			message.Subject = subject
		}

		// 处理邮件正文
		for {
			p, err := mr.NextPart()
			if err == io.EOF {
				break
			} else if err != nil {
				logging.Warn("NextPart:err ", err)
				return nil, err
			}
			// 正文消息文本
			switch h := p.Header.(type) {
			case *mail.InlineHeader:
				// 正文消息文本
				b, _ := ioutil.ReadAll(p.Body)
				message.Body = string(b)
			case *mail.AttachmentHeader:
				// 正文内附件
				filename, _ := h.Filename()
				message.Filename = filename
			}
		}
		messages = append(messages, message)
	}
	if err := <-done; err != nil {
		logging.Error("获取邮件失败", err)
		return nil, err
	}
	return messages, nil
}

func (t *ServerGmail) logout() {
	logging.Info("退出中...")
	t.client.Logout()
	logging.Info("退出成功")
}

func CheckMailConfig(MailConfig MailConfig) bool {
	if MailConfig.User == "" || MailConfig.Pass == "" || MailConfig.MailServer == "" {
		return false
	}
	return true
}
