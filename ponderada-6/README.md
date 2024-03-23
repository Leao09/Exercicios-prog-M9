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
</code></pre> 
- É importante resaltar que caso você opte por mudar os nomes nomes das variáveis de ambiente no arquivo será necessário realizar mudanças no código.
- Assim para rodar o publisher basta rodar os seguintes comando 
<pre><code>
cd publisher
go run publisher.go
</code></pre> 
- Para rodar o subscriber com a API que alimenta o banco de dados com os dados enviados pelo subscriber 
<pre><code>
go run main.go
</code></pre> 
Por fim, com o publisher e subscriber + api rodando, acesse a imagem do container do metabase com o seguinte comando:
<pre><code>
docker run -d -p 3000:3000 -v $(pwd)/ponderada.db:/ponderada.db --name metabse metabase/metabase
</code></pre> 

## Video 
[Link](https://youtu.be/bDzk1H-rOqA)