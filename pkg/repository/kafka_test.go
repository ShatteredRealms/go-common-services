package repository_test

import (
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/segmentio/kafka-go"

	"github.com/ShatteredRealms/go-common-service/pkg/config"
	"github.com/ShatteredRealms/go-common-service/pkg/repository"
	"github.com/ShatteredRealms/go-common-service/pkg/testsro"
)

var _ = Describe("Kafka repository", func() {
	Describe("ConnectKafka", func() {
		Context("valid input", func() {
			It("should work", func() {
				cleanupFunc, port, err := testsro.SetupKafkaWithDocker()
				Expect(err).NotTo(HaveOccurred())
				var kafka *kafka.Conn
				defer cleanupFunc()
				Eventually(func(g Gomega) error {
					kafka, err = repository.ConnectKafka(config.ServerAddress{
						Port: port,
						Host: "127.0.0.1",
					})
					return err
				}).Within(time.Second * 15).Should(Succeed())
				Expect(kafka).NotTo(BeNil())
			})
		})
	})
})
