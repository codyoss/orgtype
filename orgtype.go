package orgtype

import "github.com/codyoss/newtype/pb"

var AVar = pb.AVar

const AConst = pb.AConst

func AFunc() {
	pb.AFunc()
}

func AnotherFunc(a *AType) *AnotherType {
	return pb.AnotherFunc(a)
}

type AType = pb.AType

type AnotherType = pb.AnotherType

type AInterface = pb.AInterface
