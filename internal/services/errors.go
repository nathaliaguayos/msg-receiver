package services

type ServiceError string

func (e ServiceError) Error() string {
	return string(e)
}

const ()
