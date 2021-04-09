package main

import (
	"reflect"
	"testing"
)

func TestReverseListNode(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReverseListNode(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReverseListNode() = %v, want %v", got, tt.want)
			}
		})
	}
}