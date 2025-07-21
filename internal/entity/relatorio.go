package entity

import (
	"fmt"
	"time"
)

type Relatorio struct {
	TimeInicial time.Time
	TempoDecorrido time.Duration
	TempoTotalExecucaoRequest time.Duration
	QuantidadeTotalRequestsRealizadas int
	QuantidadeTotalRequestStatus200 int
	QuantidadeRequestsErro map[int]int	
}

func NewRelatorio() *Relatorio{

	return &Relatorio{
		TimeInicial: time.Now(),
		QuantidadeRequestsErro: make(map[int]int),
	}
}

func (r *Relatorio) AdicionaResposta(resposta Resposta){

	r.QuantidadeTotalRequestsRealizadas = r.QuantidadeTotalRequestsRealizadas + 1

	if (resposta.StatusCode == 200){
		r.QuantidadeTotalRequestStatus200 = r.QuantidadeTotalRequestStatus200 + 1
	} else {
		qt:= r.QuantidadeRequestsErro[resposta.StatusCode]
		r.QuantidadeRequestsErro[resposta.StatusCode] = qt + 1
	}

	r.TempoTotalExecucaoRequest = r.TempoTotalExecucaoRequest + resposta.TempoRequest
}

func (r *Relatorio) Finaliza(){

	r.TempoDecorrido = time.Since(r.TimeInicial)
}

func (r *Relatorio) ImprimirRelatorioConsole(){
	fmt.Printf(
		"Relatório de execução Stress-Test\n" +
		"Tempo decorrido = %f\n"+
		"Tempo Total executando requests = %f(.sec) \n"+
		"Total de requests realizadas = %d \n" +
		"Total de requests status 200 = %d \n"+
		"Erros de execução:\n",
		r.TempoDecorrido.Seconds(),
		r.TempoTotalExecucaoRequest.Seconds(),
		r.QuantidadeTotalRequestsRealizadas,
		r.QuantidadeTotalRequestStatus200,
	)

	for key, value := range r.QuantidadeRequestsErro {
		fmt.Printf("-Status: %d, Total: %d\n", key, value)
	}
}
