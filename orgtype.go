package orgtype

import "github.com/codyoss/newtype/newtypepb"

var AVar = newtypepb.AVar

const AConst = newtypepb.AConst

func AFunc() { newtypepb.AFunc() }

func AnotherFunc(t *AType) *AnotherType {
	return newtypepb.AnotherFunc(t)
}

type AType = newtypepb.AType

type AnotherType = newtypepb.AnotherType

type AInterface = newtypepb.AInterface
