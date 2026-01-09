package controller

type CustomerController interface {
	GetByCpf(cpf string) (*dto.GetCustomerResponseDto, error)
	Add(customer *dto.AddCustomerRequestDto) error
}
