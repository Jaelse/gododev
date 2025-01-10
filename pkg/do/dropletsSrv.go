package do

import (
	"context"
	"fmt"

	"github.com/digitalocean/godo"
)

type DropletSrv interface {
	Get(id int, ctx context.Context) (*godo.Droplet,error)
	GetAll(ctx context.Context) ([]godo.Droplet, error)
	Create(ctx context.Context) (*godo.Droplet, error)
	TakeSnapshop(id int, ctx context.Context) error
	GetSnapshotByDropletID(id int, ctx context.Context) []string
	ListSnapshotByDropeltId(id int, ctx context.Context) ([]godo.Snapshot, error)
	CreateDropletFromSnapshotID(id int, ctx context.Context) (*godo.Droplet, error)
}

func (dc DoClient) Get(id int, ctx context.Context) (*godo.Droplet,error){
	droplet, _, err := dc.client.Droplets.Get(ctx,id)

	if err != nil {
		return nil, err
	}

	return droplet, nil
}

func (dc DoClient) GetAll(ctx context.Context) ([]godo.Droplet, error) {
	list := []godo.Droplet{}

	opt := &godo.ListOptions{}

	for {
		droplets, resp, err := dc.client.Droplets.List(ctx, opt)
		if err != nil {
			return nil, err
		}

		// append the current page's droplets to our list
		list = append(list, droplets...)

		// if we are at the last page, break out the for loop
		if resp.Links == nil || resp.Links.IsLastPage() {
			break
		}

		page, err := resp.Links.CurrentPage()
		if err != nil {
			return nil, err
		}

		// set the page we want for the next request
		opt.Page = page + 1
	}

	return list, nil
}

func (dc DoClient) Create(ctx context.Context) (*godo.Droplet, error) {
	req := godo.DropletCreateRequest{
		Name:   "test",
		Region: "nyc3",
		Size:   "s-1vcpu-1gb",
		Image: godo.DropletCreateImage{
			Slug: "ubuntu-20-04-x64",
		},
	}

	drop, _, err := dc.client.Droplets.Create(ctx, &req)
	if err != nil {
		return nil, err
	}

	return drop, nil
}

func (dc DoClient) TakeSnapshop(id int, ctx context.Context) error {
	action, _, err := dc.client.DropletActions.Snapshot(ctx, id, "okokok")
	if err != nil {
		return err
	}

	fmt.Printf("Actions %s suceeded", action)
	return nil
}

func (dc DoClient) GetSnapshotByDropletID(id int, ctx context.Context) []string {
	return []string{}
}

func (dc DoClient) ListSnapshotByDropeltId(id int, ctx context.Context) ([]godo.Snapshot, error) {
	return nil, nil
}

func (dc DoClient) CreateDropletFromSnapshotID(id int, ctx context.Context) (*godo.Droplet, error) {
	_, _, err := dc.client.Snapshots.Get(ctx, "asdf")
	if err != nil {
		return nil, err
	}

	return nil, nil
}
