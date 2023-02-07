package db

import (
	"database/sql"

	"github.com/Jay-T/go-devops-advanced-diploma/internal/pb"
)

func SQLInt64(i int64) sql.NullInt64 {
	return sql.NullInt64{
		Int64: i,
		Valid: true,
	}
}

func ConvertToPBMetadata(metadata []Metadatum) []*pb.Metadata {
	pbMetadata := []*pb.Metadata{}
	for _, md := range metadata {
		pbMetadata = append(pbMetadata, &pb.Metadata{
			Key:   md.Key,
			Value: md.Value,
		})
	}
	return pbMetadata
}
