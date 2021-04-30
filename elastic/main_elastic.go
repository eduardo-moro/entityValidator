package elastic

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/olivere/elastic"

	"context"
	"os"
	"time"
)

type Elastic struct {
	Client *elastic.Client
}

func NewClient (host, schema string) (*Elastic, error) {

	client, err := elastic.NewClient(
		elastic.SetURL(host),
		elastic.SetBasicAuth(os.Getenv("AWS_ACCESS_KEY"),os.Getenv("AWS_SECRET_KEY")),
		elastic.SetSniff(false),
	)
	return &Elastic{client}, err
}

func (e *Elastic) PushRecords(data []events.KinesisEventRecord) error {
	processor, err := e.Client.BulkProcessor().
		Name("kinesisWorkers").
		Workers(2).
		BulkSize(2 << 20). // set the max bulk size to 2MB
		FlushInterval(10 * time.Second).
		Do(context.Background())
	if err != nil {
		return err
	}
	defer processor.Close()

	for _, x := range data {
		rec := elastic.NewBulkIndexRequest().Index("kinesis-apm").Type("doc").Doc(x.Kinesis.Data)
		processor.Add(rec)
	}

	return processor.Flush()
}