package repository

import (
	"go.t1ltxz.ninja/disgo-template/internal/repository/users"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"repository",
	users.Module,
)
