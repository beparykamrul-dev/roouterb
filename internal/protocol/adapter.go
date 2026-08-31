package protocol

import "context"

type State string
const ( StateUnknown State = "unknown"; StateStopped State = "stopped"; StateRunning State = "running" )

type Adapter interface {
 Name() string
 Install(ctx context.Context) error
 Start(ctx context.Context) error
 Stop(ctx context.Context) error
 Status(ctx context.Context) (State, error)
 GenerateConfig(ctx context.Context, profile map[string]any) ([]byte, error)
 ApplyConfig(ctx context.Context, config []byte) error
 Remove(ctx context.Context) error
}
