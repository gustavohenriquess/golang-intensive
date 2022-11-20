# Intensive Golang

- [Full Cycle repository](https://github.com/devfullcycle/gointensivo)

## O que é, de onde veio e os diferenciais de Go

### Vantagens

- Performance
- Simples de aprender e treinar profissionais
- Redução do poder computacional (consome menos hardware)

### Características

- Opensource
- Criada para utilizar ao máximo os recursos computacionais e de rede
- Extremamente rápida apesar de utilizar garbage collection
- Compilada com a geração de apenas um binário
- Multiplataforma

### Motivações para criação do Golang

- Limitações algumas das principais linguagens utilizadas na Google como Python, java e C++
- Python: problemas com lentidão
- C/C++: Muita complexidade e demorado para compilar
- Java: Complexidade gerada ao longo do tempo / verbosidade da linguagem
- Multithreading e Concorrência: Não nasceram nativamente pensando nisso
- Simplicidade
- Framework de testes e profiling nativos
- Detecção de Race conditions
- Deploy absurdamente simples
- Baixa curva de aprendizado

### Multithreading

- Trabalhar com Multitrheading em GO é simples
- o GO consome apenas 2kb por thread (ao contrario do padrão de 1mb)
- o GO gerência suas próprias threads

## Projeto prático

- Microsserviços de cálculo de taxas de produto
- Recebimento das transações via RabbitMQ
- Cálculo do preço final baseado em uma taxa
- Persistência no banco de dados
- Endpoint de consulta para exibir a quantidade de transações processadas

### Softwares de apoio

- MySQL
- RabbitMQ
- Prometheus
- Grafana
- Docker e Kubernets

### Observações

```
# RabbitMQ
Após subir o docker-compose lembre-se de logar no servidor do RabbitMQ (guest:guest) e:
- Criar Queue: orders
- Efetuar o bind com da Exchange com a Queue: amq.direct
```

```
# Grafana
Após logar no grafana (admin:admin) lembre-se de:
- adicionar o Prometheus como Datasource (http://prometheus:9090)
- Importar um Dashboard com o ID(10991) e adicionar o Prometheus a ele
```
