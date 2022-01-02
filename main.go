package main

import (
	"fmt"
	"github.com/2559065/cart/domain/repository"
	service2 "github.com/2559065/cart/domain/service"
	"github.com/2559065/cart/handler"
	pb "github.com/2559065/cart/proto/cart"
	"github.com/2559065/common"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/micro/go-micro/debug/log"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-plugins/registry/consul/v2"
	ratelimit "github.com/micro/go-plugins/wrapper/ratelimiter/uber/v2"
	"github.com/opentracing/opentracing-go"
)

var QPS = 100

func main() {
	// 配置中心
	consulConfig, err := common.GetConsulConfig("127.0.0.1", 8500, "/micro/config")
	if err != nil {
		log.Error(err)
	}
	// 注册中心
	consulRegistry := consul.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{
			"127.0.0.1:8500",
		}
	})

	// 链路追踪
	t, io, err := common.NewTracer("go.micro.service.cart", "localhost:6831")
	if err != nil {
		log.Error(err)
	}
	defer io.Close()
	opentracing.SetGlobalTracer(t)

	// Create service
	service := micro.NewService(
		micro.Name("go.micro.service.cart"),
		micro.Version("latest"),
		// 这里设置地址和需要暴露的端口
		micro.Address("127.0.0.1:8087"),
		// 添加consul作为注册中心
		micro.Registry(consulRegistry),
		// 添加限流
		micro.WrapHandler(ratelimit.NewHandlerWrapper(QPS)),
	)

	//只执行一次,数据表初始化
	//rp := repository.NewUserRepository(db)
	//rp.InitTable()

	// 获取mysql配置,路径中不带前缀
	mysqlInfo := common.GetMysqlFromConsul(consulConfig, "mysql")
	// 连接数据库
	db, err := gorm.Open("mysql", mysqlInfo.User + ":" + mysqlInfo.Pwd + "@/" + mysqlInfo.Database + "?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Error(err)
	}
	defer db.Close()

	// 禁止复表
	db.SingularTable(true)

	//rp := repository.NewCartRepository(db)
	//rp.InitTable()

	// 初始化服务
	service.Init()

	categoryDataService := service2.NewCartDataService(repository.NewCartRepository(db))

	// Register handler
	err = pb.RegisterCartHandler(service.Server(), &handler.Cart{categoryDataService})
	if err != nil {
		log.Error(err)
	}

	// Run service
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
