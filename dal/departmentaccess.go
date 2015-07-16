package dal

import (
	"github.com/qtzheng/SIMP/app/modules"
)

func DepInsert(dep *modules.Department) error {
	return NewDB().C(DepColl).Insert(dep)
}
