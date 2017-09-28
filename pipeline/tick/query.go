package tick

import (
	"github.com/influxdata/kapacitor/pipeline"
	"github.com/influxdata/kapacitor/tick/ast"
)

// Query converts the Query pipeline node into the TICKScript AST
type Query struct {
	Function
}

// NewQuery creates a Query function builder
func NewQuery(parents []ast.Node) *Query {
	return &Query{
		Function{
			Parents: parents,
		},
	}
}

// Build creates a Query ast.Node
func (n *Query) Build(q *pipeline.QueryNode) (ast.Node, error) {
	n.Pipe("query", q.QueryStr).
		Dot("period", q.Period).
		Dot("every", q.Every).
		DotIf("align", q.AlignFlag).
		Dot("cron", q.Cron).
		Dot("offset", q.Offset).
		DotIf("alignGroup", q.AlignGroupFlag).
		Dot("groupBy", q.Dimensions).
		DotIf("groupByMeasurement", q.GroupByMeasurementFlag).
		DotNotNil("fill", q.Fill).
		Dot("cluster", q.Cluster)

	return n.prev, n.err
}