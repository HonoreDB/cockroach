// Copyright 2018 The Cockroach Authors.
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package tree

// CreateChangefeed represents a CREATE CHANGEFEED statement.
type CreateChangefeed struct {
	Targets ChangefeedTargetList
	SinkURI Expr
	Options KVOptions
}

var _ Statement = &CreateChangefeed{}

// Format implements the NodeFormatter interface.
func (node *CreateChangefeed) Format(ctx *FmtCtx) {
	if node.SinkURI != nil {
		ctx.WriteString("CREATE ")
	} else {
		// Sinkless feeds don't really CREATE anything, so the syntax omits the
		// prefix. They're also still EXPERIMENTAL, so they get marked as such.
		ctx.WriteString("EXPERIMENTAL ")
	}
	ctx.WriteString("CHANGEFEED FOR ")
	ctx.FormatNode(&node.Targets)
	if node.SinkURI != nil {
		ctx.WriteString(" INTO ")
		ctx.FormatNode(node.SinkURI)
	}
	if node.Options != nil {
		ctx.WriteString(" WITH ")
		ctx.FormatNode(&node.Options)
	}
}

// ChangefeedTargetList represents a list of targets.
type ChangefeedTargetList struct {
	Tables  TablePatterns
	Indexes TableIndexNames
}

// Format implements the NodeFormatter interface.
func (tl *ChangefeedTargetList) Format(ctx *FmtCtx) {
	if tl.Indexes != nil {
		ctx.WriteString("INDEX ")
		ctx.FormatNode(&tl.Indexes)
	}
	if tl.Indexes != nil && tl.Tables != nil {
		ctx.WriteString(",")
	}
	if tl.Tables != nil {
		ctx.WriteString("TABLE ")
		ctx.FormatNode(&tl.Tables)
	}
}

// Merge adds the targets of tl2 to the receiver in place, and returns the receiver
func (tl1 ChangefeedTargetList) Merge(tl2 ChangefeedTargetList) ChangefeedTargetList {
	tl1.Tables = append(tl1.Tables, tl2.Tables...)
	tl1.Indexes = append(tl1.Indexes, tl2.Indexes...)
	return tl1
}
