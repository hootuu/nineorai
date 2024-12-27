package trigger

import "github.com/hootuu/nineorai/io"

type Service interface {
	Trigger(req *io.Request[Trigger]) *io.Response[Result]
}
