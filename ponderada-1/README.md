# Ponderada - Simulador de Sensor MiCS-6814 com MQTT

## Como rodar o projeto 
- É necessário primeiramente clonar o repositório utilizando o seguinte comando:
<pre><code>
 git clone  https://github.com/Leao09/Exercicios-prog-M9.git
</code></pre>
- Depois entrar no diretório do projeto
<pre><code>
 cd ponderada-1
</code></pre> 
- Instale todas as dependências do projeto com o comando:
- Obs: esse requirements foi feito utilizando o OS linux com distro Unbuntu, caso você esteja utilizando ou OS verifique se todas as libs são necessárias ou não
<pre><code>
 pip install requitements.txt
</code></pre>
- Agora dentro do diretório, é necessário rodar um comando para rodar o brocker localmente 
<pre><code>
 mosquitto -c mosquitto.conf 
</code></pre> 
- Com o brocker funcionando, bastar rodar o arquivo main.py com o seguinte comando 3
<pre><code>
 python3 main.py
</code></pre> 

## Video 
[Link](https://youtu.be/XyUxLBvz5UQ)