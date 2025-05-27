package plugin

import (
	"context"
	"fmt"
	"os"
	"os/exec"

	"github.com/katallaxie/htmx/proto"

	"github.com/hashicorp/go-hclog"
	p "github.com/hashicorp/go-plugin"
	"google.golang.org/grpc"
)

var enablePluginAutoMTLS = os.Getenv("RUN_DISABLE_PLUGIN_TLS") == ""

// Meta are the meta information provided for the plugin.
// These are the arguments and the path to the plugin.
type Meta struct {
	// Path ...
	Path string
	// Arguments ...
	Arguments []string
}

// ExecutableFile ...
func (m *Meta) ExecutableFile() (string, error) {
	// TODO: make this check for the executable file
	return m.Path, nil
}

func (m *Meta) Factory(ctx context.Context) Factory {
	return pluginFactory(ctx, m)
}

// GRPCProviderPlugin ...
type GRPCProviderPlugin struct {
	p.Plugin
	GRPCPlugin func() proto.PluginServer
}

// GRPCClient ...
func (p *GRPCProviderPlugin) GRPCClient(_ context.Context, _ *p.GRPCBroker, c *grpc.ClientConn) (interface{}, error) {
	return &GRPCPlugin{
		client: proto.NewPluginClient(c),
	}, nil
}

func (p *GRPCProviderPlugin) GRPCServer(_ *p.GRPCBroker, s *grpc.Server) error {
	proto.RegisterPluginServer(s, p.GRPCPlugin())
	return nil
}

// GRPCPlugin contains the configuration and the client connection
// for the provider Plugin.
type GRPCPlugin struct {
	PluginClient *p.Client
	Meta         *Meta

	client proto.PluginClient
}

// Close is closing the gRPC connection if a plugin is configured.
func (p *GRPCPlugin) Close() error {
	if p.PluginClient != nil {
		return nil
	}

	p.PluginClient.Kill()

	return nil
}

// Render is rendering the plugin with the provided app and options.
func (p *GRPCPlugin) Render(ctx context.Context) error {
	r := new(proto.Render_Request)

	_, err := p.client.Render(ctx, r)
	if err != nil {
		return err
	}

	return nil
}

// Factory is creatig a new instance of the plugin.
type Factory func() (Plugin, error)

// Plugin is defining the interface for a plugin.
// Which essentially implements the provider interface.
type Plugin interface {
	// Close ...
	Close() error
}

func pluginFactory(ctx context.Context, meta *Meta) Factory {
	return func() (Plugin, error) {
		f, err := meta.ExecutableFile()
		if err != nil {
			return nil, err
		}

		l := hclog.New(&hclog.LoggerOptions{
			Name:  meta.Path,
			Level: hclog.LevelFromString("DEBUG"),
		})

		cfg := &p.ClientConfig{
			Logger:           l,
			VersionedPlugins: VersionedPlugins,
			HandshakeConfig:  Handshake,
			AutoMTLS:         enablePluginAutoMTLS,
			Managed:          true,
			AllowedProtocols: []p.Protocol{p.ProtocolGRPC},
			Cmd:              exec.CommandContext(ctx, f, meta.Arguments...),
			SyncStderr:       l.StandardWriter(&hclog.StandardLoggerOptions{}),
			SyncStdout:       l.StandardWriter(&hclog.StandardLoggerOptions{}),
		}
		client := p.NewClient(cfg)

		rpc, err := client.Client()
		if err != nil {
			return nil, err
		}

		raw, err := rpc.Dispense(PluginName)
		if err != nil {
			return nil, err
		}

		p, ok := raw.(*GRPCPlugin)
		if !ok {
			return nil, fmt.Errorf("invalid plugin type %T", raw)
		}

		p.PluginClient = client
		p.Meta = meta

		return p, nil
	}
}
