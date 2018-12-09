package config

import (
	"time"
	"github.com/Shopify/sarama"
	cluster "github.com/bsm/sarama-cluster"
)

//变量
var BrokerLists = []string{"localhost:9092"}


//工厂函数
func NewProducerConfig() *sarama.Config {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.MaxMessageBytes = 10000000 //设置producer的最大消息大小，此大小要<= kafka的max.message.bytes的设置
	
	//producer认为消息发送成功时，需要的ack, 默认只需要leader回ack,即　WaitForLocal,也可以用１表示
	// sarama.WaitForLocal = 1	//只需要对应分区的leader确认
	// sarama.NoResponse = 0  //表示producer不关心消息是否发送成功
	// sarama.WaitForAll = -1	//需要所有的in-sync的副本确认，默认的min.insync.replicas 的数量是在broker中配置的。
	config.Producer.RequiredAcks = 1

	//设置producer接收ack的超时时间，单位只能是ms, 默认是10s，也就是10000ms
	config.Producer.Timeout = 10000 * time.Millisecond

	//设置producer重发消息的次数，默认是３次
	config.Producer.Retry.Max = 2

	return config
}


func NewConsumerConfig() *cluster.Config {
	config := cluster.NewConfig()
	//消费者配置
	//当消费出现错误时，是否将错误放进error chan，默认是true
	config.Consumer.Return.Errors = true
	//当consumerGroup中发生rebalance的时候，是否将通知发到notification chan中。默认是false
	config.Group.Return.Notifications = true
	//如果kafka中没有offset的记录，则从topic中最新的位置开始
	config.Consumer.Offsets.Initial = -1

	//consumer多久会将offset提交到kafka中，前提是需要手动的将已经消费的消息标记为 "已消费"
	//也就是要调用msg.MarkOffset或msg.MarkOffsets
	config.Consumer.Offsets.CommitInterval = time.Second

	return config
}