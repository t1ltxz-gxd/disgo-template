package components

import (
	"go.uber.org/fx"
)

var Module = fx.Module(
	"components",
	fx.Invoke(NewRegistry),
	//buttons.Module,
	//modals.Module,
	//selects.Module,
)
