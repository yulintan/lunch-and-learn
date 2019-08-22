package shops

import "context"

func (s *shopService) ShopSync(ctx context.Context, req *rpc.ShopSyncRequest) (*rpc.ShopSyncResponse, error) {
	exist, err := s.shopRepo.GetByIdentifier(ctx, req.ShopIdentifier) // HLxxx
	if err != nil {
		clog.FromContext(ctx).Error("error", "cannot find shop in database with err", "shopIdentifier", req.ShopIdentifier, "err", err)
		return nil, err
	}

	if req.Syncing {
		exist.SetSyncStatus(SyncStatus(req.Type)) // HLxxx
	} else {
		exist.UnsetSyncStatus(SyncStatus(req.Type)) // HLxxx

	}

	_, err = s.shopRepo.Update(ctx, exist) // HLxxx
	if err != nil {
		clog.FromContext(ctx).Error("error", "cannot update shop in database with err", "shopIdentifier", req.ShopIdentifier, "err", err)
		return nil, err
	}

	return &rpc.ShopSyncResponse{Succeed: true, TotalBitmask: uint32(exist.SyncStatus)}, nil
}
