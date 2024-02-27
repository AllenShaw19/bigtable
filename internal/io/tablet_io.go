package io

import (
	pb "github.com/allen-shaw/bigtable/internal/proto"
	"github.com/cockroachdb/pebble"
)

type TabletStatus pb.StatusCode

type TabletIO struct {
	db         *pebble.DB
	ldbOptions pebble.Options
}
