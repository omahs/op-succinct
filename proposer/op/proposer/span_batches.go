package proposer

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/succinctlabs/op-succinct-go/proposer/db/ent"
)

type Span struct {
	Start uint64
	End   uint64
}

func (l *L2OutputSubmitter) CreateSpans(start, end uint64) []Span {
	spans := []Span{}
	// Create spans of size MaxBlockRangePerSpanProof from start to end.
	// Each span starts where the previous one ended + 1.
	// Continue until we can't fit another full span before reaching end.
	for i := start; i+l.Cfg.MaxBlockRangePerSpanProof <= end; i += l.Cfg.MaxBlockRangePerSpanProof + 1 {
		spans = append(spans, Span{Start: i, End: i + l.Cfg.MaxBlockRangePerSpanProof})
	}
	return spans
}

func (l *L2OutputSubmitter) DeriveNewSpanBatches(ctx context.Context) error {
	// nextBlock is equal to the highest value in the `EndBlock` column of the DB, plus 1.
	latestL2EndBlock, err := l.db.GetLatestEndBlock()
	if err != nil {
		if ent.IsNotFound(err) {
			latestEndBlockU256, err := l.l2ooContract.LatestBlockNumber(&bind.CallOpts{Context: ctx})
			if err != nil {
				return fmt.Errorf("failed to get latest output index: %w", err)
			} else {
				latestL2EndBlock = latestEndBlockU256.Uint64()
			}
		} else {
			l.Log.Error("failed to get latest end requested", "err", err)
			return err
		}
	}
	newL2StartBlock := latestL2EndBlock + 1

	rollupClient, err := l.RollupProvider.RollupClient(ctx)
	if err != nil {
		return fmt.Errorf("failed to get rollup client: %w", err)
	}

	// Get the latest finalized L2 block.
	status, err := rollupClient.SyncStatus(ctx)
	if err != nil {
		l.Log.Error("proposer unable to get sync status", "err", err)
		return err
	}
	// Note: Originally, this used the L1 finalized block. However, to satisfy the new API, we now use the L2 finalized block.
	newL2EndBlock := status.FinalizedL2.Number

	// Create spans of size MaxBlockRangePerSpanProof from newL2StartBlock to newL2EndBlock.
	spans := l.CreateSpans(newL2StartBlock, newL2EndBlock)
	// Add each span to the DB. If there are no spans, we will not create any proofs.
	for _, span := range spans {
		err := l.db.NewEntry("SPAN", span.Start, span.End)
		l.Log.Info("New range proof request.", "start", span.Start, "end", span.End)
		if err != nil {
			l.Log.Error("failed to add span to db", "err", err)
			return err
		}
	}

	return nil
}
