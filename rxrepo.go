package main

import "github.com/restream/reindexer"

type RxRepo struct {
	db *reindexer.Reindexer
}

func (r *RxRepo) Init(dsn string) error {
	r.db = reindexer.NewReindex(dsn)

	if err := r.db.OpenNamespace("offers", reindexer.DefaultNamespaceOptions(), &YMOfferRx{}); err != nil {
		return err
	}
	if err := r.db.OpenNamespace("products", reindexer.DefaultNamespaceOptions(), &YMProductRx{}); err != nil {
		return err
	}

	return nil
}

func (r *RxRepo) PutProdct(prod YMProduct, uid string) error {
	rprod := &YMProductRx{YMProduct: prod, Uid: uid}
	return r.db.Upsert("products", rprod)
}

func (r *RxRepo) PutOffer(offer YMOffer, uid string) error {
	roffer := &YMOfferRx{YMOffer: offer, Uid: uid}
	return r.db.Upsert("offers", roffer)
}

type YMOfferRx struct {
	YMOffer
	Uid string `json:"uid" reindex:"uid,,pk"`
}

type YMProductRx struct {
	YMProduct
	Uid string `json:"uid" reindex:"uid,,pk"`
}
