# Prática de Microsserviços com gRPC
Este repositório contém três microsserviços em Go (Pedido, Pagamento, Envio) que utilizam gRPC e Protocol Buffers.

### Serviços
- **Pedido**: Gerencia a criação e o fluxo de pedidos.

- **Pagamento**: Processa pagamentos e aplica regras de negócio.

- **#Envio**: Gerencia o envio de pedidos e cálculo de dias de entrega

- **db**: Ao inicializar a imagem Mysql, o volume "init.sql" será executado, criando a base de dados e inserindo itens no "inventory" conforme instruções da parte final.

# Executando com Docker
Construa e inicie todos os serviços:

No path onde o docker-compose.yml está localizado execute: 

docker-compose up --build

Cada serviço estará disponível em sua respectiva porta:

    database: 3306

    Pedido: localhost:3010

    Pagamento: 3001

    Envio: 50051


# Exemplos de requisições (Requisições concluídas com sucesso retornam "OrderId" e "deliveryDays")
()

### Requisição válida
grpcurl -d '{\"costumer_id\": 123, \"order_items\": [{\"product_code\": \"ITEM005\", \"quantity\": 10, \"unit_price\": 10}]}' -plaintext localhost:3010 Order/Create


### Requisição inválida - Código de produto não está presente na base de dados
grpcurl -d '{\"costumer_id\": 123, \"order_items\": [{\"product_code\": \"ITEM007\", \"quantity\": 10, \"unit_price\": 10}]}' -plaintext localhost:3010 Order/Create

### Requisição inválida - Quantidade acima de 50
grpcurl -d '{\"costumer_id\": 123, \"order_items\": [{\"product_code\": \"ITEM003\", \"quantity\": 55, \"unit_price\": 10}]}' -plaintext localhost:3010 Order/Create

### Requisição inválida Preço total acima de 1000
grpcurl -d '{\"costumer_id\": 123, \"order_items\": [{\"product_code\": \"ITEM007\", \"quantity\": 49, \"unit_price\": 100}]}' -plaintext localhost:3010 Order/Create

# Adicionalmente, como o campo de unit_price ja existe na tabela, requisições sem esse campo também vão ser executadas e calculadas, porém ao longo do projeto, manipular o unit_price diretamente na requisição se mostrou mais eficiente para troubleshooting
grpcurl -d '{\"costumer_id\": 12, \"order_items\": [{\"product_code\": \"ITEM005\", \"quantity\": 20}]}'

