package db

import (
	mc "github.com/flywithbug/cache2go"
)



var table *mc.CacheTable

func MemCacheConfig()*mc.CacheTable   {
	if table== nil {
		InitCache2go()
	}
	return table
}

func InitCache2go() {
	table = mc.Cache("darkside")
}

//func Add(key interface{},value interface{})  {
//	table.Add(key,0,&value)
//}
//
//func Value(key interface{})(interface{},error)  {
//	result, err := table.Value(key)
//	if err != nil{
//		fmt.Println(err)
//	}
//	return result.Data(),err
//}


