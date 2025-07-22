package usecase

import (
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/AndersonOdilo/fullcycle-stress-test/internal/entity"
)

func StressTeste(urlRequest string, numeroRequest int, numeroRequestParalelo int){

	relatorio := entity.NewRelatorio()
	var wgRespostas sync.WaitGroup
	chRespostas := make(chan entity.Resposta, numeroRequest)
	go processaRespostas(relatorio, chRespostas, &wgRespostas)

	var wgRequests sync.WaitGroup
	chRequests := make(chan string)
	for i := 1; i <= numeroRequestParalelo; i++ {
		go worker(chRequests, chRespostas, &wgRequests, &wgRespostas)
	}

	for i := 1; i <= numeroRequest; i++ {
		wgRequests.Add(1)
		chRequests <- urlRequest
	}

	wgRequests.Wait()
	wgRespostas.Wait()

	relatorio.Finaliza()
	relatorio.ImprimirRelatorioConsole()
  }


func worker(chRequests <- chan string,  chRespostas chan <- entity.Resposta, wgRequests *sync.WaitGroup, wgRespostas *sync.WaitGroup){

	for request := range chRequests {

		processaRequest(request,chRespostas, wgRequests, wgRespostas)
	}
}

func processaRequest(urlRequest string, chRespostas chan <- entity.Resposta, wgRequests *sync.WaitGroup, wgRespostas *sync.WaitGroup){
	defer wgRequests.Done() 
	
	inicio := time.Now()
	resp, err := http.Get(urlRequest)
	duracao := time.Since(inicio)

	if err != nil{
		fmt.Println("Erro ao realizar a request")
		return
	}


	wgRespostas.Add(1)
	chRespostas <- entity.Resposta{
		TempoRequest: duracao,
		StatusCode: resp.StatusCode,
	}

}

func processaRespostas(relatorio *entity.Relatorio, chRespostas <- chan entity.Resposta, wgRespostas *sync.WaitGroup){

	for resposta := range chRespostas {

		relatorio.AdicionaResposta(resposta)

		wgRespostas.Done()
	}
}
