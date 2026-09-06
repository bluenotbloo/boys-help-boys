package nacos

import (
	"fmt"

	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/v2/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
)

type Client struct {
	namingClient naming_client.INamingClient // 服务发现客户端
	configClient config_client.IConfigClient // 配置客户端
}

func NewClient(addr string, port uint64, namespace string) (*Client, error) {
	cc := constant.ClientConfig{
		NamespaceId:         namespace,          // nacos命名空间ID，public为默认命名空间
		NotLoadCacheAtStart: true,               // 在启动的时候不读取缓存在CacheDir的service信息
		TimeoutMs:           5000,               // 超时时间，单位毫秒
		LogDir:              "/tmp/nacos/log",   // 日志目录
		CacheDir:            "/tmp/nacos/cache", // 缓存目录
		LogLevel:            "info",             // 日志级别
	}
	sc := []constant.ServerConfig{{IpAddr: addr, Port: port}}
	param := vo.NacosClientParam{ClientConfig: &cc, ServerConfigs: sc}

	namingClient, err := clients.NewNamingClient(param)
	if err != nil {
		return nil, fmt.Errorf("create naming client: %w", err)
	}
	configClient, err := clients.NewConfigClient(param)
	if err != nil {
		return nil, fmt.Errorf("create config client: %w", err)
	}
	return &Client{namingClient: namingClient, configClient: configClient}, nil
}

func (c *Client) GetConfig(dataID string, groupName string) (string, error) {
	content, err := c.configClient.GetConfig(vo.ConfigParam{
		DataId: dataID,
		Group:  groupName,
	})
	if err != nil {
		fmt.Println("get config error:", err)
		return "", err
	}
	return content, nil
}
