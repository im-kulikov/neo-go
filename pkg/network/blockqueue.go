package network

import (
	"github.com/CityOfZion/neo-go/pkg/core"
	"github.com/Workiva/go-datastructures/queue"
	log "github.com/sirupsen/logrus"
)

type blockQueue struct {
	queue       *queue.PriorityQueue
	checkBlocks chan struct{}
	chain       core.Blockchainer
}

func newBlockQueue(capacity int, bc core.Blockchainer) *blockQueue {
	return &blockQueue{
		queue:       queue.NewPriorityQueue(capacity, false),
		checkBlocks: make(chan struct{}, 1),
		chain:       bc,
	}
}

func (bq *blockQueue) run() {
	for {
		_, ok := <-bq.checkBlocks
		if !ok {
			break
		}
		for {
			item := bq.queue.Peek()
			if item == nil {
				break
			}
			minblock := item.(*core.Block)
			if minblock.Index <= bq.chain.BlockHeight()+1 {
				_, _ = bq.queue.Get(1)
				updateBlockQueueLenMetric(bq.length())
				if minblock.Index == bq.chain.BlockHeight()+1 {
					err := bq.chain.AddBlock(minblock)
					if err != nil {
						log.WithFields(log.Fields{
							"error":       err.Error(),
							"blockHeight": bq.chain.BlockHeight(),
							"nextIndex":   minblock.Index,
						}).Warn("blockQueue: failed adding block into the blockchain")
					}
				}
			} else {
				break
			}
		}
	}
}

func (bq *blockQueue) putBlock(block *core.Block) error {
	if bq.chain.BlockHeight() >= block.Index {
		// can easily happen when fetching the same blocks from
		// different peers, thus not considered as error
		return nil
	}
	err := bq.queue.Put(block)
	// update metrics
	updateBlockQueueLenMetric(bq.length())
	select {
	case bq.checkBlocks <- struct{}{}:
		// ok, signalled to goroutine processing queue
	default:
		// it's already busy processing blocks
	}
	return err
}

func (bq *blockQueue) discard() {
	close(bq.checkBlocks)
	bq.queue.Dispose()
}

func (bq *blockQueue) length() int {
	return bq.queue.Len()
}
