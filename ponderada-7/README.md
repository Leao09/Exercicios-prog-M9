# Ponderada - Testes no simulador MQTT com integração HIVEMQ

## Como rodar o projeto 
- É necessário primeiramente clonar o repositório utilizando o seguinte comando:
<pre><code>
 git clone  https://github.com/Leao09/Exercicios-prog-M9.git
</code></pre>
- Depois entrar no diretório do projeto
<pre><code>
 cd ponderada-4
</code></pre> 
- Em sequida é necessário ter instalado a linguagem [go](https://go.dev/dl/)  para a execução do projeto 
- Assim, inicie um módulo para o diretório e depois baixe as depencias para o projeto
<pre><code>
 go mod init seu-no
 go mod tidy
</code></pre>
- Como os códigos estão integrados com o Broker HIVEMQ para rodar é necessário que você crie um arquivo .env na raiz da pasta ponderada-4 e configure ele com as suas credenciais para utilizar o seu cluster, dessa forma:
<pre><code>
BROKER_ADDR='sua_url_do_cluster'
HIVE_USER='seu_usuário'
HIVE_PSWD='sua_senha'
MONGODB_USER= 'Seu_usuário'
MONGODB_PASSWORD= 'Sua_senha'
CONFLUENT_BOOTSTRAP_SERVER_SASL='seu_server'
CONFLUENT_API_KEY='sua_api_key'
CONFLUENT_API_SECRET='seu_secret'
CLUSTER_ID='seu_cluster_id'
</code></pre> 
- É importante resaltar que caso você opte por mudar os nomes nomes das variáveis de ambiente no arquivo será necessário realizar mudanças no código.
- Assim para rodar o publisher basta rodar os seguintes comando 
<pre><code>
cd publisher
go run publisher.go
</code></pre> 
- Para rodar o consumer que ira consumir a fila do kafka
<pre><code>
go run main.go
</code></pre> 
Por fim, basta olhar pela interface do confluent caso tenha duvidas se as mensagens estão sendo consumidads e para validação de teste basta olhar o arquivo publisher.txt e consumer.txt se eles estiver iguais significa que está tudo certo


## Video 
[Link](https://youtu.be/Pwgp23dzrlQ)