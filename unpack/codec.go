package unpack

import (
	"encoding/binary"
	"errors"
	"io"
)

// MsgHeader 定义模拟的header信息
const MsgHeader = "12345678"


func Encode(w io.Writer,content string) error {
	//header+content_len+content
	//写入header
	if err := binary.Write(w,binary.BigEndian,[]byte(MsgHeader));err != nil {
		return err
	}

	//写入content len
	//强转为int32 占4个字节
	clen := int32(len([]byte(content)))
	if err := binary.Write(w,binary.BigEndian,clen);err != nil {
		return err
	}

	//写入content
	if err := binary.Write(w,binary.BigEndian,[]byte(content));err != nil {
		return err
	}

	return nil
}

func Decode(r io.Reader) ([]byte,error) {
	//读取header
	headBuf := make([]byte,len(MsgHeader))
	if _,err := io.ReadFull(r,headBuf);err != nil {
		return nil,err
	}
	//校验header
	if string(headBuf) != MsgHeader {
		return nil,errors.New("header error")
	}

	//读取content len
	lenBuf := make([]byte,4)
	if _,err := io.ReadFull(r,lenBuf);err != nil {
		return nil,err
	}

	//大端字节序，解码
	length := binary.BigEndian.Uint32(lenBuf)

	//读取content
	bodyBuf := make([]byte,length)
	if _,err := io.ReadFull(r,bodyBuf);err != nil {
		return nil,err
	}

	return bodyBuf,nil
}