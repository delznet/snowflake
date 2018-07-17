package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"os"

	pb "github.com/delznet/snowflake/proto"

	"github.com/bwmarrin/snowflake"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// appConf 程序配置参数
type appConf struct {
	Port          string `json:"port"`            //tcp端口号
	Epoch         int64  `json:"epoch"`           //开始时间截
	MachineIDBits uint8  `json:"machine_id_bits"` //机器号位长度
	ServiceIDBits uint8  `json:"service_id_bits"` //业务位长度
	StepBits      uint8  `json:"step_bits"`       //序列位长度
	MachineID     int64  `json:"machine_id"`      //机器id
	ServiceID     int64  `json:"service_id"`      //业务id
}

var (
	nodeMap = make(map[int64]*snowflake.Node)
	conf    appConf
)

type idServer struct {
}

func init() {
	conf = newAppConf()
}

func newAppConf() appConf {
	return appConf{
		Port:          "",
		Epoch:         1514736000000, //默认设置为2018年1月1日 00:00:00
		MachineIDBits: 4,
		ServiceIDBits: 6,
		StepBits:      12,
		MachineID:     0,
		ServiceID:     0,
	}
}

// 根据machineID和serviceID获取Node
func getNode(machineID int64, serviceID int64) (*snowflake.Node, error) {
	nodeID := int64(machineID<<conf.ServiceIDBits | serviceID)
	if Node, ok := nodeMap[nodeID]; ok {
		return Node, nil
	}
	Node, err := snowflake.NewNode(nodeID)
	if err == nil {
		nodeMap[nodeID] = Node
	}
	return Node, err
}

func (s *idServer) Gen(ctx context.Context, in *pb.SnowflakeRequest) (*pb.SnowflakeReply, error) {
	Node, err := getNode(conf.MachineID, conf.ServiceID)
	if err != nil {
		return nil, err
	}
	reply := new(pb.SnowflakeReply)
	switch in.Format {
	case "int64":
		reply.Id = Node.Generate().String()
		break
	case "base2":
		reply.Id = Node.Generate().Base2()
		break
	case "base36":
		reply.Id = Node.Generate().Base36()
		break
	case "base32":
		reply.Id = Node.Generate().Base32()
		break
	case "base58":
		reply.Id = Node.Generate().Base58()
		break
	case "base64":
		reply.Id = Node.Generate().Base64()
		break
	default:
		reply.Id = Node.Generate().String()
		break
	}

	return reply, nil
}

func main() {
	var confData []byte
	var e error
	if confData, e = ioutil.ReadFile("app.conf"); e != nil {
		checkErr(e)
	}

	if e = json.Unmarshal(confData, &conf); e != nil {
		checkErr(e)
	}

	lis, err := net.Listen("tcp", conf.Port)
	if err != nil {
		checkErr(err)
	}

	s := grpc.NewServer()
	pb.RegisterSnowflakeServer(s, &idServer{})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		checkErr(err)
	}
}

func checkErr(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
