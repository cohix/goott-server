package service

import (
	"fmt"
	"net"

	"github.com/cohix/simplcrypto"
	"github.com/pkg/errors"

	model "github.com/cohix/goott-server/model"
	log "github.com/cohix/simplog"
	context "golang.org/x/net/context"
	"google.golang.org/grpc"
)

const goottServicePort = ":3687"

// StartGoottService starts the runner service
func StartGoottService(app *model.App, errChan chan error) {
	lis, err := net.Listen("tcp", goottServicePort)
	if err != nil {
		errChan <- err
		return
	}

	grpcServer := grpc.NewServer()

	service := &GoottService{
		App: app,
	}

	RegisterGoottServer(grpcServer, service)

	log.LogInfo(fmt.Sprintf("starting goott-server runner service on %s", goottServicePort))
	if err := grpcServer.Serve(lis); err != nil {
		errChan <- err
	}
}

// GoottService describes a goott server
type GoottService struct {
	App *model.App
}

// Auth allows a client to prove it knows the auth token in exchange for a session key
func (gs *GoottService) Auth(ctx context.Context, req *model.AuthRequest) (*model.AuthResponse, error) {
	pubkey, err := simplcrypto.KeyPairFromSerializedPubKey(req.PubKey)
	if err != nil {
		return nil, errors.Wrap(err, "failed to KeyPairFromSerializedPubKey")
	}

	log.LogInfo(fmt.Sprintf("authenticating client with pubkey %s", pubkey.KID))

	decodedToken, err := simplcrypto.Base64URLDecode(gs.App.AuthToken)
	if err != nil {
		return nil, errors.Wrap(err, "failed to Base64URLDecode")
	}

	if verified := pubkey.Verify(decodedToken, req.TokenSignature); verified != simplcrypto.SigVerified {
		return nil, errors.New("unable to Verify auth token signature")
	}

	sessionKey, err := simplcrypto.GenerateSymKey()
	if err != nil {
		return nil, errors.Wrap(err, "failed to GenerateSymKey")
	}

	gs.App.Keyset.AddSymKey(sessionKey)

	sessionKeyJSON := sessionKey.JSON()

	encSessionKey, err := pubkey.Encrypt(sessionKeyJSON)
	if err != nil {
		return nil, errors.Wrap(err, "failed to Encrypt")
	}

	resp := &model.AuthResponse{
		EncSessionKey: encSessionKey,
	}

	return resp, nil
}
