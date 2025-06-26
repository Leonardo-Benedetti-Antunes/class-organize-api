package infra

import (
	"log"
	"os"
	"strconv"

	"github.com/cristiantebaldi/class-organize-api/models"
	"github.com/resendlabs/resend-go"
)

func SendEmailOnAlocacaoSuccess(alocacao models.Alocacao) {
	apiKey := os.Getenv("RESEND_API_KEY")
	if apiKey == "" {
		log.Println("RESEND_API_KEY não está definida!")
		return
	}

	client := resend.NewClient(apiKey)

	params := &resend.SendEmailRequest{
		From:    "Teste <noreply@example.com>",
		To:      []string{"classorganizeapi@gmail.com"},
		Subject: "Nova Alocação Criada",
		Html:    "<strong>Uma nova alocação foi criada com sucesso!</strong><br>ID: " + strconv.Itoa(alocacao.ID),
	}

	_, err := client.Emails.Send(params)
	if err != nil {
		log.Println("Erro ao enviar e-mail:", err)
	} else {
		log.Println("E-mail enviado com sucesso.")
	}
}
