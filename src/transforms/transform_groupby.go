// Copyright 2020 The VectorSQL Authors.
//
// Code is licensed under Apache License, Version 2.0.

package transforms

import (
	"datablocks"
	"planners"
	"processors"
	"sync"
)

type GroupByTransform struct {
	ctx  *TransformContext
	plan *planners.GroupByPlan
	processors.BaseProcessor
}

func NewGroupByTransform(ctx *TransformContext, plan *planners.GroupByPlan) processors.IProcessor {
	return &GroupByTransform{
		ctx:           ctx,
		plan:          plan,
		BaseProcessor: processors.NewBaseProcessor("transform_groupby"),
	}
}

func (t *GroupByTransform) Execute() {
	var wg sync.WaitGroup
	var block *datablocks.DataBlock

	out := t.Out()
	defer out.Close()

	onNext := func(x interface{}) {
		switch y := x.(type) {
		case *datablocks.DataBlock:
			if block == nil {
				block = y
			} else {
				if err := block.Append(y); err != nil {
					out.Send(err)
				}
			}
		case error:
			out.Send(y)
		}
	}
	onDone := func() {
		defer wg.Done()
		if block != nil {
			if blocks, err := block.GroupByPlan(t.plan); err != nil {
				out.Send(err)
			} else {
				for _, x := range blocks {
					out.Send(x)
				}
			}
		}
	}
	wg.Add(1)
	t.Subscribe(onNext, onDone)
	wg.Wait()
}