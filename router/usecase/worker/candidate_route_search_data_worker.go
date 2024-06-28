package worker

import (
	"context"
	"fmt"
	"sync"

	"github.com/osmosis-labs/sqs/domain"
	"github.com/osmosis-labs/sqs/domain/mvc"
	"github.com/osmosis-labs/sqs/log"
	"github.com/osmosis-labs/sqs/sqsdomain"

	routerusecase "github.com/osmosis-labs/sqs/router/usecase"
)

type candidateRouteSearchDataWorker struct {
	listeners                 []domain.CandidateRouteSearchDataUpdateListener
	poolsHandler              mvc.PoolHandler
	candidateRouteDataHolders []mvc.CandidateRouteSearchDataHolder
	preferredPoolIDs          []uint64
	cosmWasmPoolConfig        domain.CosmWasmPoolRouterConfig
	hasReceivedFirstBlock     bool
	logger                    log.Logger
}

var (
	_ domain.CandidateRouteSearchDataWorker = &candidateRouteSearchDataWorker{}
)

func NewCandidateRouteSearchDataWorker(poolHandler mvc.PoolHandler, candidateRouteDataHolders []mvc.CandidateRouteSearchDataHolder, preferredPoolIDs []uint64, cosmWasmPoolConfig domain.CosmWasmPoolRouterConfig, logger log.Logger) *candidateRouteSearchDataWorker {
	return &candidateRouteSearchDataWorker{
		listeners:                 []domain.CandidateRouteSearchDataUpdateListener{},
		poolsHandler:              poolHandler,
		candidateRouteDataHolders: candidateRouteDataHolders,
		preferredPoolIDs:          preferredPoolIDs,
		cosmWasmPoolConfig:        cosmWasmPoolConfig,
		logger:                    logger,
	}
}

// ComputeSearchDataFirstBlockSync implements domain.CandidateRouteSearchDataWorker.
func (c *candidateRouteSearchDataWorker) ComputeSearchDataFirstBlockSync(ctx context.Context, height uint64, blockPoolMetaData domain.BlockPoolMetadata) error {
	// If first block, we need to have the candidate route data to be able to compute prices.
	// As a result, we syncronously block.
	if !c.hasReceivedFirstBlock {
		if err := c.compute(ctx, height, blockPoolMetaData); err != nil {
			return err
		}

		c.hasReceivedFirstBlock = true
	} else {
		// Otherwise, we can process the candidate route data asynchronously.
		go func() {
			_ = c.compute(ctx, height, blockPoolMetaData)
		}()
	}

	// Notify listeners
	for _, listener := range c.listeners {
		_ = listener.OnSearchDataUpdate(ctx, height)
	}

	return nil
}

func (c *candidateRouteSearchDataWorker) compute(ctx context.Context, height uint64, blockPoolMetaData domain.BlockPoolMetadata) error {
	mu := sync.Mutex{}
	candidateRouteData := make(map[string][]sqsdomain.PoolI, len(blockPoolMetaData.UpdatedDenoms))

	wg := sync.WaitGroup{}

	for denom := range blockPoolMetaData.UpdatedDenoms {
		wg.Add(1)

		go func(denom string) {
			defer wg.Done()

			denomLiquidityData, ok := blockPoolMetaData.DenomPoolLiquidityMap[denom]
			if !ok {
				// TODO: consider error
				return
			}

			denomPoolsIDs := domain.KeysFromMap(denomLiquidityData.Pools)

			unsortedDenomPools, err := c.poolsHandler.GetPools(denomPoolsIDs)
			if err != nil {
				// TODO: handle error
				return
			}

			// Sort pools
			sortedDenomPools := routerusecase.ValidateAndSortPools(unsortedDenomPools, c.cosmWasmPoolConfig, c.preferredPoolIDs, c.logger)

			mu.Lock()
			candidateRouteData[denom] = sortedDenomPools
			mu.Unlock()
		}(denom)
	}

	wg.Wait()

	if len(candidateRouteData) > 100 {
		fmt.Println("here")
	}

	// Update candidate route data holders
	for _, candidateRouteDataHolder := range c.candidateRouteDataHolders {
		wg.Add(1)

		go func(candidateRouteDataHolder mvc.CandidateRouteSearchDataHolder) {
			defer wg.Done()

			candidateRouteDataHolder.SetCandidateRouteSearchData(candidateRouteData)
		}(candidateRouteDataHolder)
	}

	wg.Wait()

	return nil
}

// RegisterListener implements domain.CandidateRouteSearchDataWorker.
func (c *candidateRouteSearchDataWorker) RegisterListener(listener domain.CandidateRouteSearchDataUpdateListener) {
	c.listeners = append(c.listeners, listener)
}