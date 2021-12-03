package main

import "DataStruct/dataStruct/Tree/RBTree"

//比大小
type thanInt int

func (x thanInt) Less(than RBTree.Item) bool {
	return x < than.(thanInt)
}

type UInt32 uint32

func (x UInt32) Less(than RBTree.Item) bool {
	return x < than.(UInt32)
}

type String string

func (x String) Less(than RBTree.Item) bool {
	return x < than.(String)
}
