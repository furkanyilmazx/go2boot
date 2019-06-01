package system

import "reflect"

func ObjIsEmpty(obj, ref interface{}) bool {
	return reflect.DeepEqual(ref, obj)
}
