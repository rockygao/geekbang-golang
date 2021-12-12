package main

import (
	"context"
	//"database/sql"
	"fmt"
	"log"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/examples/fs/ent"
	"go.opentelemetry.io/otel"

	// init mysql driver
	_ "github.com/go-sql-driver/mysql"
)

// Data .
type Data struct {
	db *ent.Client
	//rdb *redis.Client
}

var (
	confDatabaseDriver = "mysql"
	confDatabaseSource = "root:root@tcp(127.0.0.1:3306)/golang?parseTime=True"
)

// NewData .
func NewData() (*Data, func(), error) {
	//log := log.NewHelper(logger)
	drv, err := sql.Open(
		confDatabaseDriver,
		confDatabaseSource,
	)
	sqlDrv := dialect.DebugWithContext(drv, func(ctx context.Context, i ...interface{}) {
		//log.WithContext(ctx).Info(i...)
		tracer := otel.Tracer("ent.")
		kind := trace.SpanKindServer
		_, span := tracer.Start(ctx,
			"Query",
			trace.WithAttributes(
				attribute.String("sql", fmt.Sprint(i...)),
			),
			trace.WithSpanKind(kind),
		)
		span.End()
	})
	client := ent.NewClient(ent.Driver(sqlDrv))
	if err != nil {
		//log.Errorf("failed opening connection to sqlite: %v", err)
		fmt.Printf("failed opening connection to sqlite: %v", err)
		return nil, nil, err
	}
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		//log.Errorf("failed creating schema resources: %v", err)
		fmt.Printf("failed creating schema resources: %v", err)
		return nil, nil, err
	}
	//rdb.AddHook(redisotel.TracingHook{})
	d := &Data{
		db: client,
		//rdb: rdb,
	}
	return d, func() {
		//log.Info("message", "closing the data resources")
		fmt.Println("message", "closing the data resources")
		if err := d.db.Close(); err != nil {
			//log.Error(err)
			log.Fatal(err)
		}
		//if err := d.rdb.Close(); err != nil {
		//	//log.Error(err)
		//	log.Fatal(err)
		//}
	}, nil
}

type DB struct {
	DB *Data
}

func main() {

	_, _, err := NewData() // 调用输出化数据库的函数
	if err != nil {
		fmt.Printf("init db failed,err:%v\n", err)
		return
	}

	var DB Data
	DB.db.File.Query().All()
	fmt.Println(DB)
}
