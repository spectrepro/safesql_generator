package ast

type ColumnDef struct {
	Colname   string
	TypeName  *TypeName
	IsNotNull bool
	IsArray   bool
	Vals      *List

	// From pg.ColumnDef
	Inhcount      int
	IsLocal       bool
	IsFromType    bool
	IsFromParent  bool
	Storage       byte
	RawDefault    Node
	CookedDefault Node
	Identity      byte
	CollClause    *CollateClause
	CollOid       Oid
	Constraints   *List
	Fdwoptions    *List
	Location      int
}

func (n *ColumnDef) Pos() int {
	return n.Location
}
