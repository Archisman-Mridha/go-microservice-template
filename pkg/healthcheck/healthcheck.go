package healthcheck

import "errors"

type (
	HealthcheckFn = func() error

	Healthcheckable interface {
		Healthcheck() error
	}
)

func Healthcheck(healthcheckables []Healthcheckable) (joinedErrors error) {
	for _, healthcheckable := range healthcheckables {
		if err := healthcheckable.Healthcheck(); err != nil {
			errors.Join(joinedErrors, err)
		}
	}
	return
}
