package mongodb

import (
	"gonews/pkg/storage"
	"reflect"
	"testing"

	"go.mongodb.org/mongo-driver/mongo"
)

func TestStorage_Posts(t *testing.T) {
	type fields struct {
		Client *mongo.Client
		DB     *mongo.Database
	}
	tests := []struct {
		name    string
		fields  fields
		want    []storage.Post
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Storage{
				Client: tt.fields.Client,
				DB:     tt.fields.DB,
			}
			got, err := s.Posts()
			if (err != nil) != tt.wantErr {
				t.Errorf("Storage.Posts() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Storage.Posts() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStorage_AddPost(t *testing.T) {
	type fields struct {
		Client *mongo.Client
		DB     *mongo.Database
	}
	type args struct {
		p storage.Post
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Storage{
				Client: tt.fields.Client,
				DB:     tt.fields.DB,
			}
			if err := s.AddPost(tt.args.p); (err != nil) != tt.wantErr {
				t.Errorf("Storage.AddPost() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestStorage_UpdatePost(t *testing.T) {
	type fields struct {
		Client *mongo.Client
		DB     *mongo.Database
	}
	type args struct {
		p storage.Post
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Storage{
				Client: tt.fields.Client,
				DB:     tt.fields.DB,
			}
			if err := s.UpdatePost(tt.args.p); (err != nil) != tt.wantErr {
				t.Errorf("Storage.UpdatePost() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestStorage_DeletePost(t *testing.T) {
	type fields struct {
		Client *mongo.Client
		DB     *mongo.Database
	}
	type args struct {
		p storage.Post
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Storage{
				Client: tt.fields.Client,
				DB:     tt.fields.DB,
			}
			if err := s.DeletePost(tt.args.p); (err != nil) != tt.wantErr {
				t.Errorf("Storage.DeletePost() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
