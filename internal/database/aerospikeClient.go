package database

import (
	as "github.com/aerospike/aerospike-client-go"
)

var AerospikeClient *as.Client

func InitAerosplike(address string, port int) (*as.Client, error) {
	var err error
	//先直接返回，方便调试
	//return AerospikeClient, err
	AerospikeClient, err = as.NewClient(address, port)
	return AerospikeClient, err
}
