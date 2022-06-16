package orgtype

import "github.com/codyoss/newtype"

var AVar = newtype.AVar

const AConst = newtype.AConst

func AFunc() {
	newtype.AFunc()
}

func AnotherFunc(a *AType) *AnotherType {
	return newtype.AnotherFunc(a)
}

type AType = newtype.AType

type AnotherType = newtype.AnotherType

type AInterface = newtype.AInterface
