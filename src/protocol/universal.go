package protocol

import (
  "bytes"
  "reflect"
  "encoding/binary"
  
  "../util"
)

//See: https://wiki.vg/Server_List_Ping#Response
// C -> S
type HandshakePacket struct {
  Version int32 `type:"varint"` //Protocol version
  Address string `type:"string"`
  Port uint16 `type:"unsigned short"`
  NextState int `type:"varint"`
}

// S -> C
type ResponsePacket struct {
  JSON string `type:"string"`
}

// C -> S
type PingPacket struct {
  ClientTime int64 `type:"long"`
}

// S -> C
type PongPacket struct {
  ClientTime int64 `type:"long"`
}

//Most Important Function, parses received packets based on struct data tags
func Parse(data []byte, v interface{}) {
  r := bytes.NewReader(data);
  
  //Get settable struct val
  structVal := reflect.ValueOf(&v).Elem();
  var tag string;
  var field reflect.Value;
  for i := 0; i < structVal.NumField(); i++{
    field = structVal.Field(i)
    tag = structVal.Type().Field(i).Tag.Get("type")
    print(field.CanSet());
    if(!field.CanSet()){ continue }
    switch tag {
    case "varint":
      val, err := util.ReadVarint(r)
      if(err != nil){panic(err)}
      field.SetInt(int64(val));
    case "string":
      val, err := util.ReadString(r)
      if(err != nil){panic(err)}
      field.SetString(val);
    case "unsigned short":
      var val uint64;
      err := binary.Read(r, binary.BigEndian, val);
      if(err != nil){panic(err)}
      field.SetUint(val);
    case "long":
      var val int64;
      err := binary.Read(r, binary.BigEndian, val);
      if(err != nil){panic(err)}
      field.SetInt(val);
    }
  }
  
}