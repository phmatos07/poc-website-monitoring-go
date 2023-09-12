# POC-WEBSITE-MONITORING-GO
**Website Monitoring** é uma prova de conceito para monitorar o status code de Web site, a aplicação foi concebida em GO e, tem o principal objetivo de demostrar as principais funcionalidades que a linguagem proporciona.

## Iniciando o monitoramento
Utilize o comando abaixo para iniciar a execução da aplicação:

~~~Bash
go run src/main.go
~~~

Após a execução do comando será apresentado 4 opção do que desejar executar:

~~~Bash
Olá, esse é o aplicativo para monitoramento de Sites, selecione a opção que atende seu cenário:
1: Monitorar um Site
2: Monitorar vários sites (URL`s) de um arquivo.txt
3: Exibir Logs
4: Fechar aplicação
~~~

### Descrição
1. **Monitorar um Site:** Essa opção analisa o status code de uma única URL fornecida pelo usuário;
2. **Monitorar vários sites:** Essa opção analisa várias URLs fornecida em arquivo *sites.txt*;
3. **Exibir Logs:** A cada execução dos comandos acima a aplicação salva logs internamente e, caso deseje exibi-los no terminal basta usar essa opção;
4. **Fechar aplicação:** Essa opção dispensa explicação, de fato fecha a aplicação.

**Observação:** A cada execução dos comandos acima a aplicação salva logs em *arquivos.txt*, tanto para o monitoramento dos sites e erros que possa ocorrer internamente ao executá-los.

## Pacotes utilizados:
- https://pkg.go.dev/fmt
- https://pkg.go.dev/os
- https://pkg.go.dev/io
- https://pkg.go.dev/io/ioutil
- https://pkg.go.dev/bufio
- https://pkg.go.dev/strconv
- https://pkg.go.dev/time

## Referências
- **GO (Golang):** https://go.dev/
- **Documentação:** https://go.dev/doc/
- **Criando Módulos:** https://go.dev/doc/tutorial/create-module
- **Alura:** https://cursos.alura.com.br/